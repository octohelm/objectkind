/*
Package v1 GENERATED BY gengo:uintstr
DON'T EDIT THIS FILE
*/
package v1

import (
	strconv "strconv"
)

func (id *ProductID) UnmarshalText(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return nil
	}
	v, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return err
	}
	*id = ProductID(v)
	return nil
}

func (id ProductID) MarshalText() (text []byte, err error) {
	if id == 0 {
		return nil, nil
	}
	return []byte(strconv.FormatUint(uint64(id), 10)), nil
}

func (id ProductID) String() string {
	return strconv.FormatUint(uint64(id), 10)
}

func (id *SkuID) UnmarshalText(text []byte) error {
	str := string(text)
	if len(str) == 0 {
		return nil
	}
	v, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return err
	}
	*id = SkuID(v)
	return nil
}

func (id SkuID) MarshalText() (text []byte, err error) {
	if id == 0 {
		return nil, nil
	}
	return []byte(strconv.FormatUint(uint64(id), 10)), nil
}

func (id SkuID) String() string {
	return strconv.FormatUint(uint64(id), 10)
}
