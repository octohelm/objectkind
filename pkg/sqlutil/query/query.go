package query

import (
	"context"

	contextx "github.com/octohelm/x/context"
)

var queryModeCtx = contextx.New[int](contextx.WithDefaults(All))

func With(ctx context.Context, mode int) context.Context {
	return queryModeCtx.Inject(ctx, queryModeCtx.From(ctx)|mode)
}
