package object

import sqltypetime "github.com/octohelm/storage/pkg/sqltype/time"

type Timestamp = sqltypetime.Timestamp

type OperationTimestamps interface {
	GetCreationTimestamp() Timestamp
	SetCreationTimestamp(timestamp Timestamp)

	GetModificationTimestamp() Timestamp
	SetModificationTimestamp(timestamp Timestamp)
}
