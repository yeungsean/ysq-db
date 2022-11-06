package ysqdb

import (
	"github.com/yeungsean/ysq"
	"github.com/yeungsean/ysq-db/pkg/field"
	"github.com/yeungsean/ysq-db/pkg/order"
	"github.com/yeungsean/ysq-db/pkg/statement"
)

func columnTypeToOrder(t field.Type) *order.Order {
	return order.NewOrder(t, order.Asc)
}

func columnTypeToOrderDesc(t field.Type) *order.Order {
	return order.NewOrder(t, order.Desc)
}

// OrderAsc ORDER BY $COLUMN ASC
func (q *Query[T]) OrderAsc(cols ...field.Type) *Query[T] {
	return q.wrap(func(q *Query[T], lc *queryContext[T]) func() statement.Type {
		fields := ysq.Select(ysq.FromSlice(cols), columnTypeToOrder).ToSlice(len(cols))
		lc.orders = append(lc.orders, fields...)
		return statementSort
	})
}

// OrderDesc ORDER BY $COLUMN DESC
func (q *Query[T]) OrderDesc(cols ...field.Type) *Query[T] {
	return q.wrap(func(q *Query[T], lc *queryContext[T]) func() statement.Type {
		fields := ysq.Select(ysq.FromSlice(cols), columnTypeToOrderDesc).ToSlice(len(cols))
		lc.orders = append(lc.orders, fields...)
		return statementSort
	})
}
