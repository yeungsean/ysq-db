package str

import (
	"fmt"
	"strings"
	"sync/atomic"

	"github.com/yeungsean/ysq"
	"github.com/yeungsean/ysq-db/internal/expr/order"
	"github.com/yeungsean/ysq-db/pkg"
)

func (q *Query[T]) build() {
	if atomic.CompareAndSwapInt32(&q.buildCnt, 0, 1) {
		q.Next()
	}
}

// Args sql语句需要的所有参数
func (q *Query[T]) Args() []any {
	q.build()
	return q.ctxGetLambda().Args
}

// String 获取sql语句
func (q *Query[T]) String() string {
	q.build()
	qlCtx, provider := q.ctxGetLambda(), q.ctxGetProvider()
	q.ctx = pkg.CtxResetFilterColumnIndex(q.ctx)
	var fieldStr string
	if len(qlCtx.Fields) > 0 {
		fields := provider.SelectFields(qlCtx.Fields...)
		fieldStr = strings.Join(fields, ",")
	} else {
		fieldStr = "*"
	}

	sb := strings.Builder{}
	sb.Grow(13 + len(fieldStr))
	sb.WriteString("SELECT ")
	sb.WriteString(fieldStr)
	sb.WriteString(" FROM ")
	sb.WriteString(qlCtx.mainTable.String(q.ctx))

	if len(qlCtx.joins) > 0 {
		for _, join := range qlCtx.joins {
			sb.WriteString(join.String(q.ctx))
		}
	}

	if wStr := qlCtx.WhereClause.String(q.ctx); wStr != "" {
		sb.Grow(7 + len(wStr))
		sb.WriteString(" WHERE ")
		sb.WriteString(wStr)
	}

	if len(qlCtx.Groups) > 0 {
		groupFields := provider.OtherTypeFields(qlCtx.Groups...)
		tmp := strings.Join(groupFields, ",")
		sb.Grow(10 + len(tmp))
		sb.WriteString(" GROUP BY ")
		sb.WriteString(tmp)
	}

	if hStr := qlCtx.HavingClause.String(q.ctx); hStr != "" {
		sb.Grow(8 + len(hStr))
		sb.WriteString(" HAVING ")
		sb.WriteString(hStr)
	}

	if len(qlCtx.orders) > 0 {
		orderFields := ysq.FromSlice(qlCtx.orders).
			CastToStringBy(func(c *order.Order) string { return c.String(q.ctx) }).
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

	if qlCtx.HasForUpdate {
		sb.WriteString(" FOR UPDATE")
	}

	return sb.String()
}
