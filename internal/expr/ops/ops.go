package ops

// Type 比较操作类型
type Type string

// String ...
func (t Type) String() string {
	if t == "" {
		return string(EQ)
	}
	return string(t)
}

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
	// NotLike NOT LIKE
	NotLike Type = "NOT LIKE"
	// IsNull ...
	IsNull Type = "IS NULL"
	// IsNotNull ...
	IsNotNull Type = "IS NOT NULL"
	// In ...
	In Type = "IN"
	// NotIn ...
	NotIn Type = "NOT IN"
	// Between between ... and ...
	Between Type = "BETWEEN"
	// NotBetwee ...
	NotBetween Type = "NOT BETWEEN"
)
