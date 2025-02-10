package compose

import (
	"time"

	"github.com/octohelm/objectkind/pkg/object"
	"github.com/octohelm/storage/pkg/sqltype"
	sqltypetime "github.com/octohelm/storage/pkg/sqltype/time"
)

type Rel[ID ~uint64] struct {
	// 自增 id
	ID ID `db:"f_id,autoincrement" json:"id"`

	// 创建时间
	CreatedAt sqltypetime.Timestamp `db:"f_created_at,default='0'" json:"creationTimestamp,omitzero" sortable:""`
}

var _ sqltype.WithModificationTime = &Resource[uint64]{}

func (r *Rel[ID]) MarkModifiedAt() {
	if r.CreatedAt.IsZero() {
		r.CreatedAt = sqltypetime.Timestamp(time.Now())
	}
}

var _ interface {
	object.IDGetter[uint64]
} = &Rel[uint64]{}

func (r Rel[ID]) GetID() ID { return r.ID }

var _ object.OperationTimestamps = &Resource[uint64]{}

func (r Rel[ID]) GetCreationTimestamp() object.Timestamp {
	return r.CreatedAt
}

func (r *Rel[ID]) SetCreationTimestamp(timestamp object.Timestamp) {
	r.MarkModifiedAt()
}

func (r Rel[ID]) GetModificationTimestamp() object.Timestamp {
	return r.CreatedAt
}

func (r *Rel[ID]) SetModificationTimestamp(timestamp object.Timestamp) {
	r.MarkModifiedAt()
}
