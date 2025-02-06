package v1

import "github.com/octohelm/objectkind/pkg/object"

type Codable[Code ~string] struct {
	// 编码
	Code Code `json:"code"`
}

var _ interface {
	object.CodeGetter[string]
	object.CodeSetter[string]
	object.RefCodeConvertable
} = &Codable[string]{}

func (o Codable[Code]) GetCode() Code      { return o.Code }
func (o *Codable[Code]) SetCode(code Code) { o.Code = code }

func (o Codable[Code]) GetAsRefCode() object.RefCode {
	return object.RefCode(o.Code)
}
func (o *Codable[Code]) SetFromRefCode(refCode object.RefCode) { o.Code = Code(refCode) }
