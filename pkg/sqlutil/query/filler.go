package query

import (
	"context"

	"github.com/octohelm/objectkind/pkg/object"
	"github.com/octohelm/objectkind/pkg/sqlutil/query/internal"
	"github.com/octohelm/objectkind/pkg/sqlutil/query/internal/queryopts"
)

// Filler 查询选项，控制是否需要对指定 Object 类型执行填充操作
type Filler[O object.Type] bool

func (Filler[O]) Is(x any) bool {
	_, ok := x.(Filler[O])
	return ok
}

func (Filler[O]) QueryOptions(use internal.NotForPublicUse) {
}

// NeedFiller 从 context 中读取是否需要执行指定 Object 类型的填充操作
func NeedFiller[O object.Type](ctx context.Context) bool {
	v, ok := queryopts.GetOption(queryopts.FromContext(ctx), func(x Filler[O]) queryopts.Options {
		return x
	})
	if ok {
		return bool(v)
	}
	return true
}
