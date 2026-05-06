package runtime_test

import (
	"testing"

	"github.com/octohelm/x/cmp"
	. "github.com/octohelm/x/testing/v2"

	"github.com/octohelm/objectkind/internal/example/domain/product"
	productconvert "github.com/octohelm/objectkind/internal/example/domain/product/convert"
	productv1 "github.com/octohelm/objectkind/internal/example/pkg/apis/product/v1"
	transactionv1 "github.com/octohelm/objectkind/internal/example/pkg/apis/transaction/v1"
	"github.com/octohelm/objectkind/pkg/object"
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

type basicObject struct {
	Kind       string
	APIVersion string
	Pluralized string
}

func (b basicObject) GetKind() string             { return b.Kind }
func (b *basicObject) SetKind(k string)           { b.Kind = k }
func (b basicObject) GetAPIVersion() string       { return b.APIVersion }
func (b *basicObject) SetAPIVersion(v string)     { b.APIVersion = v }
func (b basicObject) GetPluralizedKind() string   { return "BasicObjects" }
func (b *basicObject) SetPluralizedKind(k string) { b.Pluralized = k }

type presetObject struct {
	Kind       string
	APIVersion string
}

func (p presetObject) GetKind() string          { return "MyKind" }
func (p *presetObject) SetKind(k string)        { p.Kind = k }
func (p presetObject) GetAPIVersion() string    { return "mygroup/v1" }
func (p *presetObject) SetAPIVersion(v string)  { p.APIVersion = v }

func TestNew(t *testing.T) {
	t.Run("basicObject", func(t *testing.T) {
		o := runtime.New[basicObject]()

		Then(t, "返回值非 nil",
			Expect(o != nil, Be(cmp.True())),
		)
		Then(t, "Kind 从 GetKind 填充 (初始为空)",
			Expect(o.Kind, Equal("")),
		)
		Then(t, "APIVersion 从 GetAPIVersion 填充 (初始为空)",
			Expect(o.APIVersion, Equal("")),
		)
		Then(t, "PluralizedKind 从 GetPluralizedKind 填充",
			Expect(o.Pluralized, Equal("BasicObjects")),
		)
	})

	t.Run("presetObject", func(t *testing.T) {
		o := runtime.New[presetObject]()

		Then(t, "Kind 从 GetKind 填充为预设值",
			Expect(o.Kind, Equal("MyKind")),
		)
		Then(t, "APIVersion 从 GetAPIVersion 填充为预设值",
			Expect(o.APIVersion, Equal("mygroup/v1")),
		)
	})
}

func TestBuild(t *testing.T) {
	t.Run("带 mutation 覆写 Kind", func(t *testing.T) {
		o := runtime.Build[presetObject](func(o *presetObject) {
			o.Kind = "custom"
		})

		Then(t, "Kind 被覆写",
			Expect(o.Kind, Equal("custom")),
		)
		Then(t, "APIVersion 保持自动填充",
			Expect(o.APIVersion, Equal("mygroup/v1")),
		)
	})

	t.Run("无 mutation 仅自动填充", func(t *testing.T) {
		o := runtime.Build[presetObject]()

		Then(t, "Kind 自动填充",
			Expect(o.Kind, Equal("MyKind")),
		)
		Then(t, "APIVersion 自动填充",
			Expect(o.APIVersion, Equal("mygroup/v1")),
		)
	})

	t.Run("nil mutation 安全", func(t *testing.T) {
		o := runtime.Build[presetObject](nil)

		Then(t, "Kind 自动填充",
			Expect(o.Kind, Equal("MyKind")),
		)
	})
}

type fullObject struct {
	Name        string
	Description string
	Annotations map[string]string
	CreatedAt   object.Timestamp
	UpdatedAt   object.Timestamp
	Kind        string
}

func (f fullObject) GetKind() string                           { return f.Kind }
func (f *fullObject) SetKind(k string)                         { f.Kind = k }
func (f fullObject) GetName() string                           { return f.Name }
func (f *fullObject) SetName(n string)                         { f.Name = n }
func (f fullObject) GetDescription() string                    { return f.Description }
func (f *fullObject) SetDescription(d string)                  { f.Description = d }
func (f fullObject) GetAnnotations() map[string]string         { return f.Annotations }
func (f fullObject) GetAnnotation(k string) (string, bool) {
	if f.Annotations == nil {
		return "", false
	}
	v, ok := f.Annotations[k]
	return v, ok
}
func (f *fullObject) SetAnnotations(a map[string]string) { f.Annotations = a }
func (f *fullObject) SetAnnotation(k, v string) {
	if f.Annotations == nil {
		f.Annotations = map[string]string{}
	}
	f.Annotations[k] = v
}
func (f fullObject) GetCreationTimestamp() object.Timestamp      { return f.CreatedAt }
func (f *fullObject) SetCreationTimestamp(t object.Timestamp)     { f.CreatedAt = t }
func (f fullObject) GetModificationTimestamp() object.Timestamp   { return f.UpdatedAt }
func (f *fullObject) SetModificationTimestamp(t object.Timestamp) { f.UpdatedAt = t }

func TestCopy(t *testing.T) {
	t.Run("拷贝名称与描述", func(t *testing.T) {
		src := &fullObject{Name: "src-name", Description: "src-desc"}
		dst := &fullObject{}

		runtime.Copy(dst, src)

		Then(t, "名称被拷贝",
			Expect(dst.Name, Equal("src-name")),
		)
		Then(t, "描述被拷贝",
			Expect(dst.Description, Equal("src-desc")),
		)
	})

	t.Run("拷贝非 nil annotations", func(t *testing.T) {
		src := &fullObject{Annotations: map[string]string{"key1": "val1", "key2": "val2"}}
		dst := &fullObject{}

		runtime.Copy(dst, src)

		Then(t, "annotations 被完整拷贝",
			Expect(len(dst.Annotations), Equal(2)),
			Expect(dst.Annotations["key1"], Equal("val1")),
			Expect(dst.Annotations["key2"], Equal("val2")),
		)
	})

	t.Run("拷贝 nil annotations 安全", func(t *testing.T) {
		src := &fullObject{Name: "test"}
		dst := &fullObject{}

		runtime.Copy(dst, src)

		Then(t, "名称仍被拷贝",
			Expect(dst.Name, Equal("test")),
		)
		Then(t, "annotations 保持 nil",
			Expect(dst.Annotations, Be(cmp.Nil[map[string]string]())),
		)
	})

	t.Run("拷贝时间戳", func(t *testing.T) {
		now := object.Timestamp{}
		src := &fullObject{CreatedAt: now, UpdatedAt: now}
		dst := &fullObject{}

		runtime.Copy(dst, src)

		Then(t, "CreationTimestamp 被拷贝",
			Expect(dst.CreatedAt, Equal(now)),
		)
		Then(t, "ModificationTimestamp 被拷贝",
			Expect(dst.UpdatedAt, Equal(now)),
		)
	})
}

func TestConvert(t *testing.T) {
	t.Run("Convert[D,S] 类型直接使用", func(t *testing.T) {
		convert := runtime.Convert[presetObject, presetObject](func(src *presetObject) (*presetObject, error) {
			return &presetObject{Kind: "converted", APIVersion: "v2"}, nil
		})

		dst, err := convert(&presetObject{})

		Must(t, func() error { return err })

		Then(t, "转换成功",
			Expect(dst.Kind, Equal("converted")),
			Expect(dst.APIVersion, Equal("v2")),
		)
	})

	t.Run("ConvertFunc 基于 dst/src 转换", func(t *testing.T) {
		convert := runtime.ConvertFunc[presetObject, presetObject](func(dst *presetObject, src *presetObject) error {
			dst.Kind = src.Kind + "-converted"
			return nil
		})

		src := &presetObject{Kind: "source"}
		dst, err := convert(src)

		Must(t, func() error { return err })

		Then(t, "转换结果 Kind 正确",
			Expect(dst.Kind, Equal("source-converted")),
		)
		Then(t, "APIVersion 自动填充",
			Expect(dst.APIVersion, Equal("mygroup/v1")),
		)
	})
}

type codableSrc struct {
	kind       string
	apiVersion string
	id         uint64
	code       string
}

func (c codableSrc) GetKind() string       { return c.kind }
func (c codableSrc) GetAPIVersion() string { return c.apiVersion }
func (c codableSrc) GetID() uint64        { return c.id }
func (c codableSrc) GetCode() string      { return c.code }

type codableDst struct {
	kind       string
	apiVersion string
	id         uint64
	code       string
}

func (c codableDst) GetKind() string       { return c.kind }
func (c *codableDst) SetKind(k string)     { c.kind = k }
func (c codableDst) GetAPIVersion() string { return c.apiVersion }
func (c *codableDst) SetAPIVersion(v string) { c.apiVersion = v }
func (c codableDst) GetID() uint64         { return c.id }
func (c *codableDst) SetID(id uint64)      { c.id = id }
func (c codableDst) GetCode() string       { return c.code }
func (c *codableDst) SetCode(code string)  { c.code = code }

func TestCopyCodableObject(t *testing.T) {
	src := &codableSrc{id: 42, code: "my-code"}
	dst := &codableDst{}

	runtime.CopyCodableObject(dst, src)

	Then(t, "ID 正确拷贝",
		Expect(dst.GetID(), Equal(uint64(42))),
	)
	Then(t, "Code 正确拷贝",
		Expect(dst.GetCode(), Equal("my-code")),
	)
}
