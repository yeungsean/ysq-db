package ysqdb

import (
	"context"
	"sync/atomic"

	"github.com/yeungsean/ysq"
	"github.com/yeungsean/ysq-db/pkg"
	"github.com/yeungsean/ysq-db/pkg/column"
	"github.com/yeungsean/ysq-db/pkg/common"
	"github.com/yeungsean/ysq-db/pkg/field"
	"github.com/yeungsean/ysq-db/pkg/ops"
	"github.com/yeungsean/ysq-db/pkg/statement"
)

// QueryType Query type for 4 statements
// type QueryType uint8

// const (
// QuerySelect select stmt
// QuerySelect QueryType = iota
// QueryInsert
// QueryUpdate
// QueryDelete
// )

// Query Parse analysis turns all statements into a Query tree
type Query[T pkg.ITable] struct {
	ctx      context.Context
	buildCnt int32
	Next     func() Iterator[T]
}

// NewQuery ...
func NewQuery[T pkg.ITable](entity T, opts ...LambdaOption[T]) *Query[T] {
	q := &Query[T]{
		ctx: context.TODO(),
	}
	q.Next = func() Iterator[T] {
		lambdaCtx := newQueryContext(entity, "")
		common.OptionForEach(lambdaCtx, opts)
		q.ctx = context.WithValue(q.ctx, ctxKeyLambda, lambdaCtx)
		return func() statement.Type {
			return statement.Nop
		}
	}
	return q
}

// Context ...
func (q *Query[T]) Context(ctx context.Context) *Query[T] {
	q.ctx = ctx
	return q
}

func (q *Query[T]) ctxGetLambda() *queryContext[T] {
	return q.ctx.Value(ctxKeyLambda).(*queryContext[T])
}

// Column 查询字段，带默认值
// func (q *Query[T]) Column(col field.Type, defaultValue any) *Query[T] {
// 	return q.wrap(func(q *Query[T], lc *queryContext[T]) func() statement.Type {
// 		lc.Fields = append(lc.Fields, column.NewField(col, column.WithDefaultValue(defaultValue)))
// 		return statementColumn
// 	})
// }

func columnTypeToColumn(t field.Type) *column.Column {
	return column.NewField(t)
}

// Select ...
func (q *Query[T]) Select(cols ...field.Type) *Query[T] {
	return q.wrap(func(q *Query[T], lc *queryContext[T]) func() statement.Type {
		fields := ysq.Select(ysq.FromSlice(cols), columnTypeToColumn).ToSlice(len(cols))
		lc.Fields = append(lc.Fields, fields...)
		return statementColumn
	})
}

// SelectColumn ...
func (q *Query[T]) SelectColumn(cols ...*column.Column) *Query[T] {
	return q.wrap(func(q *Query[T], lc *queryContext[T]) func() statement.Type {
		lc.Fields = append(lc.Fields, cols...)
		return statementColumn
	})
}

// Limit ...
func (q *Query[T]) Limit(count int, offset ...int) *Query[T] {
	return q.wrap(func(q *Query[T], lc *queryContext[T]) func() statement.Type {
		lc.LimitCount = count
		lc.LimitOffset = common.VarArgGetFirst(offset...)
		return func() statement.Type {
			return statement.Limit
		}
	})
}

func (q *Query[T]) getColumn(f field.Type, value any, opt ops.Type) *column.Column {
	return column.New(f,
		column.WithIsFilter(true),
		column.WithOp(opt),
		column.WithValue(value),
	)
}

func (q *Query[T]) build() {
	if atomic.CompareAndSwapInt32(&q.buildCnt, 0, 1) {
		q.Next()
	}
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

// Values ...
func (q *Query[T]) Values() []any {
	q.build()
	return q.ctxGetLambda().Values
}
