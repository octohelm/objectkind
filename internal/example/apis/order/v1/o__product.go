package v1

import metav1 "github.com/octohelm/objectkind/pkg/apis/meta/v1"

// Product 商品
// +gengo:objectkind
type Product struct {
	// 商品
	metav1.Object[ProductID]

	Status ProductStatus `json:"status"`
}

// ProductID 商品 id
// +gengo:uintstr
type ProductID uint64

// ProductStatus 商品状态
type ProductStatus struct {
	// 是否可用
	Available bool `json:"available,omitzero"`
}
