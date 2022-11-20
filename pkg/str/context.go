package str

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/yeungsean/ysq-db/internal/expr/cond"
	"github.com/yeungsean/ysq-db/internal/expr/join"
	"github.com/yeungsean/ysq-db/internal/expr/order"
	"github.com/yeungsean/ysq-db/internal/expr/statement"
	"github.com/yeungsean/ysq-db/internal/expr/table"
	"github.com/yeungsean/ysq-db/pkg"
	"github.com/yeungsean/ysq-db/pkg/dbprovider"
	"github.com/yeungsean/ysq-db/pkg/field"
)

type queryContext[T string] struct {
	HasForUpdate bool
	WhereClause  *cond.Cond
	HavingClause *cond.Cond
	Fields       []*field.Field
	Groups       []*field.Field
	orders       []*order.Order
	mainTable    table.Expr[T]
	joins        []join.Expr[T]

	LimitOffset int
	LimitCount  int
	Args        []any
}

func newQueryContext[T string]() *queryContext[T] {
	qc := &queryContext[T]{
		WhereClause:  cond.Linear(),
		HavingClause: cond.Linear(),
		Fields:       make([]*field.Field, 0, 1),
		Groups:       make([]*field.Field, 0),
		orders:       make([]*order.Order, 0),
		Args:         make([]any, 0, 1),
		joins:        make([]join.Expr[T], 0),
		mainTable:    table.Expr[T]{},
	}
	return qc
}

func (q *Query[T]) ctxGetLambda() *queryContext[T] {
	return q.ctx.Value(pkg.CtxKeyLambda).(*queryContext[T])
}

func (q *Query[T]) ctxGetProvider() dbprovider.IDBProvider {
	return q.ctx.Value(pkg.CtxKeyDBProvider).(dbprovider.IDBProvider)
}

func (q *Query[T]) ctxGetTx() *sqlx.Tx {
	tmp := q.ctx.Value(pkg.CtxKeyTx)
	if tmp == nil {
		return nil
	}
	return tmp.(*sqlx.Tx)
}

func (q *Query[T]) ctxGetDB() *sqlx.DB {
	tmp := q.ctx.Value(pkg.CtxKeyDB)
	if tmp == nil {
		return nil
	}
	return tmp.(*sqlx.DB)
}

// Context 绑定自定义context
func (q *Query[T]) Context(ctx context.Context) *Query[T] {
	q.ctx = ctx
	return q
}

// WithTx 使用事务
func (q *Query[T]) WithTx(tx *sqlx.Tx) *Query[T] {
	nextQ := &Query[T]{}
	nextQ.Next = func() statement.Type {
		q.Next()
		q.ctx = context.WithValue(q.ctx, pkg.CtxKeyTx, tx)
		nextQ.ctx = q.ctx
		return statement.Nop
	}
	return nextQ
}

// WithDB 使用db实例
func (q *Query[T]) WithDB(db *sqlx.DB) *Query[T] {
	nextQ := &Query[T]{}
	nextQ.Next = func() statement.Type {
		q.Next()
		q.ctx = context.WithValue(q.ctx, pkg.CtxKeyDB, db)
		nextQ.ctx = q.ctx
		return statement.Nop
	}
	return nextQ
}

// WithDBProvider 设置db
func (q *Query[T]) WithDBProvider(p dbprovider.IDBProvider) *Query[T] {
	nextQ := &Query[T]{}
	nextQ.Next = func() statement.Type {
		q.Next()
		q.ctx = context.WithValue(q.ctx, pkg.CtxKeyDBProvider, p)
		nextQ.ctx = q.ctx
		return statement.Nop
	}
	return nextQ
}
