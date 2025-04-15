/*
Package compose GENERATED BY gengo:runtimedoc
DON'T EDIT THIS FILE
*/
package compose

func (v *Annotatable) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "Annotations":
			return []string{
				"其他标注",
			}, true
		}

		return nil, false
	}
	return []string{}, true
}

func (*Annotations) RuntimeDoc(names ...string) ([]string, bool) {
	return []string{}, true
}

func (v *CodableResource[ID, Code]) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "Code":
			return []string{
				"编码",
				"人类可读编码",
			}, true
		}
		if doc, ok := runtimeDoc(&v.Resource, "", names...); ok {
			return doc, ok
		}

		return nil, false
	}
	return []string{}, true
}

func (v *CreationTimestamp) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "CreatedAt":
			return []string{
				"创建时间",
			}, true
		}

		return nil, false
	}
	return []string{}, true
}

func (v *DeletionTimestamp) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "DeletedAt":
			return []string{
				"删除时间",
			}, true
		}

		return nil, false
	}
	return []string{}, true
}

func (v *ModificationTimestamp) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "UpdatedAt":
			return []string{
				"更新时间",
			}, true
		}

		return nil, false
	}
	return []string{}, true
}

func (v *Rel[ID]) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "ID":
			return []string{
				"自增 id",
			}, true
		case "CreatedAt":
			return []string{
				"创建时间",
			}, true

		}

		return nil, false
	}
	return []string{}, true
}

func (v *Resource[ID]) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "ID":
			return []string{
				"id",
			}, true
		case "Name":
			return []string{
				"名称",
			}, true
		case "Description":
			return []string{
				"描述",
			}, true
		case "CreatedAt":
			return []string{
				"创建时间",
			}, true
		case "UpdatedAt":
			return []string{
				"更新时间",
			}, true

		}

		return nil, false
	}
	return []string{}, true
}

func (v *Revision[ID, Digest]) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "ID":
			return []string{
				"id",
			}, true
		case "Digest":
			return []string{
				"摘要",
			}, true
		case "CreatedAt":
			return []string{
				"创建时间",
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
