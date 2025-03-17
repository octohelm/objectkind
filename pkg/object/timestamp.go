package object

import sqltypetime "github.com/octohelm/storage/pkg/sqltype/time"

type Timestamp = sqltypetime.Timestamp

type CreationTimestampDescriber interface {
	GetCreationTimestamp() Timestamp
	SetCreationTimestamp(timestamp Timestamp)
}

type ModificationTimestampDescriber interface {
	GetModificationTimestamp() Timestamp
	SetModificationTimestamp(timestamp Timestamp)
}

type OperationTimestamps interface {
	CreationTimestampDescriber
	ModificationTimestampDescriber
}
