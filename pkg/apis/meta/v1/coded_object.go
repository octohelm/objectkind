package v1

import "github.com/octohelm/objectkind/pkg/schema"

type CodedObject[ID ~uint64, Code ~string] struct {
	Object[ID]
	// 编码
	Code Code `json:"code"`
}

var _ interface {
	schema.CodeGetter[string]
	schema.AsRefCodeGetter
	schema.CodeSetter[string]
	schema.FromRefCodeSetter
	schema.ObjectReceiver
} = &CodedObject[uint64, string]{}

func (o CodedObject[ID, Code]) GetCode() Code {
	return o.Code
}

func (o CodedObject[ID, Code]) GetAsRefCode() schema.RefCode {
	return schema.RefCode(o.Code)
}

func (o *CodedObject[ID, Code]) SetCode(code Code) {
	o.Code = code
}

func (o *CodedObject[ID, Code]) SetFromRefCode(refCode schema.RefCode) {
	o.Code = Code(refCode)
}

func (v *CodedObject[ID, Code]) CopyFromObject(o schema.Object) {
	v.Object.CopyFromObject(o)

	if x, ok := o.(schema.CodeGetter[Code]); ok {
		v.SetCode(x.GetCode())
	}
}

type CodedObjectRequest[Code ~string] struct {
	ObjectRequest
	// 编码
	Code Code `json:"code"`
}

var _ interface {
	schema.CodeGetter[string]
	schema.AsRefCodeGetter
	schema.CodeSetter[string]
	schema.FromRefCodeSetter
	schema.ObjectReceiver
} = &CodedObjectRequest[string]{}

func (v CodedObjectRequest[Code]) GetCode() Code {
	return v.Code
}

func (o CodedObjectRequest[Code]) GetAsRefCode() schema.RefCode {
	return schema.RefCode(o.Code)
}

func (v *CodedObjectRequest[Code]) SetCode(code Code) {
	v.Code = code
}

func (o *CodedObjectRequest[Code]) SetFromRefCode(refCode schema.RefCode) {
	o.Code = Code(refCode)
}

func (v *CodedObjectRequest[Code]) CopyFromObject(o schema.Object) {
	v.ObjectRequest.CopyFromObject(o)

	if x, ok := o.(schema.CodeGetter[Code]); ok {
		v.SetCode(x.GetCode())
	}
}

type CodeReference[Code ~string] struct {
	TypeMeta

	Code Code `json:"code"`
}

var _ interface {
	schema.CodeGetter[string]
	schema.AsRefCodeGetter
	schema.CodeSetter[string]
	schema.ObjectReceiver
} = &CodeReference[string]{}

func (o CodeReference[Code]) GetCode() Code {
	return o.Code
}

func (o CodeReference[Code]) GetAsRefCode() schema.RefCode {
	return schema.RefCode(o.Code)
}

func (o *CodeReference[Code]) SetCode(code Code) {
	o.Code = code
}

func (v *CodeReference[Code]) CopyFromObject(o schema.Object) {
	if x, ok := o.(schema.CodeGetter[Code]); ok {
		v.SetCode(x.GetCode())
	}
}
