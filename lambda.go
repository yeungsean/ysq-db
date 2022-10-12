package ysqdb

import (
	"context"

	"github.com/yeungsean/ysq-db/pkg"
	"github.com/yeungsean/ysq-db/pkg/statments/join"
	"github.com/yeungsean/ysq-db/pkg/statments/ops"
	"github.com/yeungsean/ysq-db/pkg/statments/order"
)

// Lambda ...
type Lambda[T any] struct {
	Root *Lambda[T]
	Prev *Lambda[T]
	Next Iterator[T]
}

// From ...
func From[T any](entity string) *Lambda[T] {
	return &Lambda[T]{
		Root: nil,
		Prev: nil,
		Next: func() T {
			var zero T
			return zero
		},
	}
}

// Context ...
func (l *Lambda[T]) Context(ctx context.Context) *Lambda[T] {
	return l
}

// Where ...
func (l *Lambda[T]) Where(column pkg.ColumnType, value interface{}, opArgs ...ops.Type) *Lambda[T] {
	// op := ops.EQ
	// if len(opArgs) > 1 {
	// 	op = opArgs[0]
	// }
	return l
}

// ColumnCompare ...
type ColumnCompare struct {
	Column pkg.ColumnType
	Value  interface{}
	Op     ops.Type
}

// And a=1 AND b=2
func (l *Lambda[T]) And(elements ...ColumnCompare) *Lambda[T] {
	return l
}

// Or a=1 OR b=2
func (l *Lambda[T]) Or(elements ...ColumnCompare) *Lambda[T] {
	return l
}

// Like a LIKE '%abc%'
func (l *Lambda[T]) Like(column pkg.ColumnType, value string) *Lambda[T] {
	return l
}

// Join ...
func (l *Lambda[T]) Join(entity string, jtArgs ...join.Type) *Lambda[T] {
	return l
}

// Order ORDER BY a
func (l *Lambda[T]) Order(column pkg.ColumnType, otArgs ...order.Type) *Lambda[T] {
	// ot := order.Asc
	// if len(otArgs) > 1 {
	// 	ot = otArgs[0]
	// }
	return l
}
