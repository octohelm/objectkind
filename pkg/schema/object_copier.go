package schema

type ObjectReceiver interface {
	CopyFromObject(o Object)
}

type ObjectProvider interface {
	CopyToObject(o ObjectSetter)
}
