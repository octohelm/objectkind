package v1

import "github.com/octohelm/objectkind/pkg/object"

type TypeMeta struct {
	// 资源类型
	Kind string `json:"kind,omitzero"`
	// 资源类型版本
	APIVersion string `json:"apiVersion,omitzero"`
}

var _ object.Type = TypeMeta{}

func (t TypeMeta) GetKind() string       { return t.Kind }
func (t TypeMeta) GetAPIVersion() string { return t.APIVersion }

var _ object.KindAndAPIVersionSetter = &TypeMeta{}

func (t *TypeMeta) SetKind(kind string)          { t.Kind = kind }
func (t *TypeMeta) SetAPIVersion(version string) { t.APIVersion = version }
