package str

import "github.com/yeungsean/ysq-db/internal/expr/statement"

// Limit ...
func (q *Query[T]) Limit(limit int) *Query[T] {
	return q.wrap(
		func(q *Query[T], qc *queryContext[T]) statement.Type {
			qc.LimitCount = limit
			return statement.Limit
		},
	)
}

// Offset ...
func (q *Query[T]) Offset(offset int) *Query[T] {
	return q.wrap(
		func(q *Query[T], qc *queryContext[T]) statement.Type {
			qc.LimitOffset = offset
			return statement.Limit
		},
	)
}

// LimitOffset ...
func (q *Query[T]) LimitOffset(limit, offset int) *Query[T] {
	return q.wrap(
		func(q *Query[T], qc *queryContext[T]) statement.Type {
			qc.LimitCount = limit
			qc.LimitOffset = offset
			return statement.Limit
		},
	)
}
