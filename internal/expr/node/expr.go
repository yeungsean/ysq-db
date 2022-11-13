package node

import "github.com/yeungsean/ysq-db/internal/expr/join"

type BoolExprType uint8

const (
	ExprAnd BoolExprType = iota
	ExprOr
	ExprNot
)

// BoolExpr ...
type BoolExpr struct {
	BoolOp BoolExprType
}

// FromExpr ...
type FromExpr struct {
	FromList List
}

// JoinExpr ...
type JoinExpr struct {
	Type join.Type
}
