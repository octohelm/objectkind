package v1

import (
	metav1 "github.com/octohelm/objectkind/pkg/apis/meta/v1"
)

// SkuReference 商品规格
// +gengo:objectkind:variant
type SkuReference struct {
	// 商品规格
	metav1.ObjectReference[Sku, SkuID]
}
