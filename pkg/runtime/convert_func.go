package runtime

import (
	"github.com/octohelm/objectkind/pkg/object"
)

type Convert[D any, S any] func(src *S) (*D, error)

func CodableObjectConvertFunc[D object.Type, ID object.Identity, Code ~string, S object.CodableObject[ID, Code]](convert func(dst *D, src *S) error) Convert[D, S] {
	return func(src *S) (*D, error) {
		dst := New[D]()

		CopyCodableObject(dst, src)

		if err := convert(dst, src); err != nil {
			return nil, err
		}
		return dst, nil
	}
}

func ObjectConvertFunc[D object.Type, ID object.Identity, S object.Object[ID]](convert func(dst *D, src *S) error) Convert[D, S] {
	return func(src *S) (*D, error) {
		dst := New[D]()

		CopyObject(dst, src)

		if err := convert(dst, src); err != nil {
			return nil, err
		}
		return dst, nil
	}
}

func ConvertFunc[D object.Type, S object.Type](convert func(dst *D, src *S) error) Convert[D, S] {
	return func(src *S) (*D, error) {
		dst := New[D]()
		copyObject(dst, src)

		if err := convert(dst, src); err != nil {
			return nil, err
		}
		return dst, nil
	}
}
