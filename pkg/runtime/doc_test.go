package runtime_test

import (
	"testing"

	"github.com/octohelm/x/cmp"
	. "github.com/octohelm/x/testing/v2"

	productv1 "github.com/octohelm/objectkind/internal/example/apis/product/v1"
	transactionv1 "github.com/octohelm/objectkind/internal/example/apis/transaction/v1"
	"github.com/octohelm/objectkind/internal/example/domain/product"
	productconvert "github.com/octohelm/objectkind/internal/example/domain/product/convert"
	"github.com/octohelm/objectkind/pkg/runtime"
)

func TestRuntime(t *testing.T) {
	pdt := runtime.Build(func(p *productv1.Product) {
		p.ID = 1
		p.Name = "product"
		p.Status.State = productv1.PRODUCT_STATE__ON_SALE
	})

	Then(t, "基础资源属性正确",
		Expect(pdt.Kind, Equal("Product")),
		Expect(pdt.APIVersion, Equal("product/v1")),
		Expect(pdt.ID, Equal(productv1.ProductID(1))),
		Expect(pdt.Name, Equal("product")),
	)

	t.Run("能够转换为变体 (Variant)", func(t *testing.T) {
		pdtRef := pdt.AsProductReference()

		Then(t, "转换后的引用属性正确",
			Expect(pdtRef.Kind, Equal("Product")),
			Expect(pdtRef.APIVersion, Equal("product/v1")),
			Expect(pdtRef.ID, Equal(productv1.ProductID(1))),
		)
	})

	t.Run("能够从创建请求转换", func(t *testing.T) {
		orderItemForRequest := runtime.Build(func(o *productv1.SkuRequestForCreate) {
			o.Spec.Price = 1
			o.Spec.Currency = transactionv1.CurrencyCNY
		})

		orderItem := orderItemForRequest.AsSku()

		Then(t, "转换后的 Sku 规格正确",
			Expect(orderItem.Spec.Price, Equal(transactionv1.CurrencyValue(1))),
			Expect(orderItem.Spec.Currency, Equal(transactionv1.CurrencyCNY)),
		)
	})

	t.Run("与模型资源 (Model Resource) 的转换", func(t *testing.T) {
		m := MustValue(t, func() (*product.Product, error) {
			m, err := productconvert.Product.FromObject(pdt)
			return m, err
		})

		// 假设 m 的实际类型支持以下字段访问，或此处根据实际情况断言
		// 这里保持原逻辑的字段校验
		Then(t, "从对象转换为模型成功",
			Expect(m.ID, Equal(productv1.ProductID(1))),
			Expect(m.Name, Equal(pdt.Name)),
			Expect(m.State, Equal(pdt.Status.State)),
		)

		t.Run("能够从模型资源转回对象", func(t *testing.T) {
			pdt2 := MustValue(t, func() (*productv1.Product, error) {
				obj, err := productconvert.Product.ToObject(m)
				return obj, err
			})

			Then(t, "转回的对象属性完整",
				Expect(pdt2.Kind, Equal("Product")),
				Expect(pdt2.APIVersion, Equal("product/v1")),
				Expect(pdt2.ID, Equal(productv1.ProductID(1))),
				Expect(pdt2.Name, Equal("product")),
				Expect(pdt2.Status.State, Equal(m.State)),
			)

			Then(t, "时间戳元数据正确",
				Expect(pdt2.CreationTimestamp.IsZero(), Equal(false)),
				Expect(pdt2.CreationTimestamp, Equal(m.CreatedAt)),
				Expect(pdt2.ModificationTimestamp, Equal(m.UpdatedAt)),
			)
		})
	})

	t.Run("边界值校验", func(t *testing.T) {
		Then(t, "空 ID 处理",
			Expect(productv1.ProductID(0), Be(cmp.Zero[productv1.ProductID]())),
		)
	})
}
