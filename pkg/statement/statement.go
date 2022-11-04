package statement

// Type ...
type Type uint8

const (
	Nop Type = iota
	// Column 列
	Column = iota
	// Table 表
	Table
	// Where 过滤
	Where
	// Sort 排序
	Sort
	// Join 连接
	Join
	// Limit 限制
	Limit
	// Offset 偏移
	Offset
	// Count 总数
	Count
	// Sum 累加
	Sum
	// Max 最大
	Max
	// Min 最小
	Min
	// GroupBy 分组
	GroupBy
	// Having ...
	Having
)
