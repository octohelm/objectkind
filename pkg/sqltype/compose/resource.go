package compose

import (
	"time"

	"github.com/octohelm/storage/pkg/sqltype"
	sqltypenullable "github.com/octohelm/storage/pkg/sqltype/nullable"
	sqltypetime "github.com/octohelm/storage/pkg/sqltype/time"

	"github.com/octohelm/objectkind/pkg/object"
)

// Resource 通用资源类型，组合 ID、名称、描述与创建/更新时间字段。
type Resource[ID ~uint64] struct {
	// id
	ID ID `db:"f_id" json:"id" sortable:""`
	// 名称
	Name string `db:"f_name,default=''" json:"name,omitzero" sortable:""`
	// 描述
	Description sqltypenullable.Text `db:"f_description,null" json:"description,omitzero"`
	// 创建时间
	CreatedAt sqltypetime.Timestamp `db:"f_created_at,default='0'" json:"creationTimestamp,omitzero" sortable:""`
	// 更新时间
	UpdatedAt sqltypetime.Timestamp `db:"f_updated_at,default='0'" json:"modificationTimestamp,omitzero" sortable:""`
}

var _ sqltype.WithCreationTime = &Resource[uint64]{}

// MarkCreatedAt 设置创建时间。仅在 CreatedAt 为零值时写入当前时间。
func (r *Resource[ID]) MarkCreatedAt() {
	if r.CreatedAt.IsZero() {
		r.CreatedAt = sqltypetime.Timestamp(time.Now())
	}
}

var _ sqltype.WithModificationTime = &Resource[uint64]{}

// MarkModifiedAt 设置修改时间。同时补齐未设置的创建时间。
func (r *Resource[ID]) MarkModifiedAt() {
	if r.UpdatedAt.IsZero() {
		r.UpdatedAt = sqltypetime.Timestamp(time.Now())
	}
	if r.CreatedAt.IsZero() {
		r.CreatedAt = r.UpdatedAt
	}
}

// ForceMarkModifiedAt 强制刷新修改时间与创建时间，不检查零值。
func (r *Resource[ID]) ForceMarkModifiedAt() {
	r.UpdatedAt = sqltypetime.Timestamp(time.Now())
	r.CreatedAt = r.UpdatedAt
}

var _ object.Describer = &Resource[uint64]{}

// GetName 返回资源名称。
func (r Resource[ID]) GetName() string {
	return r.Name
}

// SetName 设置资源名称。
func (r *Resource[ID]) SetName(name string) {
	r.Name = name
}

// GetDescription 返回资源描述。
func (r Resource[ID]) GetDescription() string {
	return string(r.Description)
}

// SetDescription 设置资源描述。
func (r *Resource[ID]) SetDescription(description string) {
	r.Description = sqltypenullable.Text(description)
}

var _ object.OperationTimestamps = &Resource[uint64]{}

// GetCreationTimestamp 返回创建时间。
func (r Resource[ID]) GetCreationTimestamp() object.Timestamp {
	return r.CreatedAt
}

// SetCreationTimestamp 设置创建时间。
func (r *Resource[ID]) SetCreationTimestamp(timestamp object.Timestamp) {
	r.CreatedAt = timestamp
	r.MarkCreatedAt()
}

// GetModificationTimestamp 返回修改时间。
func (r Resource[ID]) GetModificationTimestamp() object.Timestamp {
	return r.UpdatedAt
}

// SetModificationTimestamp 设置修改时间。
func (r *Resource[ID]) SetModificationTimestamp(timestamp object.Timestamp) {
	r.UpdatedAt = timestamp
	r.MarkModifiedAt()
}

var _ object.IDGetter[uint64] = Resource[uint64]{}

// GetID 返回资源主键 ID。
func (r Resource[ID]) GetID() ID { return r.ID }

var _ object.IDSetter[uint64] = &Resource[uint64]{}

// SetID 设置资源主键 ID。
func (r *Resource[ID]) SetID(id ID) { r.ID = id }

var _ object.AsRefIDGetter = Resource[uint64]{}

// GetAsRefID 以引用 ID 形式返回主键。
func (r Resource[ID]) GetAsRefID() object.RefID {
	return object.RefID(r.ID)
}

// CodableResource 带编码的泛型资源类型，在 Resource 基础上增加人类可读编码字段。
type CodableResource[ID ~uint64, Code ~string] struct {
	Resource[ID]
	// 编码
	// 人类可读编码
	Code Code `db:"f_code" json:"code" sortable:""`
}

var _ object.CodeGetter[string] = CodableResource[uint64, string]{}

// GetCode 返回资源编码。
func (r CodableResource[ID, Code]) GetCode() Code { return r.Code }

var _ object.CodeSetter[string] = &CodableResource[uint64, string]{}

// SetCode 设置资源编码。
func (r *CodableResource[ID, Code]) SetCode(code Code) { r.Code = code }

var _ object.AsRefCodeGetter = CodableResource[uint64, string]{}

// GetAsRefCode 以引用编码形式返回资源编码。
func (r CodableResource[ID, Code]) GetAsRefCode() object.RefCode {
	return object.RefCode(r.Code)
}
