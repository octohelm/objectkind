/*
Package v1 GENERATED BY gengo:runtimedoc
DON'T EDIT THIS FILE
*/
package v1

func (*Currency) RuntimeDoc(names ...string) ([]string, bool) {
	return []string{
		"货币单位",
	}, true
}

func (*CurrencyValue) RuntimeDoc(names ...string) ([]string, bool) {
	return []string{}, true
}

// nolint:deadcode,unused
func runtimeDoc(v any, prefix string, names ...string) ([]string, bool) {
	if c, ok := v.(interface {
		RuntimeDoc(names ...string) ([]string, bool)
	}); ok {
		doc, ok := c.RuntimeDoc(names...)
		if ok {
			if prefix != "" && len(doc) > 0 {
				doc[0] = prefix + doc[0]
				return doc, true
			}

			return doc, true
		}
	}
	return nil, false
}
