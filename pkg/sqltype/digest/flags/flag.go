package flags

import (
	"context"
)

// Flag 表示哈希时的行为标志位，支持组合与判断。
type Flag uint

// Is 判断当前标志位是否包含 mode。
func (f Flag) Is(mode Flag) bool {
	return f&mode != 0
}

// With 在当前标志位上追加 f2 并返回新值。
func (f Flag) With(f2 Flag) Flag {
	return f | f2
}

// Without 从当前标志位中移除 f2 并返回新值。
func (f Flag) Without(f2 Flag) Flag {
	return f &^ f2
}

const (
	none = -(iota + 1)
	hashSkipIfExists
)

const (
	// None 表示无任何标志位的零值。
	None Flag = 1 << -none
	// HashSkipIfExists 表示当对象已存在摘要时跳过哈希计算。
	HashSkipIfExists Flag = 1 << -hashSkipIfExists
)

type flagContext struct{}

// FromContext 从上下文中提取标志位，未注入时返回 None。
func FromContext(ctx context.Context) Flag {
	if f, ok := ctx.Value(flagContext{}).(Flag); ok {
		return f
	}
	return None
}

// InjectContext 将指定标志位注入上下文，覆盖已有值。
func InjectContext(ctx context.Context, f Flag) context.Context {
	return context.WithValue(ctx, flagContext{}, f)
}

// InjectContextWith 在上下文已有标志位基础上追加 f 并注入。
func InjectContextWith(ctx context.Context, f Flag) context.Context {
	return context.WithValue(ctx, flagContext{}, FromContext(ctx).With(f))
}

// InjectContextWithout 在上下文已有标志位基础上移除 f 并注入。
func InjectContextWithout(ctx context.Context, f Flag) context.Context {
	return context.WithValue(ctx, flagContext{}, FromContext(ctx).Without(f))
}
