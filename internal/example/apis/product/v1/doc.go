//go:generate go tool devtool gen .
package v1

import (
	metav1 "github.com/octohelm/objectkind/pkg/apis/meta/v1"
)

var SchemeGroupVersion = metav1.GroupVersion{
	Group:   "product",
	Version: "v1",
}
