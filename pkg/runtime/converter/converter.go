package converter

import (
	"github.com/octohelm/objectkind/pkg/object"
	"github.com/octohelm/objectkind/pkg/runtime"
)

type ConvertFunc[D any, S any] func(o *D, m *S) error

func NoError[D any, S any](fn func(o *D, m *S)) ConvertFunc[D, S] {
	return func(o *D, m *S) error {
		fn(o, m)
		return nil
	}
}

type Converter[O any, M any] interface {
	ToObject(m *M) (*O, error)
	FromObject(o *O) (*M, error)
}

func ForObject[ID ~uint64, M object.Object[ID], O object.Object[ID]](
	toObject ConvertFunc[O, M],
	fromObject ConvertFunc[M, O],
) Converter[O, M] {
	return &objectConverter[ID, M, O]{
		toObject:   runtime.ObjectConvertFunc(toObject),
		fromObject: runtime.ObjectConvertFunc(fromObject),
	}
}

type objectConverter[ID ~uint64, M object.Object[ID], O object.Object[ID]] struct {
	fromObject func(o *O) (*M, error)
	toObject   func(m *M) (*O, error)
}

func (c *objectConverter[ID, M, O]) ToObject(m *M) (*O, error) {
	return c.toObject(m)
}

func (c *objectConverter[ID, M, O]) FromObject(o *O) (*M, error) {
	return c.fromObject(o)
}

func ForCodableObject[ID ~uint64, Code ~string, M object.CodableObject[ID, Code], O object.CodableObject[ID, Code]](
	toObject ConvertFunc[O, M],
	fromObject ConvertFunc[M, O],
) Converter[O, M] {
	return &codableObjectConverter[ID, Code, M, O]{
		toObject:   runtime.CodableObjectConvertFunc(toObject),
		fromObject: runtime.CodableObjectConvertFunc(fromObject),
	}
}

type codableObjectConverter[ID ~uint64, Code ~string, M object.Object[ID], O object.Object[ID]] struct {
	fromObject func(o *O) (*M, error)
	toObject   func(m *M) (*O, error)
}

func (c *codableObjectConverter[ID, Code, M, O]) ToObject(m *M) (*O, error) {
	return c.toObject(m)
}

func (c *codableObjectConverter[ID, Code, M, O]) FromObject(o *O) (*M, error) {
	return c.fromObject(o)
}
