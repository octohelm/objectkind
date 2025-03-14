/*
Package v1 GENERATED BY gengo:runtimedoc
DON'T EDIT THIS FILE
*/
package v1

func (v *Codable[Code]) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "Code":
			return []string{
				"编码",
			}, true
		}

		return nil, false
	}
	return []string{}, true
}

func (v *CodableObject[ID, Code]) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		}
		if doc, ok := runtimeDoc(&v.Object, "", names...); ok {
			return doc, ok
		}
		if doc, ok := runtimeDoc(&v.Codable, "", names...); ok {
			return doc, ok
		}

		return nil, false
	}
	return []string{}, true
}

func (v *CodableReference[O, Code]) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		}
		if doc, ok := runtimeDoc(&v.TypeMeta, "", names...); ok {
			return doc, ok
		}
		if doc, ok := runtimeDoc(&v.Codable, "", names...); ok {
			return doc, ok
		}

		return nil, false
	}
	return []string{}, true
}

func (v *CodableRequest[O, Code]) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		}
		if doc, ok := runtimeDoc(&v.TypeMeta, "", names...); ok {
			return doc, ok
		}
		if doc, ok := runtimeDoc(&v.Describer, "", names...); ok {
			return doc, ok
		}
		if doc, ok := runtimeDoc(&v.Codable, "", names...); ok {
			return doc, ok
		}

		return nil, false
	}
	return []string{}, true
}

func (v *CodableResponse[O, ID, Code]) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		}
		if doc, ok := runtimeDoc(&v.Response, "", names...); ok {
			return doc, ok
		}
		if doc, ok := runtimeDoc(&v.Codable, "", names...); ok {
			return doc, ok
		}

		return nil, false
	}
	return []string{}, true
}

func (v *Describer) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "Name":
			return []string{
				"名称",
			}, true
		case "Description":
			return []string{
				"描述",
			}, true
		case "Annotations":
			return []string{
				"其他注解",
			}, true

		}

		return nil, false
	}
	return []string{}, true
}

func (v *GroupKind) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "Group":
			return []string{}, true
		case "Kind":
			return []string{}, true

		}

		return nil, false
	}
	return []string{}, true
}

func (v *GroupVersion) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "Group":
			return []string{}, true
		case "Version":
			return []string{}, true

		}

		return nil, false
	}
	return []string{}, true
}

func (v *GroupVersionKind) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "Group":
			return []string{}, true
		case "Version":
			return []string{}, true
		case "Kind":
			return []string{}, true

		}

		return nil, false
	}
	return []string{}, true
}

func (v *Identifiable[ID]) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "ID":
			return []string{
				"资源 id",
			}, true
		}

		return nil, false
	}
	return []string{}, true
}

func (v *List[T]) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "Items":
			return []string{
				"列表",
			}, true
		case "Total":
			return []string{
				"总数",
			}, true

		}

		return nil, false
	}
	return []string{}, true
}

func (v *Metadata) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		}
		if doc, ok := runtimeDoc(&v.TypeMeta, "", names...); ok {
			return doc, ok
		}
		if doc, ok := runtimeDoc(&v.Describer, "", names...); ok {
			return doc, ok
		}
		if doc, ok := runtimeDoc(&v.OperationTimestamps, "", names...); ok {
			return doc, ok
		}

		return nil, false
	}
	return []string{}, true
}

func (v *Object[ID]) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		}
		if doc, ok := runtimeDoc(&v.Metadata, "", names...); ok {
			return doc, ok
		}
		if doc, ok := runtimeDoc(&v.Identifiable, "", names...); ok {
			return doc, ok
		}

		return nil, false
	}
	return []string{}, true
}

func (v *ObjectReference[O, ID]) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		}
		if doc, ok := runtimeDoc(&v.TypeMeta, "", names...); ok {
			return doc, ok
		}
		if doc, ok := runtimeDoc(&v.Identifiable, "", names...); ok {
			return doc, ok
		}

		return nil, false
	}
	return []string{}, true
}

func (v *OperationTimestamps) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "CreationTimestamp":
			return []string{
				"创建时间",
			}, true
		case "ModificationTimestamp":
			return []string{
				"更新时间",
			}, true

		}

		return nil, false
	}
	return []string{}, true
}

func (v *Request[O]) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		}
		if doc, ok := runtimeDoc(&v.TypeMeta, "", names...); ok {
			return doc, ok
		}
		if doc, ok := runtimeDoc(&v.Describer, "", names...); ok {
			return doc, ok
		}

		return nil, false
	}
	return []string{}, true
}

func (v *Response[O, ID]) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		}
		if doc, ok := runtimeDoc(&v.Metadata, "", names...); ok {
			return doc, ok
		}
		if doc, ok := runtimeDoc(&v.Identifiable, "", names...); ok {
			return doc, ok
		}

		return nil, false
	}
	return []string{}, true
}

func (v *TypeMeta) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "Kind":
			return []string{
				"资源类型",
			}, true
		case "APIVersion":
			return []string{
				"资源类型版本",
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
