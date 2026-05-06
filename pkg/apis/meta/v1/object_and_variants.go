package v1

import (
	"github.com/octohelm/objectkind/pkg/object"
)

// Metadata 资源元数据组合，包含类型、描述与操作时间戳
type Metadata struct {
	TypeMeta
	Describer
	OperationTimestamps
}

// Object 标准资源对象，包含完整元数据与数字 ID
type Object[ID ~uint64] struct {
	Metadata
	Identifiable[ID]
}

// ObjectReference 资源引用，包含类型信息与数字 ID
type ObjectReference[O object.Object[ID], ID ~uint64] struct {
	TypeMeta
	Identifiable[ID]
}

// CodableReference 编码资源引用，包含类型信息与编码
type CodableReference[O object.Codable[Code], Code ~string] struct {
	TypeMeta
	Codable[Code]
}

// CodableObject 编码资源对象，包含完整元数据、数字 ID 与编码
type CodableObject[ID ~uint64, Code ~string] struct {
	Object[ID]
	Codable[Code]
}

// Request 通用请求结构，包含类型与描述信息
type Request[O object.Type] struct {
	TypeMeta
	Describer
}

// CodableRequest 编码请求结构，包含类型、描述与编码
type CodableRequest[O object.Codable[Code], Code ~string] struct {
	TypeMeta
	Describer
	Codable[Code]
}

// Response 通用响应结构，包含完整元数据与数字 ID
type Response[O object.Type, ID ~uint64] struct {
	Metadata
	Identifiable[ID]
}

// CodableResponse 编码响应结构，包含完整元数据、数字 ID 与编码
type CodableResponse[O object.Type, ID ~uint64, Code ~string] struct {
	Response[O, ID]
	Codable[Code]
}
