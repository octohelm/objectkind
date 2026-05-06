// +gengo:operator:register=R
//
//go:generate go tool devtool gen .
package product

import (
	"github.com/octohelm/courier/pkg/courier"
)

var R = courier.NewRouter()
