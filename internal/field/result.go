package ysqdb

import (
	"strings"

	"github.com/yeungsean/ysq"
	"github.com/yeungsean/ysq-db/pkg/column"
	"github.com/yeungsean/ysq-db/pkg/order"
)

// ToSlice ...
func (q *Query[T]) ToSlice(length ...int) ([]T, error) {
	return nil, nil
}

// One ...
func (q *Query[T]) One() (T, error) {
	var zero T
	return zero, nil
}

// ScanStruct ...
func (q *Query[T]) ScanStruct(result *T) error {
	return nil
}

// ScanSlice ...
func (q *Query[T]) ScanSlice(result []*T) error {
	return nil
}

// Count 总数
func (q *Query[T]) Count() (int64, error) {
	return 0, nil
}

// Sum 总和
func (q *Query[T]) Sum() (float64, error) {
	return 0, nil
}

// Avg 平均数
func (q *Query[T]) Avg() (float64, error) {
	return 0, nil
}

// Max 最大数
func (q *Query[T]) Max() (T, error) {
	var zero T
	return zero, nil
}

// Min 最大数
func (q *Query[T]) Min() (T, error) {
	var zero T
	return zero, nil
}

// Exists 是否存在
func (q *Query[T]) Exists() (bool, error) {
	q.build()
	qlCtx := q.ctxGetLambda()
	qlCtx.Fields = []*column.Column{column.New(`1`)}
	return false, nil
}

// String ...
func (q *Query[T]) String() string {
	q.build()
	qlCtx := q.ctxGetLambda()
	sb := strings.Builder{}
	sb.Grow(25)
	sb.WriteString("SELECT ")
	fields := ysq.FromSlice(qlCtx.Fields).
		CastToStringBy(func(c *column.Column) string { return c.String() }).
		ToSlice(len(qlCtx.Fields))

	sb.WriteString(strings.Join(fields, ","))
	sb.WriteString(" FROM ")
	sb.WriteString(qlCtx.mainEntity.String())

	if len(qlCtx.joinEntites) > 0 {

	}

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
		ToSlice(len(qlCtx.orders))
	sb.WriteString(strings.Join(orders, ","))
	return sb.String()
}
