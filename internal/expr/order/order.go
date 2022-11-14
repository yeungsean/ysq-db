package order

import (
	"context"
	"fmt"

	"github.com/yeungsean/ysq-db/internal"
	"github.com/yeungsean/ysq-db/internal/expr/common"
	"github.com/yeungsean/ysq-db/pkg/field"
)

// Type 排序类型
type Type string

const (
	// Asc 升序
	Asc Type = "ASC"
	// Desc 降序
	Desc Type = "DESC"
)

// Order 排序
type Order struct {
	field.Field
	Type Type
}

// String ...
func (o Order) String(ctx context.Context) string {
	provider := internal.CtxGetSourceProvider(ctx)
	if o.Type == "" {
		o.Type = Asc
	}
	return fmt.Sprintf(`%s %s`, provider.OtherTypeFieldQuote(&o.Field), o.Type)
}

// Option 可选参数
type Option func(*Order)

// WithAlias ...
func WithAlias(alias string) Option {
	return func(o *Order) {
		o.Alias = alias
	}
}

// NewOrder ...
func NewOrder(name field.Type, orderType Type, opts ...Option) *Order {
	o := &Order{Type: orderType}
	o.Name = name
	common.OptionForEach(o, opts)
	return o
}
