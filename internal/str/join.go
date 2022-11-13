package str

import (
	"github.com/yeungsean/ysq-db/internal/expr/common"
	"github.com/yeungsean/ysq-db/internal/expr/join"
	"github.com/yeungsean/ysq-db/internal/expr/statement"
	"github.com/yeungsean/ysq-db/internal/expr/table"
)

func (q *Query[T]) join(tb T, condition string, jt join.Type, tableAliasOpt ...string) *Query[T] {
	return q.wrap(
		func(q *Query[T], qc *queryContext[T]) statement.Type {
			qc.joins = append(qc.joins, join.Expr[T]{
				Expr: table.Expr[T]{
					Table: tb,
					Alias: common.VarArgGetFirst(tableAliasOpt...),
				},
				Type:      jt,
				Condition: condition,
			})
			return statement.Join
		},
	)
}

// LeftJoin 左连接
func (q *Query[T]) LeftJoin(table T, condition string, tableAliasOpt ...string) *Query[T] {
	return q.join(table, condition, join.Left, tableAliasOpt...)
}

// RightJoin 右连接
func (q *Query[T]) RightJoin(table T, condition string, tableAliasOpt ...string) *Query[T] {
	return q.join(table, condition, join.Right, tableAliasOpt...)
}

// InnerJoin 内连接
func (q *Query[T]) InnerJoin(table T, condition string, tableAliasOpt ...string) *Query[T] {
	return q.join(table, condition, join.Inner, tableAliasOpt...)
}
