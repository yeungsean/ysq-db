package node

import "container/list"

// List ...
type List struct {
	Head   *list.List
	Tail   *list.List
	Length int
}
