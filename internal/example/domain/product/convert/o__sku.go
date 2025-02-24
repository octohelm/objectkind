package convert

import (
	productv1 "github.com/octohelm/objectkind/internal/example/apis/product/v1"
	"github.com/octohelm/objectkind/internal/example/domain/product"
	"github.com/octohelm/objectkind/pkg/annotate"
	"github.com/octohelm/objectkind/pkg/runtime"
	runtimeconverter "github.com/octohelm/objectkind/pkg/runtime/converter"
)

var Sku = runtimeconverter.ForCodableObject(
	func(o *productv1.Sku, m *product.Sku) error {
		o.Spec.Currency = m.Currency
		o.Spec.Price = o.Spec.Currency.FromInt64(m.Price)

		if v := m.Annotations.Get(); v != nil {
			o.Annotations = *v
		}

		o.Product = runtime.Build(func(p *productv1.Product) {
			p.ID = m.ProductID
		})

		return nil
	},
	func(m *product.Sku, o *productv1.Sku) error {
		m.Annotations.Set((*annotate.Annotations)(&o.Annotations))

		m.Currency = o.Spec.Currency
		m.Price = m.Currency.ToInt64(o.Spec.Price)

		if o.Product != nil {
			m.ProductID = o.Product.ID
		}

		return nil
	},
)
