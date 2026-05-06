//go:generate go tool devtool gen .

// Package pager 提供分页操作符。
// Pager 实现 sqlpipe.Operator，支持 offset/limit 分页。
package pager
