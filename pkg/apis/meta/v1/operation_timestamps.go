package v1

import (
	"github.com/octohelm/objectkind/pkg/object"
)

// OperationTimestamps 操作时间戳，记录资源的创建与修改时间
type OperationTimestamps struct {
	// 创建时间
	CreationTimestamp object.Timestamp `json:"creationTimestamp,omitzero"`
	// 更新时间
	ModificationTimestamp object.Timestamp `json:"modificationTimestamp,omitzero"`
}

var _ object.OperationTimestamps = &OperationTimestamps{}

func (v OperationTimestamps) GetCreationTimestamp() object.Timestamp {
	return v.CreationTimestamp
}

func (v *OperationTimestamps) SetCreationTimestamp(creationTimestamp object.Timestamp) {
	v.CreationTimestamp = creationTimestamp
}

func (v OperationTimestamps) GetModificationTimestamp() object.Timestamp {
	return v.ModificationTimestamp
}

func (v *OperationTimestamps) SetModificationTimestamp(modificationTimestamp object.Timestamp) {
	v.ModificationTimestamp = modificationTimestamp
}
