package queryopts

import (
	"fmt"

	"github.com/octohelm/objectkind/pkg/sqlutil/query/internal"
	"github.com/octohelm/objectkind/pkg/sqlutil/query/internal/queryflags"
)

type Options interface {
	QueryOptions(internal.NotForPublicUse)
}

func GetOption[T any](opts Options, set func(T) Options) (T, bool) {
	structOpts, ok := opts.(*Opt)
	if !ok {
		structOpts = (&Opt{}).Join(opts)
	}

	var zero T
	switch opt := set(zero).(type) {
	case queryflags.Bools:
		v := structOpts.Flags.Get(opt)
		return any(v).(T), structOpts.Flags.Has(opt)
	case Custom:
		v, ok := structOpts.GetCustom(opt)
		return any(v).(T), ok
	default:
		panic(fmt.Errorf("unknown option: %T", opt))
	}
}

type Opt struct {
	queryflags.Flags

	customs []Custom
}

func (Opt) QueryOptions(internal.NotForPublicUse) {
}

func (dst Opt) Join(srcs ...Options) *Opt {
	for _, src := range srcs {
		switch x := src.(type) {
		case nil:
			continue
		case queryflags.Bools:
			dst.Flags.Set(x)
		case Custom:
			dst.customs = append(dst.customs, x)
		case *Opt:
			dst.Flags = dst.Flags.Join(x.Flags)
			if len(x.customs) > 0 {
				dst.customs = append(dst.customs, x.customs...)
			}
		default:
		}
	}
	return &dst
}

func (dst *Opt) GetCustom(x Custom) (Custom, bool) {
	for _, t := range dst.customs {
		if x.Is(t) {
			return t, true
		}
	}
	return x, false
}

type Custom interface {
	Is(x any) bool
}
