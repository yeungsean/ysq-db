package order

import (
	"fmt"

	"github.com/yeungsean/ysq-db/pkg/alias"
	"github.com/yeungsean/ysq-db/pkg/common"
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

// Order ...
type Order struct {
	field.Field
	Type Type
}

// String ...
func (o Order) String() string {
	if o.Type == "" {
		o.Type = Asc
	}
	return fmt.Sprintf(`%s %s`, o.OrderField(), o.Type)
}

// Option 可选参数
type Option func(*Order)

// WithAlias ...
func WithAlias(alias *alias.Alias) Option {
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
