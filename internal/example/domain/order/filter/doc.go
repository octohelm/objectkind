// Package filter 订单查询过滤条件。
//
//go:generate go tool devtool gen .
package filter

import (
	"github.com/octohelm/objectkind/internal/example/domain/order"
)

// +gengo:filterop
type filterOfOrder struct {
	order.Order
}

// +gengo:filterop
type filterOfOrderItem struct {
	order.OrderItem
}
