package filler

import (
	"context"
	"iter"
	"reflect"

	"github.com/octohelm/objectkind/pkg/object"
	sqlutilquery "github.com/octohelm/objectkind/pkg/sqlutil/query"
	sqlpipeex "github.com/octohelm/storage/pkg/sqlpipe/ex"
)

type SetFiller[ID ~uint64, O object.Object[ID]] interface {
	FillSet(ctx context.Context, itemSet sqlpipeex.Set[ID, O]) error
}

type OwnerSetFiller[ID ~uint64, O object.Object[ID]] interface {
	FillOwnerSet(ctx context.Context, itemSet sqlpipeex.Set[ID, O]) error
}

var fillers = Fillers{}

func Register[ID ~uint64, O object.Object[ID], Filler SetFiller[ID, O]](filler Filler) {
	fillers.Register(reflect.TypeFor[O](), filler)
}

func Fill[ID ~uint64, O object.Object[ID]](ctx context.Context, items ...*O) error {
	if len(items) == 0 {
		return nil
	}
	itemSet := sqlpipeex.OneToMulti[ID, O]{}
	for _, item := range items {
		itemSet.Record(any(item).(object.IDGetter[ID]).GetID(), item)
	}
	return FillSet[ID, O](ctx, itemSet)
}

func FillSeq[ID ~uint64, O object.Object[ID]](ctx context.Context, itemSeq iter.Seq[*O]) error {
	itemSet := sqlpipeex.OneToMulti[ID, O]{}
	for item := range itemSeq {
		itemSet.Record(any(item).(object.IDGetter[ID]).GetID(), item)
	}
	return FillSet[ID, O](ctx, itemSet)
}

func FillSet[ID ~uint64, O object.Object[ID], S sqlpipeex.Set[ID, O]](ctx context.Context, objects S) error {
	if any(objects) == nil || objects.IsZero() {
		return nil
	}

	for v := range fillers.Fillers() {
		switch f := v.(type) {
		case SetFiller[ID, O]:
			if err := f.FillSet(ctx, objects); err != nil {
				return err
			}
			return nil
		}
	}

	return nil
}

func FillOwnerSet[ID ~uint64, O object.Object[ID], S sqlpipeex.Set[ID, O]](ctx context.Context, objects S) error {
	return FillSet(sqlutilquery.With(ctx, sqlutilquery.SkipSubResources), objects)
}

func FillSubResourcesOfOwnerSet[OwnerID ~uint64, Owner object.Object[OwnerID], OwnerSet sqlpipeex.Set[OwnerID, Owner]](ctx context.Context, owners OwnerSet) error {
	if any(owners) == nil || owners.IsZero() {
		return nil
	}

	for v := range fillers.Fillers() {
		switch f := v.(type) {
		case OwnerSetFiller[OwnerID, Owner]:
			if err := f.FillOwnerSet(sqlutilquery.With(ctx, sqlutilquery.SkipResourceOwner), owners); err != nil {
				return err
			}
		}
	}

	return nil
}
