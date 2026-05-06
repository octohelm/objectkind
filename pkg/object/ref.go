package object

// AsRefIDGetter 获取引用 ID 的 RefID 表示。
type AsRefIDGetter interface {
	GetAsRefID() RefID
}

// FromRefIDSetter 从 RefID 设置引用 ID。
type FromRefIDSetter interface {
	SetFromRefID(refID RefID)
}

// RefIDConvertable 支持 RefID 与 ID 之间的双向转换。
type RefIDConvertable interface {
	AsRefIDGetter
	FromRefIDSetter
}

// AsRefCodeGetter 获取引用 Code 的 RefCode 表示。
type AsRefCodeGetter interface {
	GetAsRefCode() RefCode
}

// FromRefCodeSetter 从 RefCode 设置引用 Code。
type FromRefCodeSetter interface {
	SetFromRefCode(refCode RefCode)
}

// RefCodeConvertable 支持 RefCode 与 Code 之间的双向转换。
type RefCodeConvertable interface {
	AsRefCodeGetter
	FromRefCodeSetter
}

// RefID 是引用 ID 的类型表示。
//
// +gengo:uintstr
type RefID uint64

// RefCode 是引用 Code 的类型表示。
type RefCode string

// AsRefStringIDGetter 获取引用 ID 的字符串表示。
type AsRefStringIDGetter interface {
	GetAsRefStringID() string
}

// FromRefStringIDSetter 从字符串设置引用 ID。
type FromRefStringIDSetter interface {
	SetFromRefStringID(refID string)
}

// RefStringIDConvertable 支持字符串与引用 ID 之间的双向转换。
type RefStringIDConvertable interface {
	AsRefStringIDGetter
	FromRefStringIDSetter
}
