package ysqdb

import (
	"github.com/yeungsean/ysq-db/pkg"
	"github.com/yeungsean/ysq-db/pkg/column"
	"github.com/yeungsean/ysq-db/pkg/cond"
	"github.com/yeungsean/ysq-db/pkg/field"
	"github.com/yeungsean/ysq-db/pkg/join"
	"github.com/yeungsean/ysq-db/pkg/ops"
	"github.com/yeungsean/ysq-db/pkg/statement"
)

// LeftJoin ...
func (q *Query[T]) LeftJoin(entity pkg.ITable, leftColumn, rightColumn field.Type) *Query[T] {
	return q.wrap(func(q *Query[T], lc *queryContext[T]) func() statement.Type {
		je := &joinExpr[T]{
			jt: join.Left,
			jc: make([]*cond.Cond, 0, 1),
		}
		cc := cond.All().Add2(
			column.NewField(leftColumn),
			column.NewField(rightColumn),
			ops.EQ,
		)
		je.jc = append(je.jc, cc)
		lc.joinEntites = append(lc.joinEntites, je)
		return statementJoin
	})
}

// RightJoin ...
func (q *Query[T]) RightJoin(entity pkg.ITable, leftColumn, rightColumn field.Type) *Query[T] {
	return q.wrap(func(q *Query[T], lc *queryContext[T]) func() statement.Type {
		lc.joinEntites = append(lc.joinEntites, &joinExpr[T]{
			jt: join.Right,
		})
		return statementJoin
	})
}

// InnerJoin ...
func (q *Query[T]) InnerJoin(entity pkg.ITable, leftColumn, rightColumn field.Type) *Query[T] {
	return q.wrap(func(q *Query[T], lc *queryContext[T]) func() statement.Type {
		lc.joinEntites = append(lc.joinEntites, &joinExpr[T]{
			jt: join.Inner,
		})
		return statementJoin
	})
}
