package clickhouse

import (
	"fmt"

	"github.com/yeungsean/ysq"
	"github.com/yeungsean/ysq-db/pkg/dbprovider"
	"github.com/yeungsean/ysq-db/pkg/field"
)

// Provider clickhouse
type Provider struct{}

// PlaceHolder 占位符
func (m Provider) PlaceHolder(int) string {
	return "?"
}

// Type 数据库类型
func (m Provider) Type() string {
	return "clickhouse"
}

// SelectFields 包起来
func (m Provider) SelectFields(fields ...*field.Field) []string {
	return ysq.FromSlice(fields).CastToStringBy(m.SelectField).ToSlice(len(fields))
}

// SelectField 包起来
func (m Provider) SelectField(field *field.Field) string {
	c := dbprovider.SelectFieldQuote(field, "`%s`")
	if field.DefaultValue != nil {
		c = fmt.Sprintf("IFNULL(%s,%v)", c, field.DefaultValue)
	}

	if field.Alias != "" {
		c = fmt.Sprintf("%s AS %s", c, field.Alias)
	}
	return c
}

// OtherTypeFields ...
func (m Provider) OtherTypeFields(fields ...*field.Field) []string {
	return ysq.FromSlice(fields).CastToStringBy(m.OtherTypeField).ToSlice(len(fields))
}

// OtherTypeField ...
func (m Provider) OtherTypeField(f *field.Field) string {
	return dbprovider.OtherTypeField(f, func(f *field.Field) string {
		return m.Quote(string(f.Name))
	})
}

// Quote ...
func (m Provider) Quote(str string) string {
	return fmt.Sprintf("`%s`", str)
}
