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
	strs := ysq.Select(ysq.FromSlice(c.ops), func(c *column.Column) string {
		str, _ := c.String()
		return str
	}).ToSlice(uint(len(c.ops)))

	sb := strings.Builder{}
	if c.isAll {
		sb.WriteString(fmt.Sprintf("(%s)", strings.Join(strs, " AND ")))
	} else {
		sb.WriteString(fmt.Sprintf("(%s)", strings.Join(strs, " OR ")))
	}

	if len(c.childCond) == 0 {
		return sb.String()
	}

	sb.WriteString(" OR ")
	strs = ysq.Select(ysq.FromSlice(c.childCond), func(c *Cond) string {
		return c.String()
	}).ToSlice(uint(len(c.childCond)))
	if c.isAll {
		sb.WriteString(fmt.Sprintf("(%s)", strings.Join(strs, " AND ")))
	} else {
		sb.WriteString(fmt.Sprintf("(%s)", strings.Join(strs, " OR ")))
	}
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
