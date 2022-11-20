package table

import (
	"context"
	"fmt"

	"github.com/yeungsean/ysq-db/internal"
	"github.com/yeungsean/ysq-db/pkg/option"
)

// Expr table表达式
type Expr[T string] struct {
	Table T
	option.Option
}

// String ...
func (t Expr[T]) String(ctx context.Context) string {
	var tb string
	if t.Quote {
		provider := internal.CtxGetDBProvider(ctx)
		tb = provider.Quote(string(t.Table))
	} else {
		tb = string(t.Table)
	}

	if t.Alias == "" {
		return tb
	}
	return fmt.Sprintf("%s AS %s", tb, t.Alias)
}
