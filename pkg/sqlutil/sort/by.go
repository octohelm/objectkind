package sort

import (
	"cmp"
	"encoding"
	"fmt"
	"iter"
	"strings"

	"github.com/octohelm/enumeration/pkg/enumeration"
	"github.com/octohelm/storage/pkg/sort"
	"github.com/octohelm/storage/pkg/sqlbuilder"
	"github.com/octohelm/storage/pkg/sqlbuilder/modelscoped"
	"github.com/octohelm/storage/pkg/sqlpipe"
)

// 创建排序
func With[M sqlpipe.Model, E enumeration.CanEnumValues](by *sort.By[E], mapper func(e E) Sorter[M]) *By[M] {
	if by == nil {
		return nil
	}

	sorter := mapper(by.Field)
	if sorter == nil {
		return nil
	}

	return &By[M]{
		Order:      by.Order,
		Sorter:     sorter,
		underlying: by,
	}
}

// DescSort 创建降序排序
func DescSort[M sqlpipe.Model](s Sorter[M]) *By[M] {
	return &By[M]{
		Order:  sort.Desc,
		Sorter: s,
		raw:    fmt.Sprintf("%s!%s", s.Name(), sort.Desc),
	}
}

// AscSort 创建升序排序
func AscSort[M sqlpipe.Model](s Sorter[M]) *By[M] {
	return &By[M]{
		Order:  sort.Asc,
		Sorter: s,
		raw:    fmt.Sprintf("%s!%s", s.Name(), sort.Asc),
	}
}

// By 排序操作符，封装排序表达式与排序器实现
type By[M sqlpipe.Model] struct {
	Order  sort.Order
	Sorter Sorter[M]

	raw        string
	underlying encoding.TextMarshaler
}

func (a *By[M]) IsZero() bool {
	return a == nil || a.Sorter == nil
}

func (a *By[M]) OperatorType() sqlpipe.OperatorType {
	return sqlpipe.OperatorSort
}

func (a *By[M]) Next(src sqlpipe.Source[M]) sqlpipe.Source[M] {
	if a.IsZero() {
		return src
	}

	if a.Order == sort.Desc {
		return sqlpipe.Pipe(src, sqlpipe.SourceOperatorFunc[M](sqlpipe.OperatorSort, func(src sqlpipe.Source[M]) sqlpipe.Source[M] {
			return a.Sorter.Sort(src, func(col sqlbuilder.Column) sqlpipe.SourceOperator[M] {
				return sqlpipe.DescSort(
					modelscoped.CastColumn[M](col),
					sqlbuilder.NullsLast(),
				)
			})
		}))
	}

	return sqlpipe.Pipe(src, sqlpipe.SourceOperatorFunc[M](sqlpipe.OperatorSort, func(src sqlpipe.Source[M]) sqlpipe.Source[M] {
		return a.Sorter.Sort(src, func(col sqlbuilder.Column) sqlpipe.SourceOperator[M] {
			return sqlpipe.AscSort(
				modelscoped.CastColumn[M](col),
				sqlbuilder.NullsFirst(),
			)
		})
	}))
}

func (v *By[M]) MarshalText() ([]byte, error) {
	if v == nil {
		return nil, nil
	}
	if v.underlying != nil {
		return v.underlying.MarshalText()
	}
	return []byte(v.raw), nil
}

func (v *By[M]) String() string {
	raw, _ := v.MarshalText()
	return string(raw)
}

func (s *By[M]) AsEnumValues(sorters iter.Seq[Sorter[M]]) (values []any) {
	for sorter := range sorters {
		for _, order := range []sort.Order{sort.Asc, sort.Desc} {
			values = append(values, &enumValue{
				value: fmt.Sprintf("%s!%s", sorter.Name(), order),
				label: sorter.Label(),
			})
		}
	}
	return
}

func (by *By[M]) Unmarshal(text string, sorters iter.Seq[Sorter[M]]) error {
	if text == "" {
		return nil
	}

	parts := strings.Split(strings.ToLower(text), "!")

	for sorter := range sorters {
		if parts[0] == strings.ToLower(sorter.Name()) {
			by.raw = text
			by.Order = sort.Asc
			by.Sorter = sorter

			if len(parts) > 0 {
				switch strings.ToLower(parts[1]) {
				case "desc":
					by.Order = sort.Desc
				}
			}

			return nil
		}
	}

	return fmt.Errorf("invalid sort type %s", text)
}

type enumValue struct {
	value string
	label string
}

func (v enumValue) MarshalText() ([]byte, error) {
	return []byte(v.value), nil
}

func (v enumValue) Label() string {
	label := cmp.Or(v.label, v.value)
	if strings.HasSuffix(v.value, "!desc") {
		return fmt.Sprintf("%s逆序", label)
	}
	return fmt.Sprintf("%s正序", label)
}
