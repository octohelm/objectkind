package compose

import (
	"database/sql/driver"
	"time"

	"github.com/octohelm/storage/pkg/sqltype"
	sqltypetime "github.com/octohelm/storage/pkg/sqltype/time"
)

// CreationTimestamp 创建时间追踪，提供 MarkCreatedAt 方法。
type CreationTimestamp struct {
	// 创建时间
	CreatedAt sqltypetime.Timestamp `db:"f_created_at,default='0'" json:"creationTimestamp,omitzero" sortable:""`
}

// MarkCreatedAt 设置创建时间。仅在 CreatedAt 为零值时写入当前时间。
func (times *CreationTimestamp) MarkCreatedAt() {
	if times.CreatedAt.IsZero() {
		times.CreatedAt = sqltypetime.Timestamp(time.Now())
	}
}

// ModificationTimestamp 修改时间追踪，提供 MarkModifiedAt 方法。
type ModificationTimestamp struct {
	// 更新时间
	UpdatedAt sqltypetime.Timestamp `db:"f_updated_at,default='0'" json:"modificationTimestamp,omitzero" sortable:""`
}

// MarkModifiedAt 设置修改时间。仅在 UpdatedAt 为零值时写入当前时间。
func (times *ModificationTimestamp) MarkModifiedAt() {
	if times.UpdatedAt.IsZero() {
		times.UpdatedAt = sqltypetime.Timestamp(time.Now())
	}
}

var _ sqltype.WithSoftDelete = &DeletionTimestamp{}

// DeletionTimestamp 删除时间追踪，实现 sqltype.WithSoftDelete 接口。
type DeletionTimestamp struct {
	// 删除时间
	DeletedAt sqltypetime.Timestamp `db:"f_deleted_at,default='0'" json:"deletionTimestamp,omitzero"`
}

// SoftDeleteFieldAndZeroValue 返回软删除字段名及其零值，实现 sqltype.WithSoftDelete。
func (DeletionTimestamp) SoftDeleteFieldAndZeroValue() (string, driver.Value) {
	return "DeletedAt", int64(0)
}

// MarkDeletedAt 设置删除时间为当前时间。
func (times *DeletionTimestamp) MarkDeletedAt() {
	times.DeletedAt = sqltypetime.Timestamp(time.Now())
}
