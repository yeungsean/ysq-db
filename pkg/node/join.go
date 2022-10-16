package node

import "github.com/yeungsean/ysq-db/pkg/join"

// Join ...
type Join struct {
	Type join.Type
}

// Name ...
func (j Join) Name() string {
	return "JOIN"
}
