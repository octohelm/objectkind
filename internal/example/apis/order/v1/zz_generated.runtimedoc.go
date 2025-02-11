/*
Package v1 GENERATED BY gengo:runtimedoc
DON'T EDIT THIS FILE
*/
package v1

import _ "embed"

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

func (v *Order) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "Spec":
			return []string{}, true
		case "Status":
			return []string{}, true

		}
		if doc, ok := runtimeDoc(&v.Object, "", names...); ok {
			return doc, ok
		}

		return nil, false
	}
	return []string{}, true
}

func (*OrderID) RuntimeDoc(names ...string) ([]string, bool) {
	return []string{}, true
}

func (v *OrderItem) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "Spec":
			return []string{}, true
		case "Status":
			return []string{}, true

		}
		if doc, ok := runtimeDoc(&v.Object, "", names...); ok {
			return doc, ok
		}

		return nil, false
	}
	return []string{
		"订单项",
	}, true
}

func (*OrderItemID) RuntimeDoc(names ...string) ([]string, bool) {
	return []string{}, true
}

func (v *OrderItemRequestForCreate) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "Spec":
			return []string{}, true
		}
		if doc, ok := runtimeDoc(&v.Request, "订单项", names...); ok {
			return doc, ok
		}

		return nil, false
	}
	return []string{
		"订单项 (更新)",
	}, true
}

func (v *OrderItemSpec) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "Sku":
			return []string{
				"商品规格",
			}, true
		case "Quantity":
			return []string{
				"个数",
			}, true

		}

		return nil, false
	}
	return []string{}, true
}

func (v *OrderItemSpecRequestForCreate) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "Sku":
			return []string{
				"商品规格",
			}, true
		case "Quantity":
			return []string{
				"个数",
			}, true

		}

		return nil, false
	}
	return []string{}, true
}

func (v *OrderItemStatus) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "TotalPrice":
			return []string{
				"总价",
			}, true
		case "DiscountAmount":
			return []string{
				"折扣金额",
			}, true
		case "FinalPrice":
			return []string{
				"最终价格",
			}, true

		}

		return nil, false
	}
	return []string{}, true
}

func (v *OrderRequestForCreate) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "Spec":
			return []string{}, true
		}
		if doc, ok := runtimeDoc(&v.Request, "订单", names...); ok {
			return doc, ok
		}

		return nil, false
	}
	return []string{
		"订单 (创建)",
	}, true
}

func (v *OrderSpec) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "Items":
			return []string{
				"订单项",
			}, true
		}

		return nil, false
	}
	return []string{}, true
}

func (v *OrderSpecRequestForCreate) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "Items":
			return []string{
				"订单项",
			}, true
		}

		return nil, false
	}
	return []string{}, true
}

func (v *OrderStatus) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "TotalAmount":
			return []string{
				"总金额",
			}, true
		}

		return nil, false
	}
	return []string{}, true
}

func (v *Product) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "Status":
			return []string{}, true
		}
		if doc, ok := runtimeDoc(&v.Object, "商品", names...); ok {
			return doc, ok
		}

		return nil, false
	}
	return []string{
		"商品",
	}, true
}

func (*ProductID) RuntimeDoc(names ...string) ([]string, bool) {
	return []string{
		"商品 id",
	}, true
}

func (v *ProductReference) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		}
		if doc, ok := runtimeDoc(&v.ObjectReference, "商品", names...); ok {
			return doc, ok
		}

		return nil, false
	}
	return []string{
		"商品 (引用)",
	}, true
}

func (v *ProductRequestForCreate) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		}
		if doc, ok := runtimeDoc(&v.Request, "商品", names...); ok {
			return doc, ok
		}

		return nil, false
	}
	return []string{
		"商品 (创建)",
	}, true
}

func (v *ProductRequestForUpdate) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		}
		if doc, ok := runtimeDoc(&v.Request, "商品", names...); ok {
			return doc, ok
		}

		return nil, false
	}
	return []string{
		"商品 (更新)",
	}, true
}

func (v *ProductStatus) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "Available":
			return []string{
				"是否可用",
			}, true
		}

		return nil, false
	}
	return []string{
		"商品状态",
	}, true
}

func (v *Sku) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "Spec":
			return []string{
				"商品规格属性",
			}, true
		case "Product":
			return []string{
				"所属商品",
			}, true

		}
		if doc, ok := runtimeDoc(&v.CodableObject, "商品规格", names...); ok {
			return doc, ok
		}

		return nil, false
	}
	return []string{
		"商品规格",
	}, true
}

func (*SkuCode) RuntimeDoc(names ...string) ([]string, bool) {
	return []string{}, true
}

func (*SkuID) RuntimeDoc(names ...string) ([]string, bool) {
	return []string{}, true
}

func (v *SkuReference) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		}
		if doc, ok := runtimeDoc(&v.ObjectReference, "商品规格", names...); ok {
			return doc, ok
		}

		return nil, false
	}
	return []string{
		"商品规格",
	}, true
}

func (v *SkuRequestForCreate) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "Spec":
			return []string{
				"商品规格属性",
			}, true
		}
		if doc, ok := runtimeDoc(&v.CodableRequest, "商品规格", names...); ok {
			return doc, ok
		}

		return nil, false
	}
	return []string{
		"商品规格",
	}, true
}

func (v *SkuRequestForUpdate) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "Spec":
			return []string{
				"商品规格属性",
			}, true
		}
		if doc, ok := runtimeDoc(&v.Request, "商品规格", names...); ok {
			return doc, ok
		}

		return nil, false
	}
	return []string{
		"商品规格",
	}, true
}

func (v *SkuSpec) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "Price":
			return []string{
				"单价",
			}, true
		}

		return nil, false
	}
	return []string{}, true
}
