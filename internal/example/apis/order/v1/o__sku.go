package v1

import (
	metav1 "github.com/octohelm/objectkind/pkg/apis/meta/v1"
)

type SkuList = metav1.List[Sku]

// Sku 商品规格
// +gengo:objectkind
type Sku struct {
	// 商品规格
	metav1.CodableObject[SkuID, SkuCode]
	// 商品规格属性
	Spec SkuSpec `json:"spec"`
	// 所属商品
	Product *Product `json:"product,omitzero" as:"owner"`
}

// SkuID
// +gengo:uintstr
type SkuID uint64

type SkuCode string

type SkuSpec struct {
	// 单价
	Price float64 `json:"price"`
}
