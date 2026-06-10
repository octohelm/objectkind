package runtime_test

import (
	"fmt"
	"testing"

	"github.com/go-json-experiment/json"

	"github.com/octohelm/courier/pkg/validator/taggedunion"
	. "github.com/octohelm/x/testing/v2"

	metav1 "github.com/octohelm/objectkind/pkg/apis/meta/v1"
	"github.com/octohelm/objectkind/pkg/object"
	"github.com/octohelm/objectkind/pkg/runtime"
)

func TestKindTagged(t *testing.T) {
	u := &Union{}
	u.SetUnderlying(runtime.New[TypeA]())

	t.Run("Underlying 类型校验", func(t *testing.T) {
		Then(
			t, "Underlying 应该是 TypeA 指针",
			Expect(u.Underlying(), Equal(any(runtime.New[TypeA]()))),
		)
	})

	t.Run("JSON 序列化与反序列化", func(t *testing.T) {
		raw := MustValue(t, func() ([]byte, error) {
			return json.Marshal(u)
		})

		Then(
			t, "序列化结果符合预期",
			Expect(string(raw), Equal(`{"kind":"TypeA"}`)),
		)

		t.Run("反序列化", func(t *testing.T) {
			u2 := &Union{}
			Must(t, func() error {
				return json.Unmarshal(raw, u2)
			})

			Then(
				t, "反序列化后的 Underlying 保持一致",
				Expect(u2.Underlying(), Equal(any(runtime.New[TypeA]()))),
			)
		})
	})

	t.Run("IsZero 与非零值", func(t *testing.T) {
		u := &Union{}
		Then(
			t, "空 Union IsZero 返回 true",
			Expect(u.IsZero(), Equal(true)),
		)

		u.SetUnderlying(runtime.New[TypeA]())
		Then(
			t, "有 Underlying 后 IsZero 返回 false",
			Expect(u.IsZero(), Equal(false)),
		)
	})
}

var mappings = runtime.KindTaggedMapping{}

func init() {
	mappings.Register(runtime.BuildFuncFor[TypeA]())
	mappings.Register(runtime.BuildFuncFor[TypeB]())
}

type TypeA struct {
	metav1.TypeMeta
}

func (a TypeA) GetKind() string {
	return "TypeA"
}

type TypeB struct {
	metav1.TypeMeta
}

func (a TypeB) GetKind() string {
	return "TypeB"
}

type Union struct {
	runtime.KindTagged[object.Type] `json:"-"`
}

func (Union) Mapping() map[string]any {
	return mappings.AsMapping()
}

func (p *Union) UnmarshalJSON(data []byte) error {
	pp := Union{}
	if err := taggedunion.Unmarshal(data, &pp); err != nil {
		return fmt.Errorf("unmarshal failed to Union: %w", err)
	}
	*p = pp
	return nil
}
