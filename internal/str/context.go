package str

import (
	"github.com/yeungsean/ysq-db/internal/expr/cond"
	"github.com/yeungsean/ysq-db/internal/expr/join"
	"github.com/yeungsean/ysq-db/internal/expr/order"
	"github.com/yeungsean/ysq-db/internal/expr/table"
	"github.com/yeungsean/ysq-db/pkg/field"
)

type queryContext[T string] struct {
	HasForUpdate bool
	WhereClause  *cond.Cond
	HavingClause *cond.Cond
	Fields       []*field.Field
	Groups       []*field.Field
	orders       []*order.Order
	mainTable    table.Expr[T]
	joins        []join.Expr[T]

	LimitOffset int
	LimitCount  int
	Values      []any
}

func newQueryContext[T string]() *queryContext[T] {
	qc := &queryContext[T]{
		WhereClause:  cond.Linear(),
		HavingClause: cond.Linear(),
		Fields:       make([]*field.Field, 0, 1),
		Groups:       make([]*field.Field, 0),
		orders:       make([]*order.Order, 0),
		Values:       make([]any, 0, 1),
		joins:        make([]join.Expr[T], 0),
		mainTable:    table.Expr[T]{},
	}
	return qc
}
