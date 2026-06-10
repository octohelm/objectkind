package compose_test

import (
	"database/sql/driver"
	"testing"
	"time"

	"github.com/go-json-experiment/json"

	"github.com/octohelm/x/cmp"
	. "github.com/octohelm/x/testing/v2"

	"github.com/octohelm/objectkind/pkg/object"
	compose "github.com/octohelm/objectkind/pkg/sqltype/compose"
)

func TestResource(t *testing.T) {
	t.Run("初始化 Resource[uint64] 为零值", func(t *testing.T) {
		r := &compose.Resource[uint64]{}

		Then(
			t, "ID 应为 0",
			Expect(r.GetID(), Equal(uint64(0))),
		)
		Then(
			t, "Name 应为空",
			Expect(r.GetName(), Equal("")),
		)
		Then(
			t, "Description 应为空",
			Expect(r.GetDescription(), Equal("")),
		)
		Then(
			t, "CreatedAt 为零值",
			Expect(r.GetCreationTimestamp().IsZero(), Be(cmp.True())),
		)
		Then(
			t, "UpdatedAt 为零值",
			Expect(r.GetModificationTimestamp().IsZero(), Be(cmp.True())),
		)
	})

	t.Run("MarkCreatedAt 仅在 CreatedAt 为零值时设置", func(t *testing.T) {
		r := &compose.Resource[uint64]{}

		Then(
			t, "首次 MarkCreatedAt 前 CreatedAt 为零值",
			Expect(r.GetCreationTimestamp().IsZero(), Be(cmp.True())),
		)

		r.MarkCreatedAt()

		Then(
			t, "首次 MarkCreatedAt 后 CreatedAt 不再为零值",
			Expect(r.GetCreationTimestamp().IsZero(), Be(cmp.False())),
		)

		firstCreatedAt := r.GetCreationTimestamp()
		time.Sleep(1 * time.Millisecond)
		r.MarkCreatedAt()

		Then(
			t, "再次 MarkCreatedAt 后 CreatedAt 不变",
			Expect(r.GetCreationTimestamp(), Equal(firstCreatedAt)),
		)
	})

	t.Run("MarkModifiedAt 设置 UpdatedAt 并在 CreatedAt 为零值时同步", func(t *testing.T) {
		t.Run("CreatedAt 为零值时同时设置两个时间戳", func(t *testing.T) {
			r := &compose.Resource[uint64]{}

			r.MarkModifiedAt()

			Then(
				t, "UpdatedAt 不再为零值",
				Expect(r.GetModificationTimestamp().IsZero(), Be(cmp.False())),
			)
			Then(
				t, "CreatedAt 不再为零值",
				Expect(r.GetCreationTimestamp().IsZero(), Be(cmp.False())),
			)
			Then(
				t, "CreatedAt 等于 UpdatedAt",
				Expect(r.GetCreationTimestamp(), Equal(r.GetModificationTimestamp())),
			)
		})

		t.Run("CreatedAt 已设置时只更新 UpdatedAt", func(t *testing.T) {
			r := &compose.Resource[uint64]{}
			r.MarkCreatedAt()
			createdAt := r.GetCreationTimestamp()

			time.Sleep(1 * time.Millisecond)
			r.MarkModifiedAt()

			Then(
				t, "CreatedAt 保持不变",
				Expect(r.GetCreationTimestamp(), Equal(createdAt)),
			)
			Then(
				t, "UpdatedAt 不再为零值",
				Expect(r.GetModificationTimestamp().IsZero(), Be(cmp.False())),
			)
		})

		t.Run("UpdatedAt 非零值时不再更新", func(t *testing.T) {
			r := &compose.Resource[uint64]{}
			r.MarkModifiedAt()
			updatedAt := r.GetModificationTimestamp()

			time.Sleep(1 * time.Millisecond)
			r.MarkModifiedAt()

			Then(
				t, "UpdatedAt 保持不变",
				Expect(r.GetModificationTimestamp(), Equal(updatedAt)),
			)
		})
	})

	t.Run("ForceMarkModifiedAt 始终设置两个时间戳", func(t *testing.T) {
		r := &compose.Resource[uint64]{}

		r.ForceMarkModifiedAt()

		Then(
			t, "UpdatedAt 不再为零值",
			Expect(r.GetModificationTimestamp().IsZero(), Be(cmp.False())),
		)
		Then(
			t, "CreatedAt 不再为零值",
			Expect(r.GetCreationTimestamp().IsZero(), Be(cmp.False())),
		)
		Then(
			t, "CreatedAt 等于 UpdatedAt",
			Expect(r.GetCreationTimestamp(), Equal(r.GetModificationTimestamp())),
		)

		prev := r.GetModificationTimestamp()
		time.Sleep(1 * time.Millisecond)
		r.ForceMarkModifiedAt()

		Then(
			t, "再次调用后 UpdatedAt 已更新",
			Expect(r.GetModificationTimestamp() != prev, Be(cmp.True())),
		)
		Then(
			t, "CreatedAt 也被更新为新的时间",
			Expect(r.GetCreationTimestamp(), Equal(r.GetModificationTimestamp())),
		)
	})

	t.Run("GetName / SetName 回转", func(t *testing.T) {
		r := &compose.Resource[uint64]{}
		r.SetName("test-name")
		Then(
			t, "Name 应等于设置值",
			Expect(r.GetName(), Equal("test-name")),
		)
	})

	t.Run("GetDescription / SetDescription 回转", func(t *testing.T) {
		r := &compose.Resource[uint64]{}
		r.SetDescription("test-desc")
		Then(
			t, "Description 应等于设置值",
			Expect(r.GetDescription(), Equal("test-desc")),
		)
	})

	t.Run("GetID / SetID 回转", func(t *testing.T) {
		r := &compose.Resource[uint64]{}
		r.SetID(42)
		Then(
			t, "ID 应等于设置值",
			Expect(r.GetID(), Equal(uint64(42))),
		)
	})

	t.Run("GetAsRefID 返回 RefID", func(t *testing.T) {
		r := &compose.Resource[uint64]{}
		r.SetID(100)
		Then(
			t, "GetAsRefID 应返回对应 RefID",
			Expect(r.GetAsRefID(), Equal(object.RefID(100))),
		)
	})

	t.Run("时间戳 getter / setter", func(t *testing.T) {
		r := &compose.Resource[uint64]{}
		ts := object.Timestamp(time.Now())

		r.SetCreationTimestamp(ts)
		Then(
			t, "SetCreationTimestamp 后 GetCreationTimestamp 应返回设置值",
			Expect(r.GetCreationTimestamp(), Equal(ts)),
		)

		r.SetModificationTimestamp(ts)
		Then(
			t, "SetModificationTimestamp 后 GetModificationTimestamp 应返回设置值",
			Expect(r.GetModificationTimestamp(), Equal(ts)),
		)
	})
}

func TestCodableResource(t *testing.T) {
	t.Run("GetCode / SetCode 回转", func(t *testing.T) {
		r := &compose.CodableResource[uint64, string]{}
		r.SetCode("my-code")
		Then(
			t, "Code 应等于设置值",
			Expect(r.GetCode(), Equal("my-code")),
		)
	})

	t.Run("GetAsRefCode 返回 RefCode", func(t *testing.T) {
		r := &compose.CodableResource[uint64, string]{}
		r.SetCode("my-code")
		Then(
			t, "GetAsRefCode 应返回对应 RefCode",
			Expect(r.GetAsRefCode(), Equal(object.RefCode("my-code"))),
		)
	})

	t.Run("内嵌 Resource 的方法仍可用", func(t *testing.T) {
		r := &compose.CodableResource[uint64, string]{}
		r.SetName("embedded-name")
		Then(
			t, "内嵌 Resource 的 GetName 仍可用",
			Expect(r.GetName(), Equal("embedded-name")),
		)
		r.SetID(7)
		Then(
			t, "内嵌 Resource 的 GetID 仍可用",
			Expect(r.GetID(), Equal(uint64(7))),
		)
	})
}

func TestRel(t *testing.T) {
	t.Run("GetID 返回 ID", func(t *testing.T) {
		r := &compose.Rel[uint64]{ID: 99}
		Then(
			t, "GetID 应返回 99",
			Expect(r.GetID(), Equal(uint64(99))),
		)
	})

	t.Run("MarkModifiedAt 在 CreatedAt 为零值时设置", func(t *testing.T) {
		r := &compose.Rel[uint64]{}

		Then(
			t, "MarkModifiedAt 前 CreatedAt 为零值",
			Expect(r.GetCreationTimestamp().IsZero(), Be(cmp.True())),
		)

		r.MarkModifiedAt()

		Then(
			t, "MarkModifiedAt 后 CreatedAt 不再为零值",
			Expect(r.GetCreationTimestamp().IsZero(), Be(cmp.False())),
		)
	})

	t.Run("MarkModifiedAt 在 CreatedAt 非零值时保持不变", func(t *testing.T) {
		r := &compose.Rel[uint64]{}
		r.MarkModifiedAt()
		createdAt := r.GetCreationTimestamp()

		time.Sleep(1 * time.Millisecond)
		r.MarkModifiedAt()

		Then(
			t, "CreatedAt 保持不变",
			Expect(r.GetCreationTimestamp(), Equal(createdAt)),
		)
	})

	t.Run("时间戳 getter / setter", func(t *testing.T) {
		r := &compose.Rel[uint64]{}
		ts := object.Timestamp(time.Now())

		r.SetCreationTimestamp(ts)
		Then(
			t, "SetCreationTimestamp 触发 MarkModifiedAt 设置时间",
			Expect(r.GetCreationTimestamp().IsZero(), Be(cmp.False())),
		)

		thenTS := r.GetCreationTimestamp()

		Then(
			t, "GetModificationTimestamp 返回 CreatedAt",
			Expect(r.GetModificationTimestamp(), Equal(thenTS)),
		)

		r.SetModificationTimestamp(ts)
		Then(
			t, "SetModificationTimestamp 也触发 MarkModifiedAt",
			Expect(r.GetCreationTimestamp().IsZero(), Be(cmp.False())),
		)
	})
}

func TestRevision(t *testing.T) {
	t.Run("GetID / SetID 回转", func(t *testing.T) {
		r := &compose.Revision[uint64, string]{}
		r.SetID(55)
		Then(
			t, "GetID 应返回 55",
			Expect(r.GetID(), Equal(uint64(55))),
		)
	})

	t.Run("MarkModifiedAt 在 CreatedAt 为零值时设置", func(t *testing.T) {
		r := &compose.Revision[uint64, string]{}

		Then(
			t, "MarkModifiedAt 前 CreatedAt 为零值",
			Expect(r.GetCreationTimestamp().IsZero(), Be(cmp.True())),
		)

		r.MarkModifiedAt()

		Then(
			t, "MarkModifiedAt 后 CreatedAt 不再为零值",
			Expect(r.GetCreationTimestamp().IsZero(), Be(cmp.False())),
		)
	})

	t.Run("时间戳 getter / setter", func(t *testing.T) {
		r := &compose.Revision[uint64, string]{}
		ts := object.Timestamp(time.Now())

		r.SetCreationTimestamp(ts)
		Then(
			t, "SetCreationTimestamp 触发 MarkModifiedAt",
			Expect(r.GetCreationTimestamp().IsZero(), Be(cmp.False())),
		)

		thenTS := r.GetCreationTimestamp()

		Then(
			t, "GetModificationTimestamp 返回 CreatedAt",
			Expect(r.GetModificationTimestamp(), Equal(thenTS)),
		)

		r.SetModificationTimestamp(ts)
		Then(
			t, "SetModificationTimestamp 也触发 MarkModifiedAt",
			Expect(r.GetCreationTimestamp().IsZero(), Be(cmp.False())),
		)
	})
}

func TestCreationTimestamp(t *testing.T) {
	t.Run("MarkCreatedAt 在 CreatedAt 为零值时设置", func(t *testing.T) {
		ct := &compose.CreationTimestamp{}

		Then(
			t, "MarkCreatedAt 前 CreatedAt 为零值",
			Expect(ct.CreatedAt.IsZero(), Be(cmp.True())),
		)

		ct.MarkCreatedAt()

		Then(
			t, "MarkCreatedAt 后 CreatedAt 不再为零值",
			Expect(ct.CreatedAt.IsZero(), Be(cmp.False())),
		)
	})

	t.Run("MarkCreatedAt 在 CreatedAt 非零值时不变", func(t *testing.T) {
		ct := &compose.CreationTimestamp{}
		ct.MarkCreatedAt()
		createdAt := ct.CreatedAt

		time.Sleep(1 * time.Millisecond)
		ct.MarkCreatedAt()

		Then(
			t, "CreatedAt 保持不变",
			Expect(ct.CreatedAt, Equal(createdAt)),
		)
	})
}

func TestModificationTimestamp(t *testing.T) {
	t.Run("MarkModifiedAt 在 UpdatedAt 为零值时设置", func(t *testing.T) {
		mt := &compose.ModificationTimestamp{}

		Then(
			t, "MarkModifiedAt 前 UpdatedAt 为零值",
			Expect(mt.UpdatedAt.IsZero(), Be(cmp.True())),
		)

		mt.MarkModifiedAt()

		Then(
			t, "MarkModifiedAt 后 UpdatedAt 不再为零值",
			Expect(mt.UpdatedAt.IsZero(), Be(cmp.False())),
		)
	})

	t.Run("MarkModifiedAt 在 UpdatedAt 非零值时不变", func(t *testing.T) {
		mt := &compose.ModificationTimestamp{}
		mt.MarkModifiedAt()
		updatedAt := mt.UpdatedAt

		time.Sleep(1 * time.Millisecond)
		mt.MarkModifiedAt()

		Then(
			t, "UpdatedAt 保持不变",
			Expect(mt.UpdatedAt, Equal(updatedAt)),
		)
	})
}

func TestDeletionTimestamp(t *testing.T) {
	t.Run("SoftDeleteFieldAndZeroValue 返回字段名和零值", func(t *testing.T) {
		dt := compose.DeletionTimestamp{}
		field, zeroVal := dt.SoftDeleteFieldAndZeroValue()

		Then(
			t, "字段名应为 DeletedAt",
			Expect(field, Equal("DeletedAt")),
		)
		Then(t, "零值应为 int64(0)",
			Expect(zeroVal, Equal(driver.Value(int64(0)))))
	})

	t.Run("MarkDeletedAt 设置 DeletedAt", func(t *testing.T) {
		dt := &compose.DeletionTimestamp{}

		Then(
			t, "MarkDeletedAt 前 DeletedAt 为零值",
			Expect(dt.DeletedAt.IsZero(), Be(cmp.True())),
		)

		dt.MarkDeletedAt()

		Then(
			t, "MarkDeletedAt 后 DeletedAt 不再为零值",
			Expect(dt.DeletedAt.IsZero(), Be(cmp.False())),
		)
	})
}

func TestAnnotations(t *testing.T) {
	t.Run("Value 序列化", func(t *testing.T) {
		t.Run("非空 map 返回 JSON 字符串", func(t *testing.T) {
			annos := compose.Annotations{"key": "value"}
			val, err := annos.Value()

			Must(t, func() error { return err })

			expected := MustValue(t, func() ([]byte, error) {
				return json.Marshal(compose.Annotations{"key": "value"})
			})

			Then(
				t, "Value 返回序列化后的 JSON",
				Expect(val.(string), Equal(string(expected))),
			)
		})

		t.Run("空 map 返回空字符串", func(t *testing.T) {
			annos := compose.Annotations{}
			val, err := annos.Value()

			Must(t, func() error { return err })

			Then(
				t, "空 map 应返回空字符串",
				Expect(val, Equal(driver.Value(""))),
			)
		})

		t.Run("nil map 返回空字符串", func(t *testing.T) {
			var annos compose.Annotations
			val, err := annos.Value()

			Must(t, func() error { return err })

			Then(
				t, "nil map 应返回空字符串",
				Expect(val, Equal(driver.Value(""))),
			)
		})
	})

	t.Run("Scan 反序列化", func(t *testing.T) {
		t.Run("从 []byte 扫描有效 JSON", func(t *testing.T) {
			var annos compose.Annotations
			jsonBytes := MustValue(t, func() ([]byte, error) {
				return json.Marshal(compose.Annotations{"k": "v"})
			})

			Must(t, func() error {
				return annos.Scan(jsonBytes)
			})

			Then(
				t, "反序列化后应包含键值对",
				Expect(annos["k"], Equal("v")),
			)
		})

		t.Run("从空 []byte 扫描", func(t *testing.T) {
			annos := compose.Annotations{"pre": "existing"}

			Must(t, func() error {
				return annos.Scan([]byte{})
			})

			Then(
				t, "空 bytes 不修改已有值",
				Expect(annos["pre"], Equal("existing")),
			)
		})

		t.Run("从 string 扫描有效 JSON", func(t *testing.T) {
			var annos compose.Annotations

			Must(t, func() error {
				return annos.Scan(`{"x":"y"}`)
			})

			Then(
				t, "反序列化后应包含键值对",
				Expect(annos["x"], Equal("y")),
			)
		})

		t.Run("从空 string 扫描", func(t *testing.T) {
			annos := compose.Annotations{"pre": "existing"}

			Must(t, func() error {
				return annos.Scan("")
			})

			Then(
				t, "空 string 不修改已有值",
				Expect(annos["pre"], Equal("existing")),
			)
		})

		t.Run("从 nil 扫描", func(t *testing.T) {
			annos := compose.Annotations{"pre": "existing"}

			Must(t, func() error {
				return annos.Scan(nil)
			})

			Then(
				t, "nil 不修改已有值",
				Expect(annos["pre"], Equal("existing")),
			)
		})

		t.Run("无效类型返回错误", func(t *testing.T) {
			var annos compose.Annotations

			err := annos.Scan(123)

			Then(
				t, "应返回错误",
				Expect(err != nil, Be(cmp.True())),
			)
		})
	})
}

func TestAnnotatable(t *testing.T) {
	t.Run("Annotations 为空时 GetAnnotations 返回 nil map", func(t *testing.T) {
		a := compose.Annotatable{}
		annos := a.GetAnnotations()

		Then(
			t, "Annotations 应为 nil",
			Expect(annos == nil, Be(cmp.True())),
		)
	})

	t.Run("SetAnnotations 设置并读取 annotations", func(t *testing.T) {
		a := &compose.Annotatable{}
		input := map[string]string{"env": "prod", "region": "us-east"}

		a.SetAnnotations(input)
		result := a.GetAnnotations()

		Then(
			t, "annotation 值匹配",
			Expect(result["env"], Equal("prod")),
		)
		Then(
			t, "annotation 值匹配",
			Expect(result["region"], Equal("us-east")),
		)
		Then(
			t, "annotation 长度匹配",
			Expect(len(result), Equal(2)),
		)
	})

	t.Run("SetAnnotations 用新 map 覆盖旧值", func(t *testing.T) {
		a := &compose.Annotatable{}
		a.SetAnnotations(map[string]string{"old": "value"})
		a.SetAnnotations(map[string]string{"new": "data"})

		result := a.GetAnnotations()

		Then(
			t, "旧键不存在",
			Expect(result["old"], Equal("")),
		)
		Then(
			t, "新键存在",
			Expect(result["new"], Equal("data")),
		)
	})

	t.Run("GetAnnotation 按 key 取值", func(t *testing.T) {
		a := &compose.Annotatable{}
		a.SetAnnotations(map[string]string{"k1": "v1"})

		v, ok := a.GetAnnotation("k1")
		Then(
			t, "应找到 k1",
			Expect(ok, Be(cmp.True())),
		)
		Then(
			t, "k1 值应为 v1",
			Expect(v, Equal("v1")),
		)

		v2, ok2 := a.GetAnnotation("missing")
		Then(
			t, "不存在的 key 不应找到",
			Expect(ok2, Be(cmp.False())),
		)
		Then(
			t, "不存在的 key 返回空字符串",
			Expect(v2, Equal("")),
		)
	})

	t.Run("GetAnnotation 在 nil annotations 时返回 false", func(t *testing.T) {
		a := compose.Annotatable{}
		v, ok := a.GetAnnotation("any")
		Then(
			t, "nil annotations 不应找到 key",
			Expect(ok, Be(cmp.False())),
		)
		Then(
			t, "返回空字符串",
			Expect(v, Equal("")),
		)
	})

	t.Run("SetAnnotation 增量设置单个键值", func(t *testing.T) {
		a := &compose.Annotatable{}

		a.SetAnnotation("step1", "one")
		v1, ok1 := a.GetAnnotation("step1")
		Then(
			t, "应找到 step1",
			Expect(ok1, Be(cmp.True())),
		)
		Then(
			t, "step1 值为 one",
			Expect(v1, Equal("one")),
		)

		a.SetAnnotation("step2", "two")
		v2, ok2 := a.GetAnnotation("step2")
		Then(
			t, "应找到 step2",
			Expect(ok2, Be(cmp.True())),
		)
		Then(
			t, "step2 值为 two",
			Expect(v2, Equal("two")),
		)
		Then(
			t, "step1 仍存在",
			Expect(a.GetAnnotations()["step1"], Equal("one")),
		)
	})

	t.Run("SetAnnotation 在 nil annotations 上初始化", func(t *testing.T) {
		a := &compose.Annotatable{}
		a.SetAnnotation("newkey", "newval")

		annos := a.GetAnnotations()
		Then(
			t, "Annotations 不再是 nil",
			Expect(annos != nil, Be(cmp.True())),
		)
		Then(
			t, "值应正确",
			Expect(annos["newkey"], Equal("newval")),
		)
	})

	t.Run("Annotatable 无声明式满足 object.Annotater 接口", func(t *testing.T) {
		_ = object.Annotater(compose.Annotatable{})
	})

	t.Run("*Annotatable 满足 object.Annotatable 接口", func(t *testing.T) {
		_ = object.Annotatable(&compose.Annotatable{})
	})
}

func TestAnnotationsDataType(t *testing.T) {
	a := compose.Annotations{}
	dt := a.DataType("mysql")
	Then(
		t, "DataType 返回 text",
		Expect(dt, Equal("text")),
	)
}
