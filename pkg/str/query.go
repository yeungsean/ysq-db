package str

import (
	"context"

	"github.com/yeungsean/ysq-db/internal/expr/statement"
	"github.com/yeungsean/ysq-db/pkg"
	"github.com/yeungsean/ysq-db/pkg/dbprovider"
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

// NewQuery 实例化query
func NewQuery[T string](ctx context.Context, ps ...dbprovider.IDBProvider) *Query[T] {
	if ctx == nil {
		ctx = context.Background()
	}
	q := &Query[T]{
		ctx: ctx,
	}
	q.Next = func() statement.Type {
		lctx := newQueryContext()
		q.ctx = context.WithValue(q.ctx, pkg.CtxKeyLambda, lctx)
		if len(ps) > 0 {
			q.ctx = context.WithValue(q.ctx, pkg.CtxKeyDBProvider, ps[0])
		}
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

// Entity 实体名(表名)
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
