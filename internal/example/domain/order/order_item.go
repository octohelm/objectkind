package order

import (
	orderv1 "github.com/octohelm/objectkind/internal/example/pkg/apis/order/v1"
	productv1 "github.com/octohelm/objectkind/internal/example/pkg/apis/product/v1"
	sqltypecompose "github.com/octohelm/objectkind/pkg/sqltype/compose"
)

// OrderItem
// +gengo:table:register=T
// +gengo:table:group=order
// @def primary ID
// @def unique_index i_order_item OrderID SkuID
// @def index i_order_id OrderID
type OrderItem struct {
	sqltypecompose.Resource[orderv1.OrderItemID]

	OrderID  orderv1.OrderID `db:"f_order_id"`
	SkuID    productv1.SkuID `db:"f_sku_id"`
	Quantity int64           `db:"f_quantity,default=0"`

	TotalPrice     float64 `db:"f_total_price,default=0"`
	DiscountAmount float64 `db:"f_discount_amount,default=0"`
	FinalPrice     float64 `db:"f_final_price,default=0"`
}
