package object

type Type interface {
	GetKind() string
}

type APIVersionGetter interface {
	GetAPIVersion() string
}

type KindSetter interface {
	SetKind(kind string)
}

type APIVersionSetter interface {
	SetAPIVersion(version string)
}

type PluralizedKindGetter interface {
	GetPluralizedKind() string
}

type PluralizedKindSetter interface {
	SetPluralizedKind(kind string)
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

type Identity interface {
	~uint64 | ~string
}

type IDGetter[ID Identity] interface {
	GetID() ID
}

type IDSetter[ID Identity] interface {
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
