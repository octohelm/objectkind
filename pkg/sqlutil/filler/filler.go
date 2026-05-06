package filler

import (
	"context"
	"iter"
	"reflect"

	sqlpipeex "github.com/octohelm/storage/pkg/sqlpipe/ex"

	"github.com/octohelm/objectkind/pkg/object"
	sqlutilquery "github.com/octohelm/objectkind/pkg/sqlutil/query"
)

// SetFiller 集合填充接口，通过 Object ID 集合批量填充对象
type SetFiller[ID ~uint64, O object.Object[ID]] interface {
	FillSet(ctx context.Context, itemSet sqlpipeex.Set[ID, O]) error
}

// OwnerSetFiller 拥有者集合填充接口，用于填充拥有者对象的关联数据
type OwnerSetFiller[ID ~uint64, O object.Object[ID]] interface {
	FillOwnerSet(ctx context.Context, itemSet sqlpipeex.Set[ID, O]) error
}

var fillers = Fillers{}

// Register 注册指定 Object 类型的 SetFiller 实现
func Register[ID ~uint64, O object.Object[ID], Filler SetFiller[ID, O]](filler Filler) {
	fillers.Register(reflect.TypeFor[O](), filler)
}

// Fill 对指定 Object 实例执行填充操作
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

// FillSeq 对 Object 迭代器中的数据执行填充操作
func FillSeq[ID ~uint64, O object.Object[ID]](ctx context.Context, itemSeq iter.Seq[*O]) error {
	itemSet := sqlpipeex.OneToMulti[ID, O]{}
	for item := range itemSeq {
		itemSet.Record(any(item).(object.IDGetter[ID]).GetID(), item)
	}
	return FillSet[ID, O](ctx, itemSet)
}

// FillSet 对 Object Set 执行填充操作，根据注册的 SetFiller 按类型匹配填充
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

// FillOwnerSet 对拥有者 Object Set 执行填充操作，跳过子资源填充
func FillOwnerSet[ID ~uint64, O object.Object[ID], S sqlpipeex.Set[ID, O]](ctx context.Context, objects S) error {
	return FillSet(sqlutilquery.With(ctx, sqlutilquery.SkipSubResources), objects)
}

// FillSubResourcesOfOwnerSet 填充拥有者集合的子资源数据，跳过资源拥有者填充
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
