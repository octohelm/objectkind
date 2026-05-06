package query

import (
	"context"

	"github.com/octohelm/objectkind/pkg/sqlutil/query/internal/queryflags"
	"github.com/octohelm/objectkind/pkg/sqlutil/query/internal/queryopts"
)

// Background 创建一个携带完整查询标记的后台 context，与父 context 解耦
func Background(ctx context.Context, options ...Options) context.Context {
	f := queryflags.Flags{}
	f.Set(queryflags.AllFlags)
	return queryopts.InjectContext(ctx, (&queryopts.Opt{Flags: f}).Join(options...))
}

// Options 查询选项类型别名，用于控制查询行为
type Options = queryopts.Options

// With 将查询选项注入 context，用于控制后续查询的填充行为
func With(ctx context.Context, options ...Options) context.Context {
	if len(options) == 0 {
		return ctx
	}

	return queryopts.InjectContext(ctx, queryopts.FromContext(ctx).Join(options...))
}
