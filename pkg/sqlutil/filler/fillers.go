package filler

import (
	"iter"
	"maps"
	"reflect"
)

type Fillers map[reflect.Type]any

func (fillers Fillers) Register(tpe reflect.Type, filler any) {
	fillers[tpe] = filler
}

func (fillers Fillers) Fillers() iter.Seq[any] {
	return maps.Values(fillers)
}
