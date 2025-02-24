package runtime_test

import (
	"testing"

	productv1 "github.com/octohelm/objectkind/internal/example/apis/product/v1"
	transactionv1 "github.com/octohelm/objectkind/internal/example/apis/transaction/v1"
	productconvert "github.com/octohelm/objectkind/internal/example/domain/product/convert"
	"github.com/octohelm/objectkind/pkg/runtime"
	testingx "github.com/octohelm/x/testing"
)

func TestRuntime(t *testing.T) {
	pdt := runtime.Build(func(p *productv1.Product) {
		p.ID = 1
		p.Name = "product"

		p.Status.State = productv1.PRODUCT_STATE__ON_SALE
	})

	testingx.Expect(t, pdt.Kind, testingx.Be("Product"))
	testingx.Expect(t, pdt.APIVersion, testingx.Be("product/v1"))
	testingx.Expect(t, pdt.ID, testingx.Be(productv1.ProductID(1)))
	testingx.Expect(t, pdt.Name, testingx.Be("product"))

	t.Run("could convert to variant", func(t *testing.T) {
		pdtRef := pdt.AsProductReference()

		testingx.Expect(t, pdtRef.Kind, testingx.Be("Product"))
		testingx.Expect(t, pdtRef.APIVersion, testingx.Be("product/v1"))
		testingx.Expect(t, pdtRef.ID, testingx.Be(productv1.ProductID(1)))
	})

	t.Run("could convert from request for create", func(t *testing.T) {
		orderItemForRequest := runtime.Build(func(o *productv1.SkuRequestForCreate) {
			o.Spec.Price = 1
			o.Spec.Currency = transactionv1.CurrencyCNY
		})

		orderItem := orderItemForRequest.AsSku()

		testingx.Expect(t, orderItem.Spec.Price, testingx.Be(transactionv1.CurrencyValue(1)))
		testingx.Expect(t, orderItem.Spec.Currency, testingx.Be(transactionv1.CurrencyCNY))
	})

	t.Run("could convert to model resource", func(t *testing.T) {
		m, _ := productconvert.Product.FromObject(pdt)

		testingx.Expect(t, m.ID, testingx.Be(productv1.ProductID(1)))
		testingx.Expect(t, m.Name, testingx.Be(pdt.Name))

		testingx.Expect(t, m.State, testingx.Be(pdt.Status.State))

		t.Run("could convert from model resource", func(t *testing.T) {
			pdt2, _ := productconvert.Product.ToObject(m)

			testingx.Expect(t, pdt2.Kind, testingx.Be("Product"))
			testingx.Expect(t, pdt2.APIVersion, testingx.Be("product/v1"))
			testingx.Expect(t, pdt2.ID, testingx.Be(productv1.ProductID(1)))
			testingx.Expect(t, pdt2.Name, testingx.Be("product"))

			testingx.Expect(t, pdt2.Status.State, testingx.Be(m.State))

			testingx.Expect(t, pdt2.CreationTimestamp.IsZero(), testingx.BeFalse())

			testingx.Expect(t, pdt2.CreationTimestamp, testingx.Be(m.CreatedAt))
			testingx.Expect(t, pdt2.ModificationTimestamp, testingx.Be(m.UpdatedAt))
		})
	})
}
