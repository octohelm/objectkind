//go:generate go tool devtool gen .

// Package digest 提供内容寻址摘要计算能力。
// 支持对对象进行哈希与摘要生成，可跳过已有摘要以优化性能。
package digest
