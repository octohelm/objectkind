package v1

import (
	metav1 "github.com/octohelm/objectkind/pkg/apis/meta/v1"
)

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

// ProductRequestForPublish 商品 (上架)
// +gengo:objectkind:variant
type ProductRequestForPublish struct {
	// 商品
	metav1.Request[Product]
}

// ProductRequestForUnpublish 商品 (下架)
// +gengo:objectkind:variant
type ProductRequestForUnpublish struct {
	// 商品
	metav1.Request[Product]
}
