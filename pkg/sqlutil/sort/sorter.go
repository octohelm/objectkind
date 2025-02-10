package sort

import (
	"github.com/octohelm/storage/pkg/sqlbuilder"
	"github.com/octohelm/storage/pkg/sqlpipe"
)

type Sorter[M sqlpipe.Model] interface {
	Name() string
	Label() string
	Sort(src sqlpipe.Source[M], by func(col sqlbuilder.Column) sqlpipe.SourceOperator[M]) sqlpipe.Source[M]
}
