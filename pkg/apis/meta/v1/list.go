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

// 分页
type Pager struct {
	// 分页偏移
	Offset int64 `name:"offset,omitzero" in:"query"`
	// 分页数
	Limit int64 `name:"limit,omitzero"  in:"query" validate:"@int[-1,50] = 10"`
}
