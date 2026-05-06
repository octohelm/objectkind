package compose

import (
	"time"

	"github.com/octohelm/storage/pkg/sqltype"
	sqltypetime "github.com/octohelm/storage/pkg/sqltype/time"

	"github.com/octohelm/objectkind/pkg/object"
)

// Rel 泛型关联表类型，只包含自增主键与创建时间，用于多对多映射。
type Rel[ID ~uint64] struct {
	// 自增 id
	ID ID `db:"f_id,autoincrement" json:"id"`

	// 创建时间
	CreatedAt sqltypetime.Timestamp `db:"f_created_at,default='0'" json:"creationTimestamp,omitzero" sortable:""`
}

var _ sqltype.WithModificationTime = &Resource[uint64]{}

// MarkModifiedAt 设置创建时间作为修改时间。仅在 CreatedAt 为零值时写入当前时间。
func (r *Rel[ID]) MarkModifiedAt() {
	if r.CreatedAt.IsZero() {
		r.CreatedAt = sqltypetime.Timestamp(time.Now())
	}
}

var _ interface {
	object.IDGetter[uint64]
} = &Rel[uint64]{}

// GetID 返回关联主键 ID。
func (r Rel[ID]) GetID() ID { return r.ID }

var _ object.OperationTimestamps = &Resource[uint64]{}

// GetCreationTimestamp 返回创建时间。
func (r Rel[ID]) GetCreationTimestamp() object.Timestamp {
	return r.CreatedAt
}

// SetCreationTimestamp 设置创建时间。
func (r *Rel[ID]) SetCreationTimestamp(timestamp object.Timestamp) {
	r.MarkModifiedAt()
}

// GetModificationTimestamp 返回修改时间（以创建时间替代）。
func (r Rel[ID]) GetModificationTimestamp() object.Timestamp {
	return r.CreatedAt
}

// SetModificationTimestamp 设置修改时间。
func (r *Rel[ID]) SetModificationTimestamp(timestamp object.Timestamp) {
	r.MarkModifiedAt()
}
