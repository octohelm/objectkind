package order

import (
	"context"

	orderdomain "github.com/octohelm/objectkind/internal/example/domain/order"
	orderservice "github.com/octohelm/objectkind/internal/example/domain/order/service"
	endpointorderv1 "github.com/octohelm/objectkind/internal/example/pkg/endpoints/order/v1"
	"github.com/octohelm/objectkind/pkg/sqlutil/pager"
)

// +gengo:injectable
// +gengo:operator
type ListOrder struct {
	endpointorderv1.ListOrder

	orderService *orderservice.OrderService `inject:""`
}

func (op *ListOrder) Output(ctx context.Context) (any, error) {
	return op.orderService.ListOrder(
		ctx,
		&pager.Pager[orderdomain.Order]{
			Offset: op.Offset,
			Limit:  op.Limit,
		},
	)
}

// +gengo:injectable
// +gengo:operator
type GetOrderByID struct {
	endpointorderv1.GetOrderByID

	orderService *orderservice.OrderService `inject:""`
}

func (op *GetOrderByID) Output(ctx context.Context) (any, error) {
	return op.orderService.GetOrder(ctx, op.OrderID)
}

// +gengo:injectable
// +gengo:operator
type CreateOrder struct {
	endpointorderv1.CreateOrder

	orderService *orderservice.OrderService `inject:""`
}

func (op *CreateOrder) Output(ctx context.Context) (any, error) {
	return op.orderService.CreateOrder(ctx, &op.Body)
}

// +gengo:injectable
// +gengo:operator
type PayOrderByID struct {
	endpointorderv1.PayOrderByID

	orderService *orderservice.OrderService `inject:""`
}

func (op *PayOrderByID) Output(ctx context.Context) (any, error) {
	return op.orderService.PayOrder(ctx, op.OrderID, op.Body.Spec.PaymentChannel)
}

// +gengo:injectable
// +gengo:operator
type CancelOrderByID struct {
	endpointorderv1.CancelOrderByID

	orderService *orderservice.OrderService `inject:""`
}

func (op *CancelOrderByID) Output(ctx context.Context) (any, error) {
	return op.orderService.CancelOrder(ctx, op.OrderID, op.Body.Spec.Reason)
}

// +gengo:injectable
// +gengo:operator
type CompleteOrderByID struct {
	endpointorderv1.CompleteOrderByID

	orderService *orderservice.OrderService `inject:""`
}

func (op *CompleteOrderByID) Output(ctx context.Context) (any, error) {
	return op.orderService.CompleteOrder(ctx, op.OrderID)
}
