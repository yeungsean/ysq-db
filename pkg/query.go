package pkg

import (
	"github.com/yeungsean/ysq-db/pkg/column"
	"github.com/yeungsean/ysq-db/pkg/node"
	"github.com/yeungsean/ysq-db/pkg/ops"
)

// QueryType Query type for 4 statements
type QueryType uint8

const (
	QuerySelect QueryType = iota
	QueryInsert
	QueryUpdate
	QueryDelete
)

// Query Parse analysis turns all statements into a Query tree
type Query[T any] struct {
	Type         QueryType
	HasForUpdate bool

	FieldList      *node.List
	JoinList       *node.List
	GroupClause    *node.List
	SortClause     *node.List
	DistinctClause *node.List
	HavingClause   *node.List

	LimitOffset uint
	LimitCount  uint
}

// Column ...
func (q *Query[T]) Column(col column.Type, defaultValue interface{}) *Query[T] {
	return q
}

// Select ...
func (q *Query[T]) Select(columns ...column.Type) *Query[T] {
	return q
}

// Limit ...
func (q *Query[T]) Limit(count int, offset ...int) *Query[T] {
	return q
}

// GroupBy ...
func (q *Query[T]) GroupBy(columns ...column.Type) *Query[T] {
	return q
}

// Having ...
func (q *Query[T]) Having(columns ...column.Type) *Query[T] {
	return q
}

// Where ...
func (q *Query[T]) Where(col column.Type, value interface{}, opType ...ops.Type) *Query[T] {
	return q
}

// IsNull ...
func (q *Query[T]) IsNull(col column.Type) *Query[T] {
	return q
}

// IsNotNull ...
func (q *Query[T]) IsNotNull(col column.Type) *Query[T] {
	return q
}

// Like ...
func (q *Query[T]) Like(col column.Type, value interface{}) *Query[T] {
	return q
}

// And ...
func (q *Query[T]) And() *Query[T] {
	return q
}

// Or ...
func (q *Query[T]) Or() *Query[T] {
	return q
}

// SortAsc ...
func (q *Query[T]) SortAsc(columns ...column.Type) *Query[T] {
	return q
}

// SortDesc ...
func (q *Query[T]) SortDesc(columns ...column.Type) *Query[T] {
	return q
}
