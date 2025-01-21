package v1

import "github.com/octohelm/objectkind/pkg/schema"

type ObjectKind = schema.ObjectKind

type TypeMeta struct {
	// 资源类型
	Kind string `json:"kind,omitzero"`
	// 资源类型版本
	APIVersion string `json:"apiVersion,omitzero"`
}

func (v TypeMeta) GetObjectKind() ObjectKind {
	return &v
}

func (v TypeMeta) GroupVersionKind() schema.GroupVersionKind {
	return schema.FromAPIVersionAndKind(v.APIVersion, v.Kind)
}

func (v *TypeMeta) SetGroupVersionKind(gvk schema.GroupVersionKind) {
	v.APIVersion, v.Kind = gvk.ToAPIVersionAndKind()
}
