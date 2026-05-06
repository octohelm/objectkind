package compose

import (
	"database/sql/driver"
	"fmt"

	"github.com/go-json-experiment/json"

	"github.com/octohelm/objectkind/pkg/object"
)

// Annotations 键值对注解，实现 sql.Scanner 与 driver.Valuer，以 JSON 文本形式存储。
type Annotations map[string]string

// IsZero 判断是否为空注解。
func (annos Annotations) IsZero() bool {
	return len(annos) == 0
}

// DataType 返回数据库列类型 text。
func (Annotations) DataType(driverName string) string {
	return "text"
}

// Value 实现 driver.Valuer，序列化为 JSON 文本。
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

// Scan 实现 sql.Scanner，从字节或字符串反序列化 JSON。
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

// Annotatable 可注解类型，组合 Annotations 并提供读写便利方法。
type Annotatable struct {
	// 其他标注
	Annotations Annotations `db:"f_annotations,null"`
}

// GetAnnotations 返回所有注解。
func (a Annotatable) GetAnnotations() map[string]string {
	return a.Annotations
}

// GetAnnotation 返回指定 key 对应的注解值。
func (a Annotatable) GetAnnotation(k string) (string, bool) {
	if a.Annotations == nil {
		return "", false
	}
	v, ok := a.Annotations[k]
	return v, ok
}

var _ object.Annotater = Annotatable{}

var _ object.Annotatable = &Annotatable{}

// SetAnnotations 批量覆盖注解。
func (a *Annotatable) SetAnnotations(annotations map[string]string) {
	a.Annotations = annotations
}

// SetAnnotation 设置指定 key 的注解值。
func (a *Annotatable) SetAnnotation(key string, value string) {
	if a.Annotations == nil {
		a.Annotations = Annotations{}
	}
	a.Annotations[key] = value
}
