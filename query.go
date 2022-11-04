package ysqdb

import (
	"context"
	"strings"
	"sync/atomic"

	"github.com/yeungsean/ysq"
	"github.com/yeungsean/ysq-db/pkg"
	"github.com/yeungsean/ysq-db/pkg/column"
	"github.com/yeungsean/ysq-db/pkg/common"
	"github.com/yeungsean/ysq-db/pkg/cond"
	"github.com/yeungsean/ysq-db/pkg/field"
	"github.com/yeungsean/ysq-db/pkg/join"
	"github.com/yeungsean/ysq-db/pkg/ops"
	"github.com/yeungsean/ysq-db/pkg/order"
	"github.com/yeungsean/ysq-db/pkg/statement"
	"github.com/yeungsean/ysq/pkg/delegate"
)

// QueryType Query type for 4 statements
type QueryType uint8

const (
	// QuerySelect select stmt
	QuerySelect QueryType = iota
	// QueryInsert
	// QueryUpdate
	// QueryDelete
)

// Query Parse analysis turns all statements into a Query tree
type Query[T pkg.ITable] struct {
	ctx      context.Context
	buildCnt int32
	Next     func() Iterator[T]
}

// NewQuery ...
func NewQuery[T pkg.ITable](entity T, opts ...LambdaOption[T]) *Query[T] {
	q := &Query[T]{
		ctx: context.TODO(),
	}
	q.Next = func() Iterator[T] {
		lambdaCtx := newlambdaContext(entity)
		common.OptionForEach(lambdaCtx, opts)
		q.ctx = context.WithValue(q.ctx, ctxKeyLambda, lambdaCtx)
		return func() statement.Type {
			return statement.Nop
		}
	}
	return q
}

// Context ...
func (q *Query[T]) Context(ctx context.Context) *Query[T] {
	q.ctx = ctx
	return q
}

func (q *Query[T]) ctxGetLambda() *lambdaContext[T] {
	return q.ctx.Value(ctxKeyLambda).(*lambdaContext[T])
}

// Column 查询字段，带默认值
func (q *Query[T]) Column(col field.Type, defaultValue any) *Query[T] {
	return q.wrap(func(q *Query[T], lc *lambdaContext[T]) func() statement.Type {
		lc.Fields = append(lc.Fields, column.NewField(col, column.WithDefaultValue(defaultValue)))
		return statementColumn
	})
}

func columnTypeToColumn(t field.Type) *column.Column {
	return column.NewField(t)
}

// Select ...
func (q *Query[T]) Select(cols ...field.Type) *Query[T] {
	return q.wrap(func(q *Query[T], lc *lambdaContext[T]) func() statement.Type {
		fields := ysq.Select(ysq.FromSlice(cols), columnTypeToColumn).ToSlice(len(cols))
		lc.Fields = append(lc.Fields, fields...)
		return statementColumn
	})
}

// Limit ...
func (q *Query[T]) Limit(count int, offset ...int) *Query[T] {
	return q.wrap(func(q *Query[T], lc *lambdaContext[T]) func() statement.Type {
		lc.LimitCount = count
		lc.LimitOffset = common.VarArgGetFirst(offset...)
		return func() statement.Type {
			return statement.Limit
		}
	})
}

// GroupBy 分组
func (q *Query[T]) GroupBy(cols ...field.Type) *Query[T] {
	return q.wrap(func(q *Query[T], lc *lambdaContext[T]) func() statement.Type {
		fields := ysq.Select(ysq.FromSlice(cols), columnTypeToColumn).ToSlice(len(cols))
		lc.Groups = append(lc.Groups, fields...)
		return func() statement.Type {
			return statement.GroupBy
		}
	})
}

// Having 分组查询后的过滤
func (q *Query[T]) Having(col field.Type, value any, opTypes ...ops.Type) *Query[T] {
	return q.wrap(func(q *Query[T], lc *lambdaContext[T]) func() statement.Type {
		ot := common.VarArgGetFirst(opTypes...)
		lc.HavingClause.Add(column.New(col,
			column.WithOp(ot),
			column.WithValue(value),
			column.WithIsFilter(true),
		))
		return func() statement.Type {
			return statement.Having
		}
	})
}

// Where ...
func (q *Query[T]) Where(conds ...*cond.Cond) *Query[T] {
	return q.wrap(func(q *Query[T], lc *lambdaContext[T]) func() statement.Type {
		lc.WhereClause.AddChildren(conds...)
		return statementWhere
	})
}

func (q *Query[T]) getColumn(f field.Type, value any, opt ops.Type) *column.Column {
	return column.New(f,
		column.WithIsFilter(true),
		column.WithOp(opt),
		column.WithValue(value),
	)
}

// IsNull 是否为null
func (q *Query[T]) IsNull(col field.Type, lts ...cond.LogicType) *Query[T] {
	return q.wrap(func(q *Query[T], lc *lambdaContext[T]) func() statement.Type {
		lc.WhereClause.Add(column.New(col).IsNull(), lts...)
		return statementWhere
	})
}

// IsNotNull 是否为not null
func (q *Query[T]) IsNotNull(col field.Type, lts ...cond.LogicType) *Query[T] {
	return q.wrap(func(q *Query[T], lc *lambdaContext[T]) func() statement.Type {
		lc.WhereClause.Add(column.New(col).IsNotNull(), lts...)
		return statementWhere
	})
}

// Like ...
func (q *Query[T]) Like(col field.Type, value any, lts ...cond.LogicType) *Query[T] {
	return q.wrap(func(q *Query[T], lc *lambdaContext[T]) func() statement.Type {
		lc.WhereClause.Add(column.New(col).Like(value), lts...)
		return statementWhere
	})
}

func columnTypeToOrder(t field.Type) *order.Order {
	return order.NewOrder(t, order.Asc)
}

func columnTypeToOrderDesc(t field.Type) *order.Order {
	return order.NewOrder(t, order.Desc)
}

// OrderAsc ORDER BY $COLUMN ASC
func (q *Query[T]) OrderAsc(cols ...field.Type) *Query[T] {
	return q.wrap(func(q *Query[T], lc *lambdaContext[T]) func() statement.Type {
		fields := ysq.Select(ysq.FromSlice(cols), columnTypeToOrder).ToSlice(len(cols))
		lc.orders = append(lc.orders, fields...)
		return statementSort
	})
}

// OrderDesc ORDER BY $COLUMN DESC
func (q *Query[T]) OrderDesc(cols ...field.Type) *Query[T] {
	return q.wrap(func(q *Query[T], lc *lambdaContext[T]) func() statement.Type {
		fields := ysq.Select(ysq.FromSlice(cols), columnTypeToOrderDesc).ToSlice(len(cols))
		lc.orders = append(lc.orders, fields...)
		return statementSort
	})
}

// Join https://www.tektutorialshub.com/entity-framework/join-query-entity-framework/
func Join[TLeft, TRight, TResult pkg.ITable, TKey field.Type](
	left *Query[TLeft],
	right *Query[TRight],
	leftKeySelector delegate.Func1[TLeft, TKey],
	rightKeySelector delegate.Func1[TRight, TKey],
	jtArgs ...join.Type,
) *Query[TResult] {
	return nil
}

// Join ...
func (q *Query[T]) Join(entity pkg.ITable, leftColumn, rightColumn field.Type, jtArgs ...join.Type) *Query[T] {
	return q
}

func (q *Query[T]) build() {
	if atomic.CompareAndSwapInt32(&q.buildCnt, 0, 1) {
		q.Next()
	}
}

// Values ...
func (q *Query[T]) Values() []any {
	q.build()
	return q.ctxGetLambda().Values
}

// String ...
func (q *Query[T]) String() string {
	q.build()
	qlCtx := q.ctxGetLambda()
	sb := strings.Builder{}
	sb.Grow(25)
	sb.WriteString("SELECT ")
	fields := ysq.FromSlice(qlCtx.Fields).CastToStringBy(func(c *column.Column) string {
		return c.String()
	}).ToSlice(len(qlCtx.Fields))

	sb.WriteString(strings.Join(fields, ","))
	sb.WriteString(" FROM ")
	sb.WriteString(qlCtx.Entity.String())
	if wStr := qlCtx.WhereClause.String(); wStr != "" {
		sb.WriteString(" WHERE ")
		sb.WriteString(wStr)
	}

	if len(qlCtx.orders) == 0 {
		return sb.String()
	} else if len(qlCtx.orders) > 0 {
		sb.WriteString(" ORDER BY ")
	}

	orders := ysq.FromSlice(qlCtx.orders).
		CastToStringBy(func(o *order.Order) string { return o.String() }).
		Reduce("", func(s1, s2 string) string {
			if s1 == "" {
				return s2
			}
			return s1 + "," + s2
		})
	sb.WriteString(orders)
	return sb.String()
}
