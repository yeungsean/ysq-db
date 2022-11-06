package str

import (
	"github.com/yeungsean/ysq-db/pkg/common"
	"github.com/yeungsean/ysq-db/pkg/join"
	"github.com/yeungsean/ysq-db/pkg/statement"
)

type tableExpr[T string] struct {
	table T
	alias string
}

type joinExpr[T string] struct {
	tableExpr[T]
	jt        join.Type
	condition string
}

func (q *Query[T]) join(table T, condition string, jt join.Type, tableAliasOpt ...string) *Query[T] {
	return q.wrap(
		func(q *Query[T], qc *queryContext[T]) func() statement.Type {
			qc.joins = append(qc.joins, joinExpr[T]{
				tableExpr: tableExpr[T]{
					table: table,
					alias: common.VarArgGetFirst(tableAliasOpt...),
				},
				jt:        jt,
				condition: condition,
			})
			return func() statement.Type {
				return statement.Join
			}
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
