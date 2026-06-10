// Package v1 定义交易相关的通用契约类型，例如货币单位与金额表达。
// +gengo:runtimedoc
package v1

import (
	metav1 "github.com/octohelm/objectkind/pkg/apis/meta/v1"
)

var SchemeGroupVersion = metav1.GroupVersion{
	Group:   "transaction",
	Version: "v1",
}
