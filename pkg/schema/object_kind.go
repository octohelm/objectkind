package schema

type ObjectKind interface {
	GroupVersionKind() GroupVersionKind
}

type ObjectKindSetter interface {
	SetGroupVersionKind(kind GroupVersionKind)
}
