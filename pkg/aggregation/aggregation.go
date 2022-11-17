package aggregation

// Type 聚合类型
type Type string

// IsNone ...
func (t Type) IsNone() bool {
	return t == None
}

const (
	// None ...
	None = ""
	// Count ...
	Count Type = "COUNT"
	// Sum ...
	Sum Type = "SUM"
	// Min ...
	Min Type = "MIN"
	// Max ...
	Max Type = "MAX"
	// AVG ...
	Avg Type = "AVG"
)
