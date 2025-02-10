package converter

import (
	"github.com/octohelm/objectkind/pkg/object"
	"github.com/octohelm/objectkind/pkg/runtime"
)

func For[M object.Type, O object.Type](
	toObject func(o *O, m *M) error,
	fromObject func(m *M, o *O) error,
) Converter[O, M] {
	return &converter[M, O]{
		toObject:   runtime.ConvertFunc(toObject),
		fromObject: runtime.ConvertFunc(fromObject),
	}
}

type converter[M object.Type, O object.Type] struct {
	fromObject func(o *O) (*M, error)
	toObject   func(m *M) (*O, error)
}

func (c *converter[M, O]) ToObject(m *M) (*O, error) {
	return c.toObject(m)
}

func (c *converter[M, O]) FromObject(o *O) (*M, error) {
	return c.fromObject(o)
}
