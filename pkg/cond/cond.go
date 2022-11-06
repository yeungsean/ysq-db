package cond

import (
	"fmt"
	"strings"

	"github.com/yeungsean/ysq"
	"github.com/yeungsean/ysq-db/pkg/column"
	"github.com/yeungsean/ysq-db/pkg/common"
	"github.com/yeungsean/ysq-db/pkg/ops"
)

// LogicType 逻辑类型
type LogicType uint8

const (
	// And 逻辑 &&
	And LogicType = iota
	// Or 逻辑 ||
	Or
	// Line 线形
	Line
)

// String ...
func (lt LogicType) String() string {
	switch lt {
	case And:
		return " AND "
	case Or:
		return " OR "
	}
	return ""
}

// CondColumn ...
type CondColumn struct {
	LogicType
	*column.Column
}

// Cond ...
type Cond struct {
	logic     LogicType
	ops       []*CondColumn
	childCond []*Cond
}

// AddChildren ...
func (c *Cond) AddChildren(others ...*Cond) *Cond {
	c.childCond = append(c.childCond, others...)
	return c
}

// Add ...
func (c *Cond) Add(col *column.Column, lts ...LogicType) *Cond {
	lt := common.VarArgGetFirst(lts...)
	c.ops = append(c.ops, &CondColumn{Column: col, LogicType: lt})
	return c
}

// Add2 ...
func (c *Cond) Add2(col1, col2 *column.Column, ot ops.Type, lts ...LogicType) *Cond {
	lt := common.VarArgGetFirst(lts...)
	col1.Set(col2, ot)
	c.ops = append(c.ops, &CondColumn{
		Column:    col1,
		LogicType: lt,
	})
	return c
}

// stringLine ...
func (c *Cond) stringLine() string {
	sb := strings.Builder{}
	if len(c.ops) == 0 {
		return sb.String()
	}

	sb.WriteString(c.ops[0].Column.String())
	if len(c.ops) == 1 {
		return sb.String()
	}

	for _, op := range c.ops[1:] {
		sb.WriteString(fmt.Sprintf("%s%s", op.LogicType, op.Column.String()))
	}
	return sb.String()
}

// String ...
func (c *Cond) String() string {
	if len(c.ops) == 0 && len(c.childCond) == 0 {
		return ""
	} else if c.logic == Line {
		return c.stringLine()
	}

	var join string
	switch c.logic {
	case Or:
		join = Or.String()
	case And:
		join = And.String()
	}
	strs := ysq.FromSlice(c.ops).CastToStringBy(func(c *CondColumn) string { return c.Column.String() }).
		Reduce("", func(s1, s2 string) string {
			if s1 == "" {
				return "(" + s2
			} else if s2 == "" {
				return s1
			}
			return s1 + join + s2
		})
	strs += ")"

	sb := strings.Builder{}
	if len(c.ops) > 0 {
		sb.WriteString(strs)
	}
	if len(c.childCond) == 0 {
		return sb.String()
	}

	cc := ysq.FromSlice(c.childCond).CastToStringBy(func(c *Cond) string { return c.String() }).
		Reduce("", func(s1, s2 string) string {
			if s1 == "" {
				return "(" + s2
			} else if s2 == "" {
				return s1
			}
			return s1 + join + s2
		})
	cc += ")"
	sb.WriteString(cc)
	return sb.String()
}

// Linear ...
func Linear() *Cond {
	cond := &Cond{
		ops:       make([]*CondColumn, 0),
		childCond: make([]*Cond, 0),
		logic:     Line,
	}
	return cond
}

// Any 任何一个条件满足
func Any() *Cond {
	c := Linear()
	c.logic = Or
	return c
}

// All 所有条件满足
func All() *Cond {
	c := Linear()
	c.logic = And
	return c
}
