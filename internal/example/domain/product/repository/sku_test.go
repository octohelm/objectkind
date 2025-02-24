package repository_test

import (
	"testing"

	transactionv1 "github.com/octohelm/objectkind/internal/example/apis/transaction/v1"

	"github.com/innoai-tech/infra/pkg/otel"
	productv1 "github.com/octohelm/objectkind/internal/example/apis/product/v1"
	"github.com/octohelm/objectkind/internal/example/domain/product"
	productfilter "github.com/octohelm/objectkind/internal/example/domain/product/filter"
	productrepository "github.com/octohelm/objectkind/internal/example/domain/product/repository"
	"github.com/octohelm/objectkind/internal/pkg/testingutil"
	"github.com/octohelm/objectkind/pkg/idgen"
	"github.com/octohelm/objectkind/pkg/runtime"
	"github.com/octohelm/storage/pkg/filter"
	sessiondb "github.com/octohelm/storage/pkg/session/db"
	testingx "github.com/octohelm/x/testing"
)

func TestSkuRepository(t *testing.T) {
	d := &struct {
		otel.Otel
		idgen.IDGen

		sessiondb.Database

		productrepository.ProductRepository
		productrepository.SkuRepository
	}{}

	d.LogLevel = "debug"
	d.LogFormat = "text"
	d.EnableMigrate = true

	d.ApplyCatalog("test", product.T)

	ctx := testingutil.NewContext(t, d)

	pdt := runtime.Build(func(pdt *productv1.Product) {
		pdt.Name = "测试产品"
	})

	t.Run("put product", func(t *testing.T) {
		err := d.PutProducts(ctx, pdt)
		testingx.Expect(t, err, testingx.BeNil[error]())
	})

	t.Run("put sku", func(t *testing.T) {
		sku := runtime.Build(func(sku *productv1.Sku) {
			sku.Code = "pdt-1122"
			sku.Name = "pdt-1122"

			sku.Spec.Price = 1
			sku.Spec.Currency = transactionv1.CurrencyCNY

			sku.Product = pdt
		})

		err := d.PutSkus(ctx, sku)
		testingx.Expect(t, err, testingx.BeNil[error]())

		t.Run("then skus should be listed", func(t *testing.T) {
			skuList, err := d.ListSku(ctx, &productfilter.SkuByCode{
				Code: filter.Eq(sku.Code),
			})
			testingx.Expect(t, err, testingx.BeNil[error]())
			testingx.Expect(t, len(skuList.Items), testingx.Be(1))

			pdtOfSku := skuList.Items[0]
			testingx.Expect(t, pdtOfSku.Product.ID, testingx.Be(pdt.ID))
			testingx.Expect(t, pdtOfSku.Product.Name, testingx.Be(pdt.Name))
		})

		t.Run("then product listed should be includes sku", func(t *testing.T) {
			productList, err := d.ListProduct(ctx, &productfilter.ProductByID{
				ID: filter.Eq(pdt.ID),
			})
			testingx.Expect(t, err, testingx.BeNil[error]())
			testingx.Expect(t, len(productList.Items), testingx.Be(1))

			pdtWithSkus := productList.Items[0]
			testingx.Expect(t, len(pdtWithSkus.Skus), testingx.Be(1))
			testingx.Expect(t, pdtWithSkus.Skus[0].Code, testingx.Be(sku.Code))
		})
	})
}
