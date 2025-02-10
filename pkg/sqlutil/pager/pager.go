package pager

import (
	"github.com/octohelm/storage/pkg/sqlpipe"
)

type Pager[M sqlpipe.Model] struct {
	// 分页偏移
	Offset int64 `name:"offset,omitzero" in:"query"`
	// 分页条数
	Limit int64 `name:"limit,omitzero" validate:"@int[-1,50] = 10" in:"query"`
}

func (Pager[M]) OperatorType() sqlpipe.OperatorType {
	return sqlpipe.OperatorLimit
}

func (p *Pager[M]) Next(src sqlpipe.Source[M]) sqlpipe.Source[M] {
	if p.Limit > 50 {
		p.Limit = 50
	}

	if offset := p.Offset; offset > 0 {
		return src.Pipe(sqlpipe.Limit[M](p.Limit, sqlpipe.Offset(offset)))
	}

	return src.Pipe(sqlpipe.Limit[M](p.Limit))
}
