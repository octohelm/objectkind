package digest

import (
	"github.com/octohelm/courier/pkg/validator"
	"github.com/octohelm/exp/xiter"
	metav1 "github.com/octohelm/objectkind/pkg/apis/meta/v1"
	"github.com/octohelm/objectkind/pkg/object"
	"github.com/octohelm/x/anyjson"
	"github.com/opencontainers/go-digest"
)

var annotationKeysShouldOmit = []string{
	string(metav1.AnnotationSpecDigest),
	string(metav1.AnnotationRevisionID),
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

func HashTo(dgst *Digest, v any) error {
	o, err := fromValue(v)
	if err != nil {
		return nil
	}
	d := digest.SHA256.Digester()
	if err := validator.MarshalWrite(d.Hash(), anyjson.Sorted(o)); err != nil {
		return err
	}
	*dgst = d.Digest()
	return nil
}

func fromValue(v any) (anyjson.Valuer, error) {
	if vv, ok := v.(anyjson.Valuer); ok {
		return vv, nil
	}
	return anyjson.FromValue(v)
}
