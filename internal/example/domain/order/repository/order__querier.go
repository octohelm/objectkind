package repository

import (
	"context"
	"iter"

	"github.com/octohelm/storage/pkg/dberr"
	"github.com/octohelm/storage/pkg/filter"
	"github.com/octohelm/storage/pkg/sqlpipe"
	sqlpipeex "github.com/octohelm/storage/pkg/sqlpipe/ex"
	iterx "github.com/octohelm/x/iter"

	"github.com/octohelm/objectkind/internal/example/domain/order"
	orderconvert "github.com/octohelm/objectkind/internal/example/domain/order/convert"
	orderfilter "github.com/octohelm/objectkind/internal/example/domain/order/filter"
	orderv1 "github.com/octohelm/objectkind/internal/example/pkg/apis/order/v1"
	"github.com/octohelm/objectkind/pkg/sqlutil"
	sqlutilfiller "github.com/octohelm/objectkind/pkg/sqlutil/filler"
	sqlutilquery "github.com/octohelm/objectkind/pkg/sqlutil/query"
)

type OrderQuerier struct {
	Order sqlpipeex.Executor[order.Order]
}

func (q *OrderQuerier) Orders(ctx context.Context, ex sqlpipeex.SourceExecutor[order.Order]) iter.Seq2[*orderv1.Order, error] {
	return iterx.Action(func(yield func(*orderv1.Order) bool) error {
		orders := sqlpipeex.OneToOne[orderv1.OrderID, orderv1.Order]{}

		for m, err := range ex.Items(ctx) {
			if err != nil {
				return err
			}

			o, err := orderconvert.Order.ToObject(m)
			if err != nil {
				return err
			}

			if !yield(o) {
				return nil
			}

			orders.Record(o.ID, o)
		}

		if sqlutilquery.NeedSubResources(ctx) {
			if err := sqlutilfiller.FillSubResourcesOfOwnerSet(ctx, orders); err != nil {
				return err
			}
		}

		return nil
	})
}

func (q *OrderQuerier) ListOrder(ctx context.Context, operators ...sqlpipe.SourceOperator[order.Order]) (*orderv1.OrderList, error) {
	return sqlutil.List(ctx, q.Order.PipeE(operators...), q.Orders)
}

func (q *OrderQuerier) FindOneOrder(ctx context.Context, operators ...sqlpipe.SourceOperator[order.Order]) (*orderv1.Order, error) {
	o, err := sqlutil.FindOne(ctx, q.Order.PipeE(operators...), q.Orders)
	if err != nil {
		if dberr.IsErrNotFound(err) {
			return nil, &orderv1.ErrOrderNotFound{}
		}
		return nil, err
	}
	return o, nil
}

func init() {
	sqlutilfiller.Register(&OrderQuerier{})
}

func (q *OrderQuerier) FillSet(ctx context.Context, orders sqlpipeex.Set[orderv1.OrderID, orderv1.Order]) error {
	ex := q.Order.PipeE(&orderfilter.OrderByID{
		ID: filter.InSeq(orders.Keys()),
	})
	return sqlutil.FillSet(ctx, orders, ex, q.Orders)
}
