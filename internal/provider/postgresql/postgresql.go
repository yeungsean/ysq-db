package postgresql

import (
	"fmt"

	"github.com/yeungsean/ysq"
	"github.com/yeungsean/ysq-db/internal/provider"
	"github.com/yeungsean/ysq-db/pkg/field"
)

const quote = `"%s"`

// Provider postgresql
type Provider struct{}

// PlaceHolder 占位符
func (m Provider) PlaceHolder(i int) string {
	return fmt.Sprintf("$%d", i)
}

// SelectFields 包起来
func (m Provider) SelectFields(fields ...*field.Field) []string {
	return ysq.FromSlice(fields).CastToStringBy(m.SelectField).ToSlice(len(fields))
}

// SelectField 包起来
func (m Provider) SelectField(field *field.Field) (c string) {
	c = provider.SelectFieldQuote(field, quote)
	if field.DefaultValue != nil {
		c = fmt.Sprintf("COALESCE(%s,%v)", c, field.DefaultValue)
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
