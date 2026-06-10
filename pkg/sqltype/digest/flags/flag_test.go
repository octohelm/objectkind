package flags_test

import (
	"context"
	"testing"

	"github.com/octohelm/x/cmp"
	. "github.com/octohelm/x/testing/v2"

	"github.com/octohelm/objectkind/pkg/sqltype/digest/flags"
)

func TestFlagOperations(t *testing.T) {
	zero := flags.Flag(0)

	t.Run("None", func(t *testing.T) {
		Then(
			t, "None.Is(None) 为 true",
			Expect(flags.None.Is(flags.None), Be(cmp.True())),
		)
		Then(
			t, "None.Is(HashSkipIfExists) 为 false",
			Expect(flags.None.Is(flags.HashSkipIfExists), Be(cmp.False())),
		)
	})

	t.Run("HashSkipIfExists", func(t *testing.T) {
		Then(
			t, "HashSkipIfExists.Is(HashSkipIfExists) 为 true",
			Expect(flags.HashSkipIfExists.Is(flags.HashSkipIfExists), Be(cmp.True())),
		)
		Then(
			t, "HashSkipIfExists.Is(None) 为 false",
			Expect(flags.HashSkipIfExists.Is(flags.None), Be(cmp.False())),
		)
	})

	t.Run("With 追加标志", func(t *testing.T) {
		f := zero.With(flags.HashSkipIfExists)
		Then(
			t, "zero.With(HashSkipIfExists) 包含 HashSkipIfExists",
			Expect(f.Is(flags.HashSkipIfExists), Be(cmp.True())),
		)
	})

	t.Run("Without 移除标志", func(t *testing.T) {
		f := flags.HashSkipIfExists
		f = f.Without(flags.HashSkipIfExists)
		Then(
			t, "移除后不再包含 HashSkipIfExists",
			Expect(f.Is(flags.HashSkipIfExists), Be(cmp.False())),
		)
	})

	t.Run("复合 With 后检测", func(t *testing.T) {
		f := flags.HashSkipIfExists.With(flags.None)
		Then(
			t, "With(None) 后仍包含 HashSkipIfExists",
			Expect(f.Is(flags.HashSkipIfExists), Be(cmp.True())),
		)
		Then(
			t, "With(None) 后也包含 None",
			Expect(f.Is(flags.None), Be(cmp.True())),
		)
	})
}

func TestFlagContext(t *testing.T) {
	t.Run("默认上下文中获取 None", func(t *testing.T) {
		ctx := context.Background()
		Then(
			t, "FromContext 返回 None",
			Expect(flags.FromContext(ctx), Equal(flags.None)),
		)
	})

	t.Run("注入与获取", func(t *testing.T) {
		ctx := flags.InjectContext(context.Background(), flags.HashSkipIfExists)
		Then(
			t, "FromContext 返回 HashSkipIfExists",
			Expect(flags.FromContext(ctx), Equal(flags.HashSkipIfExists)),
		)
	})

	t.Run("InjectContextWith 追加标志", func(t *testing.T) {
		ctx := flags.InjectContext(context.Background(), flags.None)
		ctx = flags.InjectContextWith(ctx, flags.HashSkipIfExists)
		Then(
			t, "追加后包含 HashSkipIfExists",
			Expect(flags.FromContext(ctx).Is(flags.HashSkipIfExists), Be(cmp.True())),
		)
	})

	t.Run("InjectContextWithout 移除标志", func(t *testing.T) {
		ctx := flags.InjectContext(context.Background(), flags.HashSkipIfExists)
		ctx = flags.InjectContextWithout(ctx, flags.HashSkipIfExists)
		Then(
			t, "移除后不再包含 HashSkipIfExists",
			Expect(flags.FromContext(ctx).Is(flags.HashSkipIfExists), Be(cmp.False())),
		)
	})
}
