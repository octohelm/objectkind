package object_test

import (
	"testing"

	"github.com/octohelm/x/cmp"
	. "github.com/octohelm/x/testing/v2"

	"github.com/octohelm/objectkind/pkg/object"
)

func TestRefID(t *testing.T) {
	t.Run("object.RefID(0) 为零值", func(t *testing.T) {
		var id object.RefID
		Then(t, "零值为 0", Expect(uint64(id) == 0, Be(cmp.True())))
	})

	t.Run("object.RefID(100) 相等", func(t *testing.T) {
		id := object.RefID(100)
		Then(t, "值为 100", Expect(uint64(id), Equal(uint64(100))))
	})

	t.Run("object.RefID 可从 uint64 赋值", func(t *testing.T) {
		var id object.RefID = 42
		Then(t, "赋值为 42", Expect(uint64(id), Equal(uint64(42))))
	})
}

func TestRefCode(t *testing.T) {
	t.Run("object.RefCode(\"\") 为零值", func(t *testing.T) {
		var code object.RefCode
		Then(t, "零值为空字符串", Expect(string(code), Equal("")))
	})

	t.Run("object.RefCode(\"my-code\") 相等", func(t *testing.T) {
		code := object.RefCode("my-code")
		Then(t, "值为 my-code", Expect(string(code), Equal("my-code")))
	})
}

type MyID uint64

func (m MyID) String() string { return "" }

type MyCode string

type myIDGetter struct {
	ID MyID
}

func (m myIDGetter) GetID() MyID { return m.ID }

type myCodeGetter struct {
	Code MyCode
}

func (m myCodeGetter) GetCode() MyCode { return m.Code }

func identityType[ID object.Identity](id ID) ID { return id }

func TestIdentity(t *testing.T) {
	t.Run("~uint64 覆盖 uint64 和 MyID", func(t *testing.T) {
		_ = identityType(uint64(1))
		_ = identityType(MyID(2))
	})

	t.Run("~string 覆盖 string 和 MyCode", func(t *testing.T) {
		_ = identityType("x")
		_ = identityType(MyCode("y"))
	})

	t.Run("MyID 满足 IDGetter[MyID]", func(t *testing.T) {
		var g object.IDGetter[MyID] = myIDGetter{ID: 10}
		Then(t, "GetID 返回 10", Expect(g.GetID(), Equal(MyID(10))))
	})

	t.Run("MyCode 满足 CodeGetter[MyCode]", func(t *testing.T) {
		var g object.CodeGetter[MyCode] = myCodeGetter{Code: "test-code"}
		Then(t, "GetCode 返回 test-code", Expect(g.GetCode(), Equal(MyCode("test-code"))))
	})
}

type testObj struct {
	kind        string
	name        string
	description string
	annotations map[string]string
	createdAt   object.Timestamp
	updatedAt   object.Timestamp
}

func (t testObj) GetKind() string                              { return t.kind }
func (t testObj) GetName() string                              { return t.name }
func (t testObj) GetDescription() string                       { return t.description }
func (t testObj) GetAnnotations() map[string]string            { return t.annotations }
func (t testObj) GetAnnotation(k string) (string, bool)        { v, ok := t.annotations[k]; return v, ok }
func (t testObj) GetCreationTimestamp() object.Timestamp       { return t.createdAt }
func (t testObj) SetCreationTimestamp(ts object.Timestamp)     { t.createdAt = ts }
func (t testObj) GetModificationTimestamp() object.Timestamp   { return t.updatedAt }
func (t testObj) SetModificationTimestamp(ts object.Timestamp) { t.updatedAt = ts }

func TestInterfaces(t *testing.T) {
	t.Run("实现 object.Type", func(t *testing.T) {
		var _ object.Type = testObj{kind: "TestKind"}
	})

	t.Run("实现 object.Describer", func(t *testing.T) {
		var _ object.Describer = testObj{name: "test", description: "desc"}
	})

	t.Run("实现 object.Annotater", func(t *testing.T) {
		var _ object.Annotater = testObj{annotations: map[string]string{"k": "v"}}
	})

	t.Run("实现 object.OperationTimestamps", func(t *testing.T) {
		var _ object.OperationTimestamps = testObj{}
	})
}
