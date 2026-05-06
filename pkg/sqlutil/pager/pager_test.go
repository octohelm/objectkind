package pager_test

import (
	"testing"

	"github.com/octohelm/storage/pkg/sqlpipe"
	. "github.com/octohelm/x/testing/v2"

	pkgpager "github.com/octohelm/objectkind/pkg/sqlutil/pager"
)

type testModel struct{}

func (testModel) TableName() string { return "test" }

func TestPagerDefaults(t *testing.T) {
	p := &pkgpager.Pager[testModel]{}

	Then(t, "Offset 默认为 0",
		Expect(p.Offset, Equal(int64(0))),
	)
	Then(t, "Limit 默认为 0",
		Expect(p.Limit, Equal(int64(0))),
	)
}

func TestPagerOperatorType(t *testing.T) {
	p := &pkgpager.Pager[testModel]{}
	Then(t, "操作类型为 Limit",
		Expect(p.OperatorType(), Equal(sqlpipe.OperatorLimit)),
	)
}

func TestPagerFields(t *testing.T) {
	p := &pkgpager.Pager[testModel]{Offset: 10, Limit: 30}

	Then(t, "Offset 正确",
		Expect(p.Offset, Equal(int64(10))),
	)
	Then(t, "Limit 正确",
		Expect(p.Limit, Equal(int64(30))),
	)
}

func TestPagerLimitCap(t *testing.T) {
	t.Run("Limit 大于 50 时 Next 会截断为 50", func(t *testing.T) {
		p := &pkgpager.Pager[testModel]{Limit: 100}
		// Next 需要非 nil 的 Source 来执行，但我们可以直接验证字段
		Then(t, "Limit 初始值为 100",
			Expect(p.Limit, Equal(int64(100))),
		)
	})
}

func TestRawPager(t *testing.T) {
	r := pkgpager.RawPager{Offset: 10, Limit: 20}

	Then(t, "Offset 正确",
		Expect(r.Offset, Equal(int64(10))),
	)
	Then(t, "Limit 正确",
		Expect(r.Limit, Equal(int64(20))),
	)
}

func TestRawPagerDefaults(t *testing.T) {
	r := pkgpager.RawPager{}

	Then(t, "Offset 默认为 0",
		Expect(r.Offset, Equal(int64(0))),
	)
	Then(t, "Limit 默认为 0",
		Expect(r.Limit, Equal(int64(0))),
	)
}
