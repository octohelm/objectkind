package converter

import (
	"github.com/octohelm/objectkind/pkg/object"
	"github.com/octohelm/objectkind/pkg/runtime"
)

func ForObject[ID object.Identity, M object.Object[ID], O object.Object[ID]](
	toObject func(o *O, m *M) error,
	fromObject func(m *M, o *O) error,
) Converter[O, M] {
	return &objectConverter[ID, M, O]{
		toObject:   runtime.ObjectConvertFunc(toObject),
		fromObject: runtime.ObjectConvertFunc(fromObject),
	}
}

type objectConverter[ID object.Identity, M object.Object[ID], O object.Object[ID]] struct {
	fromObject func(o *O) (*M, error)
	toObject   func(m *M) (*O, error)
}

func (c *objectConverter[ID, M, O]) ToObject(m *M) (*O, error) {
	return c.toObject(m)
}

func (c *objectConverter[ID, M, O]) FromObject(o *O) (*M, error) {
	return c.fromObject(o)
}
