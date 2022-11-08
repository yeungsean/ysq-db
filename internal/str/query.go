package str

import (
	"context"

	"github.com/yeungsean/ysq-db/pkg/column"
	"github.com/yeungsean/ysq-db/pkg/common"
	"github.com/yeungsean/ysq-db/pkg/cond"
	"github.com/yeungsean/ysq-db/pkg/field"
	"github.com/yeungsean/ysq-db/pkg/order"
	"github.com/yeungsean/ysq-db/pkg/statement"
)

type ctxKey uint8

const (
	ctxKeyLambda ctxKey = iota
)

type (
	queryContext[T string] struct {
		HasForUpdate bool
		WhereClause  *cond.Cond
		HavingClause *cond.Cond
		Fields       []*field.Field
		Groups       []*column.Column
		orders       []*order.Order
		mainTable    tableExpr[T]
		joins        []joinExpr[T]

		LimitOffset int
		LimitCount  int
		Values      []any
	}

	// Iterable ...
	Iterable[T string] interface {
		Next() Iterator[T]
	}

	// Iterator 迭代器
	Iterator[T string] statement.Type

	// Query ...
	Query[T string] struct {
		ctx      context.Context
		buildCnt int32
		Next     func() statement.Type
	}
)

func newQueryContext[T string]() *queryContext[T] {
	qc := &queryContext[T]{
		WhereClause:  cond.Linear(),
		HavingClause: cond.Linear(),
		Fields:       make([]*field.Field, 0, 1),
		Groups:       make([]*column.Column, 0),
		orders:       make([]*order.Order, 0),
		Values:       make([]any, 0, 1),
		joins:        make([]joinExpr[T], 0),
		mainTable:    tableExpr[T]{},
	}
	return qc
}

// NewQuery ...
func NewQuery[T string]() *Query[T] {
	q := &Query[T]{
		ctx: context.TODO(),
	}
	q.Next = func() statement.Type {
		q.ctx = context.WithValue(q.ctx, ctxKeyLambda, newQueryContext())
		return statement.Nop
	}
	return q
}

// Context ...
func (q *Query[T]) Context(ctx context.Context) *Query[T] {
	q.ctx = ctx
	return q
}

// Entity ...
func (q *Query[T]) Entity(e T, aliasOpt ...string) *Query[T] {
	nextQ := &Query[T]{}
	nextQ.Next = func() statement.Type {
		q.Next()
		alias := common.VarArgGetFirst(aliasOpt...)
		lctx := q.ctxGetLambda()
		if alias != "" {
			lctx.mainTable.table = e
			lctx.mainTable.alias = alias
		} else {
			lctx.mainTable.table = e
		}
		nextQ.ctx = context.WithValue(q.ctx, ctxKeyLambda, lctx)
		return statement.Table
	}
	return nextQ
}

func (q *Query[T]) ctxGetLambda() *queryContext[T] {
	return q.ctx.Value(ctxKeyLambda).(*queryContext[T])
}

func (q *Query[T]) wrap(f func(*Query[T], *queryContext[T]) statement.Type) *Query[T] {
	nextQ := &Query[T]{}
	nextQ.Next = func() statement.Type {
		q.Next()
		lctx := q.ctxGetLambda()
		f1 := f(q, lctx)
		nextQ.ctx = context.WithValue(q.ctx, ctxKeyLambda, lctx)
		return f1
	}
	return nextQ
}
