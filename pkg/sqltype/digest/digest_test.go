package digest_test

import (
	"context"
	"crypto/sha256"
	"testing"

	"github.com/octohelm/x/cmp"
	. "github.com/octohelm/x/testing/v2"

	"github.com/octohelm/objectkind/pkg/sqltype/digest"
	"github.com/octohelm/objectkind/pkg/sqltype/digest/flags"
)

func TestDigestible(t *testing.T) {
	t.Run("零值 Digestible", func(t *testing.T) {
		d := digest.Digestible{}
		Then(
			t, "Digest 为空字符串",
			Expect(d.GetDigest(), Equal(digest.Digest(""))),
		)
	})

	t.Run("有值 Digestible", func(t *testing.T) {
		d := digest.Digestible{Digest: "sha256:abc123"}
		Then(
			t, "GetDigest 返回正确值",
			Expect(d.GetDigest(), Equal(digest.Digest("sha256:abc123"))),
		)
	})
}

func TestNewDigester(t *testing.T) {
	d := digest.NewDigester("sha256", sha256.New())

	t.Run("Digest 返回非空摘要", func(t *testing.T) {
		h := d.Hash()
		h.Write([]byte("hello"))
		dgst := d.Digest()
		Then(
			t, "摘要非空",
			Expect(dgst != "", Be(cmp.True())),
		)
	})

	t.Run("Digest 幂等", func(t *testing.T) {
		d1 := digest.NewDigester("sha256", sha256.New())
		d1.Hash().Write([]byte("hello"))
		first := d1.Digest()

		d2 := digest.NewDigester("sha256", sha256.New())
		d2.Hash().Write([]byte("hello"))
		second := d2.Digest()

		Then(
			t, "相同输入相同算法摘要一致",
			Expect(first, Equal(second)),
		)
	})
}

func TestHashTo(t *testing.T) {
	var d digest.Digest
	err := digest.HashTo(&d, map[string]any{"key": "value"})
	Must(t, func() error { return err })

	Then(
		t, "摘要非空",
		Expect(d != "", Be(cmp.True())),
	)

	t.Run("相同输入摘要一致", func(t *testing.T) {
		var d2 digest.Digest
		err := digest.HashTo(&d2, map[string]any{"key": "value"})
		Must(t, func() error { return err })

		Then(
			t, "摘要相同",
			Expect(d, Equal(d2)),
		)
	})

	t.Run("不同输入摘要不同", func(t *testing.T) {
		var d3 digest.Digest
		err := digest.HashTo(&d3, map[string]any{"key": "different"})
		Must(t, func() error { return err })

		Then(
			t, "摘要不同",
			Expect(d != d3, Be(cmp.True())),
		)
	})
}

type testAnn struct {
	annotations map[string]string
}

func (a testAnn) GetAnnotations() map[string]string { return a.annotations }
func (a testAnn) GetAnnotation(k string) (string, bool) {
	v, ok := a.annotations[k]
	return v, ok
}
func (a *testAnn) SetAnnotations(annotations map[string]string) { a.annotations = annotations }
func (a *testAnn) SetAnnotation(k, v string) {
	if a.annotations == nil {
		a.annotations = map[string]string{}
	}
	a.annotations[k] = v
}

func TestOmitAnnotations(t *testing.T) {
	t.Run("移除 known annotations", func(t *testing.T) {
		src := &testAnn{
			annotations: map[string]string{
				"spec/digest":     "abc",
				"revision/id":     "1",
				"revision/digest": "xyz",
				"custom/key":      "val",
			},
		}

		digest.OmitAnnotations(src)

		_, ok := src.GetAnnotation("spec/digest")
		Then(
			t, "spec/digest 被移除",
			Expect(ok, Be(cmp.False())),
		)
		_, ok = src.GetAnnotation("revision/id")
		Then(
			t, "revision/id 被移除",
			Expect(ok, Be(cmp.False())),
		)
		_, ok = src.GetAnnotation("revision/digest")
		Then(
			t, "revision/digest 被移除",
			Expect(ok, Be(cmp.False())),
		)
	})

	t.Run("保留 custom annotations", func(t *testing.T) {
		src := &testAnn{
			annotations: map[string]string{
				"spec/digest": "abc",
				"custom/key":  "val",
			},
		}

		digest.OmitAnnotations(src)

		v, ok := src.GetAnnotation("custom/key")
		Then(
			t, "custom/key 保留",
			Expect(ok, Be(cmp.True())),
		)
		Then(
			t, "custom/key 值正确",
			Expect(v, Equal("val")),
		)
	})

	t.Run("自定义 omitKeys", func(t *testing.T) {
		src := &testAnn{
			annotations: map[string]string{
				"my-key":  "remove-me",
				"keep-me": "keep",
			},
		}

		digest.OmitAnnotations(src, "my-key")

		_, ok := src.GetAnnotation("my-key")
		Then(
			t, "my-key 被移除",
			Expect(ok, Be(cmp.False())),
		)

		v, ok := src.GetAnnotation("keep-me")
		Then(
			t, "keep-me 保留",
			Expect(ok, Be(cmp.True())),
		)
		Then(
			t, "keep-me 值正确",
			Expect(v, Equal("keep")),
		)
	})
}

func TestNewHasher(t *testing.T) {
	ctx := context.Background()

	t.Run("基本哈希", func(t *testing.T) {
		h := digest.NewHasher(ctx, nil)
		err := h.Hash(map[string]any{"x": 1})
		Must(t, func() error { return err })

		Then(
			t, "Digest 非空",
			Expect(h.Digest() != "", Be(cmp.True())),
		)
	})

	t.Run("相同内容哈希一致", func(t *testing.T) {
		h1 := digest.NewHasher(ctx, nil)
		Must(t, func() error { return h1.Hash(map[string]any{"x": 1}) })

		h2 := digest.NewHasher(ctx, nil)
		Must(t, func() error { return h2.Hash(map[string]any{"x": 1}) })

		Then(
			t, "摘要相同",
			Expect(h1.Digest(), Equal(h2.Digest())),
		)
	})
}

func TestHasherSkipIfExists(t *testing.T) {
	target := &testAnn{
		annotations: map[string]string{
			"revision/digest": "sha256:precomputed",
		},
	}

	ctx := flags.InjectContextWith(context.Background(), flags.HashSkipIfExists)

	h := digest.NewHasher(ctx, target)
	BeforeHash := h.Digest()

	err := h.Hash(map[string]any{"different": "content"})
	Must(t, func() error { return err })

	Then(
		t, "跳过时摘要保持不变",
		Expect(h.Digest(), Equal(BeforeHash)),
	)
}
