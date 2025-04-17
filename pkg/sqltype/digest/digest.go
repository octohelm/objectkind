package digest

import (
	"github.com/octohelm/exp/xiter"
	metav1 "github.com/octohelm/objectkind/pkg/apis/meta/v1"
	"github.com/octohelm/objectkind/pkg/object"
	"github.com/opencontainers/go-digest"
)

var annotationKeysShouldOmit = []string{
	string(metav1.AnnotationRevisionID),
	string(metav1.AnnotationRevisionDigest),

	string(metav1.AnnotationSpecDigest),
}

func OmitAnnotations[O object.Annotater](src O, omitKeys ...string) {
	if target, ok := any(src).(object.Annotatable); ok {
		// omit internals keys
		for key := range xiter.Concat(xiter.Of(annotationKeysShouldOmit...), xiter.Of(omitKeys...)) {
			if _, ok := src.GetAnnotation(key); ok {
				annotations := src.GetAnnotations()
				delete(annotations, key)
				target.SetAnnotations(annotations)
			}
		}
	}
}

type Digest = digest.Digest
