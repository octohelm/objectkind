package annotate

type Annotations map[string]string

func (a Annotations) SetAnnotation(k string, value string) {
	a[k] = value
}

func (a Annotations) GetAnnotation(k string) (value string, ok bool) {
	v, ok := a[k]
	return v, ok
}
