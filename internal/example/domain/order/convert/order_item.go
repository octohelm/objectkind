package convert

import (
	"github.com/octohelm/objectkind/internal/example/domain/order"
	orderv1 "github.com/octohelm/objectkind/internal/example/pkg/apis/order/v1"
	productv1 "github.com/octohelm/objectkind/internal/example/pkg/apis/product/v1"
	"github.com/octohelm/objectkind/pkg/runtime"
	runtimeconverter "github.com/octohelm/objectkind/pkg/runtime/converter"
)

var OrderItem = runtimeconverter.ForObject(
	func(o *orderv1.OrderItem, m *order.OrderItem) error {
		o.Spec.Quantity = m.Quantity
		o.Status.TotalPrice = m.TotalPrice
		o.Status.DiscountAmount = m.DiscountAmount
		o.Status.FinalPrice = m.FinalPrice

		o.Spec.Sku = runtime.Build(func(sku *productv1.Sku) {
			sku.ID = m.SkuID
		})

		return nil
	},
	func(m *order.OrderItem, o *orderv1.OrderItem) error {
		m.Quantity = o.Spec.Quantity
		m.TotalPrice = o.Status.TotalPrice
		m.DiscountAmount = o.Status.DiscountAmount
		m.FinalPrice = o.Status.FinalPrice

		if o.Spec.Sku != nil {
			m.SkuID = o.Spec.Sku.ID
		}

		return nil
	},
)
