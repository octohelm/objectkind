package runtime

import (
	"github.com/octohelm/courier/pkg/validator"
	"github.com/octohelm/objectkind/pkg/object"
	"github.com/octohelm/x/anyjson"
)

type KindTaggedMapping map[string]func() object.Type

func (m KindTaggedMapping) Register(build func() object.Type) {
	o := build()
	m[o.GetKind()] = build
}

func (m KindTaggedMapping) AsMapping() map[string]any {
	mm := map[string]any{}
	for k, build := range m {
		mm[k] = build()
	}
	return mm
}

type KindTagged[O object.Type] struct {
	o *O
}

func (p KindTagged[O]) GetKind() string {
	if p.o == nil {
		return ""
	}
	return any(p.o).(object.Type).GetKind()
}

func (p KindTagged[O]) IsZero() bool {
	return p.o == nil
}

func (p KindTagged[O]) MarshalJSON() ([]byte, error) {
	if p.o == nil {
		return []byte("null"), nil
	}
	return validator.Marshal(p.o)
}

func (KindTagged[O]) Discriminator() string {
	return "kind"
}

func (p *KindTagged[O]) SetUnderlying(v any) {
	switch x := v.(type) {
	case O:
		p.o = &x
	default:
	}
}

func (p KindTagged[O]) Underlying() any {
	if p.o == nil {
		return nil
	}
	return *p.o
}

func (v *KindTagged[O]) As(target any) error {
	if v.o != nil {
		valuer, err := anyjson.FromValue(v.o)
		if err != nil {
			return err
		}
		return anyjson.As(valuer, target)
	}
	return nil
}

func (d *KindTagged[O]) AsObject() (*anyjson.Object, bool) {
	if d.o != nil {
		valuer, err := anyjson.FromValue(d.o)
		if err != nil {
			return nil, false
		}
		o, ok := valuer.(*anyjson.Object)
		return o, ok
	}
	return nil, false
}
