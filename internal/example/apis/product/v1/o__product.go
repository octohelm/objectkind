package v1

import metav1 "github.com/octohelm/objectkind/pkg/apis/meta/v1"

type ProductList = metav1.List[Product]

// Product 商品
// +gengo:objectkind
type Product struct {
	// 商品
	metav1.Object[ProductID]

	Status ProductStatus `json:"status"`

	Skus []*Sku `json:"skus,omitempty"`
}

// ProductID 商品 id
// +gengo:uintstr
type ProductID uint64

// ProductStatus 商品状态
type ProductStatus struct {
	State ProductState `json:"state,omitzero"`
}

// ProductState
// +gengo:enum
type ProductState uint8

const (
	PRODUCT_STATE_UNKNOWN   ProductState = iota
	PRODUCT_STATE__ON_SALE               // 上架
	PRODUCT_STATE__OFF_SALE              // 下架
)
