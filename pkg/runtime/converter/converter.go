package converter

type Converter[O any, M any] interface {
	ToObject(m *M) (*O, error)
	FromObject(o *O) (*M, error)
}

func NoError[D any, S any](fn func(o *D, m *S)) func(o *D, m *S) error {
	return func(o *D, m *S) error {
		fn(o, m)
		return nil
	}
}
