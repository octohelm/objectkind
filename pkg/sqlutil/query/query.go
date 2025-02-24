package query

import (
	"context"

	"github.com/octohelm/objectkind/pkg/sqlutil/query/internal/queryopts"
)

type Options = queryopts.Options

func With(ctx context.Context, options ...Options) context.Context {
	if len(options) == 0 {
		return ctx
	}

	return queryopts.InjectContext(ctx, queryopts.FromContext(ctx).Join(options...))
}
