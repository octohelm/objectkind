package repository_test

import (
	"testing"

	"github.com/innoai-tech/infra/pkg/otel"
	"github.com/octohelm/storage/pkg/filter"
	sessiondb "github.com/octohelm/storage/pkg/session/db"
	"github.com/octohelm/x/cmp"
	. "github.com/octohelm/x/testing/v2"

	productv1 "github.com/octohelm/objectkind/internal/example/apis/product/v1"
	transactionv1 "github.com/octohelm/objectkind/internal/example/apis/transaction/v1"
	"github.com/octohelm/objectkind/internal/example/domain/product"
	productfilter "github.com/octohelm/objectkind/internal/example/domain/product/filter"
	productrepository "github.com/octohelm/objectkind/internal/example/domain/product/repository"
	"github.com/octohelm/objectkind/internal/pkg/testingutil"
	"github.com/octohelm/objectkind/pkg/idgen"
	"github.com/octohelm/objectkind/pkg/runtime"
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

	Must(t, func() error {
		return d.PutProducts(ctx, pdt)
	})

	t.Run("保存并查询 SKU", func(t *testing.T) {
		sku := runtime.Build(func(sku *productv1.Sku) {
			sku.Code = "pdt-1122"
			sku.Name = "pdt-1122"
			sku.Spec.Price = 1
			sku.Spec.Currency = transactionv1.CurrencyCNY
			sku.Product = pdt
		})

		Must(t, func() error {
			return d.PutSkus(ctx, sku)
		})

		t.Run("能够通过 Code 列表查询到该 SKU", func(t *testing.T) {
			skuList := MustValue(t, func() (*productv1.SkuList, error) {
				return d.ListSku(ctx, &productfilter.SkuByCode{
					Code: filter.Eq(sku.Code),
				})
			})

			Then(t, "SKU 列表属性正确",
				Expect(skuList.Items, Be(cmp.Len[[]*productv1.Sku](1))),
				Expect(skuList.Items[0].Product.ID, Equal(pdt.ID)),
				Expect(skuList.Items[0].Product.Name, Equal(pdt.Name)),
			)
		})

		t.Run("查询产品时应包含关联的 SKU", func(t *testing.T) {
			productList := MustValue(t, func() (*productv1.ProductList, error) {
				return d.ListProduct(ctx, &productfilter.ProductByID{
					ID: filter.Eq(pdt.ID),
				})
			})

			Then(t, "产品及其关联 SKU 校验",
				Expect(productList.Items, Be(cmp.Len[[]*productv1.Product](1))),
			)

			pdtWithSkus := productList.Items[0]
			Then(t, "SKU 列表符合预期",
				Expect(pdtWithSkus.Skus, Be(cmp.Len[[]*productv1.Sku](1))),
				Expect(pdtWithSkus.Skus[0].Code, Equal(sku.Code)),
			)
		})
	})
}
