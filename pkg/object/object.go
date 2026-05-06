package object

// Object 是具有 Kind 和 ID 的基础资源。
type Object[ID Identity] interface {
	Type
	IDGetter[ID]
}

// Codable 是具有 Kind 和 Code 的基础资源。
type Codable[Code ~string] interface {
	Type
	CodeGetter[Code]
}

// CodableObject 是同时具备 ID 和 Code 的资源。
type CodableObject[ID Identity, Code ~string] interface {
	Type
	IDGetter[ID]
	CodeGetter[Code]
}
