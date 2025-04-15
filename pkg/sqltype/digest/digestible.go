package digest

import "github.com/opencontainers/go-digest"

type Digestible struct {
	// 摘要
	Digest Digest `db:"f_digest,default=''" json:"digest"`
}

func (d Digestible) GetDigest() digest.Digest {
	return d.Digest
}
