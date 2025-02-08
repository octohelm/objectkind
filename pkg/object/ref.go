package object

type AsRefIDGetter interface {
	GetAsRefID() RefID
}

type FromRefIDSetter interface {
	SetFromRefID(refID RefID)
}

type RefIDConvertable interface {
	AsRefIDGetter
	FromRefIDSetter
}

type AsRefCodeGetter interface {
	GetAsRefCode() RefCode
}

type FromRefCodeSetter interface {
	SetFromRefCode(refCode RefCode)
}

type RefCodeConvertable interface {
	AsRefCodeGetter
	FromRefCodeSetter
}

// RefID
// +gengo:uintstr
type RefID uint64

type RefCode string

type AsRefStringIDGetter interface {
	GetAsRefStringID() string
}

type FromRefStringIDSetter interface {
	SetFromRefStringID(refID string)
}

type RefStringIDConvertable interface {
	AsRefStringIDGetter
	FromRefStringIDSetter
}
