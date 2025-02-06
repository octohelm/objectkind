package v1

import "github.com/octohelm/objectkind/pkg/object"

type Describer struct {
	// 名称
	Name string `json:"name,omitzero"`
	// 描述
	Description string `json:"description,omitzero"`
	// 其他注解
	Annotations map[string]string `json:"annotations,omitzero"`
}

var _ object.Describer = Describer{}

func (v Describer) GetName() string        { return v.Name }
func (v Describer) GetDescription() string { return v.Description }

var _ object.Describable = &Describer{}

func (v *Describer) SetName(name string)        { v.Name = name }
func (v *Describer) SetDescription(desc string) { v.Description = desc }

var _ object.Annotater = Describer{}

func (v Describer) GetAnnotations() map[string]string { return v.Annotations }
func (v Describer) GetAnnotation(k string) (value string, ok bool) {
	if v.Annotations == nil {
		return "", false
	}
	vv, ok := v.Annotations[k]
	return vv, ok
}

var _ object.Annotatable = &Describer{}

func (v *Describer) SetAnnotations(annotations map[string]string) { v.Annotations = annotations }
func (v *Describer) SetAnnotation(k string, value string) {
	if v.Annotations == nil {
		v.Annotations = map[string]string{}
	}
	v.Annotations[k] = value
}
