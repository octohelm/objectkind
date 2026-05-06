package v1

// List 泛型资源列表，包含条目与总数
type List[T any] struct {
	// 列表
	Items []*T `json:"items,omitzero"`
	// 总数
	Total int64 `json:"total,omitzero"`
}

func (v *List[T]) Add(item *T) {
	v.Items = append(v.Items, item)
}
