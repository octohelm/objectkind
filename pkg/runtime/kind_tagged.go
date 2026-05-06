package runtime

import (
	"github.com/octohelm/courier/pkg/validator"
	"github.com/octohelm/x/anyjson"

	"github.com/octohelm/objectkind/pkg/object"
)

// KindTaggedMapping 维护 Kind 到对象工厂函数的映射。
type KindTaggedMapping map[string]func() object.Type

// Register 注册一个对象工厂函数到映射中。
func (m KindTaggedMapping) Register(build func() object.Type) {
	o := build()
	m[o.GetKind()] = build
}

// AsMapping 将映射转换为示例对象的映射。
func (m KindTaggedMapping) AsMapping() map[string]any {
	mm := map[string]any{}
	for k, build := range m {
		mm[k] = build()
	}
	return mm
}

// KindTagged 包装支持 Kind 辨别的泛型对象。
type KindTagged[O object.Type] struct {
	o *O
}

// GetKind 返回底层对象的 Kind。
func (p KindTagged[O]) GetKind() string {
	if p.o == nil {
		return ""
	}
	return any(p.o).(object.Type).GetKind()
}

// IsZero 判断底层对象是否为空。
func (p KindTagged[O]) IsZero() bool {
	return p.o == nil
}

// MarshalJSON 实现 json.Marshaler 接口。
func (p KindTagged[O]) MarshalJSON() ([]byte, error) {
	if p.o == nil {
		return []byte("null"), nil
	}
	return validator.Marshal(p.o)
}

// Discriminator 返回用于辨别的字段名。
func (KindTagged[O]) Discriminator() string {
	return "kind"
}

// SetUnderlying 设置底层对象。
func (p *KindTagged[O]) SetUnderlying(v any) {
	switch x := v.(type) {
	case O:
		p.o = &x
	default:
	}
}

// Underlying 返回底层对象。
func (p KindTagged[O]) Underlying() any {
	if p.o == nil {
		return nil
	}
	return *p.o
}

// As 将底层对象转换为目标类型。
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

// AsObject 将底层对象转换为 anyjson.Object。
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
