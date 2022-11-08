package str

import (
	"github.com/yeungsean/ysq"
	"github.com/yeungsean/ysq-db/pkg/column"
	"github.com/yeungsean/ysq-db/pkg/common"
	"github.com/yeungsean/ysq-db/pkg/cond"
	"github.com/yeungsean/ysq-db/pkg/field"
	"github.com/yeungsean/ysq-db/pkg/ops"
	"github.com/yeungsean/ysq-db/pkg/statement"
)

// GroupBy 分组
func (q *Query[T]) GroupBy(cols ...field.Type) *Query[T] {
	return q.wrap(func(q *Query[T], lc *queryContext[T]) statement.Type {
		fields := ysq.Select(ysq.FromSlice(cols),
			func(t field.Type) *column.Column {
				return column.NewField(t)
			}).ToSlice(len(cols))
		lc.Groups = append(lc.Groups, fields...)
		return statement.GroupBy
	})
}

// Having 分组查询后的过滤
func (q *Query[T]) Having(col field.Type, value any, opTypes ...ops.Type) *Query[T] {
	return q.wrap(func(q *Query[T], lc *queryContext[T]) statement.Type {
		ot := common.VarArgGetFirst(opTypes...)
		lc.HavingClause.Add(column.New(col,
			column.WithOp(ot),
			column.WithValue(value),
			column.WithIsFilter(true),
		), cond.And)
		return statement.Having
	})
}
