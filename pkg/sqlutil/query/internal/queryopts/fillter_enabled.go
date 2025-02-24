package queryopts

import (
	"github.com/octohelm/objectkind/pkg/object"
	"github.com/octohelm/objectkind/pkg/sqlutil/query/internal"
)

type FillerEnabled[O object.Type] bool

func (FillerEnabled[O]) Is(x any) bool {
	_, ok := x.(FillerEnabled[O])
	return ok
}

func (FillerEnabled[O]) QueryOptions(use internal.NotForPublicUse) {
}
