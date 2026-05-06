package v1

import (
	"github.com/octohelm/objectkind/pkg/annotate"
)

// AnnotationSpecDigest 规格摘要注解
const (
	AnnotationSpecDigest annotate.Annotation = "spec/digest"
)

// AnnotationRevisionID 修订版本 ID 注解
// AnnotationRevisionDigest 修订版本摘要注解
const (
	AnnotationRevisionID     annotate.Annotation = "revision/id"
	AnnotationRevisionDigest annotate.Annotation = "revision/digest"
)
