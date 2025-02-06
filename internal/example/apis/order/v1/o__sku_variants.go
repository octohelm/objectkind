package v1

import metav1 "github.com/octohelm/objectkind/pkg/apis/meta/v1"

// SkuReference 商品规格
// +gengo:objectkind:variant
type SkuReference struct {
	// 商品规格
	metav1.ObjectReference[Sku, SkuID]
}

// SkuRequestForCreate 商品规格
// +gengo:objectkind:variant
type SkuRequestForCreate struct {
	// 商品规格
	metav1.CodableRequest[Sku, SkuCode]
	// 商品规格属性
	Spec SkuSpec `json:"spec"`
}

// SkuRequestForUpdate 商品规格
// +gengo:objectkind:variant
type SkuRequestForUpdate struct {
	// 商品规格
	metav1.Request[Sku]
	// 商品规格属性
	Spec SkuSpec `json:"spec"`
}
