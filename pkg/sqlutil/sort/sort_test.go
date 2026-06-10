package sort_test

import (
	"testing"

	"github.com/octohelm/storage/pkg/sqlbuilder"
	"github.com/octohelm/storage/pkg/sqlpipe"
	"github.com/octohelm/x/cmp"
	. "github.com/octohelm/x/testing/v2"

	sortpkg "github.com/octohelm/objectkind/pkg/sqlutil/sort"
)

type testModel struct{}

func (testModel) TableName() string { return "test" }

type testSorter struct{}

func (testSorter) Name() string  { return "created_at" }
func (testSorter) Label() string { return "创建时间" }
func (testSorter) Sort(src sqlpipe.Source[testModel], by func(col sqlbuilder.Column) sqlpipe.SourceOperator[testModel]) sqlpipe.Source[testModel] {
	return src
}

func sorterSeq(sorters ...sortpkg.Sorter[testModel]) func(func(sortpkg.Sorter[testModel]) bool) {
	return func(yield func(sortpkg.Sorter[testModel]) bool) {
		for _, s := range sorters {
			if !yield(s) {
				return
			}
		}
	}
}

func TestAscSort(t *testing.T) {
	by := sortpkg.AscSort(testSorter{})

	Then(
		t, "By 字符串以 asc 结尾",
		Expect(by.String(), Equal("created_at!asc")),
	)
	Then(
		t, "Sorter 非空",
		Expect(by.IsZero(), Be(cmp.False())),
	)
	Then(
		t, "操作类型为排序",
		Expect(by.OperatorType(), Equal(sqlpipe.OperatorSort)),
	)
}

func TestDescSort(t *testing.T) {
	by := sortpkg.DescSort(testSorter{})

	Then(
		t, "By 字符串以 desc 结尾",
		Expect(by.String(), Equal("created_at!desc")),
	)
	Then(
		t, "Sorter 非空",
		Expect(by.IsZero(), Be(cmp.False())),
	)
}

func TestByIsZero(t *testing.T) {
	t.Run("Sorter 为 nil 时为零值", func(t *testing.T) {
		by := &sortpkg.By[testModel]{}
		Then(
			t, "IsZero 返回 true",
			Expect(by.IsZero(), Be(cmp.True())),
		)
	})

	t.Run("Sorter 非 nil 时非零值", func(t *testing.T) {
		by := sortpkg.AscSort(testSorter{})
		Then(
			t, "IsZero 返回 false",
			Expect(by.IsZero(), Be(cmp.False())),
		)
	})
}

func TestByOperatorType(t *testing.T) {
	by := sortpkg.DescSort(testSorter{})
	Then(
		t, "返回 OperatorSort",
		Expect(by.OperatorType(), Equal(sqlpipe.OperatorSort)),
	)
}

func TestByMarshalText(t *testing.T) {
	by := sortpkg.DescSort(testSorter{})
	b, err := by.MarshalText()
	Must(t, func() error { return err })
	Then(
		t, "序列化为 By 字符串",
		Expect(string(b), Equal("created_at!desc")),
	)
}

func TestByAsEnumValues(t *testing.T) {
	s1 := testSorter{}
	s2 := &customSorter{}

	by := &sortpkg.By[testModel]{}
	values := by.AsEnumValues(sorterSeq(s1, s2))

	// Should generate both asc and desc for each sorter
	Then(
		t, "生成 4 个枚举值 (2 sorters × 2 directions)",
		Expect(len(values), Equal(4)),
	)
}

type customSorter struct{}

func (customSorter) Name() string  { return "updated_at" }
func (customSorter) Label() string { return "" }
func (customSorter) Sort(src sqlpipe.Source[testModel], by func(col sqlbuilder.Column) sqlpipe.SourceOperator[testModel]) sqlpipe.Source[testModel] {
	return src
}

func TestByUnmarshal(t *testing.T) {
	t.Run("解析 asc", func(t *testing.T) {
		by := &sortpkg.By[testModel]{}
		err := by.Unmarshal("created_at!asc", sorterSeq(testSorter{}))
		Must(t, func() error { return err })

		Then(
			t, "By 字段匹配",
			Expect(by.String(), Equal("created_at!asc")),
		)
		Then(
			t, "Sorter 非空",
			Expect(by.IsZero(), Be(cmp.False())),
		)
	})

	t.Run("解析 desc", func(t *testing.T) {
		by := &sortpkg.By[testModel]{}
		err := by.Unmarshal("created_at!desc", sorterSeq(testSorter{}))
		Must(t, func() error { return err })

		Then(
			t, "By 字段匹配",
			Expect(by.String(), Equal("created_at!desc")),
		)
		Then(
			t, "Sorter 非空",
			Expect(by.IsZero(), Be(cmp.False())),
		)
	})

	t.Run("空字符串不做任何操作", func(t *testing.T) {
		by := &sortpkg.By[testModel]{}
		err := by.Unmarshal("", sorterSeq(testSorter{}))
		Must(t, func() error { return err })

		Then(
			t, "Sorter 保持 nil",
			Expect(by.IsZero(), Be(cmp.True())),
		)
	})

	t.Run("无效排序类型返回错误", func(t *testing.T) {
		by := &sortpkg.By[testModel]{}
		err := by.Unmarshal("invalid!asc", sorterSeq(testSorter{}))
		Then(
			t, "应返回错误",
			Expect(err != nil, Be(cmp.True())),
		)
	})
}

type labeler interface {
	Label() string
}

func TestEnumValueLabel(t *testing.T) {
	by := sortpkg.AscSort(testSorter{})
	values := by.AsEnumValues(sorterSeq(testSorter{}))

	labels := make([]string, 0, len(values))
	for _, v := range values {
		if l, ok := v.(labeler); ok {
			labels = append(labels, l.Label())
		}
	}

	Then(
		t, "包含 asc label",
		Expect(labels[0], Equal("创建时间正序")),
	)
	Then(
		t, "包含 desc label",
		Expect(labels[1], Equal("创建时间逆序")),
	)
}
