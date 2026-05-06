package repository

import (
	"context"

	"github.com/octohelm/storage/pkg/sqlpipe"

	"github.com/octohelm/objectkind/internal/example/domain/order"
	orderconvert "github.com/octohelm/objectkind/internal/example/domain/order/convert"
	orderv1 "github.com/octohelm/objectkind/internal/example/pkg/apis/order/v1"
	"github.com/octohelm/objectkind/pkg/idgen"
)

// +gengo:injectable
type OrderItemRepository struct {
	OrderItemQuerier

	orderItemID idgen.Typed[orderv1.OrderItemID]
}

func (repo *OrderItemRepository) PutOrderItems(ctx context.Context, orderID orderv1.OrderID, items ...*orderv1.OrderItem) error {
	if len(items) == 0 {
		return nil
	}

	models := make([]*order.OrderItem, 0, len(items))

	for _, item := range items {
		if item.ID == 0 {
			if err := repo.orderItemID.NewTo(&item.ID); err != nil {
				return err
			}
		}

		m, err := orderconvert.OrderItem.FromObject(item)
		if err != nil {
			return err
		}
		m.OrderID = orderID
		models = append(models, m)
	}

	ex := repo.OrderItem.From(sqlpipe.Values(models)).PipeE(
		sqlpipe.OnConflictDoUpdateSet(
			order.OrderItemT.I.IOrderItem,
			order.OrderItemT.Quantity,
			order.OrderItemT.TotalPrice,
			order.OrderItemT.DiscountAmount,
			order.OrderItemT.FinalPrice,
			order.OrderItemT.UpdatedAt,
		),
		sqlpipe.Returning(order.OrderItemT.ID),
	)

	i := 0
	for m, err := range ex.Items(ctx) {
		if err != nil {
			return err
		}
		items[i].ID = m.ID
		i++
	}

	return nil
}
