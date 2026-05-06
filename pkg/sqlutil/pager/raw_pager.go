package pager

// RawPager 原始分页结构，用于传递偏移量和条数限制参数
type RawPager struct {
	// 分页偏移
	Offset int64 `name:"offset,omitzero" in:"query"`
	// 分页条数
	Limit int64 `name:"limit,omitzero" validate:"@int[-1,50] = 10" in:"query"`
}
