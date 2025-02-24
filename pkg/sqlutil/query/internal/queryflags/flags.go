package queryflags

import "github.com/octohelm/objectkind/pkg/sqlutil/query/internal"

type Bools uint64

func (Bools) QueryOptions(internal.NotForPublicUse) {}

const (
	initFlag                      Bools = 1 << iota // reserved for the boolean value itself
	RequestCount                                    // 查询 Count
	RequestResourceStatus                           // 查询资源状态类信息
	RequestResourceOwner                            // 查询资源直接归属
	RequestResourceSecondaryOwner                   // 查询资源次要归属
	RequestSubResources                             // 查询子资源详细信息
	maxFlag
)

const (
	AllFlags = 1 |
		RequestCount |
		RequestResourceStatus |
		RequestResourceOwner |
		RequestResourceSecondaryOwner |
		RequestSubResources
)

type Flags struct{ Presence, Values uint64 }

func (flags Flags) Join(src Flags) Flags {
	flags.Presence |= src.Presence
	flags.Values &= ^src.Presence
	flags.Values |= src.Values
	return flags
}

func (flags *Flags) Set(f Bools) {
	id := uint64(f) &^ uint64(1)
	flags.Presence |= id
	flags.Values &= ^id
	flags.Values |= uint64(f&1) * id
}

func (flags Flags) Get(f Bools) bool {
	return flags.Values&uint64(f) > 0
}

func (flags Flags) Has(f Bools) bool {
	return flags.Presence&uint64(f) > 0
}

func (flags *Flags) Clear(f Bools) {
	mask := uint64(^f)
	flags.Presence &= mask
	flags.Values &= mask
}
