/*
Package filter GENERATED BY gengo:runtimedoc
DON'T EDIT THIS FILE
*/
package filter

func (v *ProductByID) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "ID":
			return []string{
				"通过 id 筛选",
			}, true
		}

		return nil, false
	}
	return []string{}, true
}

func (v *ProductByState) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "State":
			return []string{}, true
		}

		return nil, false
	}
	return []string{}, true
}

func (v *SkuByCode) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "Code":
			return []string{
				"通过 编码 筛选",
			}, true
		}

		return nil, false
	}
	return []string{}, true
}

func (v *SkuByID) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "ID":
			return []string{
				"通过 id 筛选",
			}, true
		}

		return nil, false
	}
	return []string{}, true
}

func (v *SkuByProductID) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "ProductID":
			return []string{
				"通过 所属产品 筛选",
			}, true
		}

		return nil, false
	}
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
