package str

import (
	"github.com/yeungsean/ysq-db/pkg/column"
	"github.com/yeungsean/ysq-db/pkg/common"
	"github.com/yeungsean/ysq-db/pkg/cond"
	"github.com/yeungsean/ysq-db/pkg/field"
	"github.com/yeungsean/ysq-db/pkg/ops"
	"github.com/yeungsean/ysq-db/pkg/statement"
)

func (q *Query[T]) wrapStatementWhere(f func() *column.Column,
	value any, lts ...cond.LogicType) *Query[T] {
	return q.wrap(func(iq *Query[T], lc *queryContext[T]) statement.Type {
		col := f()
		lt := common.VarArgGetFirst(lts...)
		lc.WhereClause.Add(col, lt)
		lc.Values = append(lc.Values, value)
		return statement.Where
	})
}

// Equal =
func (q *Query[T]) Equal(f field.Type, value any, lts ...cond.LogicType) *Query[T] {
	return q.wrapStatementWhere(
		func() *column.Column {
			return column.New(f,
				column.WithValue(value),
				column.WithOp(ops.EQ),
			)
		}, value, lts...,
	)
}

// NotEqual <>
func (q *Query[T]) NotEqual(f field.Type, value any, lts ...cond.LogicType) *Query[T] {
	return q.wrapStatementWhere(
		func() *column.Column {
			return column.New(f,
				column.WithValue(value),
				column.WithOp(ops.NEQ),
			)
		}, value, lts...,
	)
}

// Greater >
func (q *Query[T]) Greater(f field.Type, value any, lts ...cond.LogicType) *Query[T] {
	return q.wrapStatementWhere(
		func() *column.Column {
			return column.New(f,
				column.WithValue(value),
				column.WithOp(ops.GT),
			)
		}, value, lts...,
	)
}

// GreaterOrEqual >=
func (q *Query[T]) GreaterOrEqual(f field.Type, value any, lts ...cond.LogicType) *Query[T] {
	return q.wrapStatementWhere(
		func() *column.Column {
			return column.New(f,
				column.WithValue(value),
				column.WithOp(ops.GTE),
			)
		}, value, lts...,
	)
}

// Less <
func (q *Query[T]) Less(f field.Type, value any, lts ...cond.LogicType) *Query[T] {
	return q.wrapStatementWhere(
		func() *column.Column {
			return column.New(f,
				column.WithValue(value),
				column.WithOp(ops.LT),
			)
		}, value, lts...,
	)
}

// LessOrEqual <=
func (q *Query[T]) LessOrEqual(f field.Type, value any, lts ...cond.LogicType) *Query[T] {
	return q.wrapStatementWhere(
		func() *column.Column {
			return column.New(f,
				column.WithValue(value),
				column.WithOp(ops.LTE),
			)
		}, value, lts...,
	)
}

// IsNull 是否为null
func (q *Query[T]) IsNull(col field.Type, lts ...cond.LogicType) *Query[T] {
	return q.wrap(func(q *Query[T], lc *queryContext[T]) statement.Type {
		lt := common.VarArgGetFirst(lts...)
		lc.WhereClause.Add(column.New(col).IsNull(), lt)
		return statement.Where
	})
}

// IsNotNull 是否为not null
func (q *Query[T]) IsNotNull(col field.Type, lts ...cond.LogicType) *Query[T] {
	return q.wrap(func(q *Query[T], lc *queryContext[T]) statement.Type {
		lt := common.VarArgGetFirst(lts...)
		lc.WhereClause.Add(column.New(col).IsNotNull(), lt)
		return statement.Where
	})
}

// Like ...
func (q *Query[T]) Like(col field.Type, value any, lts ...cond.LogicType) *Query[T] {
	return q.wrap(func(q *Query[T], lc *queryContext[T]) statement.Type {
		lt := common.VarArgGetFirst(lts...)
		lc.WhereClause.Add(column.New(col).Like(value), lt)
		return statement.Where
	})
}

// Where ...
func (q *Query[T]) Where(conds ...*cond.Cond) *Query[T] {
	return q.wrap(func(q *Query[T], lc *queryContext[T]) statement.Type {
		lc.WhereClause.AddChildren(conds...)
		return statement.Where
	})
}
