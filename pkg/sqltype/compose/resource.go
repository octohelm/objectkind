package compose

import (
	"time"

	"github.com/octohelm/objectkind/pkg/schema"
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

type CodedResource[ID ~uint64, Code ~string] struct {
	Resource[ID]

	// 编码
	// 人类可读编码
	Code Code `db:"f_code" json:"code"`
}

var _ interface {
	schema.ObjectReceiver
	schema.ObjectProvider
} = &Resource[uint64]{}

func (r *Resource[ID]) CopyFromObject(o schema.Object) {
	if canGetID, ok := o.(schema.IDGetter[ID]); ok {
		r.ID = canGetID.GetID()
	}

	if x, ok := o.(schema.ObjectDescriptor); ok {
		r.Name = x.GetName()
		r.Description = sqltypenullable.Text(x.GetDescription())
	}

	if x, ok := o.(schema.CreationTimestampGetter); ok {
		r.CreatedAt = x.GetCreationTimestamp()
	}

	if x, ok := o.(schema.ModificationTimestampGetter); ok {
		r.UpdatedAt = x.GetModificationTimestamp()
	}

	r.MarkModifiedAt()
}

func (r Resource[ID]) CopyToObject(o schema.ObjectSetter) {
	if canSetID, ok := o.(schema.IDSetter[ID]); ok {
		canSetID.SetID(r.ID)
	}

	if x, ok := o.(schema.ObjectDescriptorSetter); ok {
		x.SetName(r.Name)
		x.SetDescription(string(r.Description))
	}

	if x, ok := o.(schema.CreationTimestampSetter); ok {
		x.SetCreationTimestamp(r.CreatedAt)
	}

	if x, ok := o.(schema.ModificationTimestampSetter); ok {
		x.SetModificationTimestamp(r.UpdatedAt)
	}

}

var _ interface {
	schema.ObjectReceiver
	schema.ObjectProvider
} = &CodedResource[uint64, string]{}

func (r CodedResource[ID, Code]) CopyToObject(o schema.ObjectSetter) {
	r.Resource.CopyToObject(o)

	if canSetCode, ok := o.(schema.CodeSetter[Code]); ok {
		canSetCode.SetCode(r.Code)
	}
}

func (r *CodedResource[ID, Code]) CopyFromObject(o schema.Object) {
	if canGetID, ok := o.(schema.CodeGetter[Code]); ok {
		r.Code = canGetID.GetCode()
	}

	r.Resource.CopyFromObject(o)
}
