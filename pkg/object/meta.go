package object

type Type interface {
	GetKind() string
}

type APIVersionGetter interface {
	GetAPIVersion() string
}

type KindAndAPIVersionSetter interface {
	SetKind(kind string)
	SetAPIVersion(version string)
}

type Describer interface {
	GetName() string
	GetDescription() string
}

type Describable interface {
	SetName(name string)
	SetDescription(description string)
}

type Annotater interface {
	GetAnnotations() map[string]string
	GetAnnotation(k string) (value string, ok bool)
}

type Annotatable interface {
	SetAnnotations(annotations map[string]string)
	SetAnnotation(key string, value string)
}

type IDGetter[ID ~uint64] interface {
	GetID() ID
}

type IDSetter[ID ~uint64] interface {
	SetID(id ID)
}

type CodeGetter[Code ~string] interface {
	GetCode() Code
}

type CodeSetter[Code ~string] interface {
	SetCode(code Code)
}

type Copier interface {
	CopyFrom(o Type)
}
