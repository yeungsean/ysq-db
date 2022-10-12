package ops

// Type 比较操作类型
type Type string

const (
	// EQ equal
	EQ Type = "="
	// NEQ not equal
	NEQ Type = "<>"
	// GTE >=
	GTE Type = ">="
	// GT >
	GT Type = ">"
	// LTE <=
	LTE Type = "<="
	// LT <
	LT Type = "<"
	// Like LIKE
	Like Type = "LIKE"
)
