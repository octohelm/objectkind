package idgen

import (
	"context"
	"time"

	"github.com/octohelm/idx/pkg/snowflake"
	"github.com/octohelm/idx/pkg/workerid"

	"github.com/octohelm/objectkind/pkg/idgen/internal"
)

var (
	startTime, _ = time.Parse(time.RFC3339, "2020-01-01T00:00:00Z")
	sff          = snowflake.NewFactory(16, 8, 5, startTime)
)

// +gengo:injectable:provider
// Gen 唯一 ID 生成器接口。
type Gen interface {
	ID() (uint64, error)
}

// TypedGen 类型安全的 ID 生成器接口。
type TypedGen[ID ~uint64] interface {
	NewTo(v *ID) error
}

// +gengo:injectable
// Typed 类型安全的 ID 生成器实现，基于 Gen 为指定 ID 类型生成唯一 ID。
type Typed[ID ~uint64] struct {
	g Gen `inject:""`
}

func (t *Typed[ID]) NewTo(v *ID) error {
	u, err := t.g.ID()
	if err != nil {
		return err
	}
	*v = ID(u)
	return nil
}

// +gengo:injectable:provider
// IDGen 基于雪花算法的全局唯一 ID 生成器。
type IDGen struct {
	gen Gen `provide:""`
}

func (i *IDGen) afterInit(ctx context.Context) error {
	g, err := sff.NewSnowflake(workerid.FromIP(internal.ResolveExposedIP()))
	if err != nil {
		return err
	}
	i.gen = g
	return nil
}
