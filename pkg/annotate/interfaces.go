package annotate

// Provider 注解值获取接口。
type Provider interface {
	GetAnnotation(k string) (value string, ok bool)
}

// Setter 注解值设置接口。
type Setter interface {
	SetAnnotation(k string, value string)
}
