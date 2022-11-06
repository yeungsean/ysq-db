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
		tables       []tableExpr[T]
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
	Iterator[T string] func() statement.Type

	// Query ...
	Query[T string] struct {
		ctx      context.Context
		buildCnt int32
		Next     func() Iterator[T]
	}
)

func statementWhere() statement.Type {
	return statement.Where
}

// NewQuery ...
func NewQuery[T string]() *Query[T] {
	return &Query[T]{
		ctx: context.TODO(),
	}
}

// Context ...
func (q *Query[T]) Context(ctx context.Context) *Query[T] {
	q.ctx = ctx
	return q
}

// Entity ...
func (q *Query[T]) Entity(e T, aliasOpt ...string) *Query[T] {
	alias := common.VarArgGetFirst(aliasOpt...)
	ctx := q.ctxGetLambda()
	if alias != "" {
		ctx.tables = append(ctx.tables, tableExpr[T]{table: e, alias: alias})
	} else {
		ctx.tables = append(ctx.tables, tableExpr[T]{table: e})
	}
	return q
}

func (q *Query[T]) ctxGetLambda() *queryContext[T] {
	return q.ctx.Value(ctxKeyLambda).(*queryContext[T])
}

func (q *Query[T]) wrap(f func(*Query[T], *queryContext[T]) func() statement.Type) *Query[T] {
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
