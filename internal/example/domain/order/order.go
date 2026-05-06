package order

import (
	orderv1 "github.com/octohelm/objectkind/internal/example/pkg/apis/order/v1"
	sqltypecompose "github.com/octohelm/objectkind/pkg/sqltype/compose"
)

// Order
// +gengo:table:register=T
// +gengo:table:group=order
// @def primary ID
// @def index i_state State
type Order struct {
	sqltypecompose.Resource[orderv1.OrderID]

	State       orderv1.OrderState `db:"state,default=1"`
	TotalAmount int                `db:"f_total_amount,default=0"`
}
