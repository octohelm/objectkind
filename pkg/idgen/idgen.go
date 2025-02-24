package idgen

import (
	"context"
	"time"

	"github.com/go-courier/snowflakeid"
	"github.com/go-courier/snowflakeid/workeridutil"
	"github.com/octohelm/objectkind/pkg/idgen/internal"
)

var (
	startTime, _ = time.Parse(time.RFC3339, "2020-01-01T00:00:00Z")
	sff          = snowflakeid.NewSnowflakeFactory(16, 8, 5, startTime)
)

// +gengo:injectable:provider
type Gen interface {
	ID() (uint64, error)
}

type TypedGen[ID ~uint64] interface {
	NewTo(v *ID) error
}

// +gengo:injectable
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

// +gengo:injectable:provider Gen
type IDGen struct {
	Gen
}

func (i *IDGen) afterInit(ctx context.Context) error {
	g, err := sff.NewSnowflake(workeridutil.WorkerIDFromIP(internal.ResolveExposedIP()))
	if err != nil {
		return err
	}
	i.Gen = g
	return nil
}
