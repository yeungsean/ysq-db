package mysql

import (
	"fmt"

	"github.com/yeungsean/ysq"
	"github.com/yeungsean/ysq-db/internal/provider"
	"github.com/yeungsean/ysq-db/pkg/field"
)

const quote = "`%s`"

// Provider mysql
type Provider struct{}

// PlaceHolder 占位符
func (m Provider) PlaceHolder(int) string {
	return "?"
}

// SelectFieldsQuote 包起来
func (m Provider) SelectFieldsQuote(fields ...*field.Field) []string {
	return ysq.FromSlice(fields).CastToStringBy(m.SelectFieldQuote).ToSlice(len(fields))
}

// SelectFieldQuote 包起来
func (m Provider) SelectFieldQuote(field *field.Field) (c string) {
	c = provider.SelectFieldQuote(field, quote)
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
func (m Provider) OtherTypeField(f *field.Field) (c string) {
	return provider.OtherTypeField(f, func(f *field.Field) string {
		return m.Quote(string(f.Name))
	})
}

// Quote ...
func (m Provider) Quote(str string) string {
	return fmt.Sprintf(quote, str)
}
