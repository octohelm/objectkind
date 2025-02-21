package compose

import (
	"database/sql/driver"
	"fmt"

	"github.com/go-json-experiment/json"
	"github.com/octohelm/objectkind/pkg/object"
)

type Annotations map[string]string

func (annos Annotations) IsZero() bool {
	return len(annos) == 0
}

func (Annotations) DataType(driverName string) string {
	return "text"
}

func (annos Annotations) Value() (driver.Value, error) {
	if annos.IsZero() {
		return "", nil
	}

	bytes, err := json.Marshal(annos)
	if err != nil {
		return "", err
	}
	str := string(bytes)
	if str == "null" {
		return "", nil
	}
	return str, nil
}

func (annos *Annotations) Scan(src any) error {
	switch v := src.(type) {
	case []byte:
		bytes := v
		if len(bytes) > 0 {
			return json.Unmarshal(bytes, annos)
		}
		return nil
	case string:
		str := v
		if str == "" {
			return nil
		}
		return json.Unmarshal([]byte(str), annos)
	case nil:
		return nil
	default:
		return fmt.Errorf("cannot sql.Scan() from: %#v", annos)
	}
}

type Annotatable struct {
	// 其他标注
	Annotations Annotations `db:"f_annotations,null"`
}

func (a Annotatable) GetAnnotations() map[string]string {
	return a.Annotations
}

func (a Annotatable) GetAnnotation(k string) (string, bool) {
	if a.Annotations == nil {
		return "", false
	}
	v, ok := a.Annotations[k]
	return v, ok
}

var _ object.Annotater = Annotatable{}

var _ object.Annotatable = &Annotatable{}

func (a *Annotatable) SetAnnotations(annotations map[string]string) {
	a.Annotations = annotations
}

func (a *Annotatable) SetAnnotation(key string, value string) {
	if a.Annotations == nil {
		a.Annotations = Annotations{}
	}
	a.Annotations[key] = value
}
