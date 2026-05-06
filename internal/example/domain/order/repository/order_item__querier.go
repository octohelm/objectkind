package repository

import (
	"context"

	"github.com/octohelm/storage/pkg/filter"
	sqlpipeex "github.com/octohelm/storage/pkg/sqlpipe/ex"

	"github.com/octohelm/objectkind/internal/example/domain/order"
	orderconvert "github.com/octohelm/objectkind/internal/example/domain/order/convert"
	orderfilter "github.com/octohelm/objectkind/internal/example/domain/order/filter"
	orderv1 "github.com/octohelm/objectkind/internal/example/pkg/apis/order/v1"
	sqlutilfiller "github.com/octohelm/objectkind/pkg/sqlutil/filler"
)

type OrderItemQuerier struct {
	OrderItem sqlpipeex.Executor[order.OrderItem]
}

func init() {
	sqlutilfiller.Register(&OrderItemQuerier{})
}

func (q *OrderItemQuerier) FillSet(ctx context.Context, items sqlpipeex.Set[orderv1.OrderItemID, orderv1.OrderItem]) error {
	ex := q.OrderItem.PipeE(&orderfilter.OrderItemByID{
		ID: filter.InSeq(items.Keys()),
	})

	for m, err := range ex.Items(ctx) {
		if err != nil {
			return err
		}

		item, err := orderconvert.OrderItem.ToObject(m)
		if err != nil {
			return err
		}

		for x := range items.Records(m.ID) {
			*x = *item
		}
	}

	return nil
}

func (q *OrderItemQuerier) FillOwnerSet(ctx context.Context, orders sqlpipeex.Set[orderv1.OrderID, orderv1.Order]) error {
	if orders == nil || orders.IsZero() {
		return nil
	}

	ex := q.OrderItem.PipeE(&orderfilter.OrderItemByOrderID{
		OrderID: filter.InSeq(orders.Keys()),
	})

	for m, err := range ex.Items(ctx) {
		if err != nil {
			return err
		}

		item, err := orderconvert.OrderItem.ToObject(m)
		if err != nil {
			return err
		}

		for o := range orders.Records(m.OrderID) {
			o.Spec.Items = append(o.Spec.Items, item)
		}
	}

	return nil
}
