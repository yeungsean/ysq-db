package str

import (
	"context"

	"github.com/yeungsean/ysq-db/internal"
	"github.com/yeungsean/ysq-db/internal/expr/common"
	"github.com/yeungsean/ysq-db/internal/expr/statement"
	"github.com/yeungsean/ysq-db/internal/provider"
	"github.com/yeungsean/ysq-db/internal/provider/mysql"
	"github.com/yeungsean/ysq-db/pkg"
)

type (
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

// NewQuery ...
func NewQuery[T string](e ...T) *Query[T] {
	q := &Query[T]{
		ctx: context.TODO(),
	}
	q.ctx = context.WithValue(q.ctx, internal.CtxKeySourceProvider, &mysql.Provider{})
	q.Next = func() statement.Type {
		lctx := newQueryContext()
		if len(e) > 0 {
			lctx.mainTable.Table = string(e[0])
		}
		q.ctx = context.WithValue(q.ctx, internal.CtxKeyLambda, lctx)
		return statement.Nop
	}
	return q
}

// As 别名
func (q *Query[T]) As(alias string) *Query[T] {
	nextQ := &Query[T]{}
	nextQ.Next = func() statement.Type {
		q.Next()
		lctx := q.ctxGetLambda()
		lctx.mainTable.Alias = alias
		nextQ.ctx = q.ctx
		return statement.Table
	}
	return nextQ
}

// Context ...
func (q *Query[T]) Context(ctx context.Context) *Query[T] {
	q.ctx = ctx
	return q
}

// Entity ...
func (q *Query[T]) Entity(e T, opts ...pkg.Options) *Query[T] {
	nextQ := &Query[T]{}
	nextQ.Next = func() statement.Type {
		q.Next()
		var opt pkg.Option
		common.OptionForEach(&opt, opts)
		lctx := q.ctxGetLambda()
		lctx.mainTable.Option = opt
		lctx.mainTable.Table = e
		nextQ.ctx = q.ctx
		return statement.Table
	}
	return nextQ
}

func (q *Query[T]) ctxGetLambda() *queryContext[T] {
	return q.ctx.Value(internal.CtxKeyLambda).(*queryContext[T])
}

func (q *Query[T]) ctxGetProvider() provider.IProvider {
	return q.ctx.Value(internal.CtxKeySourceProvider).(provider.IProvider)
}

func (q *Query[T]) wrap(f func(*Query[T], *queryContext[T]) statement.Type) *Query[T] {
	nextQ := &Query[T]{}
	nextQ.Next = func() statement.Type {
		q.Next()
		lctx := q.ctxGetLambda()
		f1 := f(q, lctx)
		nextQ.ctx = q.ctx
		return f1
	}
	return nextQ
}
