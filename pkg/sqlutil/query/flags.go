package query

import (
	"context"

	"github.com/octohelm/objectkind/pkg/sqlutil/query/internal/queryflags"
	"github.com/octohelm/objectkind/pkg/sqlutil/query/internal/queryopts"
)

// FillCount 生成控制是否填充总数的查询选项
func FillCount(v bool) Options {
	if v {
		return queryflags.RequestCount | 1
	}
	return queryflags.RequestCount | 0
}

// FillResourceStatus 生成控制是否填充资源状态的查询选项
func FillResourceStatus(v bool) Options {
	if v {
		return queryflags.RequestResourceStatus | 1
	}
	return queryflags.RequestResourceStatus | 0
}

// FillSubResources 生成控制是否填充子资源的查询选项
func FillSubResources(v bool) Options {
	if v {
		return queryflags.RequestSubResources | 1
	}
	return queryflags.RequestSubResources | 0
}

// FillResourceOwner 生成控制是否填充资源拥有者的查询选项
func FillResourceOwner(v bool) Options {
	if v {
		return queryflags.RequestResourceOwner | 1
	}
	return queryflags.RequestResourceOwner | 0
}

// FillResourceSecondaryOwner 生成控制是否填充资源次级拥有者的查询选项
func FillResourceSecondaryOwner(v bool) Options {
	if v {
		return queryflags.RequestResourceSecondaryOwner | 1
	}
	return queryflags.RequestResourceSecondaryOwner | 0
}

var (
	// SkipCount 跳过总数填充的查询选项
	SkipCount = FillCount(false)
	// SkipResourceStatus 跳过资源状态填充的查询选项
	SkipResourceStatus = FillResourceStatus(false)
	// SkipSubResources 跳过子资源填充的查询选项
	SkipSubResources = FillSubResources(false)
	// SkipResourceOwner 跳过资源拥有者填充的查询选项
	SkipResourceOwner = FillResourceOwner(false)
	// SkipResourceSecondaryOwner 跳过资源次级拥有者填充的查询选项
	SkipResourceSecondaryOwner = FillResourceSecondaryOwner(false)
)

// NeedCount 从 context 中读取是否需要填充总数
func NeedCount(ctx context.Context) bool {
	v, _ := queryopts.GetOption(queryopts.FromContext(ctx), FillCount)
	return v
}

// NeedResourceStatus 从 context 中读取是否需要填充资源状态
func NeedResourceStatus(ctx context.Context) bool {
	v, _ := queryopts.GetOption(queryopts.FromContext(ctx), FillResourceStatus)
	return v
}

// NeedSubResources 从 context 中读取是否需要填充子资源
func NeedSubResources(ctx context.Context) bool {
	v, _ := queryopts.GetOption(queryopts.FromContext(ctx), FillSubResources)
	return v
}

// NeedResourceOwner 从 context 中读取是否需要填充资源拥有者
func NeedResourceOwner(ctx context.Context) bool {
	v, _ := queryopts.GetOption(queryopts.FromContext(ctx), FillResourceOwner)
	return v
}

// NeedResourceSecondaryOwner 从 context 中读取是否需要填充资源次级拥有者
func NeedResourceSecondaryOwner(ctx context.Context) bool {
	v, _ := queryopts.GetOption(queryopts.FromContext(ctx), FillResourceSecondaryOwner)
	return v
}
