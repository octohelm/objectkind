package v1

import "github.com/octohelm/objectkind/pkg/schema"

type Object[ID ~uint64] struct {
	Metadata
	// id
	ID ID `json:"id"`
}

var _ interface {
	schema.IDGetter[uint64]
	schema.AsRefIDGetter
	schema.IDSetter[uint64]
	schema.FromRefIDSetter
	schema.ObjectReceiver
} = &Object[uint64]{}

func (o Object[ID]) GetID() ID {
	return o.ID
}

func (o Object[ID]) GetAsRefID() schema.RefID {
	return schema.RefID(o.ID)
}

func (o *Object[ID]) SetID(id ID) {
	o.ID = id
}

func (o *Object[ID]) SetFromRefID(refID schema.RefID) {
	o.ID = ID(refID)
}

func (v *Object[ID]) CopyFromObject(o schema.Object) {
	v.Metadata.CopyFromObject(o)

	if x, ok := o.(schema.IDGetter[ID]); ok {
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
	schema.IDGetter[uint64]
	schema.AsRefIDGetter
	schema.IDSetter[uint64]
	schema.FromRefIDSetter
	schema.ObjectReceiver
} = &ObjectReference[uint64]{}

func (o ObjectReference[ID]) GetID() ID {
	return o.ID
}

func (o ObjectReference[ID]) GetAsRefID() schema.RefID {
	return schema.RefID(o.ID)
}

func (o *ObjectReference[ID]) SetID(id ID) {
	o.ID = id
}

func (o *ObjectReference[ID]) SetFromRefID(refID schema.RefID) {
	o.ID = ID(refID)
}

func (v *ObjectReference[ID]) CopyFromObject(o schema.Object) {
	if x, ok := o.(schema.IDGetter[ID]); ok {
		v.SetID(x.GetID())
	}
}
