package str

import "sync/atomic"

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
