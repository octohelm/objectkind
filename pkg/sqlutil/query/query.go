package query

import (
	"context"

	"github.com/octohelm/objectkind/pkg/sqlutil/query/internal/queryflags"
	"github.com/octohelm/objectkind/pkg/sqlutil/query/internal/queryopts"
)

func Background(ctx context.Context, options ...Options) context.Context {
	f := queryflags.Flags{}
	f.Set(queryflags.AllFlags)
	return queryopts.InjectContext(ctx, (&queryopts.Opt{Flags: f}).Join(options...))
}

type Options = queryopts.Options

func With(ctx context.Context, options ...Options) context.Context {
	if len(options) == 0 {
		return ctx
	}

	return queryopts.InjectContext(ctx, queryopts.FromContext(ctx).Join(options...))
}
