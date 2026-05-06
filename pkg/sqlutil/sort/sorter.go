package sort

import (
	"github.com/octohelm/storage/pkg/sqlbuilder"
	"github.com/octohelm/storage/pkg/sqlpipe"
)

// Sorter 排序接口，定义排序器的名称、标签及排序逻辑
type Sorter[M sqlpipe.Model] interface {
	Name() string
	Label() string
	Sort(src sqlpipe.Source[M], by func(col sqlbuilder.Column) sqlpipe.SourceOperator[M]) sqlpipe.Source[M]
}
