package digest

import (
	"context"

	"github.com/octohelm/courier/pkg/validator"
	metav1 "github.com/octohelm/objectkind/pkg/apis/meta/v1"
	"github.com/octohelm/objectkind/pkg/object"
	"github.com/octohelm/objectkind/pkg/sqltype/digest/flags"
	"github.com/octohelm/x/anyjson"
	"github.com/opencontainers/go-digest"
)

type Hasher interface {
	Hash(target any) error
	Digest() Digest
}

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
