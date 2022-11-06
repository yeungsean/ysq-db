package ysqdb

import (
	"fmt"

	"github.com/yeungsean/ysq-db/pkg"
	"github.com/yeungsean/ysq-db/pkg/column"
	"github.com/yeungsean/ysq-db/pkg/cond"
	"github.com/yeungsean/ysq-db/pkg/join"
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

	entity[T pkg.ITable] struct {
		Entity T
		Alias  string
	}

	joinExpr[T pkg.ITable] struct {
		entity[T]
		jt join.Type
		jc []*cond.Cond
	}

	queryContext[T pkg.ITable] struct {
		HasForUpdate bool
		WhereClause  *cond.Cond
		HavingClause *cond.Cond
		Fields       []*column.Column
		Groups       []*column.Column
		orders       []*order.Order
		mainEntity   *entity[T]
		joinEntites  []*joinExpr[T]

		LimitOffset int
		LimitCount  int
		Values      []any
	}

	// LambdaOption ...
	LambdaOption[T pkg.ITable] func(*queryContext[T])
)

// String ...
func (e entity[T]) String() string {
	str := e.Entity.String()
	if e.Alias == "" {
		return str
	}
	return fmt.Sprintf(`%s AS %s`, str, e.Alias)
}

// WithAlias 别名
func WithAlias[T pkg.ITable](alias string) func(*queryContext[T]) {
	return func(lc *queryContext[T]) {
		lc.mainEntity.Alias = alias
	}
}

func newQueryContext[T pkg.ITable](value T, alias string) *queryContext[T] {
	qc := &queryContext[T]{
		WhereClause:  cond.Linear(),
		HavingClause: cond.Linear(),
		Fields:       make([]*column.Column, 0, 1),
		Groups:       make([]*column.Column, 0),
		orders:       make([]*order.Order, 0),
		Values:       make([]any, 0, 1),
		joinEntites:  make([]*joinExpr[T], 0),
		mainEntity:   &entity[T]{},
	}
	qc.mainEntity.Entity = value
	qc.mainEntity.Alias = alias
	return qc
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

func statementJoin() statement.Type {
	return statement.Join
}
