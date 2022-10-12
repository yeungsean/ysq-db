package join

// Type 连接类型
type Type string

const (
	// Left 左连接
	Left Type = "LEFT"
	// Right 右连接
	Right Type = "RIGHT"
	// Inner 内连接
	Inner Type = "INNER"
)
