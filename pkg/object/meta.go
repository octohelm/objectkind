package object

import (
	"iter"
)

// Type 表示具有 Kind 的资源类型。
type Type interface {
	GetKind() string
}

// ParentIter 支持遍历父级资源。
type ParentIter interface {
	Parents() iter.Seq[Type]
}

// APIVersionGetter 获取资源的 API 版本。
type APIVersionGetter interface {
	GetAPIVersion() string
}

// KindSetter 设置资源的 Kind。
type KindSetter interface {
	SetKind(kind string)
}

// APIVersionSetter 设置资源的 API 版本。
type APIVersionSetter interface {
	SetAPIVersion(version string)
}

// PluralizedKindGetter 获取资源的复数形式 Kind。
type PluralizedKindGetter interface {
	GetPluralizedKind() string
}

// PluralizedKindSetter 设置资源的复数形式 Kind。
type PluralizedKindSetter interface {
	SetPluralizedKind(kind string)
}

// Describer 支持获取资源的名称和描述。
type Describer interface {
	GetName() string
	GetDescription() string
}

// Describable 支持设置资源的名称和描述。
type Describable interface {
	SetName(name string)
	SetDescription(description string)
}

// Annotater 支持获取资源的注解。
type Annotater interface {
	GetAnnotations() map[string]string
	GetAnnotation(k string) (value string, ok bool)
}

// Annotatable 支持设置资源的注解。
type Annotatable interface {
	SetAnnotations(annotations map[string]string)
	SetAnnotation(key string, value string)
}

// Identity 是 ID 的约束类型，可为 uint64 或 string。
type Identity interface {
	~uint64 | ~string
}

// IDGetter 获取资源的唯一标识。
type IDGetter[ID Identity] interface {
	GetID() ID
}

// IDSetter 设置资源的唯一标识。
type IDSetter[ID Identity] interface {
	SetID(id ID)
}

// CodeGetter 获取资源的编码。
type CodeGetter[Code ~string] interface {
	GetCode() Code
}

// CodeSetter 设置资源的编码。
type CodeSetter[Code ~string] interface {
	SetCode(code Code)
}

// Copier 支持从其他 Type 复制字段。
type Copier interface {
	CopyFrom(o Type)
}
