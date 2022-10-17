package cond

import (
	"fmt"
	"strings"

	"github.com/yeungsean/ysq"
	"github.com/yeungsean/ysq-db/pkg/column"
)

// Cond ...
type Cond struct {
	isAll     bool
	ops       []*column.Column
	childCond []*Cond
}

// AddChildren ...
func (c *Cond) AddChildren(others ...*Cond) *Cond {
	c.childCond = append(c.childCond, others...)
	return c
}

// Add ...
func (c *Cond) Add(col *column.Column) *Cond {
	c.ops = append(c.ops, col)
	return c
}

// String ...
func (c *Cond) String() string {
	strs := ysq.FromSlice(c.ops).CastToStringBy(func(c *column.Column) string {
		return c.String()
	}).ToSlice(uint(len(c.ops)))
	join := " OR "
	if c.isAll {
		join = " AND "
	}

	sb := strings.Builder{}
	if len(c.ops) > 0 {
		sb.WriteString(fmt.Sprintf("(%s)", strings.Join(strs, join)))
	}
	if len(c.childCond) == 0 {
		return sb.String()
	}

	strs = ysq.FromSlice(c.childCond).CastToStringBy(func(c *Cond) string {
		return c.String()
	}).ToSlice(uint(len(c.childCond)))
	sb.WriteString(fmt.Sprintf("(%s)", strings.Join(strs, join)))
	return sb.String()
}

// Or ...
func Or() *Cond {
	return Any()
}

// And ...
func And() *Cond {
	return All()
}

// Any ...
func Any() *Cond {
	return &Cond{
		ops:       make([]*column.Column, 0),
		childCond: make([]*Cond, 0),
	}
}

// All ...
func All() *Cond {
	c := Any()
	c.isAll = true
	return c
}
