package flags

import "context"

type Flag uint

func (f Flag) Is(mode Flag) bool {
	return f&mode != 0
}

func (f Flag) With(f2 Flag) Flag {
	return f | f2
}

func (f Flag) Without(f2 Flag) Flag {
	return f &^ f2
}

const (
	none = -(iota + 1)
	hashSkipIfExists
)

const (
	None             Flag = 1 << -none
	HashSkipIfExists Flag = 1 << -hashSkipIfExists
)

type flagContext struct{}

func FromContext(ctx context.Context) Flag {
	if f, ok := ctx.Value(flagContext{}).(Flag); ok {
		return f
	}
	return None
}

func InjectContext(ctx context.Context, f Flag) context.Context {
	return context.WithValue(ctx, flagContext{}, f)
}

func InjectContextWith(ctx context.Context, f Flag) context.Context {
	return context.WithValue(ctx, flagContext{}, FromContext(ctx).With(f))
}

func InjectContextWithout(ctx context.Context, f Flag) context.Context {
	return context.WithValue(ctx, flagContext{}, FromContext(ctx).Without(f))
}
