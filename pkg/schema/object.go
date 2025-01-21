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

type ObjectWithID[ID ~uint64] interface {
	GetID() ID
}

type ObjectIDSetter[ID ~uint64] interface {
	SetID(id ID)
}

type ObjectWithCode[Code ~string] interface {
	GetCode() Code
}

type ObjectCodeSetter[Code ~string] interface {
	SetCode(code Code)
}

type ObjectWithCreationTimestamp interface {
	GetCreationTimestamp() sqltypetime.Timestamp
}

type ObjectCreationTimestampSetter interface {
	SetCreationTimestamp(creationTimestamp sqltypetime.Timestamp)
}

type ObjectWithModificationTimestamp interface {
	GetModificationTimestamp() sqltypetime.Timestamp
}

type ObjectModificationTimestampSetter interface {
	SetModificationTimestamp(modificationTimestamp sqltypetime.Timestamp)
}
