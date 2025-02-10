package compose

import (
	"time"

	"github.com/octohelm/objectkind/pkg/object"
	"github.com/octohelm/storage/pkg/sqltype"

	sqltypetime "github.com/octohelm/storage/pkg/sqltype/time"
)

type Revision[ID ~uint64, Digest ~string] struct {
	// id
	ID ID `db:"f_id" json:"id" sortable:""`
	// 摘要
	Digest Digest `db:"f_digest" json:"digest"`
	// 创建时间
	CreatedAt sqltypetime.Timestamp `db:"f_created_at,default='0'" json:"creationTimestamp,omitzero" sortable:""`
}

var _ sqltype.WithModificationTime = &Resource[uint64]{}

func (r *Revision[ID, Digest]) MarkModifiedAt() {
	if r.CreatedAt.IsZero() {
		r.CreatedAt = sqltypetime.Timestamp(time.Now())
	}
}

var _ interface {
	object.IDGetter[uint64]
	object.IDSetter[uint64]
} = &Revision[uint64, string]{}

func (r Revision[ID, Digest]) GetID() ID { return r.ID }

func (r *Revision[ID, Digest]) SetID(id ID) { r.ID = id }

var _ object.OperationTimestamps = &Resource[uint64]{}

func (r Revision[ID, Digest]) GetCreationTimestamp() object.Timestamp {
	return r.CreatedAt
}

func (r *Revision[ID, Digest]) SetCreationTimestamp(timestamp object.Timestamp) {
	r.MarkModifiedAt()
}

func (r Revision[ID, Digest]) GetModificationTimestamp() object.Timestamp {
	return r.CreatedAt
}

func (r *Revision[ID, Digest]) SetModificationTimestamp(timestamp object.Timestamp) {
	r.MarkModifiedAt()
}
