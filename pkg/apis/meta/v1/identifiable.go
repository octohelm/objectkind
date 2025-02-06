package v1

import "github.com/octohelm/objectkind/pkg/object"

type Identifiable[ID ~uint64] struct {
	// 资源 id
	ID ID `json:"id"`
}

var _ interface {
	object.IDGetter[uint64]
	object.IDSetter[uint64]
	object.RefIDConvertable
} = &Identifiable[uint64]{}

func (o Identifiable[ID]) GetID() ID      { return o.ID }
func (o *Identifiable[ID]) SetID(code ID) { o.ID = code }

func (o Identifiable[ID]) GetAsRefID() object.RefID         { return object.RefID(o.ID) }
func (o *Identifiable[ID]) SetFromRefID(refID object.RefID) { o.ID = ID(refID) }
