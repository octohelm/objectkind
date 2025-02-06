package compose

import (
	"database/sql/driver"
	"time"

	"github.com/octohelm/storage/pkg/sqltype"
	sqltypetime "github.com/octohelm/storage/pkg/sqltype/time"
)

type CreationTimestamp struct {
	// 创建时间
	CreatedAt sqltypetime.Timestamp `db:"f_created_at,default='0'" json:"creationTimestamp,omitzero"`
}

func (times *CreationTimestamp) MarkCreatedAt() {
	if times.CreatedAt.IsZero() {
		times.CreatedAt = sqltypetime.Timestamp(time.Now())
	}
}

type ModificationTimestamp struct {
	// 更新时间
	UpdatedAt sqltypetime.Timestamp `db:"f_updated_at,default='0'" json:"modificationTimestamp,omitzero"`
}

func (times *ModificationTimestamp) MarkModifiedAt() {
	if times.UpdatedAt.IsZero() {
		times.UpdatedAt = sqltypetime.Timestamp(time.Now())
	}
}

var _ sqltype.WithSoftDelete = &DeletionTimestamp{}

type DeletionTimestamp struct {
	// 删除时间
	DeletedAt sqltypetime.Timestamp `db:"f_deleted_at,default='0'" json:"deletionTimestamp,omitempty"`
}

func (DeletionTimestamp) SoftDeleteFieldAndZeroValue() (string, driver.Value) {
	return "DeletedAt", int64(0)
}

func (times *DeletionTimestamp) MarkDeletedAt() {
	times.DeletedAt = sqltypetime.Timestamp(time.Now())
}
