package v1

import "github.com/octohelm/objectkind/pkg/schema"

type Metadata struct {
	TypeMeta
	Descriptor
	OperationTimes
}

var _ schema.ObjectReceiver = &Metadata{}

func (v *Metadata) CopyFromObject(o schema.Object) {
	if x, ok := o.(schema.ObjectDescriptor); ok {
		v.SetName(x.GetName())
		v.SetDescription(x.GetDescription())

		for k, ann := range x.GetAnnotations() {
			v.SetAnnotation(k, ann)
		}
	}

	if x, ok := o.(schema.ObjectWithCreationTimestamp); ok {
		v.SetCreationTimestamp(x.GetCreationTimestamp())
	}

	if x, ok := o.(schema.ObjectWithModificationTimestamp); ok {
		v.SetModificationTimestamp(x.GetModificationTimestamp())
	}
}
