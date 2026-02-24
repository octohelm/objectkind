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

type X struct {
	Str string `json:"str"`
}
