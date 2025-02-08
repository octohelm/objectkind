package sqlutil

import (
	"context"
	"iter"

	metav1 "github.com/octohelm/objectkind/pkg/apis/meta/v1"
	"github.com/octohelm/objectkind/pkg/object"
	sqlutilquery "github.com/octohelm/objectkind/pkg/sqlutil/query"
	"github.com/octohelm/storage/pkg/dberr"
	"github.com/octohelm/storage/pkg/sqlpipe"
	sqlpipeex "github.com/octohelm/storage/pkg/sqlpipe/ex"
)

func Objects[M sqlpipe.Model, O object.Type](ctx context.Context, ex sqlpipeex.SourceExecutor[M], convert func(m *M) (*O, error)) iter.Seq2[*O, error] {
	return func(yield func(*O, error) bool) {
		for m, err := range ex.Item(ctx) {
			if err != nil {
				yield(nil, err)
				return
			}
			if !yield(convert(m)) {
				return
			}
		}
	}
}

type BuildObjectSeq[M sqlpipe.Model, O object.Type] func(ctx context.Context, ex sqlpipeex.SourceExecutor[M]) iter.Seq2[*O, error]

func List[M sqlpipe.Model, O object.Type](ctx context.Context, ex sqlpipeex.SourceExecutor[M], create BuildObjectSeq[M, O]) (*metav1.List[O], error) {
	list := &metav1.List[O]{
		Items: make([]*O, 0),
	}

	if sqlutilquery.NeedCount(ctx) {
		if err := ex.CountTo(ctx, &list.Total); err != nil {
			return nil, err
		}
	}

	values, err := Collect(create(ctx, ex))
	if err != nil {
		return nil, err
	}

	list.Items = values

	return list, nil
}

func Collect[O object.Type](itemSeq iter.Seq2[*O, error]) ([]*O, error) {
	items := make([]*O, 0)

	for item, err := range itemSeq {
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func FindOne[M sqlpipe.Model, O object.Type](ctx context.Context, ex sqlpipeex.SourceExecutor[M], createSeq BuildObjectSeq[M, O]) (*O, error) {
	var v *O

	for item, err := range createSeq(ctx, ex.PipeE(sqlpipe.Limit[M](1))) {
		if err != nil {
			return nil, err
		}
		if v == nil {
			v = item
		}
	}

	if v == nil {
		return nil, &dberr.SqlError{Type: dberr.ErrTypeNotFound}
	}

	return v, nil
}
