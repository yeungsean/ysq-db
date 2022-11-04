package field

import (
	"fmt"

	"github.com/yeungsean/ysq-db/pkg/alias"
)

// Type 列类型
type Type string

// Field ...
type Field struct {
	Table string
	Name  Type
	Alias *alias.Alias
}

// SelectField ...
func (f Field) SelectField() string {
	if f.Alias == nil {
		if f.Table == "" {
			return string(f.Name)
		}
		return fmt.Sprintf(`%s.%s`, f.Table, f.Name)
	}

	if f.Table == "" {
		return fmt.Sprintf(`%s AS %s`, f.Name, f.Alias.Name)
	}
	return fmt.Sprintf(`%s.%s AS %s`, f.Table, f.Name, f.Alias.Name)
}

// OrderField ...
func (f Field) OrderField() string {
	if f.Alias == nil {
		if f.Table == "" {
			return string(f.Name)
		}
		return fmt.Sprintf(`%s.%s`, f.Table, f.Name)
	}
	return f.Alias.Name
}
