package pkg

import (
	"github.com/yeungsean/ysq-db/pkg/field"
)

// ITable 获取数据表名和Schema
type ITable interface {
	Table() string
	Schema() string
	String() string
}

// IColumn 获取列
type IColumn interface {
	Columns() []field.Type
}
