package str

import "github.com/yeungsean/ysq-db/pkg/statement"

// Limit ...
func (q *Query[T]) Limit(limit int) *Query[T] {
	return q.wrap(
		func(q *Query[T], qc *queryContext[T]) func() statement.Type {
			qc.LimitCount = limit
			return func() statement.Type {
				return statement.Limit
			}
		},
	)
}

// Offset ...
func (q *Query[T]) Offset(offset int) *Query[T] {
	return q.wrap(
		func(q *Query[T], qc *queryContext[T]) func() statement.Type {
			qc.LimitOffset = offset
			return func() statement.Type {
				return statement.Limit
			}
		},
	)
}

// LimitOffset ...
func (q *Query[T]) LimitOffset(limit, offset int) *Query[T] {
	return q.wrap(
		func(q *Query[T], qc *queryContext[T]) func() statement.Type {
			qc.LimitCount = limit
			qc.LimitOffset = offset
			return func() statement.Type {
				return statement.Limit
			}
		},
	)
}
