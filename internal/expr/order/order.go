package order

import (
	"context"
	"fmt"

	provider "github.com/yeungsean/ysq-db/pkg/dbprovider"
	"github.com/yeungsean/ysq-db/pkg/field"
	"github.com/yeungsean/ysq-db/pkg/option"
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
	provider := provider.CtxGet(ctx)
	if o.Type == "" {
		o.Type = Asc
	}
	return fmt.Sprintf(`%s %s`, provider.OtherTypeField(&o.Field), o.Type)
}

// Options 可选参数
type Options func(*Order)

// WithAlias ...
func WithAlias(alias string) Options {
	return func(o *Order) {
		o.Alias = alias
	}
}

// NewOrder ...
func NewOrder(name field.Type, orderType Type, opts ...Options) *Order {
	o := &Order{Type: orderType}
	o.Name = name
	option.ForEach(o, opts)
	return o
}
