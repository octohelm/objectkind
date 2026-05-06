package digest

import (
	"context"

	"github.com/opencontainers/go-digest"

	"github.com/octohelm/courier/pkg/validator"
	"github.com/octohelm/x/anyjson"

	metav1 "github.com/octohelm/objectkind/pkg/apis/meta/v1"
	"github.com/octohelm/objectkind/pkg/object"
	"github.com/octohelm/objectkind/pkg/sqltype/digest/flags"
)

// Hasher 定义内容哈希器，可对目标进行哈希计算并返回摘要。
type Hasher interface {
	Hash(target any) error
	Digest() Digest
}

// NewHasher 创建哈希器，从上下文读取 Flags 以控制跳过行为，若 src 已带摘要则预先载入。
func NewHasher(ctx context.Context, src any) Hasher {
	h := &hasher{}

	if target, ok := src.(object.Annotater); ok {
		_ = metav1.AnnotationRevisionDigest.UnmarshalFrom(target, &h.digest)
	}

	h.skipIfExists = flags.FromContext(ctx).Is(flags.HashSkipIfExists)

	return h
}

type hasher struct {
	digest       Digest
	skipIfExists bool
}

func (h *hasher) Digest() Digest {
	return h.digest
}

func (h *hasher) Hash(target any) error {
	if h.skipIfExists && h.digest != "" {
		return nil
	}

	o, err := fromValue(target)
	if err != nil {
		return nil
	}
	d := digest.SHA256.Digester()
	if err := validator.MarshalWrite(d.Hash(), anyjson.Sorted(o)); err != nil {
		return err
	}
	h.digest = d.Digest()
	return nil
}

// HashTo 对目标计算 SHA256 摘要并写入 dgst，等价于无状态的快速哈希。
func HashTo(dgst *Digest, target any) error {
	h := &hasher{}

	if err := h.Hash(target); err != nil {
		return err
	}

	*dgst = h.Digest()

	return nil
}

func fromValue(v any) (anyjson.Valuer, error) {
	if vv, ok := v.(anyjson.Valuer); ok {
		return vv, nil
	}
	return anyjson.FromValue(v)
}
