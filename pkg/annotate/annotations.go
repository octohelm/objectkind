package annotate

// Annotations 注解键值对映射。
type Annotations map[string]string

// SetAnnotation 设置注解键对应的值。
func (a Annotations) SetAnnotation(k string, value string) {
	a[k] = value
}

// GetAnnotation 获取注解键对应的值。
func (a Annotations) GetAnnotation(k string) (value string, ok bool) {
	v, ok := a[k]
	return v, ok
}
