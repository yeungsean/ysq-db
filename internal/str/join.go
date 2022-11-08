package str

import (
	"fmt"
	"strings"

	"github.com/yeungsean/ysq-db/pkg/common"
	"github.com/yeungsean/ysq-db/pkg/join"
	"github.com/yeungsean/ysq-db/pkg/statement"
)

type tableExpr[T string] struct {
	table T
	alias string
}

// String ...
func (t tableExpr[T]) String() string {
	if t.alias == "" {
		return string(t.table)
	}
	return fmt.Sprintf("%s AS %s", t.table, t.alias)
}

type joinExpr[T string] struct {
	tableExpr[T]
	jt        join.Type
	condition string
}

// String ...
func (j joinExpr[T]) String() string {
	sb := strings.Builder{}
	tb := j.tableExpr.String()
	sb.Grow(len(tb) + 10 + len(j.condition) + len(j.jt))
	sb.WriteString(string(j.jt))
	sb.WriteString(" JOIN ")
	sb.WriteString(tb)
	sb.WriteString(" ON ")
	sb.WriteString(j.condition)
	return sb.String()
}

func (q *Query[T]) join(table T, condition string, jt join.Type, tableAliasOpt ...string) *Query[T] {
	return q.wrap(
		func(q *Query[T], qc *queryContext[T]) statement.Type {
			qc.joins = append(qc.joins, joinExpr[T]{
				tableExpr: tableExpr[T]{
					table: table,
					alias: common.VarArgGetFirst(tableAliasOpt...),
				},
				jt:        jt,
				condition: condition,
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
