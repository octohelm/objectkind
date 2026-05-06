package v1

import (
	metav1 "github.com/octohelm/objectkind/pkg/apis/meta/v1"
)

// OrderRequestForCreate 订单 (创建)
// +gengo:objectkind:variant
type OrderRequestForCreate struct {
	// 订单
	metav1.Request[Order]

	Spec OrderSpecRequestForCreate `json:"spec"`
}

type OrderSpecRequestForCreate struct {
	// 订单项
	Items []*OrderItemRequestForCreate `json:"items" validate:"@slice[1,]"`
}

// OrderRequestForPay 订单 (支付)
type OrderRequestForPay struct {
	// 订单
	metav1.Request[Order]

	Spec OrderPaymentSpec `json:"spec"`
}

type OrderPaymentSpec struct {
	// 支付渠道
	PaymentChannel OrderPaymentChannel `json:"paymentChannel" validate:"@string[1,32]"`
}

type OrderPaymentChannel string

// OrderRequestForCancel 订单 (取消)
type OrderRequestForCancel struct {
	// 订单
	metav1.Request[Order]

	Spec OrderCancelSpec `json:"spec"`
}

type OrderCancelSpec struct {
	// 取消原因
	Reason OrderCancelReason `json:"reason" validate:"@string[1,128]"`
}

type OrderCancelReason string

// OrderRequestForComplete 订单 (完成)
type OrderRequestForComplete struct {
	// 订单
	metav1.Request[Order]
}
