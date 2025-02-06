package annotate

type Provider interface {
	GetAnnotation(k string) (value string, ok bool)
}

type Setter interface {
	SetAnnotation(k string, value string)
}
