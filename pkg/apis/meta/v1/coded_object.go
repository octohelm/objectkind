package v1

import "github.com/octohelm/objectkind/pkg/schema"

type CodedObject[ID ~uint64, Code ~string] struct {
	Object[ID]
	// 编码
	Code Code `json:"code"`
}

var _ interface {
	schema.ObjectWithCode[string]
	schema.ObjectCodeSetter[string]
	schema.ObjectReceiver
} = &CodedObject[uint64, string]{}

func (o CodedObject[ID, Code]) GetCode() Code {
	return o.Code
}

func (o *CodedObject[ID, Code]) SetCode(code Code) {
	o.Code = code
}

func (v *CodedObject[ID, Code]) CopyFromObject(o schema.Object) {
	v.Object.CopyFromObject(o)

	if x, ok := o.(schema.ObjectWithCode[Code]); ok {
		v.SetCode(x.GetCode())
	}
}

type CodedObjectRequest[Code ~string] struct {
	ObjectRequest
	// 编码
	Code Code `json:"code"`
}

var _ interface {
	schema.ObjectWithCode[string]
	schema.ObjectCodeSetter[string]
	schema.ObjectReceiver
} = &CodedObjectRequest[string]{}

func (v CodedObjectRequest[Code]) GetCode() Code {
	return v.Code
}

func (v *CodedObjectRequest[Code]) SetCode(code Code) {
	v.Code = code
}

func (v *CodedObjectRequest[Code]) CopyFromObject(o schema.Object) {
	v.ObjectRequest.CopyFromObject(o)

	if x, ok := o.(schema.ObjectWithCode[Code]); ok {
		v.SetCode(x.GetCode())
	}
}

type CodeReference[Code ~string] struct {
	TypeMeta

	Code Code `json:"code"`
}

var _ interface {
	schema.ObjectWithCode[string]
	schema.ObjectCodeSetter[string]
	schema.ObjectReceiver
} = &CodeReference[string]{}

func (o CodeReference[Code]) GetCode() Code {
	return o.Code
}

func (o *CodeReference[Code]) SetCode(code Code) {
	o.Code = code
}

func (v *CodeReference[Code]) CopyFromObject(o schema.Object) {
	if x, ok := o.(schema.ObjectWithCode[Code]); ok {
		v.SetCode(x.GetCode())
	}
}
