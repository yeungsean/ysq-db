package str

import (
	"github.com/yeungsean/ysq"
	"github.com/yeungsean/ysq-db/internal/expr/statement"
	"github.com/yeungsean/ysq-db/pkg/field"
	"github.com/yeungsean/ysq-db/pkg/option"
)

// Select 查询字段
func (q *Query[T]) Select(fields ...string) *Query[T] {
	return q.wrap(
		func(q *Query[T], qc *queryContext[T]) statement.Type {
			res := ysq.Select(
				ysq.FromSlice(fields),
				func(s string) *field.Field {
					return field.New(field.Type(s))
				}).ToSlice(len(fields))
			qc.Fields = append(qc.Fields, res...)
			return statement.Column
		})
}

// SelectPrefix 查询字段，统一前缀
func (q *Query[T]) SelectPrefix(prefix string, fields ...string) *Query[T] {
	return q.wrap(func(q *Query[T], qc *queryContext[T]) statement.Type {
		res := ysq.Select(
			ysq.FromSlice(fields),
			func(s string) *field.Field {
				return &field.Field{
					Name: field.Type(s),
					Option: field.Option{
						Prefix: prefix,
					},
				}
			}).ToSlice(len(fields))
		qc.Fields = append(qc.Fields, res...)
		return statement.Column
	})
}

// Field 字段
func (q *Query[T]) Field(fieldName string, opts ...field.Options) *Query[T] {
	return q.wrap(
		func(q *Query[T], qc *queryContext[T]) statement.Type {
			f := field.New(field.Type(fieldName))
			option.ForEach(&f.Option, opts)
			if f.DefaultValue != nil && f.Alias == "" {
				f.Alias = string(f.Name)
			}
			qc.Fields = append(qc.Fields, f)
			return statement.Column
		})
}
