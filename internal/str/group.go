package str

import (
	"github.com/yeungsean/ysq"
	"github.com/yeungsean/ysq-db/internal/expr/column"
	"github.com/yeungsean/ysq-db/internal/expr/common"
	"github.com/yeungsean/ysq-db/internal/expr/cond"
	"github.com/yeungsean/ysq-db/internal/expr/field"
	"github.com/yeungsean/ysq-db/internal/expr/ops"
	"github.com/yeungsean/ysq-db/internal/expr/statement"
)

// GroupBy 分组
func (q *Query[T]) GroupBy(cols ...field.Type) *Query[T] {
	return q.wrap(func(q *Query[T], lc *queryContext[T]) statement.Type {
		fields := ysq.Select(ysq.FromSlice(cols),
			func(t field.Type) *field.Field { return &field.Field{Name: t} }).
			ToSlice(len(cols))
		lc.Groups = append(lc.Groups, fields...)
		return statement.GroupBy
	})
}

// HavingOr 分组查询后的过滤
func (q *Query[T]) HavingOr(col field.Type, value any, opTypes ...ops.Type) *Query[T] {
	return q.wrap(func(q *Query[T], lc *queryContext[T]) statement.Type {
		ot := common.VarArgGetFirst(opTypes...)
		lc.HavingClause.Add(column.New(col,
			column.WithOp(ot),
			column.WithValue(value),
		), cond.Or)
		return statement.Having
	})
}

// HavingAnd 分组查询后的过滤
func (q *Query[T]) HavingAnd(col field.Type, value any, opTypes ...ops.Type) *Query[T] {
	return q.wrap(func(q *Query[T], lc *queryContext[T]) statement.Type {
		ot := common.VarArgGetFirst(opTypes...)
		lc.HavingClause.Add(column.New(col,
			column.WithOp(ot),
			column.WithValue(value),
		), cond.And)
		return statement.Having
	})
}
