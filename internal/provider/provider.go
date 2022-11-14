package provider

import (
	"fmt"

	"github.com/yeungsean/ysq-db/pkg/field"
)

// IProvider ...
type IProvider interface {
	// PlaceHolder 占位符
	PlaceHolder(int) string
	// DefaultValue 默认值的sql
	DefaultValue(field, value string) string
	// Quote ...
	Quote(string) string
	// SelectFieldsQuote 包起来
	SelectFieldsQuote(...*field.Field) []string
	// SelectFieldQuote 包起来
	SelectFieldQuote(*field.Field) string
	// OtherTypeFieldsQuote ...
	OtherTypeFieldsQuote(...*field.Field) []string
	// OtherTypeFieldQuote ...
	OtherTypeFieldQuote(*field.Field) string
}

// SelectFieldQuote ...
func SelectFieldQuote(f *field.Field, quoteFmt string) string {
	if f.Prefix == "" {
		if !f.GetAggregation().IsNone() {
			return fmt.Sprintf("%s(%s)", f.GetAggregation(), f.Name)
		} else {
			return fmt.Sprintf(quoteFmt, f.Name)
		}
	} else {
		if !f.GetAggregation().IsNone() {
			return fmt.Sprintf("%s(%s.%s)", f.GetAggregation(), f.Prefix, f.Name)
		} else {
			return fmt.Sprintf("%s.%s", f.Prefix, f.Name)
		}
	}
}
