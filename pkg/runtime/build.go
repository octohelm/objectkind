package runtime

import (
	"github.com/octohelm/objectkind/pkg/object"
)

func Build[T object.Type](mutations ...func(t *T)) *T {
	o := New[T]()
	for _, mut := range mutations {
		if mut != nil {
			mut(o)
		}
	}
	return o
}

func New[O object.Type]() *O {
	o := new(O)
	if x, ok := any(o).(object.KindAndAPIVersionSetter); ok {
		if kinder, ok := any(o).(object.Type); ok {
			x.SetKind(kinder.GetKind())
		}
		if apiVersioner, ok := any(o).(object.APIVersionGetter); ok {
			x.SetAPIVersion(apiVersioner.GetAPIVersion())
		}
	}
	return o
}
