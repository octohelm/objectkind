package digest

import (
	"crypto/sha256"
	"fmt"
	"hash"
	"hash/fnv"
	"math/rand/v2"
	"strconv"
	"testing"

	"github.com/opencontainers/go-digest"
)

const (
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func generate(n int) []byte {
	prefix := []byte("txt,")
	charsets := letterBytes + strconv.FormatUint(10, 10)
	b := make([]byte, n)
	for i := range b {
		if i < len(prefix) {
			b[i] = prefix[i]
			continue
		}
		b[i] = charsets[rand.IntN(len(charsets))]
	}
	return b
}

func fromBytes(alg string, hash hash.Hash, data []byte) digest.Digest {
	d := NewDigester(alg, hash)
	d.Hash().Write(data)
	return d.Digest()
}

func runBenchmark(b *testing.B, alg string, createHash func() hash.Hash) {
	for i := range 6 {
		size := 2 << (4 * i)

		data := generate(size)

		b.Run(fmt.Sprintf("%s/%d", alg, size), func(b *testing.B) {
			b.Log(fromBytes(alg, createHash(), data))

			for b.Loop() {
				_ = fromBytes(alg, createHash(), data)
			}
		})
	}
}

func Benchmark(b *testing.B) {
	runBenchmark(b, "fnv128", func() hash.Hash { return fnv.New128() })
	runBenchmark(b, "fnv128a", func() hash.Hash { return fnv.New128a() })
	runBenchmark(b, "sha256", func() hash.Hash { return sha256.New() })
}
