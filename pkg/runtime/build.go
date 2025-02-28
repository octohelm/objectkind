package runtime

import (
	"github.com/octohelm/objectkind/pkg/object"
)

func New[O any]() *O {
	v := new(O)

	if o, ok := any(v).(object.Type); ok {
		if x, ok := o.(object.PluralizedKindSetter); ok {
			if kinder, ok := o.(object.PluralizedKindGetter); ok {
				x.SetPluralizedKind(kinder.GetPluralizedKind())
			}
		}

		if x, ok := o.(object.KindSetter); ok {
			if kinder, ok := o.(object.Type); ok {
				x.SetKind(kinder.GetKind())
			}
		}

		if x, ok := o.(object.APIVersionSetter); ok {
			if apiVersioner, ok := o.(object.APIVersionGetter); ok {
				x.SetAPIVersion(apiVersioner.GetAPIVersion())
			}
		}
	}

	return v
}

func Build[T any](mutations ...func(t *T)) *T {
	o := New[T]()
	for _, mut := range mutations {
		if mut != nil {
			mut(o)
		}
	}
	return o
}

func BuildFuncFor[M object.Type]() func() object.Type {
	return func() object.Type {
		return any(New[M]()).(object.Type)
	}
}
