package digest

import (
	"github.com/opencontainers/go-digest"
)

// Digestible 提供摘要字段，嵌入结构体后可获取对象的内容摘要。
type Digestible struct {
	// 摘要
	Digest Digest `db:"f_digest,default=''" json:"digest"`
}

// GetDigest 返回当前对象的摘要值。
func (d Digestible) GetDigest() digest.Digest {
	return d.Digest
}
