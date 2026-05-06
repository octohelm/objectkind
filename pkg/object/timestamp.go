package object

import (
	sqltypetime "github.com/octohelm/storage/pkg/sqltype/time"
)

// Timestamp 是时间戳类型别名。
type Timestamp = sqltypetime.Timestamp

// CreationTimestampDescriber 支持获取和设置创建时间戳。
type CreationTimestampDescriber interface {
	GetCreationTimestamp() Timestamp
	SetCreationTimestamp(timestamp Timestamp)
}

// ModificationTimestampDescriber 支持获取和设置修改时间戳。
type ModificationTimestampDescriber interface {
	GetModificationTimestamp() Timestamp
	SetModificationTimestamp(timestamp Timestamp)
}

// OperationTimestamps 聚合创建和修改时间戳接口。
type OperationTimestamps interface {
	CreationTimestampDescriber
	ModificationTimestampDescriber
}
