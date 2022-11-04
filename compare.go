package ysqdb

import (
	"context"

	"github.com/yeungsean/ysq-db/pkg/column"
	"github.com/yeungsean/ysq-db/pkg/cond"
	"github.com/yeungsean/ysq-db/pkg/field"
	"github.com/yeungsean/ysq-db/pkg/ops"
	"github.com/yeungsean/ysq-db/pkg/statement"
)

func (q *Query[T]) wrap(f func(*Query[T], *lambdaContext[T]) func() statement.Type) *Query[T] {
	nextQ := &Query[T]{}
	nextQ.Next = func() Iterator[T] {
		q.Next()
		lctx := q.ctxGetLambda()
		f1 := f(q, lctx)
		nextQ.ctx = context.WithValue(q.ctx, ctxKeyLambda, lctx)
		return f1
	}
	return nextQ
}

func (q *Query[T]) wrapStatementWhere(f func(*Query[T], *lambdaContext[T]) *column.Column,
	value any, lts ...cond.LogicType) *Query[T] {
	return q.wrap(func(iq *Query[T], lc *lambdaContext[T]) func() statement.Type {
		col := f(iq, lc)
		lc.WhereClause.Add(col, lts...)
		lc.Values = append(lc.Values, value)
		return statementWhere
	})
}

// Equal =
func (q *Query[T]) Equal(f field.Type, value any, lts ...cond.LogicType) *Query[T] {
	return q.wrapStatementWhere(func(q *Query[T], lc *lambdaContext[T]) *column.Column {
		return q.getColumn(f, value, ops.EQ)
	}, value, lts...)
}

// NotEqual <>
func (q *Query[T]) NotEqual(f field.Type, value any, lts ...cond.LogicType) *Query[T] {
	return q.wrapStatementWhere(func(q *Query[T], lc *lambdaContext[T]) *column.Column {
		return q.getColumn(f, value, ops.NEQ)
	}, value, lts...)
}

// Greater >
func (q *Query[T]) Greater(f field.Type, value any, lts ...cond.LogicType) *Query[T] {
	return q.wrapStatementWhere(func(q *Query[T], lc *lambdaContext[T]) *column.Column {
		return q.getColumn(f, value, ops.GT)
	}, value, lts...)
}

// GreaterOrEqual >=
func (q *Query[T]) GreaterOrEqual(f field.Type, value any, lts ...cond.LogicType) *Query[T] {
	return q.wrapStatementWhere(func(q *Query[T], lc *lambdaContext[T]) *column.Column {
		return q.getColumn(f, value, ops.GTE)
	}, value, lts...)
}

// Less <
func (q *Query[T]) Less(f field.Type, value any, lts ...cond.LogicType) *Query[T] {
	return q.wrapStatementWhere(func(q *Query[T], lc *lambdaContext[T]) *column.Column {
		return q.getColumn(f, value, ops.LT)
	}, value, lts...)
}

// LessOrEqual <=
func (q *Query[T]) LessOrEqual(f field.Type, value any, lts ...cond.LogicType) *Query[T] {
	return q.wrapStatementWhere(func(q *Query[T], lc *lambdaContext[T]) *column.Column {
		return q.getColumn(f, value, ops.LTE)
	}, value, lts...)
}
