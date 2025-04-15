package digest

import (
	"hash"
	"sync"

	"github.com/opencontainers/go-digest"
)

func NewDigester(alg string, h hash.Hash) digest.Digester {
	return &digester{
		alg:  digest.Algorithm(alg),
		hash: h,
	}
}

type digester struct {
	alg  digest.Algorithm
	hash hash.Hash
	once sync.Once
	dgst digest.Digest
}

func (d *digester) Digest() digest.Digest {
	if d.dgst == "" {
		d.once.Do(func() {
			d.dgst = digest.NewDigest(d.alg, d.hash)
		})
	}

	return d.dgst
}

func (d *digester) Hash() hash.Hash {
	return d.hash
}
