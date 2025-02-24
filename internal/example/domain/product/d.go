//go:generate go tool devtool gen .
package product

import (
	"github.com/octohelm/storage/pkg/sqlbuilder"
)

var T = &sqlbuilder.Tables{}
