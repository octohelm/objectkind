package service

import (
	"context"
	"time"

	"github.com/octohelm/storage/pkg/filter"
	"github.com/octohelm/storage/pkg/sqlpipe"
	sqltypetime "github.com/octohelm/storage/pkg/sqltype/time"

	"github.com/octohelm/objectkind/internal/example/domain/order"
	orderfilter "github.com/octohelm/objectkind/internal/example/domain/order/filter"
	orderrepository "github.com/octohelm/objectkind/internal/example/domain/order/repository"
	productservice "github.com/octohelm/objectkind/internal/example/domain/product/service"
	orderv1 "github.com/octohelm/objectkind/internal/example/pkg/apis/order/v1"
	productv1 "github.com/octohelm/objectkind/internal/example/pkg/apis/product/v1"
	sqlutilfiller "github.com/octohelm/objectkind/pkg/sqlutil/filler"
)

// +gengo:injectable:provider
type OrderService struct {
	orderrepository.OrderRepository
	orderrepository.OrderItemRepository
	Products productservice.ProductService
}

func (svc *OrderService) CreateOrder(ctx context.Context, req *orderv1.OrderRequestForCreate) (*orderv1.Order, error) {
	order := req.AsOrder()
	order.Status.State = orderv1.ORDER_STATE__CREATED
	order.Status.TotalAmount = 0

	if err := fillOrderSkus(ctx, order); err != nil {
		return nil, err
	}

	for _, item := range order.Spec.Items {
		if item.Spec.Sku == nil {
			continue
		}

		item.Status.TotalPrice = float64(item.Spec.Sku.Spec.Price) * float64(item.Spec.Quantity)
		item.Status.DiscountAmount = 0
		item.Status.FinalPrice = item.Status.TotalPrice

		order.Status.TotalAmount += int(item.Status.FinalPrice)
	}

	if err := svc.OrderRepository.PutOrders(ctx, order); err != nil {
		return nil, err
	}

	if err := svc.OrderItemRepository.PutOrderItems(ctx, order.ID, order.Spec.Items...); err != nil {
		return nil, err
	}

	return svc.GetOrder(ctx, order.ID)
}

func (svc *OrderService) GetOrder(ctx context.Context, id orderv1.OrderID) (*orderv1.Order, error) {
	order, err := svc.OrderRepository.FindOneOrder(ctx, &orderfilter.OrderByID{
		ID: filter.Eq(id),
	})
	if err != nil {
		return nil, err
	}

	if err := fillOrderSkus(ctx, order); err != nil {
		return nil, err
	}

	return order, nil
}

func (svc *OrderService) ListOrder(ctx context.Context, operators ...sqlpipe.SourceOperator[order.Order]) (*orderv1.OrderList, error) {
	orders, err := svc.OrderRepository.ListOrder(ctx, operators...)
	if err != nil {
		return nil, err
	}

	if err := fillOrderSkus(ctx, orders.Items...); err != nil {
		return nil, err
	}

	return orders, nil
}

func (svc *OrderService) PayOrder(ctx context.Context, id orderv1.OrderID, _ orderv1.OrderPaymentChannel) (*orderv1.Order, error) {
	order, err := svc.GetOrder(ctx, id)
	if err != nil {
		return nil, err
	}

	if order.Status.State != orderv1.ORDER_STATE__CREATED {
		return nil, &orderv1.ErrOrderStateConflict{}
	}

	order.Status.State = orderv1.ORDER_STATE__PAID
	if err := svc.OrderRepository.PutOrders(ctx, order); err != nil {
		return nil, err
	}

	return order, nil
}

func (svc *OrderService) CancelOrder(ctx context.Context, id orderv1.OrderID, _ orderv1.OrderCancelReason) (*orderv1.Order, error) {
	order, err := svc.GetOrder(ctx, id)
	if err != nil {
		return nil, err
	}

	if order.Status.State != orderv1.ORDER_STATE__CREATED {
		return nil, &orderv1.ErrOrderStateConflict{}
	}

	order.Status.State = orderv1.ORDER_STATE__CANCELED
	if err := svc.OrderRepository.PutOrders(ctx, order); err != nil {
		return nil, err
	}

	return order, nil
}

func (svc *OrderService) CompleteOrder(ctx context.Context, id orderv1.OrderID) (*orderv1.Order, error) {
	order, err := svc.GetOrder(ctx, id)
	if err != nil {
		return nil, err
	}

	if order.Status.State != orderv1.ORDER_STATE__PAID {
		return nil, &orderv1.ErrOrderStateConflict{}
	}

	order.Status.State = orderv1.ORDER_STATE__COMPLETED
	if err := svc.OrderRepository.PutOrders(ctx, order); err != nil {
		return nil, err
	}

	return order, nil
}

func (svc *OrderService) CloseExpiredOrders(ctx context.Context, expireBefore time.Time) (int, error) {
	orders, err := svc.OrderRepository.ListExpiredCreatedOrders(ctx, sqltypetime.Timestamp(expireBefore))
	if err != nil {
		return 0, err
	}

	closed := 0

	for _, order := range orders.Items {
		if _, err := svc.CancelOrder(ctx, order.ID, orderv1.OrderCancelReason("expired")); err != nil {
			return closed, err
		}

		closed++
	}

	return closed, nil
}

func fillOrderSkus(ctx context.Context, orders ...*orderv1.Order) error {
	skus := make([]*productv1.Sku, 0)

	for _, order := range orders {
		if order == nil {
			continue
		}

		for _, item := range order.Spec.Items {
			if item.Spec.Sku == nil {
				continue
			}

			skus = append(skus, item.Spec.Sku)
		}
	}

	return sqlutilfiller.Fill(ctx, skus...)
}
