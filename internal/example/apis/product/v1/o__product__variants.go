package v1

import metav1 "github.com/octohelm/objectkind/pkg/apis/meta/v1"

// ProductReference 商品 (引用)
// +gengo:objectkind:variant
type ProductReference struct {
	// 商品
	metav1.ObjectReference[Product, ProductID]
}

// ProductRequestForCreate 商品 (创建)
// +gengo:objectkind:variant
type ProductRequestForCreate struct {
	// 商品
	metav1.Request[Product]
}

// ProductRequestForUpdate 商品 (更新)
// +gengo:objectkind:variant
type ProductRequestForUpdate struct {
	// 商品
	metav1.Request[Product]
}
