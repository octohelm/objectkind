// Package runtime 提供运行时对象构造、拷贝和类型感知转换工具。
// 包括 Build（构造对象并自动填充 Kind/APIVersion）、Copy（对象间拷贝）等。
package runtime

// R 是运行时包的占位类型，用于包级引用。
type R struct{}
