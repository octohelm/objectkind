package schema

type ObjectMapping map[string]func() any

func (m ObjectMapping) Register(build func() any) {
	if o, ok := build().(Object); ok {
		m[o.GetObjectKind().GroupVersionKind().Kind] = build
	}
}

func (m ObjectMapping) AsMapping() map[string]any {
	mm := map[string]any{}
	for k, build := range m {
		mm[k] = build()
	}
	return mm
}
