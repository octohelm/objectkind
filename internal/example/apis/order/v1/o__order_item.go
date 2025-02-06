package v1

import (
	metav1 "github.com/octohelm/objectkind/pkg/apis/meta/v1"
)

// OrderItem 订单项
// +gengo:objectkind
type OrderItem struct {
	metav1.Object[OrderItemID]

	Spec   OrderItemSpec   `json:"spec"`
	Status OrderItemStatus `json:"status"`
}

// OrderItemID
// +gengo:uintstr
type OrderItemID uint64

type OrderItemSpec struct {
	// 商品规格
	Sku *Sku `json:"sku"`
	// 个数
	Quantity int64 `json:"quantity"`
}

type OrderItemStatus struct {
	// 总价
	TotalPrice float64 `json:"totalPrice"`
	// 折扣金额
	DiscountAmount float64 `json:"discountAmount"`
	// 最终价格
	FinalPrice float64 `json:"finalPrice"`
}
