package field

import (
	"fmt"

	"github.com/yeungsean/ysq-db/pkg/alias"
)

// Type 列类型
type Type string

// Field ...
type Field struct {
	Prefix       string
	Name         Type
	Alias        *alias.Alias
	DefaultValue any
}

// Option 可选参数
type Option func(*Field)

// WithDefaultValue ...
func WithDefaultValue(value any) Option {
	return func(c *Field) {
		c.DefaultValue = value
	}
}

// SelectField ...
func (f Field) SelectField() string {
	if f.Alias == nil {
		if f.Prefix == "" {
			return string(f.Name)
		}
		return fmt.Sprintf(`%s.%s`, f.Prefix, f.Name)
	}

	if f.Prefix == "" {
		return fmt.Sprintf(`%s AS %s`, f.Name, f.Alias.Name)
	}
	return fmt.Sprintf(`%s.%s AS %s`, f.Prefix, f.Name, f.Alias.Name)
}

// OrderField ...
func (f Field) OrderField() string {
	if f.Alias == nil {
		if f.Prefix == "" {
			return string(f.Name)
		}
		return fmt.Sprintf(`%s.%s`, f.Prefix, f.Name)
	}
	return f.Alias.Name
}
