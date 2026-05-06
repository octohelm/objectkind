package convert

import (
	"github.com/octohelm/objectkind/internal/example/domain/order"
	orderv1 "github.com/octohelm/objectkind/internal/example/pkg/apis/order/v1"
	runtimeconverter "github.com/octohelm/objectkind/pkg/runtime/converter"
)

var Order = runtimeconverter.ForObject(
	func(o *orderv1.Order, m *order.Order) error {
		o.Status.State = m.State
		o.Status.TotalAmount = m.TotalAmount
		return nil
	},
	func(m *order.Order, o *orderv1.Order) error {
		m.State = o.Status.State
		m.TotalAmount = o.Status.TotalAmount
		return nil
	},
)
