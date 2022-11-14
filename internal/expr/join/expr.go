package join

import (
	"context"
	"strings"

	"github.com/yeungsean/ysq-db/internal/expr/table"
)

// Expr ...
type Expr[T string] struct {
	table.Expr[T]
	Type
	// 暂没支持多db，写法请带上table name
	Condition string
}

// String ...
func (j Expr[T]) String(ctx context.Context) string {
	sb := strings.Builder{}
	tb := j.Expr.String(ctx)
	sb.Grow(len(tb) + 11 + len(j.Condition) + len(j.Type))
	sb.WriteString(" ")
	sb.WriteString(string(j.Type))
	sb.WriteString(" JOIN ")
	sb.WriteString(tb)
	sb.WriteString(" ON ")
	sb.WriteString(j.Condition)
	return sb.String()
}
