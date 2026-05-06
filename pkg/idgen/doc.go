// +gengo:runtimedoc=false
//
//go:generate go tool devtool gen .

// Package idgen 提供全局唯一 ID 生成能力。
// IDGen 基于雪花算法，结合本机 IP 作为 worker ID，通过 injectable 生命周期接入。
package idgen
