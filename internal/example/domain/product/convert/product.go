package convert

import (
	"github.com/octohelm/objectkind/internal/example/domain/product"
	productv1 "github.com/octohelm/objectkind/internal/example/pkg/apis/product/v1"
	runtimeconverter "github.com/octohelm/objectkind/pkg/runtime/converter"
)

var Product = runtimeconverter.ForObject(
	func(o *productv1.Product, m *product.Product) error {
		o.Status.State = m.State
		return nil
	},
	func(m *product.Product, o *productv1.Product) error {
		m.State = o.Status.State
		return nil
	},
)
