package str

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/yeungsean/ysq-db/internal"
	"github.com/yeungsean/ysq-db/internal/expr/statement"
	"github.com/yeungsean/ysq-db/internal/provider"
	"github.com/yeungsean/ysq-db/internal/provider/mysql"
	"github.com/yeungsean/ysq-db/pkg/option"
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
	q.ctx = context.WithValue(q.ctx, internal.CtxKeyDBProvider, &mysql.Provider{})
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

// Context 绑定自定义context
func (q *Query[T]) Context(ctx context.Context) *Query[T] {
	q.ctx = ctx
	return q
}

// WithTx 使用事务
func (q *Query[T]) WithTx(tx *sqlx.Tx) *Query[T] {
	q.ctx = context.WithValue(q.ctx, internal.CtxKeyTx, tx)
	return q
}

// WithDBProvider ...
func (q *Query[T]) WithDBProvider(p provider.IProvider) *Query[T] {
	return q
}

// Entity ...
func (q *Query[T]) Entity(e T, opts ...option.Options) *Query[T] {
	nextQ := &Query[T]{}
	nextQ.Next = func() statement.Type {
		q.Next()
		var opt option.Option
		option.ForEach(&opt, opts)
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
	return q.ctx.Value(internal.CtxKeyDBProvider).(provider.IProvider)
}

func (q *Query[T]) ctxGetTx() *sqlx.Tx {
	tmp := q.ctx.Value(internal.CtxKeyTx)
	if tmp == nil {
		return nil
	}
	return tmp.(*sqlx.Tx)
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
