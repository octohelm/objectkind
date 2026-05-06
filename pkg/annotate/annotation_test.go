package annotate_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/octohelm/x/cmp"
	. "github.com/octohelm/x/testing/v2"

	"github.com/octohelm/objectkind/pkg/annotate"
)

var AnnotationX = annotate.Annotation("x.io/x")

func TestAnnotation(t *testing.T) {
	t.Run("简单值处理", func(t *testing.T) {
		values := []any{"1", true, false, "abcd"}

		for _, x := range values {
			t.Run(fmt.Sprintf("当值为 %v 时", x), func(t *testing.T) {
				annotations := annotate.Annotations{}

				Must(t, func() error {
					return AnnotationX.MarshalTo(annotations, x)
				})

				Then(t, "Annotations 映射正确",
					Expect(annotations, Equal(annotate.Annotations{
						"x.io/x": fmt.Sprintf("%v", x),
					})),
				)

				// Unmarshal
				ret := reflect.New(reflect.TypeOf(x)).Elem().Interface()
				Must(t, func() error {
					return AnnotationX.UnmarshalFrom(annotations, &ret)
				})

				Then(t, "反序列化后的值相等",
					Expect(ret == x, Be(cmp.True())),
				)
			})
		}
	})

	t.Run("JSON 对象处理", func(t *testing.T) {
		annotations := annotate.Annotations{}
		val := X{Str: "xxx"}

		t.Run("序列化对象", func(t *testing.T) {
			Must(t, func() error {
				return AnnotationX.MarshalTo(annotations, val)
			})

			Then(t, "序列化为 JSON 字符串",
				Expect(annotations, Equal(annotate.Annotations{
					"x.io/x": `{"str":"xxx"}`,
				})),
			)
		})

		t.Run("反序列化对象", func(t *testing.T) {
			var decoded X
			Must(t, func() error {
				return AnnotationX.UnmarshalFrom(annotations, &decoded)
			})

			Then(t, "字段值恢复正确",
				Expect(decoded.Str, Equal("xxx")),
			)
		})
	})
}

func TestAnnotationGet(t *testing.T) {
	annotations := annotate.Annotations{"x.io/x": "hello"}
	v, ok := AnnotationX.Get(annotations)
	Then(t, "值正确", Expect(v, Equal("hello")))
	Then(t, "key 存在", Expect(ok, Equal(true)))

	v, ok = annotate.Annotation("nonexistent").Get(annotations)
	Then(t, "key 不存在", Expect(ok, Equal(false)))
	Then(t, "值为空", Expect(v, Equal("")))
}

type X struct {
	Str string `json:"str"`
}

func TestAnnotations(t *testing.T) {
	t.Run("零值 Annotations{} 长度为 0", func(t *testing.T) {
		a := annotate.Annotations{}
		Then(t, "长度为 0", Expect(len(a), Equal(0)))
	})

	t.Run("SetAnnotation 添加键值对", func(t *testing.T) {
		a := annotate.Annotations{}
		a.SetAnnotation("key1", "value1")
		Then(t, "映射包含该键", Expect(a, Equal(annotate.Annotations{"key1": "value1"})))
	})

	t.Run("GetAnnotation 返回存在的值", func(t *testing.T) {
		a := annotate.Annotations{"key1": "value1"}
		v, ok := a.GetAnnotation("key1")
		Then(t, "返回值和 ok=true",
			Expect(v, Equal("value1")),
			Expect(ok, Be(cmp.True())),
		)
	})

	t.Run("GetAnnotation 对不存在的键返回空和 false", func(t *testing.T) {
		a := annotate.Annotations{}
		v, ok := a.GetAnnotation("missing")
		Then(t, "返回空字符串和 false",
			Expect(v, Equal("")),
			Expect(ok, Be(cmp.False())),
		)
	})

	t.Run("覆盖已有键更新值", func(t *testing.T) {
		a := annotate.Annotations{"key1": "value1"}
		a.SetAnnotation("key1", "newvalue")
		Then(t, "值被更新", Expect(a, Equal(annotate.Annotations{"key1": "newvalue"})))
	})

	t.Run("设置多个键", func(t *testing.T) {
		a := annotate.Annotations{}
		a.SetAnnotation("k1", "v1")
		a.SetAnnotation("k2", "v2")
		Then(t, "包含所有键", Expect(a, Equal(annotate.Annotations{"k1": "v1", "k2": "v2"})))
	})
}

func TestAnnotationString(t *testing.T) {
	ann := annotate.Annotation("test.io/test")
	Then(t, "字符串值为 test.io/test", Expect(string(ann), Equal("test.io/test")))
}

func TestAnnotationMarshalUnmarshalComposite(t *testing.T) {
	t.Run("int 值 42 读写往返", func(t *testing.T) {
		a := annotate.Annotations{}
		ann := annotate.Annotation("test.io/int")
		Must(t, func() error { return ann.MarshalTo(a, 42) })
		var v int
		Must(t, func() error { return ann.UnmarshalFrom(a, &v) })
		Then(t, "反序列化值为 42", Expect(v, Equal(42)))
	})

	t.Run("float64 值 3.14 读写往返", func(t *testing.T) {
		a := annotate.Annotations{}
		ann := annotate.Annotation("test.io/float")
		Must(t, func() error { return ann.MarshalTo(a, 3.14) })
		var v float64
		Must(t, func() error { return ann.UnmarshalFrom(a, &v) })
		Then(t, "反序列化值为 3.14", Expect(v, Equal(3.14)))
	})

	t.Run("嵌套结构体读写往返", func(t *testing.T) {
		a := annotate.Annotations{}
		ann := annotate.Annotation("test.io/struct")
		val := X{Str: "hello"}
		Must(t, func() error { return ann.MarshalTo(a, val) })
		var v X
		Must(t, func() error { return ann.UnmarshalFrom(a, &v) })
		Then(t, "反序列化后结构体字段一致", Expect(v, Equal(val)))
	})

	t.Run("数组/切片读写往返", func(t *testing.T) {
		a := annotate.Annotations{}
		ann := annotate.Annotation("test.io/slice")
		val := []string{"a", "b", "c"}
		Must(t, func() error { return ann.MarshalTo(a, val) })
		var v []string
		Must(t, func() error { return ann.UnmarshalFrom(a, &v) })
		Then(t, "反序列化后切片一致", Expect(v, Equal(val)))
	})
}
