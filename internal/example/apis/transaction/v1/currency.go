package v1

import (
	"encoding/json"
	"fmt"
)

// Currency 货币单位
// +gengo:enum
type Currency string

const (
	CurrencyCNY Currency = "CNY" // 人民币
)

func (c Currency) ToInt64(v CurrencyValue) int64 {
	return int64(v * 100)
}

func (c Currency) FromInt64(i int64) CurrencyValue {
	if i == 0 {
		return 0
	}
	return CurrencyValue(i) / 100
}

type CurrencyValue float64

var _ json.Marshaler = CurrencyValue(0)

func (v CurrencyValue) String() string {
	return fmt.Sprintf("%.2f", v)
}

func (v CurrencyValue) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%.2f", v)), nil
}

func (v CurrencyValue) Mul(quantity int64) CurrencyValue {
	return v * CurrencyValue(quantity)
}

func (v CurrencyValue) Add(a CurrencyValue) CurrencyValue {
	return v + a
}
