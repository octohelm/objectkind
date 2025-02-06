package v1

import metav1 "github.com/octohelm/objectkind/pkg/apis/meta/v1"

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
