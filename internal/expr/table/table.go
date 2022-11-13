package table

import (
	"context"
	"fmt"

	"github.com/yeungsean/ysq-db/internal"
)

// Expr ...
type Expr[T string] struct {
	Table T
	Alias string
}

// String ...
func (t Expr[T]) String(ctx context.Context) string {
	provider := internal.CtxGetSourceProvider(ctx)
	tb := provider.Quote(string(t.Table))
	if t.Alias == "" {
		return tb
	}
	return fmt.Sprintf("%s AS %s", tb, t.Alias)
}
