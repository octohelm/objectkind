package query

import (
	"context"
	"testing"

	testingx "github.com/octohelm/x/testing"
)

func TestOptions(t *testing.T) {
	t.Run("flags", func(t *testing.T) {
		ctx := With(context.Background())

		testingx.Expect(t, NeedCount(ctx), testingx.BeTrue())
		testingx.Expect(t, NeedResourceStatus(ctx), testingx.BeTrue())
		testingx.Expect(t, NeedSubResources(ctx), testingx.BeTrue())
		testingx.Expect(t, NeedResourceOwner(ctx), testingx.BeTrue())
		testingx.Expect(t, NeedResourceSecondaryOwner(ctx), testingx.BeTrue())

		t.Run("skip count", func(t *testing.T) {
			ctx2 := With(ctx, SkipCount)
			testingx.Expect(t, NeedCount(ctx2), testingx.BeFalse())

			t.Run("then enabled", func(t *testing.T) {
				ctx3 := With(ctx, FillCount(true))
				testingx.Expect(t, NeedCount(ctx3), testingx.BeTrue())
			})
		})

		t.Run("skip resource status", func(t *testing.T) {
			ctx2 := With(ctx, SkipResourceStatus)
			testingx.Expect(t, NeedResourceStatus(ctx2), testingx.BeFalse())
		})

		t.Run("skip sub resources", func(t *testing.T) {
			ctx2 := With(ctx, SkipSubResources)
			testingx.Expect(t, NeedSubResources(ctx2), testingx.BeFalse())
		})

		t.Run("skip resource owner", func(t *testing.T) {
			ctx2 := With(ctx, SkipResourceOwner)
			testingx.Expect(t, NeedResourceOwner(ctx2), testingx.BeFalse())
		})

		t.Run("skip resource secondary owner", func(t *testing.T) {
			ctx2 := With(ctx, SkipResourceSecondaryOwner)
			testingx.Expect(t, NeedResourceSecondaryOwner(ctx2), testingx.BeFalse())
		})
	})

	t.Run("with filler", func(t *testing.T) {
		ctx := With(context.Background())
		testingx.Expect(t, NeedFiller[Object](ctx), testingx.BeTrue())

		t.Run("skip filler for object", func(t *testing.T) {
			ctx2 := With(ctx, Filler[Object](false))

			testingx.Expect(t, NeedFiller[Object](ctx2), testingx.BeFalse())
		})
	})
}

type Object struct{}

func (o Object) GetKind() string {
	return "Object"
}
