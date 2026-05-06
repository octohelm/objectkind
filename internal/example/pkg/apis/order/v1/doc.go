// Package v1 定义订单域的版本化对象契约，包括订单、订单项、请求变体、列表和业务错误。
//
//go:generate go tool devtool gen .
package v1

import (
	metav1 "github.com/octohelm/objectkind/pkg/apis/meta/v1"
)

var SchemeGroupVersion = metav1.GroupVersion{
	Group:   "order",
	Version: "v1",
}
