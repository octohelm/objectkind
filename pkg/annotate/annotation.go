package annotate

import (
	"bytes"
	"strconv"

	"github.com/go-json-experiment/json"
	"github.com/go-json-experiment/json/jsontext"
	jsonv1 "github.com/go-json-experiment/json/v1"
	"github.com/octohelm/courier/pkg/validator"
)

type Annotation string

func (ann Annotation) Get(accessor Provider) (string, bool) {
	return accessor.GetAnnotation(string(ann))
}

func (ann Annotation) MarshalTo(accessor Accessor, v any) error {
	b := &bytes.Buffer{}

	switch x := v.(type) {
	default:
		if err := json.MarshalWrite(b, x, jsonv1.OmitEmptyWithLegacyDefinition(true)); err != nil {
			return err
		}
	}

	if b.Len() > 0 {
		str := string(b.Bytes())
		if str[0] == '"' {
			s, err := strconv.Unquote(str)
			if err != nil {
				return err
			}
			str = s
		}
		accessor.SetAnnotation(string(ann), str)
	}
	return nil
}

func (ann Annotation) UnmarshalFrom(provider Provider, v any) error {
	if str, ok := provider.GetAnnotation(string(ann)); ok {
		dec := jsontext.NewDecoder(bytes.NewBufferString(str))
		t := dec.PeekKind()
		switch t {
		case '{', '[':
		default:
			shouldQuote := true
			switch t {
			case 'n':
				if str == "null" {
					shouldQuote = false
				}
			case 't':
				if str == "true" {
					shouldQuote = false
				}
			case 'f':
				if str == "false" {
					shouldQuote = false
				}
			}

			if shouldQuote {
				// force quote
				dec = jsontext.NewDecoder(bytes.NewBufferString(strconv.Quote(str)))
			}
		}
		return validator.UnmarshalDecode(dec, v)
	}
	return nil
}
