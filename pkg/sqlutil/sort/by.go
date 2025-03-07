package sort

import (
	"cmp"
	"fmt"
	"iter"
	"strings"

	"github.com/octohelm/storage/pkg/sqlbuilder"
	"github.com/octohelm/storage/pkg/sqlbuilder/modelscoped"
	"github.com/octohelm/storage/pkg/sqlpipe"
)

func DescSort[M sqlpipe.Model](s Sorter[M]) *By[M] {
	return &By[M]{
		By:     fmt.Sprintf("%s!%s", s.Name(), "desc"),
		Sorter: s,
	}
}

func AscSort[M sqlpipe.Model](s Sorter[M]) *By[M] {
	return &By[M]{
		By:     fmt.Sprintf("%s!%s", s.Name(), "asc"),
		Sorter: s,
	}
}

type By[M sqlpipe.Model] struct {
	By     string
	Sorter Sorter[M]
}

func (a *By[M]) IsZero() bool {
	return a.Sorter == nil
}

func (a *By[M]) OperatorType() sqlpipe.OperatorType {
	return sqlpipe.OperatorSort
}

func (a *By[M]) Next(src sqlpipe.Source[M]) sqlpipe.Source[M] {
	if a.Sorter == nil {
		return src
	}

	if strings.HasSuffix(a.By, "!desc") {
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
	return []byte(v.By), nil
}

func (s *By[M]) AsEnumValues(sorters iter.Seq[Sorter[M]]) (values []any) {
	for sorter := range sorters {
		for _, order := range []string{"asc", "desc"} {
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
			by.By = text
			by.Sorter = sorter
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
