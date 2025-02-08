package object

type Object[ID Identity] interface {
	Type
	IDGetter[ID]
}

type Codable[Code ~string] interface {
	Type
	CodeGetter[Code]
}

type CodableObject[ID Identity, Code ~string] interface {
	Type
	IDGetter[ID]
	CodeGetter[Code]
}
