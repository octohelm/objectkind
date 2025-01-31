package schema

import sqltypetime "github.com/octohelm/storage/pkg/sqltype/time"

type Object interface {
	GetObjectKind() ObjectKind
}

type ObjectSetter interface {
	SetGroupVersionKind(gvk GroupVersionKind)
}

type ObjectDescriptor interface {
	GetName() string
	GetDescription() string
	GetAnnotations() map[string]string
	GetAnnotation(k string) (value string, ok bool)
}

type ObjectDescriptorSetter interface {
	SetName(name string)
	SetDescription(description string)
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

type CreationTimestampGetter interface {
	GetCreationTimestamp() sqltypetime.Timestamp
}

type CreationTimestampSetter interface {
	SetCreationTimestamp(creationTimestamp sqltypetime.Timestamp)
}

type ModificationTimestampGetter interface {
	GetModificationTimestamp() sqltypetime.Timestamp
}

type ModificationTimestampSetter interface {
	SetModificationTimestamp(modificationTimestamp sqltypetime.Timestamp)
}
