package pager

type RawPager struct {
	// 分页偏移
	Offset int64 `name:"offset,omitzero" in:"query"`
	// 分页条数
	Limit int64 `name:"limit,omitzero" validate:"@int[-1,50] = 10" in:"query"`
}
