package ysqdb

import (
	"github.com/yeungsean/ysq-db/pkg"
	"github.com/yeungsean/ysq-db/pkg/column"
	"github.com/yeungsean/ysq-db/pkg/cond"
	"github.com/yeungsean/ysq-db/pkg/order"
	"github.com/yeungsean/ysq-db/pkg/statement"
)

type ctxKey uint8

const (
	ctxKeyLambda ctxKey = iota
)

type (
	// Iterable ...
	Iterable[T pkg.ITable] interface {
		Next() Iterator[T]
	}

	// Iterator 迭代器
	Iterator[T pkg.ITable] func() statement.Type

	lambdaContext[T pkg.ITable] struct {
		HasForUpdate bool
		WhereClause  *cond.Cond
		HavingClause *cond.Cond
		Fields       []*column.Column
		Groups       []*column.Column
		orders       []*order.Order
		Entity       T

		LimitOffset int
		LimitCount  int
		Values      []any
	}

	// LambdaOption ...
	LambdaOption[T pkg.ITable] func(*lambdaContext[T])
)

func newlambdaContext[T pkg.ITable](entity T) *lambdaContext[T] {
	return &lambdaContext[T]{
		WhereClause:  cond.Linear(),
		HavingClause: cond.Linear(),
		Fields:       make([]*column.Column, 0, 1),
		Groups:       make([]*column.Column, 0),
		orders:       make([]*order.Order, 0),
		Entity:       entity,
		Values:       make([]interface{}, 0, 1),
	}
}

func statementWhere() statement.Type {
	return statement.Where
}

func statementSort() statement.Type {
	return statement.Sort
}

func statementColumn() statement.Type {
	return statement.Column
}
