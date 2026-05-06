package runtime

import (
	"github.com/octohelm/objectkind/pkg/object"
)

// Convert 定义从源类型 S 到目标类型 D 的转换函数。
type Convert[D any, S any] func(src *S) (*D, error)

// CodableObjectConvertFunc 创建 CodableObject 类型的转换函数，自动处理 Code 和 Object 信息的拷贝。
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

// ObjectConvertFunc 创建 Object 类型的转换函数，自动处理 ID 和 Type 信息的拷贝。
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

// ConvertFunc 创建普通 Type 类型的转换函数，自动处理 Type 信息的拷贝。
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
