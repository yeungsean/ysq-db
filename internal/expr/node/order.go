package node

import "github.com/yeungsean/ysq-db/internal/expr/order"

// OrderBy ...
type OrderBy struct {
	Type order.Type
}

// Name ...
func (o OrderBy) Name() string {
	return "ORDER BY"
}
