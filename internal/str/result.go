package str

import (
	"fmt"
	"strings"
	"sync/atomic"

	"github.com/yeungsean/ysq"
	"github.com/yeungsean/ysq-db/pkg/column"
	"github.com/yeungsean/ysq-db/pkg/field"
	"github.com/yeungsean/ysq-db/pkg/order"
)

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
	fields := ysq.FromSlice(qlCtx.Fields).
		CastToStringBy(func(c *field.Field) string { return c.SelectField() }).
		ToSlice(len(qlCtx.Fields))
	sb := strings.Builder{}
	fieldStr := strings.Join(fields, ",")
	if fieldStr == "" {
		fieldStr = "*"
	}
	sb.Grow(13 + len(fieldStr))
	sb.WriteString("SELECT ")
	sb.WriteString(fieldStr)
	sb.WriteString(" FROM ")
	sb.WriteString(qlCtx.mainTable.String())

	if len(qlCtx.joins) > 0 {
		for _, join := range qlCtx.joins {
			sb.WriteString(join.String())
		}
	}

	if wStr := qlCtx.WhereClause.String(); wStr != "" {
		sb.Grow(7 + len(wStr))
		sb.WriteString(" WHERE ")
		sb.WriteString(wStr)
	}

	if len(qlCtx.Groups) > 0 {
		groupFields := ysq.FromSlice(qlCtx.Groups).
			CastToStringBy(func(c *column.Column) string { return c.String() }).
			ToSlice(len(qlCtx.Groups))
		tmp := strings.Join(groupFields, ",")
		sb.Grow(10 + len(tmp))
		sb.WriteString(" GROUP BY ")
		sb.WriteString(tmp)
	}

	if hStr := qlCtx.HavingClause.String(); hStr != "" {
		sb.Grow(8 + len(hStr))
		sb.WriteString(" HAVING ")
		sb.WriteString(hStr)
	}

	if len(qlCtx.orders) > 0 {
		orderFields := ysq.FromSlice(qlCtx.orders).
			CastToStringBy(func(c *order.Order) string { return c.String() }).
			ToSlice(len(qlCtx.orders))
		tmp := strings.Join(orderFields, ",")
		sb.Grow(10 + len(tmp))
		sb.WriteString(" ORDER BY ")
		sb.WriteString(tmp)
	}

	if qlCtx.LimitCount > 0 {
		sb.WriteString(fmt.Sprintf(" LIMIT %d", qlCtx.LimitCount))
	}

	if qlCtx.LimitOffset > 0 {
		sb.WriteString(fmt.Sprintf(" OFFSET %d", qlCtx.LimitOffset))
	}
	return sb.String()
}
