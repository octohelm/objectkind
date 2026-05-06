package filler

import (
	"iter"
	"maps"
	"reflect"
)

// Fillers 填充器注册表，按 Object 类型维护填充器映射
type Fillers map[reflect.Type]any

func (fillers Fillers) Register(tpe reflect.Type, filler any) {
	fillers[tpe] = filler
}

func (fillers Fillers) Fillers() iter.Seq[any] {
	return maps.Values(fillers)
}
