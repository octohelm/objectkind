package v1_test

import (
	"testing"
	"time"

	"github.com/octohelm/x/cmp"
	. "github.com/octohelm/x/testing/v2"

	metav1 "github.com/octohelm/objectkind/pkg/apis/meta/v1"
	"github.com/octohelm/objectkind/pkg/object"
)

func TestParseGroupVersion(t *testing.T) {
	t.Run("空字符串", func(t *testing.T) {
		gv, err := metav1.ParseGroupVersion("")
		Must(t, func() error { return err })
		Then(
			t, "返回零值",
			Expect(gv, Equal(metav1.GroupVersion{})),
		)
	})

	t.Run("斜杠字符串", func(t *testing.T) {
		gv, err := metav1.ParseGroupVersion("/")
		Must(t, func() error { return err })
		Then(
			t, "返回零值",
			Expect(gv, Equal(metav1.GroupVersion{})),
		)
	})

	t.Run("仅版本", func(t *testing.T) {
		gv, err := metav1.ParseGroupVersion("v1")
		Must(t, func() error { return err })
		Then(
			t, "Group 为空，Version 为 v1",
			Expect(gv, Equal(metav1.GroupVersion{Version: "v1"})),
		)
	})

	t.Run("Group 和版本", func(t *testing.T) {
		gv, err := metav1.ParseGroupVersion("example.com/v1")
		Must(t, func() error { return err })
		Then(
			t, "Group 为 example.com，Version 为 v1",
			Expect(gv, Equal(metav1.GroupVersion{Group: "example.com", Version: "v1"})),
		)
	})

	t.Run("多斜杠错误", func(t *testing.T) {
		_, err := metav1.ParseGroupVersion("too/many/slashes")
		Then(
			t, "返回错误",
			Expect(err == nil, Be(cmp.False())),
		)
	})
}

func TestGroupVersion(t *testing.T) {
	t.Run("IsZero", func(t *testing.T) {
		t.Run("空值返回 true", func(t *testing.T) {
			Then(
				t, "零值 IsZero 为 true",
				Expect(metav1.GroupVersion{}.IsZero(), Be(cmp.True())),
			)
		})
		t.Run("有 Version 返回 false", func(t *testing.T) {
			Then(
				t, "非零值 IsZero 为 false",
				Expect(metav1.GroupVersion{Version: "v1"}.IsZero(), Be(cmp.False())),
			)
		})
	})

	t.Run("String", func(t *testing.T) {
		t.Run("仅 Version", func(t *testing.T) {
			gv := metav1.GroupVersion{Version: "v1"}
			Then(
				t, "仅返回版本",
				Expect(gv.String(), Equal("v1")),
			)
		})
		t.Run("Group 和 Version", func(t *testing.T) {
			gv := metav1.GroupVersion{Group: "example.com", Version: "v1"}
			Then(
				t, "返回 group/version",
				Expect(gv.String(), Equal("example.com/v1")),
			)
		})
	})

	t.Run("WithKind", func(t *testing.T) {
		gv := metav1.GroupVersion{Group: "example.com", Version: "v1"}
		gvk := gv.WithKind("Foo")
		Then(
			t, "生成 GroupVersionKind",
			Expect(gvk, Equal(metav1.GroupVersionKind{
				Group: "example.com", Version: "v1", Kind: "Foo",
			})),
		)
	})
}

func TestGroupKind(t *testing.T) {
	t.Run("IsZero", func(t *testing.T) {
		t.Run("空值返回 true", func(t *testing.T) {
			Then(
				t, "零值 IsZero 为 true",
				Expect(metav1.GroupKind{}.IsZero(), Be(cmp.True())),
			)
		})
		t.Run("有 Group 返回 false", func(t *testing.T) {
			Then(
				t, "非零值 IsZero 为 false",
				Expect(metav1.GroupKind{Group: "example.com"}.IsZero(), Be(cmp.False())),
			)
		})
	})

	t.Run("WithVersion", func(t *testing.T) {
		gk := metav1.GroupKind{Group: "example.com", Kind: "Foo"}
		gvk := gk.WithVersion("v1")
		Then(
			t, "生成 GroupVersionKind",
			Expect(gvk, Equal(metav1.GroupVersionKind{
				Group: "example.com", Version: "v1", Kind: "Foo",
			})),
		)
	})
}

func TestFromAPIVersionAndKind(t *testing.T) {
	t.Run("正常解析", func(t *testing.T) {
		gvk := metav1.FromAPIVersionAndKind("example.com/v1", "Foo")
		Then(
			t, "返回完整 GroupVersionKind",
			Expect(gvk, Equal(metav1.GroupVersionKind{
				Group: "example.com", Version: "v1", Kind: "Foo",
			})),
		)
	})

	t.Run("空 apiVersion", func(t *testing.T) {
		gvk := metav1.FromAPIVersionAndKind("", "Foo")
		Then(
			t, "仅保留 Kind",
			Expect(gvk, Equal(metav1.GroupVersionKind{Kind: "Foo"})),
		)
	})

	t.Run("斜杠 apiVersion", func(t *testing.T) {
		gvk := metav1.FromAPIVersionAndKind("/", "Foo")
		Then(
			t, "仅保留 Kind",
			Expect(gvk, Equal(metav1.GroupVersionKind{Kind: "Foo"})),
		)
	})

	t.Run("错误格式回退", func(t *testing.T) {
		gvk := metav1.FromAPIVersionAndKind("too/many/slashes", "Foo")
		Then(
			t, "解析失败时仅保留 Kind",
			Expect(gvk, Equal(metav1.GroupVersionKind{Kind: "Foo"})),
		)
	})
}

func TestGroupVersionKind(t *testing.T) {
	t.Run("IsZero", func(t *testing.T) {
		t.Run("空值返回 true", func(t *testing.T) {
			Then(
				t, "零值 IsZero 为 true",
				Expect(metav1.GroupVersionKind{}.IsZero(), Be(cmp.True())),
			)
		})
		t.Run("有 Kind 返回 false", func(t *testing.T) {
			Then(
				t, "非零值 IsZero 为 false",
				Expect(metav1.GroupVersionKind{Kind: "Foo"}.IsZero(), Be(cmp.False())),
			)
		})
	})

	t.Run("GroupKind", func(t *testing.T) {
		gvk := metav1.GroupVersionKind{Group: "example.com", Version: "v1", Kind: "Foo"}
		Then(
			t, "提取 GroupKind",
			Expect(gvk.GroupKind(), Equal(metav1.GroupKind{Group: "example.com", Kind: "Foo"})),
		)
	})

	t.Run("GroupVersion", func(t *testing.T) {
		gvk := metav1.GroupVersionKind{Group: "example.com", Version: "v1", Kind: "Foo"}
		Then(
			t, "提取 GroupVersion",
			Expect(gvk.GroupVersion(), Equal(metav1.GroupVersion{Group: "example.com", Version: "v1"})),
		)
	})

	t.Run("ToAPIVersionAndKind", func(t *testing.T) {
		t.Run("完整值", func(t *testing.T) {
			gvk := metav1.GroupVersionKind{Group: "example.com", Version: "v1", Kind: "Foo"}
			apiVersion, kind := gvk.ToAPIVersionAndKind()
			Then(
				t, "返回 apiVersion",
				Expect(apiVersion, Equal("example.com/v1")),
			)
			Then(
				t, "返回 kind",
				Expect(kind, Equal("Foo")),
			)
		})

		t.Run("零值返回空字符串", func(t *testing.T) {
			apiVersion, kind := metav1.GroupVersionKind{}.ToAPIVersionAndKind()
			Then(
				t, "apiVersion 为空",
				Expect(apiVersion, Equal("")),
			)
			Then(
				t, "kind 为空",
				Expect(kind, Equal("")),
			)
		})
	})
}

func TestList(t *testing.T) {
	t.Run("零值", func(t *testing.T) {
		l := metav1.List[string]{}
		Then(
			t, "Items 为 nil",
			Expect(l.Items == nil, Be(cmp.True())),
		)
		Then(
			t, "Total 为 0",
			Expect(l.Total, Equal(int64(0))),
		)
	})

	t.Run("Add", func(t *testing.T) {
		l := &metav1.List[string]{}
		a, b := "a", "b"
		l.Add(&a)
		l.Add(&b)
		Then(
			t, "Items 长度",
			Expect(len(l.Items), Equal(2)),
		)
	})
}

func TestDescriber(t *testing.T) {
	t.Run("Name 读写", func(t *testing.T) {
		d := &metav1.Describer{}
		d.SetName("test-name")
		Then(
			t, "GetName 返回设置的值",
			Expect(d.GetName(), Equal("test-name")),
		)
	})

	t.Run("Description 读写", func(t *testing.T) {
		d := &metav1.Describer{}
		d.SetDescription("test-desc")
		Then(
			t, "GetDescription 返回设置的值",
			Expect(d.GetDescription(), Equal("test-desc")),
		)
	})

	t.Run("Annotations", func(t *testing.T) {
		t.Run("零值 GetAnnotation 返回 ok=false", func(t *testing.T) {
			d := metav1.Describer{}
			_, ok := d.GetAnnotation("key")
			Then(
				t, "ok 为 false",
				Expect(ok, Be(cmp.False())),
			)
		})

		t.Run("SetAnnotation 自动初始化 map", func(t *testing.T) {
			d := &metav1.Describer{}
			d.SetAnnotation("key", "value")
			Then(
				t, "GetAnnotations 非空",
				Expect(d.GetAnnotations() == nil, Be(cmp.False())),
			)
			v, ok := d.GetAnnotation("key")
			Then(
				t, "ok 为 true",
				Expect(ok, Be(cmp.True())),
			)
			Then(
				t, "值正确",
				Expect(v, Equal("value")),
			)
		})

		t.Run("SetAnnotations 整体替换", func(t *testing.T) {
			d := &metav1.Describer{}
			d.SetAnnotations(map[string]string{"a": "1", "b": "2"})
			v, ok := d.GetAnnotation("a")
			Then(
				t, "ok 为 true",
				Expect(ok, Be(cmp.True())),
			)
			Then(
				t, "值正确",
				Expect(v, Equal("1")),
			)
		})
	})
}

func TestIdentifiable(t *testing.T) {
	t.Run("ID 读写", func(t *testing.T) {
		o := &metav1.Identifiable[uint64]{}
		o.SetID(42)
		Then(
			t, "GetID 返回设置的值",
			Expect(o.GetID(), Equal(uint64(42))),
		)
	})

	t.Run("RefID 转换", func(t *testing.T) {
		o := &metav1.Identifiable[uint64]{}
		o.SetFromRefID(99)
		Then(
			t, "GetID 返回转换后的值",
			Expect(o.GetID(), Equal(uint64(99))),
		)
		Then(
			t, "GetAsRefID 返回 RefID",
			Expect(uint64(o.GetAsRefID()), Equal(uint64(99))),
		)
	})
}

func TestCodable(t *testing.T) {
	t.Run("Code 读写", func(t *testing.T) {
		o := &metav1.Codable[string]{}
		o.SetCode("abc")
		Then(
			t, "GetCode 返回设置的值",
			Expect(o.GetCode(), Equal("abc")),
		)
	})

	t.Run("RefCode 转换", func(t *testing.T) {
		o := &metav1.Codable[string]{}
		o.SetFromRefCode("xyz")
		Then(
			t, "GetCode 返回转换后的值",
			Expect(o.GetCode(), Equal("xyz")),
		)
		Then(
			t, "GetAsRefCode 返回 RefCode",
			Expect(string(o.GetAsRefCode()), Equal("xyz")),
		)
	})
}

func TestTypeMeta(t *testing.T) {
	t.Run("Kind 读写", func(t *testing.T) {
		tm := &metav1.TypeMeta{}
		tm.SetKind("Pod")
		Then(
			t, "GetKind 返回设置的值",
			Expect(tm.GetKind(), Equal("Pod")),
		)
	})

	t.Run("APIVersion 读写", func(t *testing.T) {
		tm := &metav1.TypeMeta{}
		tm.SetAPIVersion("v1")
		Then(
			t, "GetAPIVersion 返回设置的值",
			Expect(tm.GetAPIVersion(), Equal("v1")),
		)
	})
}

func TestObject(t *testing.T) {
	t.Run("Object 嵌入 Metadata 和 Identifiable", func(t *testing.T) {
		o := &metav1.Object[uint64]{}
		o.SetKind("TestKind")
		o.SetName("test-object")
		o.SetID(1)

		Then(
			t, "Kind 可访问",
			Expect(o.GetKind(), Equal("TestKind")),
		)
		Then(
			t, "Name 可访问",
			Expect(o.GetName(), Equal("test-object")),
		)
		Then(
			t, "ID 可访问",
			Expect(o.GetID(), Equal(uint64(1))),
		)
	})
}

func TestCodableObject(t *testing.T) {
	t.Run("CodableObject 嵌入 Object 和 Codable", func(t *testing.T) {
		o := &metav1.CodableObject[uint64, string]{}
		o.SetKind("TestKind")
		o.SetName("test")
		o.SetID(2)
		o.SetCode("code-001")

		Then(
			t, "Kind 可访问",
			Expect(o.GetKind(), Equal("TestKind")),
		)
		Then(
			t, "ID 可访问",
			Expect(o.GetID(), Equal(uint64(2))),
		)
		Then(
			t, "Code 可访问",
			Expect(o.GetCode(), Equal("code-001")),
		)
	})
}

func TestRequest(t *testing.T) {
	t.Run("Request 嵌入 TypeMeta 和 Describer", func(t *testing.T) {
		r := &metav1.Request[metav1.TypeMeta]{}
		r.SetKind("TestKind")
		r.SetName("req-name")
		r.SetDescription("req-desc")

		Then(
			t, "Kind 可访问",
			Expect(r.GetKind(), Equal("TestKind")),
		)
		Then(
			t, "Name 可访问",
			Expect(r.GetName(), Equal("req-name")),
		)
		Then(
			t, "Description 可访问",
			Expect(r.GetDescription(), Equal("req-desc")),
		)
	})
}

func TestOperationTimestamps(t *testing.T) {
	t.Run("CreationTimestamp 读写", func(t *testing.T) {
		ot := &metav1.OperationTimestamps{}
		ts := ot.GetCreationTimestamp()
		Then(
			t, "零值 GetCreationTimestamp 不为 nil",
			Expect(ts.IsZero(), Be(cmp.True())),
		)

		now := object.Timestamp(time.Unix(99999, 0))
		ot.SetCreationTimestamp(now)
		Then(
			t, "SetCreationTimestamp 设置成功",
			Expect(ot.GetCreationTimestamp(), Equal(now)),
		)
	})

	t.Run("ModificationTimestamp 读写", func(t *testing.T) {
		ot := &metav1.OperationTimestamps{}
		ts := ot.GetModificationTimestamp()
		Then(
			t, "零值 GetModificationTimestamp 不为 nil",
			Expect(ts.IsZero(), Be(cmp.True())),
		)

		now := object.Timestamp(time.Unix(88888, 0))
		ot.SetModificationTimestamp(now)
		Then(
			t, "SetModificationTimestamp 设置成功",
			Expect(ot.GetModificationTimestamp(), Equal(now)),
		)
	})
}

func TestMetadata(t *testing.T) {
	t.Run("Metadata 嵌入 TypeMeta + Describer + OperationTimestamps", func(t *testing.T) {
		m := &metav1.Metadata{}
		m.SetKind("MetaKind")
		m.SetName("meta-name")
		m.SetDescription("meta-desc")

		Then(
			t, "Kind 可访问",
			Expect(m.GetKind(), Equal("MetaKind")),
		)
		Then(
			t, "Name 可访问",
			Expect(m.GetName(), Equal("meta-name")),
		)
		Then(
			t, "Description 可访问",
			Expect(m.GetDescription(), Equal("meta-desc")),
		)
	})
}
