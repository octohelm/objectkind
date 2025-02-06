package v1

import metav1 "github.com/octohelm/objectkind/pkg/apis/meta/v1"

// OrderItemRequestForCreate 订单项 (更新)
// +gengo:objectkind:variant
type OrderItemRequestForCreate struct {
	// 订单项
	metav1.Request[OrderItem]

	Spec OrderItemSpecRequestForCreate `json:"spec"`
}

type OrderItemSpecRequestForCreate struct {
	// 商品规格
	Sku SkuReference `json:"sku"`
	// 个数
	Quantity int64 `json:"quantity"`
}
