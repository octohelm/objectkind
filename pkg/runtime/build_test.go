package runtime_test

import (
	"testing"

	orderv1 "github.com/octohelm/objectkind/internal/example/apis/order/v1"
	"github.com/octohelm/objectkind/pkg/runtime"
	sqltypecompose "github.com/octohelm/objectkind/pkg/sqltype/compose"
	testingx "github.com/octohelm/x/testing"
)

func TestRuntime(t *testing.T) {
	pdt := runtime.Build(func(p *orderv1.Product) {
		p.ID = 1
		p.Name = "product"
	})

	testingx.Expect(t, pdt.Kind, testingx.Be("Product"))
	testingx.Expect(t, pdt.APIVersion, testingx.Be("order/v1"))
	testingx.Expect(t, pdt.ID, testingx.Be(orderv1.ProductID(1)))
	testingx.Expect(t, pdt.Name, testingx.Be("product"))

	t.Run("could convert to variant", func(t *testing.T) {
		pdtRef := pdt.AsProductReference()

		testingx.Expect(t, pdtRef.Kind, testingx.Be("Product"))
		testingx.Expect(t, pdtRef.APIVersion, testingx.Be("order/v1"))
		testingx.Expect(t, pdtRef.ID, testingx.Be(orderv1.ProductID(1)))
	})

	t.Run("could convert from request for create", func(t *testing.T) {
		orderItemForRequest := runtime.Build(func(o *orderv1.OrderItemRequestForCreate) {
			o.Spec.Sku.ID = 1
			o.Spec.Quantity = 10
		})

		orderItem := orderItemForRequest.AsOrderItem()

		testingx.Expect(t, orderItem.Spec.Sku.ID, testingx.Be(orderv1.SkuID(1)))
		testingx.Expect(t, orderItem.Spec.Quantity, testingx.Be(int64(10)))
	})

	t.Run("could convert to model resource", func(t *testing.T) {
		m := &MProduct{}
		runtime.CopyObject(m, pdt)

		testingx.Expect(t, m.ID, testingx.Be(orderv1.ProductID(1)))
		testingx.Expect(t, m.Name, testingx.Be(pdt.Name))

		t.Run("could convert from model resource", func(t *testing.T) {
			pdt2 := runtime.Build(func(p *orderv1.Product) {
				runtime.CopyObject(p, m)
			})

			testingx.Expect(t, pdt2.Kind, testingx.Be("Product"))
			testingx.Expect(t, pdt2.APIVersion, testingx.Be("order/v1"))
			testingx.Expect(t, pdt2.ID, testingx.Be(orderv1.ProductID(1)))
			testingx.Expect(t, pdt2.Name, testingx.Be("product"))

			testingx.Expect(t, pdt2.CreationTimestamp, testingx.Be(m.CreatedAt))
			testingx.Expect(t, pdt2.ModificationTimestamp, testingx.Be(m.UpdatedAt))
		})
	})
}

type MProduct struct {
	sqltypecompose.Resource[orderv1.ProductID]
}

func (MProduct) GetKind() string {
	return "MProduct"
}
