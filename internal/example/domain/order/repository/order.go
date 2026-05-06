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
type OrderRepository struct {
	OrderQuerier

	orderID idgen.Typed[orderv1.OrderID]
}

func (repo *OrderRepository) PutOrders(ctx context.Context, orders ...*orderv1.Order) error {
	if len(orders) == 0 {
		return nil
	}

	mOrders := make([]*order.Order, 0, len(orders))

	for _, item := range orders {
		if item.ID == 0 {
			if err := repo.orderID.NewTo(&item.ID); err != nil {
				return err
			}
		}

		m, err := orderconvert.Order.FromObject(item)
		if err != nil {
			return err
		}
		mOrders = append(mOrders, m)
	}

	ex := repo.Order.From(sqlpipe.Values(
		mOrders,
		order.OrderT.ID,
		order.OrderT.Name,
		order.OrderT.Description,
		order.OrderT.CreatedAt,
		order.OrderT.UpdatedAt,
		order.OrderT.State,
		order.OrderT.TotalAmount,
	)).PipeE(
		sqlpipe.OnConflictDoUpdateSet(
			order.OrderT.I.Primary,
			order.OrderT.Name,
			order.OrderT.Description,
			order.OrderT.State,
			order.OrderT.TotalAmount,
			order.OrderT.UpdatedAt,
		),
		sqlpipe.Returning(
			order.OrderT.ID,
			order.OrderT.CreatedAt,
		),
	)

	i := 0
	for m, err := range ex.Items(ctx) {
		if err != nil {
			return err
		}
		orders[i].ID = m.ID
		orders[i].CreationTimestamp = m.CreatedAt
		i++
	}

	return nil
}
