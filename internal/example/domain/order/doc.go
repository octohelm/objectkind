// Package order 订单领域模型定义。
//
//go:generate go tool devtool gen .
package order

import (
	"github.com/octohelm/storage/pkg/sqlbuilder"
)

var T = &sqlbuilder.Tables{}
