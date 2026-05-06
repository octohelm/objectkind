package compose

import (
	"time"

	"github.com/octohelm/storage/pkg/sqltype"
	sqltypetime "github.com/octohelm/storage/pkg/sqltype/time"

	"github.com/octohelm/objectkind/pkg/object"
)

// Revision 泛型版本记录类型，用于记录每次变更的快照摘要。
type Revision[ID ~uint64, Digest ~string] struct {
	// id
	ID ID `db:"f_id" json:"id" sortable:""`
	// 摘要
	Digest Digest `db:"f_digest" json:"digest"`
	// 创建时间
	CreatedAt sqltypetime.Timestamp `db:"f_created_at,default='0'" json:"creationTimestamp,omitzero" sortable:""`
}

var _ sqltype.WithModificationTime = &Resource[uint64]{}

// MarkModifiedAt 设置创建时间作为修改时间。仅在 CreatedAt 为零值时写入当前时间。
func (r *Revision[ID, Digest]) MarkModifiedAt() {
	if r.CreatedAt.IsZero() {
		r.CreatedAt = sqltypetime.Timestamp(time.Now())
	}
}

var _ interface {
	object.IDGetter[uint64]
	object.IDSetter[uint64]
} = &Revision[uint64, string]{}

// GetID 返回版本记录主键 ID。
func (r Revision[ID, Digest]) GetID() ID { return r.ID }

// SetID 设置版本记录主键 ID。
func (r *Revision[ID, Digest]) SetID(id ID) { r.ID = id }

var _ object.OperationTimestamps = &Resource[uint64]{}

// GetCreationTimestamp 返回创建时间。
func (r Revision[ID, Digest]) GetCreationTimestamp() object.Timestamp {
	return r.CreatedAt
}

// SetCreationTimestamp 设置创建时间。
func (r *Revision[ID, Digest]) SetCreationTimestamp(timestamp object.Timestamp) {
	r.MarkModifiedAt()
}

// GetModificationTimestamp 返回修改时间（以创建时间替代）。
func (r Revision[ID, Digest]) GetModificationTimestamp() object.Timestamp {
	return r.CreatedAt
}

// SetModificationTimestamp 设置修改时间。
func (r *Revision[ID, Digest]) SetModificationTimestamp(timestamp object.Timestamp) {
	r.MarkModifiedAt()
}
