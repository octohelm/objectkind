package object

type Object[ID ~uint64] interface {
	Type
	IDGetter[ID]
}

type Codable[Code ~string] interface {
	Type
	CodeGetter[Code]
}

type CodableObject[ID ~uint64, Code ~string] interface {
	Type
	IDGetter[ID]
	CodeGetter[Code]
}
