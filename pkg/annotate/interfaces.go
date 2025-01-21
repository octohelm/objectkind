package annotate

type Provider interface {
	GetAnnotation(k string) (value string, ok bool)
}

type Accessor interface {
	SetAnnotation(k string, value string)
}
