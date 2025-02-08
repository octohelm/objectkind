package compose

import (
	"time"

	"github.com/octohelm/objectkind/pkg/object"
	"github.com/octohelm/storage/pkg/sqltype"
	sqltypenullable "github.com/octohelm/storage/pkg/sqltype/nullable"
	sqltypetime "github.com/octohelm/storage/pkg/sqltype/time"
)

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

func (r *Resource[ID]) MarkCreatedAt() {
	if r.CreatedAt.IsZero() {
		r.CreatedAt = sqltypetime.Timestamp(time.Now())
	}
}

var _ sqltype.WithModificationTime = &Resource[uint64]{}

func (r *Resource[ID]) MarkModifiedAt() {
	if r.UpdatedAt.IsZero() {
		r.UpdatedAt = sqltypetime.Timestamp(time.Now())
	}
	if r.CreatedAt.IsZero() {
		r.CreatedAt = r.UpdatedAt
	}
}

func (r *Resource[ID]) ForceMarkModifiedAt() {
	r.UpdatedAt = sqltypetime.Timestamp(time.Now())
	r.CreatedAt = r.UpdatedAt
}

var (
	_ object.Describer = &Resource[uint64]{}
)

func (r *Resource[ID]) GetName() string {
	return r.Name
}

func (r *Resource[ID]) SetName(name string) {
	r.Name = name
}

func (r Resource[ID]) GetDescription() string {
	return string(r.Description)
}

func (r *Resource[ID]) SetDescription(description string) {
	r.Description = sqltypenullable.Text(description)
}

var _ object.OperationTimestamps = &Resource[uint64]{}

func (r Resource[ID]) GetCreationTimestamp() object.Timestamp {
	return r.CreatedAt
}

func (r *Resource[ID]) SetCreationTimestamp(timestamp object.Timestamp) {
	r.CreatedAt = timestamp
	r.MarkCreatedAt()
}

func (r Resource[ID]) GetModificationTimestamp() object.Timestamp {
	return r.UpdatedAt
}

func (r *Resource[ID]) SetModificationTimestamp(timestamp object.Timestamp) {
	r.UpdatedAt = timestamp
	r.MarkModifiedAt()
}

var _ interface {
	object.IDGetter[uint64]
	object.IDSetter[uint64]
} = &Resource[uint64]{}

func (r Resource[ID]) GetID() ID    { return r.ID }
func (r *Resource[ID]) SetID(id ID) { r.ID = id }

type CodableResource[ID ~uint64, Code ~string] struct {
	Resource[ID]
	// 编码
	// 人类可读编码
	Code Code `db:"f_code" json:"code"`
}

var _ interface {
	object.CodeGetter[string]
	object.CodeSetter[string]
} = &CodableResource[uint64, string]{}

func (r CodableResource[ID, Code]) GetCode() Code      { return r.Code }
func (r *CodableResource[ID, Code]) SetCode(code Code) { r.Code = code }
