package column

import (
	"context"
	"fmt"
	"strings"

	"github.com/yeungsean/ysq"
	"github.com/yeungsean/ysq-db/internal"
	"github.com/yeungsean/ysq-db/internal/expr/common"
	"github.com/yeungsean/ysq-db/internal/expr/ops"
	"github.com/yeungsean/ysq-db/pkg/field"
)

// Column 筛选列
type Column struct {
	field.Field
	value any
	op    ops.Type
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
func WithAlias(a string) Option {
	return func(c *Column) {
		c.Alias = a
	}
}

// New 实例化
func New(name field.Type, options ...Option) *Column {
	column := &Column{
		op: ops.EQ,
	}
	column.Name = name
	common.OptionForEach(column, options)
	return column
}

// NewField 实例化字段
func NewField(name field.Type, options ...Option) *Column {
	return New(name, options...)
}

// String ...
func (c Column) String(ctx context.Context) string {
	idx := ctx.Value(internal.CtxKeyFilterColumnIndex).(*int)
	provider := internal.CtxGetSourceProvider(ctx)
	name := provider.OtherTypeFieldQuote(&c.Field)
	switch c.op {
	case ops.IsNull, ops.IsNotNull:
		return fmt.Sprintf(`%s %s`, name, c.op)
	case ops.Like, ops.NotLike:
		ph := provider.PlaceHolder(*idx)
		return fmt.Sprintf(`%s %s %s`, name, c.op, ph)
	case ops.Between, ops.NotBetween:
		ph1 := provider.PlaceHolder(*idx)
		ph2 := provider.PlaceHolder(*idx + 1)
		*idx = *idx + 2
		return fmt.Sprintf("%s %s %s AND %s", name, c.op, ph1, ph2)
	case ops.In, ops.NotIn:
		lst := c.value.([]any)
		strs := ysq.FromSlice(lst).
			CastToStringBy(func(any) string {
				res := provider.PlaceHolder(*idx)
				*idx = *idx + 1
				return res
			}).ToSlice(len(lst))
		return fmt.Sprintf(`%s %s(%s)`, name, c.op, strings.Join(strs, ","))
	default:
		if tmp, ok := c.value.(*Column); ok {
			tmpName := provider.OtherTypeFieldQuote(&tmp.Field)
			return fmt.Sprintf("%s%s%s", name, c.op, tmpName)
		}
		res := fmt.Sprintf("%s%s%s", name, c.op, provider.PlaceHolder(*idx))
		*idx = *idx + 1
		return res
	}
}

// Set ...
func (c *Column) Set(value any, op ops.Type) *Column {
	c.value = value
	c.op = op
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
	return c
}

// IsNotNull 不为空
func (c *Column) IsNotNull() *Column {
	c.op = ops.IsNotNull
	return c
}

// In ...
func (c *Column) In(values ...any) *Column {
	return c.Set(values, ops.In)
}

// NotIn ...
func (c *Column) NotIn(values ...any) *Column {
	return c.Set(values, ops.NotIn)
}

// Between ...
func (c *Column) Between(min, max any) *Column {
	return c.Set([]any{min, max}, ops.Between)
}

// NotBetween ...
func (c *Column) NotBetween(min, max any) *Column {
	return c.Set([]any{min, max}, ops.NotBetween)
}
