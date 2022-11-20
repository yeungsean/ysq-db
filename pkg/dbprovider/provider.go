package dbprovider

import (
	"context"
	"fmt"

	"github.com/yeungsean/ysq-db/pkg"
	"github.com/yeungsean/ysq-db/pkg/field"
)

// IDBProvider ...
type IDBProvider interface {
	// PlaceHolder 占位符
	PlaceHolder(int) string
	// Quote ...
	Quote(string) string
	// SelectFields 包起来
	SelectFields(...*field.Field) []string
	// SelectField 包起来
	SelectField(*field.Field) string
	// OtherTypeFields ...
	OtherTypeFields(...*field.Field) []string
	// OtherTypeField ...
	OtherTypeField(*field.Field) string
	// Type 数据库类型
	Type() string
}

// SelectFieldQuote ...
func SelectFieldQuote(f *field.Field, quoteFmt string) (c string) {
	if f.Prefix == "" {
		if f.Quote {
			c = fmt.Sprintf(quoteFmt, f.Name)
		} else {
			c = string(f.Name)
		}

		if f.GetAggregation().IsNone() {
			return
		} else {
			return fmt.Sprintf("%s(%s)", f.GetAggregation(), c)
		}
	} else {
		if f.GetAggregation().IsNone() {
			return fmt.Sprintf("%s.%s", f.Prefix, f.Name)
		} else {
			return fmt.Sprintf("%s(%s.%s)", f.GetAggregation(), f.Prefix, f.Name)
		}
	}
}

// OtherTypeField ...
func OtherTypeField(field *field.Field, quoteFn func(*field.Field) string) (c string) {
	defer func() {
		if !field.GetAggregation().IsNone() {
			if field.Alias == "" {
				c = fmt.Sprintf(`%s(%s)`, field.GetAggregation(), c)
			} else {
				c = fmt.Sprintf(`%s(%s) AS %s`, field.GetAggregation(), c, field.Alias)
			}
		}
	}()
	switch {
	case field.Alias != "":
		return field.Alias
	case field.Prefix != "":
		return fmt.Sprintf("%s.%s", field.Prefix, field.Name)
	case field.Quote:
		return quoteFn(field)
	default:
		return string(field.Name)
	}
}

// CtxGet 获取数据源provider
func CtxGet(ctx context.Context) IDBProvider {
	return ctx.Value(pkg.CtxKeyDBProvider).(IDBProvider)
}
