package query_test

import (
	"context"
	"testing"

	"github.com/octohelm/x/cmp"
	. "github.com/octohelm/x/testing/v2"

	"github.com/octohelm/objectkind/pkg/sqlutil/query"
)

func TestOptions(t *testing.T) {
	t.Run("Flags 状态校验", func(t *testing.T) {
		ctx := query.With(context.Background())

		Then(t, "默认情况下所有 Flag 均为开启状态",
			Expect(query.NeedCount(ctx), Be(cmp.True())),
			Expect(query.NeedResourceStatus(ctx), Be(cmp.True())),
			Expect(query.NeedSubResources(ctx), Be(cmp.True())),
			Expect(query.NeedResourceOwner(ctx), Be(cmp.True())),
			Expect(query.NeedResourceSecondaryOwner(ctx), Be(cmp.True())),
		)

		t.Run("跳过总数统计 (SkipCount)", func(t *testing.T) {
			ctx2 := query.With(ctx, query.SkipCount)

			Then(t, "NeedCount 应该关闭",
				Expect(query.NeedCount(ctx2), Be(cmp.False())),
			)

			t.Run("重新开启总数统计", func(t *testing.T) {
				ctx3 := query.With(ctx2, query.FillCount(true))

				Then(t, "NeedCount 应该恢复开启",
					Expect(query.NeedCount(ctx3), Be(cmp.True())),
				)
			})
		})

		t.Run("各种 Skip 选项覆盖", func(t *testing.T) {
			Then(t, "SkipResourceStatus 生效",
				Expect(query.NeedResourceStatus(query.With(ctx, query.SkipResourceStatus)), Be(cmp.False())),
			)

			Then(t, "SkipSubResources 生效",
				Expect(query.NeedSubResources(query.With(ctx, query.SkipSubResources)), Be(cmp.False())),
			)

			Then(t, "SkipResourceOwner 生效",
				Expect(query.NeedResourceOwner(query.With(ctx, query.SkipResourceOwner)), Be(cmp.False())),
			)

			Then(t, "SkipResourceSecondaryOwner 生效",
				Expect(query.NeedResourceSecondaryOwner(query.With(ctx, query.SkipResourceSecondaryOwner)), Be(cmp.False())),
			)
		})
	})

	t.Run("Filler 泛型配置校验", func(t *testing.T) {
		ctx := query.With(context.Background())

		Then(t, "默认开启 Filler",
			Expect(query.NeedFiller[Object](ctx), Be(cmp.True())),
		)

		t.Run("禁用特定类型的 Filler", func(t *testing.T) {
			ctx2 := query.With(ctx, query.Filler[Object](false))

			Then(t, "NeedFiller 应该关闭",
				Expect(query.NeedFiller[Object](ctx2), Be(cmp.False())),
			)
		})
	})
}

type Object struct{}

func (o Object) GetKind() string {
	return "Object"
}
