package str

import (
	"github.com/yeungsean/ysq-db/internal/expr/column"
	"github.com/yeungsean/ysq-db/internal/expr/cond"
	"github.com/yeungsean/ysq-db/internal/expr/ops"
	"github.com/yeungsean/ysq-db/internal/expr/statement"
	"github.com/yeungsean/ysq-db/pkg/field"
)

func (q *Query[T]) wrapStatementWhere(ft field.Type, op ops.Type, value any, lt cond.LogicType,
	opts ...column.Option) *Query[T] {
	return q.wrap(func(iq *Query[T], qc *queryContext[T]) statement.Type {
		col := buildColumn(ft, value, op, opts...)
		qc.WhereClause.Add(col, lt)
		if value == nil {
			return statement.Where
		}
		switch v := value.(type) {
		case []any:
			qc.Args = append(qc.Args, v...)
		case *column.Column:
		case *Query[T]:
			panic("Unsupport")
		default:
			qc.Args = append(qc.Args, value)
		}
		return statement.Where
	})
}

func (q *Query[T]) wrapStatementWhereAnd(f field.Type, op ops.Type, value any, opts ...column.Option) *Query[T] {
	return q.wrapStatementWhere(f, op, value, cond.And, opts...)
}

func (q *Query[T]) wrapStatementWhereOr(f field.Type, op ops.Type, value any, opts ...column.Option) *Query[T] {
	return q.wrapStatementWhere(f, op, value, cond.Or, opts...)
}

// AndIn ...
func (q *Query[T]) AndIn(f field.Type, value []any, opts ...column.Option) *Query[T] {
	return q.wrapStatementWhereAnd(f, ops.In, value, opts...)
}

// OrIn ...
func (q *Query[T]) OrIn(f field.Type, value []any, opts ...column.Option) *Query[T] {
	return q.wrapStatementWhereOr(f, ops.In, value, opts...)
}

// AndNotIn ...
func (q *Query[T]) AndNotIn(f field.Type, value []any, opts ...column.Option) *Query[T] {
	return q.wrapStatementWhereAnd(f, ops.NotIn, value, opts...)
}

// OrNotIn ...
func (q *Query[T]) OrNotIn(f field.Type, value []any, opts ...column.Option) *Query[T] {
	return q.wrapStatementWhereOr(f, ops.NotIn, value, opts...)
}

func buildColumn(f field.Type, value any, op ops.Type, opts ...column.Option) *column.Column {
	opts = append(opts, column.WithOp(op))
	if op != ops.IsNull && op != ops.IsNotNull {
		opts = append(opts, column.WithValue(value))
	}
	return column.New(f, opts...)
}

// AndEqual =
func (q *Query[T]) AndEqual(f field.Type, value any, opts ...column.Option) *Query[T] {
	return q.wrapStatementWhereAnd(f, ops.EQ, value, opts...)
}

// OrEqual =
func (q *Query[T]) OrEqual(f field.Type, value any, opts ...column.Option) *Query[T] {
	return q.wrapStatementWhereOr(f, ops.EQ, value, opts...)
}

// AndNotEqual <>
func (q *Query[T]) AndNotEqual(f field.Type, value any, opts ...column.Option) *Query[T] {
	return q.wrapStatementWhereAnd(f, ops.NEQ, value, opts...)
}

// OrNotEqual <>
func (q *Query[T]) OrNotEqual(f field.Type, value any, opts ...column.Option) *Query[T] {
	return q.wrapStatementWhereOr(f, ops.NEQ, value, opts...)
}

// AndGreater >
func (q *Query[T]) AndGreater(f field.Type, value any, opts ...column.Option) *Query[T] {
	return q.wrapStatementWhereAnd(f, ops.GT, value, opts...)
}

// OrGreater >
func (q *Query[T]) OrGreater(f field.Type, value any, opts ...column.Option) *Query[T] {
	return q.wrapStatementWhereOr(f, ops.GT, value, opts...)
}

// AndGreaterOrEqual >=
func (q *Query[T]) AndGreaterOrEqual(f field.Type, value any, opts ...column.Option) *Query[T] {
	return q.wrapStatementWhereAnd(f, ops.GTE, value, opts...)
}

// OrGreaterOrEqual >=
func (q *Query[T]) OrGreaterOrEqual(f field.Type, value any, opts ...column.Option) *Query[T] {
	return q.wrapStatementWhereOr(f, ops.GTE, value, opts...)
}

// AndLess <
func (q *Query[T]) AndLess(f field.Type, value any, opts ...column.Option) *Query[T] {
	return q.wrapStatementWhereAnd(f, ops.LT, value, opts...)
}

// OrLess <
func (q *Query[T]) OrLess(f field.Type, value any, opts ...column.Option) *Query[T] {
	return q.wrapStatementWhereOr(f, ops.LT, value, opts...)
}

// AndLessOrEqual <=
func (q *Query[T]) AndLessOrEqual(f field.Type, value any, opts ...column.Option) *Query[T] {
	return q.wrapStatementWhereAnd(f, ops.LTE, value, opts...)
}

// OrLessOrEqual <=
func (q *Query[T]) OrLessOrEqual(f field.Type, value any, opts ...column.Option) *Query[T] {
	return q.wrapStatementWhereOr(f, ops.LTE, value, opts...)
}

// AndIsNull 是否为null
func (q *Query[T]) AndIsNull(f field.Type, opts ...column.Option) *Query[T] {
	return q.wrapStatementWhereAnd(f, ops.IsNull, nil, opts...)
}

// OrIsNull 是否为null
func (q *Query[T]) OrIsNull(f field.Type, opts ...column.Option) *Query[T] {
	return q.wrapStatementWhereOr(f, ops.IsNull, nil, opts...)
}

// AndIsNotNull 是否为not null
func (q *Query[T]) AndIsNotNull(f field.Type, opts ...column.Option) *Query[T] {
	return q.wrapStatementWhereAnd(f, ops.IsNotNull, nil, opts...)
}

// OrIsNotNull 是否为not null
func (q *Query[T]) OrIsNotNull(f field.Type, opts ...column.Option) *Query[T] {
	return q.wrapStatementWhereOr(f, ops.IsNotNull, nil, opts...)
}

// AndLike AND LIKE ..
func (q *Query[T]) AndLike(f field.Type, value any, opts ...column.Option) *Query[T] {
	return q.wrapStatementWhereAnd(f, ops.Like, value, opts...)
}

// OrLike OR LIKE ..
func (q *Query[T]) OrLike(f field.Type, value any, opts ...column.Option) *Query[T] {
	return q.wrapStatementWhereOr(f, ops.Like, value, opts...)
}

// AndBetween AND BETWEEN .. AND ..
func (q *Query[T]) AndBetween(f field.Type, min, max any, opts ...column.Option) *Query[T] {
	return q.wrapStatementWhereAnd(f, ops.Between, []any{min, max}, opts...)
}

// OrBetween OR BETWEEN .. AND ..
func (q *Query[T]) OrBetween(f field.Type, min, max any, opts ...column.Option) *Query[T] {
	return q.wrapStatementWhereOr(f, ops.Between, []any{min, max}, opts...)
}

// AndNotBetween AND NOT BETWEEN .. AND ..
func (q *Query[T]) AndNotBetween(f field.Type, min, max any, opts ...column.Option) *Query[T] {
	return q.wrapStatementWhereAnd(f, ops.NotBetween, []any{min, max}, opts...)
}

// OrNotBetween OR NOT BETWEEN .. AND ..
func (q *Query[T]) OrNotBetween(f field.Type, min, max any, opts ...column.Option) *Query[T] {
	return q.wrapStatementWhereOr(f, ops.NotBetween, []any{min, max}, opts...)
}

// Where ...
func (q *Query[T]) Where(conds ...*cond.Cond) *Query[T] {
	return q.wrap(func(q *Query[T], lc *queryContext[T]) statement.Type {
		lc.WhereClause.AddChildren(conds...)
		return statement.Where
	})
}
