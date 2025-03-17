package runtime

import (
	"github.com/octohelm/objectkind/pkg/object"
)

func CopyCodableObject[D object.Type, ID object.Identity, Code ~string, S object.CodableObject[ID, Code]](dst *D, src *S) {
	if d, ok := any(dst).(object.CodeSetter[Code]); ok {
		if s, ok := any(src).(object.CodeGetter[Code]); ok {
			d.SetCode(s.GetCode())
		}
	}

	CopyObject(dst, src)
}

func CopyObject[D object.Type, ID object.Identity, S object.Object[ID]](dst *D, src *S) {
	if d, ok := any(dst).(object.IDSetter[ID]); ok {
		if s, ok := any(src).(object.IDGetter[ID]); ok {
			d.SetID(s.GetID())
		}
	}

	Copy(dst, src)
}

func CopyCodable[D object.Type, Code ~string, S object.Codable[Code]](dst *D, src *S) {
	if d, ok := any(dst).(object.CodeSetter[Code]); ok {
		if s, ok := any(src).(object.CodeGetter[Code]); ok {
			d.SetCode(s.GetCode())
		}
	}

	Copy(dst, src)
}

func Copy[D object.Type, S object.Type](dst *D, src *S) {
	copyObject(dst, src)
}

func copyObject(dst any, src any) {
	if d, ok := dst.(object.Describable); ok {
		if s, ok := src.(object.Describer); ok {
			d.SetName(s.GetName())
			d.SetDescription(s.GetDescription())
		}
	}

	if d, ok := dst.(object.Annotatable); ok {
		if s, ok := src.(object.Annotater); ok {
			for k, v := range s.GetAnnotations() {
				d.SetAnnotation(k, v)
			}
		}
	}

	if d, ok := dst.(object.CreationTimestampDescriber); ok {
		if s, ok := src.(object.CreationTimestampDescriber); ok {
			d.SetCreationTimestamp(s.GetCreationTimestamp())
		}
	}

	if d, ok := dst.(object.ModificationTimestampDescriber); ok {
		if s, ok := src.(object.ModificationTimestampDescriber); ok {
			d.SetModificationTimestamp(s.GetModificationTimestamp())
		}
	}
}
