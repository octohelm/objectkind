package query

import (
	"context"

	"github.com/octohelm/objectkind/pkg/sqlutil/query/internal/queryflags"
	"github.com/octohelm/objectkind/pkg/sqlutil/query/internal/queryopts"
)

func FillCount(v bool) Options {
	if v {
		return queryflags.RequestCount | 1
	}
	return queryflags.RequestCount | 0
}

func FillResourceStatus(v bool) Options {
	if v {
		return queryflags.RequestResourceStatus | 1
	}
	return queryflags.RequestResourceStatus | 0
}

func FillSubResources(v bool) Options {
	if v {
		return queryflags.RequestSubResources | 1
	}
	return queryflags.RequestSubResources | 0
}

func FillResourceOwner(v bool) Options {
	if v {
		return queryflags.RequestResourceOwner | 1
	}
	return queryflags.RequestResourceOwner | 0
}

func FillResourceSecondaryOwner(v bool) Options {
	if v {
		return queryflags.RequestResourceSecondaryOwner | 1
	}
	return queryflags.RequestResourceSecondaryOwner | 0
}

var (
	SkipCount                  = FillCount(false)
	SkipResourceStatus         = FillResourceStatus(false)
	SkipSubResources           = FillSubResources(false)
	SkipResourceOwner          = FillResourceOwner(false)
	SkipResourceSecondaryOwner = FillResourceSecondaryOwner(false)
)

func NeedCount(ctx context.Context) bool {
	v, _ := queryopts.GetOption(queryopts.FromContext(ctx), FillCount)
	return v
}

func NeedResourceStatus(ctx context.Context) bool {
	v, _ := queryopts.GetOption(queryopts.FromContext(ctx), FillResourceStatus)
	return v
}

func NeedSubResources(ctx context.Context) bool {
	v, _ := queryopts.GetOption(queryopts.FromContext(ctx), FillSubResources)
	return v
}

func NeedResourceOwner(ctx context.Context) bool {
	v, _ := queryopts.GetOption(queryopts.FromContext(ctx), FillResourceOwner)
	return v
}

func NeedResourceSecondaryOwner(ctx context.Context) bool {
	v, _ := queryopts.GetOption(queryopts.FromContext(ctx), FillResourceSecondaryOwner)
	return v
}
