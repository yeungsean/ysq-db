package node

// Node 节点
type Node interface {
	Name() string
}

// AliasType ...
type AliasType uint8

const (
	AliasTable AliasType = iota
	AliasColumn
)

// Alias ...
type Alias struct {
	Type        AliasType
	Name        string
	ColumnNames []string
}
