package query

import "context"

const (
	all = -(iota + 1)

	skipCount
	skipResourceStatus

	skipResourceOwner
	skipResourceSecondaryOwner

	skipSubResources
)

const (
	All = 1 << -all
	// 不 Count
	SkipCount = 1 << -skipCount
	// 不查询资源状态类信息
	SkipResourceStatus = 1 << -skipResourceStatus
	// 不查询资源直接归属
	SkipResourceOwner = 1 << -skipResourceOwner
	// 不查询资源次要归属
	SkipResourceSecondaryOwner = 1 << -skipResourceSecondaryOwner
	// 不查询子资源详细信息
	SkipSubResources = 1 << -skipSubResources
)

func NeedCount(ctx context.Context) bool {
	mode := queryModeCtx.From(ctx)

	if mode&SkipCount != 0 {
		return false
	}

	return true
}

func NeedResourceStatus(ctx context.Context) bool {
	mode := queryModeCtx.From(ctx)

	if mode&SkipResourceStatus != 0 {
		return false
	}

	return true
}

func NeedResourceOwner(ctx context.Context) bool {
	mode := queryModeCtx.From(ctx)

	if mode&SkipResourceOwner != 0 {
		return false
	}

	return true
}

func NeedResourceSecondaryOwner(ctx context.Context) bool {
	mode := queryModeCtx.From(ctx)

	if mode&SkipResourceSecondaryOwner != 0 {
		return false
	}

	return true
}

func NeedSubResources(ctx context.Context) bool {
	mode := queryModeCtx.From(ctx)

	if mode&SkipSubResources != 0 {
		return false
	}

	return true
}
