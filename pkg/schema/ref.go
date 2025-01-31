package schema

import "strconv"

type AsRefIDGetter interface {
	GetAsRefID() RefID
}

type FromRefIDSetter interface {
	SetFromRefID(refID RefID)
}

type AsRefCodeGetter interface {
	GetAsRefCode() RefCode
}

type FromRefCodeSetter interface {
	SetFromRefCode(refCode RefCode)
}

type RefID uint64

func (id *RefID) UnmarshalText(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return nil
	}
	v, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return err
	}
	*id = RefID(v)
	return nil
}

func (id RefID) MarshalText() (text []byte, err error) {
	if id == 0 {
		return nil, nil
	}
	return []byte(strconv.FormatUint(uint64(id), 10)), nil
}

func (id RefID) String() string {
	return strconv.FormatUint(uint64(id), 10)
}

type RefCode string
