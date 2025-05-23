package filler

import (
	"context"
	"iter"

	"github.com/octohelm/objectkind/pkg/object"
	"github.com/octohelm/objectkind/pkg/sqlutil"
	"github.com/octohelm/storage/pkg/sqlpipe"
	sqlpipeex "github.com/octohelm/storage/pkg/sqlpipe/ex"
)

// Filler
// Deprecated use FillSet, Fill, FillSeq, FillOwnerSet, FillSubResourcesOfOwnerSet instead
type Filler[ID ~uint64, O object.Object[ID], M sqlpipe.Model] interface {
	FillSet(ctx context.Context, itemSet sqlpipeex.Set[ID, O]) error
	FillSeq(ctx context.Context, itemSeq iter.Seq[*O]) error
	Fill(ctx context.Context, items ...*O) error
}

func For[ID ~uint64, O object.Object[ID], M sqlpipe.Model](
	buildExecutor func(ids iter.Seq[ID]) sqlpipeex.SourceExecutor[M],
	buildObjectSeq sqlutil.BuildObjectSeq[M, O],
) Filler[ID, O, M] {
	return &filler[ID, O, M]{
		buildExecutor: buildExecutor,
		createSeq:     buildObjectSeq,
	}
}

type filler[ID ~uint64, O object.Object[ID], M sqlpipe.Model] struct {
	buildExecutor func(ids iter.Seq[ID]) sqlpipeex.SourceExecutor[M]
	createSeq     sqlutil.BuildObjectSeq[M, O]
}

func (f *filler[ID, O, M]) FillSet(ctx context.Context, itemSet sqlpipeex.Set[ID, O]) error {
	if itemSet == nil || itemSet.IsZero() {
		return nil
	}

	return sqlutil.FillSet(ctx, itemSet, f.buildExecutor(itemSet.Keys()), f.createSeq)
}

func (f *filler[ID, O, M]) FillSeq(ctx context.Context, itemSeq iter.Seq[*O]) error {
	itemSet := sqlpipeex.OneToMulti[ID, O]{}
	for item := range itemSeq {
		itemSet.Record(any(item).(object.IDGetter[ID]).GetID(), item)
	}
	return f.FillSet(ctx, itemSet)
}

func (f *filler[ID, O, M]) Fill(ctx context.Context, items ...*O) error {
	if len(items) == 0 {
		return nil
	}
	itemSet := sqlpipeex.OneToMulti[ID, O]{}
	for _, item := range items {
		itemSet.Record(any(item).(object.IDGetter[ID]).GetID(), item)
	}
	return f.FillSet(ctx, itemSet)
}
