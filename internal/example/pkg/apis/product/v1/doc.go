// Package v1 定义商品域的版本化对象契约，包括资源对象、请求变体、引用、列表和业务错误。
// +gengo:runtimedoc
package v1

import (
	metav1 "github.com/octohelm/objectkind/pkg/apis/meta/v1"
)

var SchemeGroupVersion = metav1.GroupVersion{
	Group:   "product",
	Version: "v1",
}
