package column

import (
	"fmt"
	"strings"

	"github.com/yeungsean/ysq"
	"github.com/yeungsean/ysq-db/pkg/ops"
)

// Type 列类型
type Type string

// Column 列
type Column struct {
	name  Type
	value interface{}
	op    ops.Type
}

// New 实例化
func New(name Type) *Column {
	return &Column{
		name: name,
		op:   ops.EQ,
	}
}

// String ...
func (c Column) String() string {
	switch c.op {
	case ops.IsNull, ops.IsNotNull:
		return fmt.Sprintf(`%s %s`, c.name, c.op)
	case ops.Like:
		return fmt.Sprintf(`%s LIKE ?`, c.name)
	case ops.In:
		lst := c.value.([]interface{})
		strs := ysq.FromSlice(lst).CastToStringBy(func(i interface{}) string {
			return "?"
		}).ToSlice(uint(len(lst)))
		return fmt.Sprintf(`%s IN(%s)`, c.name, strings.Join(strs, ","))
	default:
		return fmt.Sprintf("%s%s?", c.name, c.op)
	}
}

// Default ...
func (c *Column) Default(val interface{}) *Column {
	c.value = val
	return c
}

func (c *Column) set(value interface{}, op ops.Type) {
	c.value = value
	c.op = op
}

// Equal =
func (c *Column) Equal(val interface{}) *Column {
	c.set(val, ops.EQ)
	return c
}

// NotEqual <>
func (c *Column) NotEqual(val interface{}) *Column {
	c.set(val, ops.NEQ)
	return c
}

// GreaterThan >
func (c *Column) GreaterThan(val interface{}) *Column {
	c.set(val, ops.GT)
	return c
}

// LessThan <
func (c *Column) LessThan(val interface{}) *Column {
	c.set(val, ops.LT)
	return c
}

// GreaterEqual >=
func (c *Column) GreaterEqual(val interface{}) *Column {
	c.set(val, ops.GTE)
	return c
}

// LessEqual <=
func (c *Column) LessEqual(val interface{}) *Column {
	c.set(val, ops.LTE)
	return c
}

// Like ...
func (c *Column) Like(val interface{}) *Column {
	c.set(val, ops.Like)
	return c
}

// IsNull 是否为空
func (c *Column) IsNull() *Column {
	c.op = ops.IsNull
	return c
}

// IsNotNull 不为空
func (c *Column) IsNotNull() *Column {
	c.op = ops.IsNotNull
	return c
}

// In ...
func (c *Column) In(vals ...interface{}) *Column {
	c.op = ops.In
	c.value = vals
	return c
}
