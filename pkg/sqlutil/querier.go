package sqlutil

import (
	"context"
	"iter"

	"github.com/octohelm/storage/pkg/dberr"
	"github.com/octohelm/storage/pkg/sqlpipe"
	sqlpipeex "github.com/octohelm/storage/pkg/sqlpipe/ex"

	metav1 "github.com/octohelm/objectkind/pkg/apis/meta/v1"
	"github.com/octohelm/objectkind/pkg/object"
	sqlutilquery "github.com/octohelm/objectkind/pkg/sqlutil/query"
)

// Objects 将 SourceExecutor 中的模型通过 convert 转换为 object.Type 的迭代器
func Objects[M sqlpipe.Model, O object.Type](ctx context.Context, ex sqlpipeex.SourceExecutor[M], convert func(m *M) (*O, error)) iter.Seq2[*O, error] {
	return func(yield func(*O, error) bool) {
		for m, err := range ex.Items(ctx) {
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

// BuildObjectSeq 构建 Object 序列的工厂函数类型，接收 SourceExecutor 返回 object.Type 迭代器
type BuildObjectSeq[M sqlpipe.Model, O object.Type] func(ctx context.Context, ex sqlpipeex.SourceExecutor[M]) iter.Seq2[*O, error]

// List 从 SourceExecutor 中获取模型数据并转换为 Object 列表，支持自动计数
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

// Collect 将 Object 迭代器中的数据收集为切片
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

// FindOne 从 SourceExecutor 中获取单条模型数据并转换为 Object，未找到时返回 NotFound 错误
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

// FillSet 从 SourceExecutor 中批量获取数据并按 ID 填充到 Set 中的目标对象上
func FillSet[ID ~uint64, O object.Object[ID], M sqlpipe.Model](ctx context.Context, targets sqlpipeex.Set[ID, O], ex sqlpipeex.SourceExecutor[M], createSeq BuildObjectSeq[M, O]) error {
	items := sqlpipeex.OneToOne[ID, O]{}

	for x, err := range createSeq(ctx, ex) {
		if err != nil {
			return err
		}

		items.Record(any(x).(object.IDGetter[ID]).GetID(), x)
	}

	// have to wait side data filled
	for id, item := range items {
		for t := range targets.Records(id) {
			*t = *item
		}
	}

	return nil
}
