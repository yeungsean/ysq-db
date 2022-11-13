package field

import (
	"github.com/yeungsean/ysq-db/internal/expr/aggregation"
)

// Type 列类型
type Type string

// FieldOption ...
type FieldOption struct {
	Prefix       string
	DefaultValue any
	Alias        string
	agg          aggregation.Type
}

// Field ...
type Field struct {
	FieldOption
	Name Type
}

// GetAggregation ...
func (f Field) GetAggregation() aggregation.Type {
	return f.agg
}

// Option 可选参数
type Option func(*FieldOption)

// WithDefaultValue ...
func WithDefaultValue(value any) Option {
	return func(c *FieldOption) {
		c.DefaultValue = value
	}
}

// WithAggregation ...
func WithAggregation(agg aggregation.Type) Option {
	return func(c *FieldOption) {
		c.agg = agg
	}
}
