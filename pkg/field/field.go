package field

import (
	"strings"

	"github.com/yeungsean/ysq-db/pkg"
	"github.com/yeungsean/ysq-db/pkg/aggregation"
)

// Type 列类型
type Type string

// Option ...
type Option struct {
	pkg.Option
	Prefix       string
	DefaultValue any
	agg          aggregation.Type
}

// Field 字段
type Field struct {
	Option
	Name Type
}

// New ...
func New(t Type) *Field {
	str := string(t)
	f := &Field{}
	if idx := strings.Index(str, "."); idx > -1 {
		f.Prefix = str[0:idx]
		f.Name = Type(str[idx+1:])
		return f
	}

	if idx := strings.Index(str, "("); idx > -1 && str[len(str)-1] == ')' {
		switch fn := strings.ToUpper(str[0:idx]); fn {
		case string(aggregation.Avg), string(aggregation.Count),
			string(aggregation.Max), string(aggregation.Min), string(aggregation.Sum):
			f.agg = aggregation.Type(fn)
			f.Name = Type(str[idx+1 : len(str)-1])
			return f
		}
	}

	f.Name = t
	return f
}

// GetAggregation ...
func (f Field) GetAggregation() aggregation.Type {
	return f.agg
}

// SetAggregation ...
func (f *Field) SetAggregation(at aggregation.Type) {
	f.agg = at
}

// Options 可选参数
type Options func(*Option)

// WithQuote ...
func WithQuote(qs ...bool) Options {
	return func(c *Option) {
		v := true
		if len(qs) > 0 {
			v = qs[0]
		}
		c.Quote = v
	}
}

// WithDefaultValue ...
func WithDefaultValue(value any) Options {
	return func(c *Option) {
		c.DefaultValue = value
	}
}

// WithAggregation ...
func WithAggregation(agg aggregation.Type) Options {
	return func(c *Option) {
		c.agg = agg
	}
}

// WithAlias ...
func WithAlias(alias string) Options {
	return func(c *Option) {
		c.Alias = alias
	}
}
