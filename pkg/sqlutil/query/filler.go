package query

import (
	"context"

	"github.com/octohelm/objectkind/pkg/object"
	"github.com/octohelm/objectkind/pkg/sqlutil/query/internal"
	"github.com/octohelm/objectkind/pkg/sqlutil/query/internal/queryopts"
)

type Filler[O object.Type] bool

func (Filler[O]) Is(x any) bool {
	_, ok := x.(Filler[O])
	return ok
}

func (Filler[O]) QueryOptions(use internal.NotForPublicUse) {
}

func NeedFiller[O object.Type](ctx context.Context) bool {
	v, ok := queryopts.GetOption(queryopts.FromContext(ctx), func(x Filler[O]) queryopts.Options {
		return x
	})
	if ok {
		return bool(v)
	}
	return true
}
