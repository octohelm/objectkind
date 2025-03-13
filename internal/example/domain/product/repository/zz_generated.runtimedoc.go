/*
Package repository GENERATED BY gengo:runtimedoc
DON'T EDIT THIS FILE
*/
package repository

func (v *ProductQuerier) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "Product":
			return []string{}, true
		}

		return nil, false
	}
	return []string{}, true
}

func (v *ProductRepository) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		}
		if doc, ok := runtimeDoc(&v.ProductQuerier, "", names...); ok {
			return doc, ok
		}

		return nil, false
	}
	return []string{}, true
}

func (v *SkuQuerier) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "Sku":
			return []string{}, true
		}

		return nil, false
	}
	return []string{}, true
}

func (v *SkuRepository) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		}
		if doc, ok := runtimeDoc(&v.SkuQuerier, "", names...); ok {
			return doc, ok
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
