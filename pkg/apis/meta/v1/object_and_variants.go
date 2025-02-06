package v1

import "github.com/octohelm/objectkind/pkg/object"

type Metadata struct {
	TypeMeta
	Describer
	OperationTimestamps
}

type Object[ID ~uint64] struct {
	Metadata
	Identifiable[ID]
}

type ObjectReference[O object.Object[ID], ID ~uint64] struct {
	TypeMeta
	Identifiable[ID]
}

type CodableReference[O object.Codable[Code], Code ~string] struct {
	TypeMeta
	Codable[Code]
}

type CodableObject[ID ~uint64, Code ~string] struct {
	Object[ID]
	Codable[Code]
}

type Request[O object.Type] struct {
	TypeMeta
	Describer
}

type CodableRequest[O object.Codable[Code], Code ~string] struct {
	TypeMeta
	Describer
	Codable[Code]
}
