package repository

import (
	"context"

	"github.com/octohelm/storage/pkg/filter"
	sqlpipefilter "github.com/octohelm/storage/pkg/sqlpipe/filter"
	sqltypetime "github.com/octohelm/storage/pkg/sqltype/time"

	"github.com/octohelm/objectkind/internal/example/domain/order"
	orderv1 "github.com/octohelm/objectkind/internal/example/pkg/apis/order/v1"
)

func (repo *OrderRepository) ListExpiredCreatedOrders(ctx context.Context, before sqltypetime.Timestamp) (*orderv1.OrderList, error) {
	return repo.ListOrder(
		ctx,
		sqlpipefilter.AsWhere(order.OrderT.State, filter.Eq(orderv1.ORDER_STATE__CREATED)),
		sqlpipefilter.AsWhere(order.OrderT.CreatedAt, filter.Lt(before)),
	)
}
