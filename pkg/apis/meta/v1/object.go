package v1

import "github.com/octohelm/objectkind/pkg/schema"

type Object[ID ~uint64] struct {
	Metadata
	// id
	ID ID `json:"id"`
}

var _ interface {
	schema.ObjectWithID[uint64]
	schema.ObjectIDSetter[uint64]
	schema.ObjectReceiver
} = &Object[uint64]{}

func (o Object[ID]) GetID() ID {
	return o.ID
}

func (o *Object[ID]) SetID(id ID) {
	o.ID = id
}

func (v *Object[ID]) CopyFromObject(o schema.Object) {
	v.Metadata.CopyFromObject(o)

	if x, ok := o.(schema.ObjectWithID[ID]); ok {
		v.SetID(x.GetID())
	}
}

type ObjectRequest struct {
	Metadata
}

type ObjectReference[ID ~uint64] struct {
	TypeMeta
	// id
	ID ID `json:"id"`
}

var _ interface {
	schema.ObjectWithID[uint64]
	schema.ObjectIDSetter[uint64]
	schema.ObjectReceiver
} = &ObjectReference[uint64]{}

func (o ObjectReference[ID]) GetID() ID {
	return o.ID
}

func (o *ObjectReference[ID]) SetID(id ID) {
	o.ID = id
}

func (v *ObjectReference[ID]) CopyFromObject(o schema.Object) {
	if x, ok := o.(schema.ObjectWithID[ID]); ok {
		v.SetID(x.GetID())
	}
}
