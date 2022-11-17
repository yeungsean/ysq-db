package str

import (
	"github.com/yeungsean/ysq-db/internal/expr/common"
	"github.com/yeungsean/ysq-db/internal/expr/join"
	"github.com/yeungsean/ysq-db/internal/expr/statement"
	"github.com/yeungsean/ysq-db/internal/expr/table"
	"github.com/yeungsean/ysq-db/pkg"
)

func (q *Query[T]) join(tb T, condition string, jt join.Type, opts ...pkg.Options) *Query[T] {
	return q.wrap(
		func(q *Query[T], qc *queryContext[T]) statement.Type {
			var opt pkg.Option
			common.OptionForEach(&opt, opts)
			qc.joins = append(qc.joins, join.Expr[T]{
				Expr: table.Expr[T]{
					Table:  tb,
					Option: opt,
				},
				Type:      jt,
				Condition: condition,
			})
			return statement.Join
		},
	)
}

// LeftJoin 左连接
func (q *Query[T]) LeftJoin(table T, condition string, opts ...pkg.Options) *Query[T] {
	return q.join(table, condition, join.Left, opts...)
}

// RightJoin 右连接
func (q *Query[T]) RightJoin(table T, condition string, opts ...pkg.Options) *Query[T] {
	return q.join(table, condition, join.Right, opts...)
}

// InnerJoin 内连接
func (q *Query[T]) InnerJoin(table T, condition string, opts ...pkg.Options) *Query[T] {
	return q.join(table, condition, join.Inner, opts...)
}
