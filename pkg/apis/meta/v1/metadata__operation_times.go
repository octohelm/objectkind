package v1

import (
	"github.com/octohelm/objectkind/pkg/schema"
	sqltypetime "github.com/octohelm/storage/pkg/sqltype/time"
)

type OperationTimes struct {
	// 创建时间
	CreationTimestamp sqltypetime.Timestamp `json:"creationTimestamp,omitzero"`
	// 更新时间
	ModificationTimestamp sqltypetime.Timestamp `json:"modificationTimestamp,omitzero"`
}

var _ schema.CreationTimestampGetter = &OperationTimes{}

func (v OperationTimes) GetCreationTimestamp() sqltypetime.Timestamp {
	return v.CreationTimestamp
}

var _ schema.CreationTimestampSetter = &OperationTimes{}

func (v *OperationTimes) SetCreationTimestamp(creationTimestamp sqltypetime.Timestamp) {
	v.CreationTimestamp = creationTimestamp
}

var _ schema.ModificationTimestampGetter = &OperationTimes{}

func (v OperationTimes) GetModificationTimestamp() sqltypetime.Timestamp {
	return v.ModificationTimestamp
}

var _ schema.ModificationTimestampSetter = &OperationTimes{}

func (v *OperationTimes) SetModificationTimestamp(modificationTimestamp sqltypetime.Timestamp) {
	v.ModificationTimestamp = modificationTimestamp
}
