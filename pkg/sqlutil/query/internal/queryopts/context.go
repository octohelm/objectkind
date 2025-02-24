package queryopts

import (
	"context"

	"github.com/octohelm/objectkind/pkg/sqlutil/query/internal/queryflags"
)

type structContext struct{}

func InjectContext(ctx context.Context, x *Opt) context.Context {
	return context.WithValue(ctx, structContext{}, x)
}

func FromContext(ctx context.Context) *Opt {
	if v, ok := ctx.Value(structContext{}).(*Opt); ok {
		return v
	}
	f := queryflags.Flags{}
	f.Set(queryflags.AllFlags)
	return &Opt{Flags: f}
}
