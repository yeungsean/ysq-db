package postgresql

import (
	"fmt"

	"github.com/yeungsean/ysq"
	"github.com/yeungsean/ysq-db/internal/expr/field"
	"github.com/yeungsean/ysq-db/internal/provider"
)

// Provider postgresql
type Provider struct{}

// PlaceHolder 占位符
func (m Provider) PlaceHolder(i int) string {
	return fmt.Sprintf("$%d", i)
}

// DefaultValue 默认值
func (m Provider) DefaultValue(field, value string) string {
	return fmt.Sprintf(`COALESCE(%s,%s)`, field, value)
}

// SelectFieldsQuote 包起来
func (m Provider) SelectFieldsQuote(fields ...*field.Field) []string {
	return ysq.FromSlice(fields).CastToStringBy(m.SelectFieldQuote).ToSlice(len(fields))
}

// SelectFieldQuote 包起来
func (m Provider) SelectFieldQuote(field *field.Field) string {
	c := provider.SelectFieldQuote(field, `"%s"`)
	if field.DefaultValue != nil {
		c = fmt.Sprintf("COALESCE(%s,%s)", c, field.DefaultValue)
	}

	if field.Alias != "" {
		c = fmt.Sprintf("%s AS %s", c, field.Alias)
	}
	return c
}

// OtherTypeFieldsQuote ...
func (m Provider) OtherTypeFieldsQuote(fields ...*field.Field) []string {
	return ysq.FromSlice(fields).CastToStringBy(m.OtherTypeFieldQuote).ToSlice(len(fields))
}

// OtherTypeFieldQuote ...
func (m Provider) OtherTypeFieldQuote(field *field.Field) string {
	if field.Alias != "" {
		return field.Alias
	} else if field.Prefix != "" {
		return fmt.Sprintf("%s.%s", field.Prefix, field.Name)
	} else {
		return m.Quote(string(field.Name))
	}
}

// Quote ...
func (m Provider) Quote(str string) string {
	return fmt.Sprintf(`"%s"`, str)
}