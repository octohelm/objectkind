package runtime

import (
	"github.com/octohelm/courier/pkg/validator"
	"github.com/octohelm/objectkind/pkg/object"
	"github.com/octohelm/x/anyjson"
)

func BuildFuncFor[M object.Type]() func() object.Type {
	return func() object.Type {
		return any(New[M]()).(object.Type)
	}
}

type KindTaggedUnionMapping map[string]func() object.Type

func (m KindTaggedUnionMapping) Register(build func() object.Type) {
	o := build()
	m[o.GetKind()] = build
}

func (m KindTaggedUnionMapping) AsMapping() map[string]any {
	mm := map[string]any{}
	for k, build := range m {
		mm[k] = build()
	}
	return mm
}

type KindTaggedUnion struct {
	Object object.Type
}

func (p KindTaggedUnion) GetKind() string {
	if p.Object == nil {
		return ""
	}
	return p.Object.GetKind()
}

func (p KindTaggedUnion) IsZero() bool {
	return p.Object == nil
}

func (d *KindTaggedUnion) SetUnderlying(u any) {
	d.Object = u.(object.Type)
}

func (p KindTaggedUnion) MarshalJSON() ([]byte, error) {
	if p.Object == nil {
		return []byte("null"), nil
	}
	return validator.Marshal(p.Object)
}

func (KindTaggedUnion) Discriminator() string {
	return "kind"
}

func (v *KindTaggedUnion) As(target any) error {
	if v.Object != nil {
		valuer, err := anyjson.FromValue(v.Object)
		if err != nil {
			return err
		}
		return anyjson.As(valuer, target)
	}
	return nil
}

func (d *KindTaggedUnion) AsObject() (*anyjson.Object, bool) {
	if d.Object != nil {
		valuer, err := anyjson.FromValue(d.Object)
		if err != nil {
			return nil, false
		}
		o, ok := valuer.(*anyjson.Object)
		return o, ok
	}
	return nil, false
}
