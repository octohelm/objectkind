package v1

import (
	metav1 "github.com/octohelm/objectkind/pkg/apis/meta/v1"
)

// Order
// +gengo:objectkind
type Order struct {
	metav1.Object[OrderID]

	Spec   OrderSpec   `json:"spec"`
	Status OrderStatus `json:"status"`
}

// OrderID
// +gengo:uintstr
type OrderID uint64

type OrderSpec struct {
	// 订单项
	Items []*OrderItem `json:"items" validate:"@slice[1,]"`
}

type OrderStatus struct {
	// 总金额
	TotalAmount int `json:"totalAmount,omitzero"`
}
