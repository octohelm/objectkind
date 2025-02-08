package runtime_test

import (
	"fmt"
	"github.com/go-json-experiment/json"
	"github.com/octohelm/courier/pkg/validator/taggedunion"
	metav1 "github.com/octohelm/objectkind/pkg/apis/meta/v1"
	"github.com/octohelm/objectkind/pkg/object"
	"github.com/octohelm/objectkind/pkg/runtime"
	testingx "github.com/octohelm/x/testing"
	"testing"
)

func TestKindTagged(t *testing.T) {
	u := &Union{}
	u.SetUnderlying(runtime.New[TypeA]())

	t.Run("underlying should be TypeA pointer", func(t *testing.T) {
		testingx.Expect(t, u.Underlying(), testingx.Equal(any(runtime.New[TypeA]())))
	})

	t.Run("should be json marshaled", func(t *testing.T) {
		raw, err := json.Marshal(u)
		testingx.Expect(t, err, testingx.BeNil[error]())
		testingx.Expect(t, string(raw), testingx.Be(`{"kind":"TypeA"}`))

		t.Run("should be json unmarshalled", func(t *testing.T) {
			u2 := &Union{}
			err := json.Unmarshal(raw, u2)
			testingx.Expect(t, err, testingx.BeNil[error]())
			testingx.Expect(t, u2.Underlying(), testingx.Equal(any(runtime.New[TypeA]())))
		})
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
