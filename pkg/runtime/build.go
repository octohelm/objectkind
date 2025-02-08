package runtime

import (
	"github.com/octohelm/objectkind/pkg/object"
)

func New[O object.Type]() *O {
	o := new(O)

	if x, ok := any(o).(object.PluralizedKindSetter); ok {
		if kinder, ok := any(o).(object.PluralizedKindGetter); ok {
			x.SetPluralizedKind(kinder.GetPluralizedKind())
		}
	}

	if x, ok := any(o).(object.KindSetter); ok {
		if kinder, ok := any(o).(object.Type); ok {
			x.SetKind(kinder.GetKind())
		}
	}

	if x, ok := any(o).(object.APIVersionSetter); ok {
		if apiVersioner, ok := any(o).(object.APIVersionGetter); ok {
			x.SetAPIVersion(apiVersioner.GetAPIVersion())
		}
	}

	return o
}

func Build[T object.Type](mutations ...func(t *T)) *T {
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
