package service_test

import (
	"testing"

	"github.com/innoai-tech/infra/pkg/configuration/testingutil"
	"github.com/innoai-tech/infra/pkg/otel"
	"github.com/octohelm/storage/pkg/filter"
	sessiondb "github.com/octohelm/storage/pkg/session/db"
	"github.com/octohelm/x/cmp"
	. "github.com/octohelm/x/testing/v2"

	"github.com/octohelm/objectkind/internal/example/domain/product"
	productfilter "github.com/octohelm/objectkind/internal/example/domain/product/filter"
	productrepository "github.com/octohelm/objectkind/internal/example/domain/product/repository"
	productservice "github.com/octohelm/objectkind/internal/example/domain/product/service"
	productv1 "github.com/octohelm/objectkind/internal/example/pkg/apis/product/v1"
	transactionv1 "github.com/octohelm/objectkind/internal/example/pkg/apis/transaction/v1"
	"github.com/octohelm/objectkind/pkg/idgen"
	"github.com/octohelm/objectkind/pkg/runtime"
	"github.com/octohelm/objectkind/pkg/sqlutil/pager"
)

func TestProductService(t *testing.T) {
	ctx, d := testingutil.BuildContext(t, func(d *struct {
		otel.Otel
		idgen.IDGen
		sessiondb.Database
		productrepository.ProductRepository
		productrepository.SkuRepository

		productservice.ProductService
	},
	) {
		d.LogLevel = "debug"
		d.LogFormat = "text"
		d.EnableMigrate = true
		d.ApplyCatalog("test", product.T)
	})

	svc := &d.ProductService

	pdt := runtime.Build(func(p *productv1.Product) {
		p.Name = "服务层产品"
	})

	sku := runtime.Build(func(s *productv1.Sku) {
		s.Code = productv1.SkuCode("svc-001")
		s.Name = "svc-001"
		s.Spec.Price = 99
		s.Spec.Currency = transactionv1.CurrencyCNY
	})

	Must(t, func() error {
		return svc.UpsertProduct(ctx, pdt, sku)
	})

	t.Run("获取商品详情", func(t *testing.T) {
		got := MustValue(t, func() (*productv1.Product, error) {
			return svc.FindOneProduct(ctx, &productfilter.ProductByID{
				ID: filter.Eq(pdt.ID),
			})
		})

		Then(t, "商品及规格正确",
			Expect(got.ID, Equal(pdt.ID)),
			Expect(got.Skus, Be(cmp.Len[[]*productv1.Sku](1))),
			Expect(got.Skus[0].Code, Equal(sku.Code)),
		)
	})

	t.Run("获取规格详情", func(t *testing.T) {
		got := MustValue(t, func() (*productv1.Sku, error) {
			return svc.FindOneSku(ctx, &productfilter.SkuByID{
				ID: filter.Eq(sku.ID),
			})
		})

		Then(t, "规格及所属商品正确",
			Expect(got.ID, Equal(sku.ID)),
			Expect(got.Product.ID, Equal(pdt.ID)),
		)
	})

	t.Run("商品上下架与删除流程", func(t *testing.T) {
		created := MustValue(t, func() (*productv1.Product, error) {
			return svc.CreateProduct(ctx, runtime.Build(func(p *productv1.Product) {
				p.Name = "待上架商品"
				p.Status.State = productv1.PRODUCT_STATE__ON_SALE
			}))
		})

		unpublished := MustValue(t, func() (*productv1.Product, error) {
			return svc.UnpublishProduct(ctx, created.ID)
		})

		Then(t, "下架后状态正确",
			Expect(unpublished.Status.State, Equal(productv1.PRODUCT_STATE__OFF_SALE)),
		)

		published := MustValue(t, func() (*productv1.Product, error) {
			return svc.PublishProduct(ctx, created.ID)
		})

		Then(t, "重新上架后状态正确",
			Expect(published.Status.State, Equal(productv1.PRODUCT_STATE__ON_SALE)),
		)

		createdSku := MustValue(t, func() (*productv1.Sku, error) {
			return svc.CreateSku(ctx, created.ID, runtime.Build(func(s *productv1.Sku) {
				s.Code = productv1.SkuCode("flow-001")
				s.Name = "flow-001"
				s.Spec.Price = 88
				s.Spec.Currency = transactionv1.CurrencyCNY
			}))
		})

		list := MustValue(t, func() (*productv1.SkuList, error) {
			return svc.ListSkuByProductID(ctx, created.ID, pager.Pager[product.Sku]{
				Offset: 0,
				Limit:  10,
			})
		})

		Then(t, "商品规格列表正确",
			Expect(list.Items, Be(cmp.Len[[]*productv1.Sku](1))),
			Expect(list.Items[0].ID, Equal(createdSku.ID)),
		)

		Must(t, func() error {
			return svc.DeleteSku(ctx, createdSku.ID)
		})

		Must(t, func() error {
			return svc.DeleteProduct(ctx, created.ID)
		})

		_, err := svc.FindOneProduct(ctx, &productfilter.ProductByID{
			ID: filter.Eq(created.ID),
		})

		Then(t, "删除后商品不存在",
			Expect(err, ErrorAsType[*productv1.ErrProductNotFound]()),
		)
	})
}
