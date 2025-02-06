package object

type RefIDConvertable interface {
	GetAsRefID() RefID
	SetFromRefID(refID RefID)
}

type RefCodeConvertable interface {
	GetAsRefCode() RefCode
	SetFromRefCode(refCode RefCode)
}

// RefID
// +gengo:uintstr
type RefID uint64

type RefCode string
