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
	// 订单状态
	State OrderState `json:"state,omitzero"`
	// 总金额
	TotalAmount int `json:"totalAmount,omitzero"`
}

// OrderState
// +gengo:enum
type OrderState uint8

const (
	ORDER_STATE_UNKNOWN    OrderState = iota
	ORDER_STATE__CREATED              // 已创建
	ORDER_STATE__PAID                 // 已支付
	ORDER_STATE__CANCELED             // 已取消
	ORDER_STATE__COMPLETED            // 已完成
)
