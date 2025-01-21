package annotate

import (
	"fmt"
	"reflect"
	"testing"

	testingx "github.com/octohelm/x/testing"
)

var AnnotationX = Annotation("x.io/x")

func TestAnnotation(t *testing.T) {
	t.Run("simple value", func(t *testing.T) {
		values := []any{"1", true, false, "abcd"}

		for _, x := range values {
			t.Run(fmt.Sprintf("value %v", x), func(t *testing.T) {
				annotations := Annotations{}

				err := AnnotationX.MarshalTo(annotations, x)
				testingx.Expect(t, err, testingx.BeNil[error]())
				testingx.Expect(t, annotations, testingx.Equal(Annotations{
					"x.io/x": fmt.Sprintf("%v", x),
				}))

				ret := reflect.New(reflect.TypeOf(x)).Elem().Interface()

				err = AnnotationX.UnmarshalFrom(annotations, &ret)
				testingx.Expect(t, ret == x, testingx.BeTrue())
			})
		}
	})

	t.Run("json value", func(t *testing.T) {
		annotations := Annotations{}

		err := AnnotationX.MarshalTo(annotations, X{
			Str: "xxx",
		})
		testingx.Expect(t, err, testingx.BeNil[error]())
		testingx.Expect(t, annotations, testingx.Equal(Annotations{
			"x.io/x": `{"str":"xxx"}`,
		}))

		var value X
		err = AnnotationX.UnmarshalFrom(annotations, &value)
		testingx.Expect(t, err, testingx.BeNil[error]())
		testingx.Expect(t, value.Str, testingx.Equal("xxx"))
	})
}

type X struct {
	Str string `json:"str"`
}
