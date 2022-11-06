package column

import (
	"fmt"
	"strings"

	"github.com/yeungsean/ysq"
	"github.com/yeungsean/ysq-db/pkg/alias"
	"github.com/yeungsean/ysq-db/pkg/common"
	"github.com/yeungsean/ysq-db/pkg/field"
	"github.com/yeungsean/ysq-db/pkg/ops"
)

// Column 筛选列
type Column struct {
	field.Field
	isFilter bool
	value    any
	op       ops.Type
}

// Option 可选参数
type Option func(*Column)

// WithPrefix 前缀
func WithPrefix(value string) Option {
	return func(c *Column) {
		c.Prefix = value
	}
}

// WithValue ...
func WithValue(value any) Option {
	return func(c *Column) {
		c.value = value
	}
}

// WithOp ...
func WithOp(op ops.Type) Option {
	return func(c *Column) {
		c.op = op
	}
}

// WithAlias 别名
func WithAlias(a *alias.Alias) Option {
	return func(c *Column) {
		c.Alias = a
	}
}

// WithIsFilter ...
func WithIsFilter(filter bool) Option {
	return func(c *Column) {
		c.isFilter = filter
	}
}

// New 实例化
func New(name field.Type, options ...Option) *Column {
	column := &Column{
		op:       ops.EQ,
		isFilter: true,
	}
	column.Name = name
	common.OptionForEach(column, options)
	return column
}

// NewField 实例化字段
func NewField(name field.Type, options ...Option) *Column {
	column := New(name, options...)
	column.isFilter = false
	return column
}

// String ...
func (c Column) String() string {
	if !c.isFilter {
		return c.GetName()
	}

	switch c.op {
	case ops.IsNull, ops.IsNotNull:
		return fmt.Sprintf(`%s %s`, c.GetName(), c.op)
	case ops.Like:
		return fmt.Sprintf(`%s LIKE ?`, c.GetName())
	case ops.In:
		lst := c.value.([]any)
		strs := ysq.FromSlice(lst).
			CastToStringBy(func(any) string { return "?" }).
			ToSlice(len(lst))
		return fmt.Sprintf(`%s IN(%s)`, c.GetName(), strings.Join(strs, ","))
	default:
		if tmp, ok := c.value.(*Column); ok {
			return fmt.Sprintf("%s%s%s", c.GetName(), c.op, tmp.GetName())
		}
		return fmt.Sprintf("%s%s?", c.GetName(), c.op)
	}
}

// GetName ...
func (c *Column) GetName() string {
	field := c.Field.SelectField()
	if !c.isFilter && c.DefaultValue != nil {
		return fmt.Sprintf(`IFNULL(%s,%s)`, field, c.DefaultValue)
	}
	return field
}

// Set ...
func (c *Column) Set(value any, op ops.Type) *Column {
	c.value = value
	c.op = op
	c.isFilter = true
	return c
}

// Equal =
func (c *Column) Equal(value any) *Column {
	return c.Set(value, ops.EQ)
}

// NotEqual <>
func (c *Column) NotEqual(value any) *Column {
	return c.Set(value, ops.NEQ)
}

// GreaterThan >
func (c *Column) GreaterThan(value any) *Column {
	return c.Set(value, ops.GT)
}

// LessThan <
func (c *Column) LessThan(value any) *Column {
	return c.Set(value, ops.LT)
}

// GreaterEqual >=
func (c *Column) GreaterEqual(value any) *Column {
	return c.Set(value, ops.GTE)
}

// LessEqual <=
func (c *Column) LessEqual(value any) *Column {
	return c.Set(value, ops.LTE)
}

// Like ...
func (c *Column) Like(value any) *Column {
	return c.Set(value, ops.Like)
}

// IsNull 是否为空
func (c *Column) IsNull() *Column {
	c.op = ops.IsNull
	c.isFilter = true
	return c
}

// IsNotNull 不为空
func (c *Column) IsNotNull() *Column {
	c.op = ops.IsNotNull
	c.isFilter = true
	return c
}

// In ...
func (c *Column) In(values ...any) *Column {
	return c.Set(values, ops.In)
}
