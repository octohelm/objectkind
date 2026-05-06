// +gengo:operator:register=R
//
//go:generate go tool devtool gen .
package order

import (
	"github.com/octohelm/courier/pkg/courier"
)

var R = courier.NewRouter()
