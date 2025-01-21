package v1

import "github.com/octohelm/objectkind/pkg/schema"

type Descriptor struct {
	// 名称
	Name string `json:"name,omitzero"`
	// 描述
	Description string `json:"description,omitzero"`
	// 其他注解
	Annotations map[string]string `json:"annotations,omitzero"`
}

var _ interface {
	schema.ObjectDescriptor
	schema.ObjectDescriptorSetter
} = &Descriptor{}

func (v Descriptor) GetName() string {
	return v.Name
}

func (v Descriptor) GetDescription() string {
	return v.Description
}

func (v *Descriptor) SetName(name string) {
	v.Name = name
}

func (v *Descriptor) SetDescription(desc string) {
	v.Description = desc
}

func (v Descriptor) GetAnnotations() map[string]string {
	return v.Annotations
}

func (v Descriptor) GetAnnotation(k string) (value string, ok bool) {
	if v.Annotations == nil {
		return "", false
	}
	vv, ok := v.Annotations[k]
	return vv, ok
}

func (v *Descriptor) SetAnnotation(k string, value string) {
	if v.Annotations == nil {
		v.Annotations = map[string]string{}
	}
	v.Annotations[k] = value
}
