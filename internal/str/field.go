package str

import (
	"github.com/yeungsean/ysq"
	"github.com/yeungsean/ysq-db/internal/expr/statement"
	"github.com/yeungsean/ysq-db/pkg/field"
)

// Select ...
func (q *Query[T]) Select(fields ...string) *Query[T] {
	return q.wrap(
		func(q *Query[T], qc *queryContext[T]) statement.Type {
			res := ysq.Select(
				ysq.FromSlice(fields),
				func(s string) *field.Field {
					return &field.Field{
						Name: field.Type(s),
					}
				}).ToSlice(len(fields))
			qc.Fields = append(qc.Fields, res...)
			return statement.Column
		})
}

// SelectPrefix ...
func (q *Query[T]) SelectPrefix(prefix string, fields ...string) *Query[T] {
	return q.wrap(func(q *Query[T], qc *queryContext[T]) statement.Type {
		res := ysq.Select(
			ysq.FromSlice(fields),
			func(s string) *field.Field {
				return &field.Field{
					Name: field.Type(s),
					FieldOption: field.FieldOption{
						Prefix: prefix,
					},
				}
			}).ToSlice(len(fields))
		qc.Fields = append(qc.Fields, res...)
		return statement.Column
	})
}

// Field ...
func (q *Query[T]) Field(fieldName string, opts ...field.Option) *Query[T] {
	return q.wrap(
		func(q *Query[T], qc *queryContext[T]) statement.Type {
			f := &field.Field{
				Name: field.Type(fieldName),
			}
			for _, opt := range opts {
				opt(&f.FieldOption)
			}
			if f.Alias == "" {
				f.Alias = fieldName
			}
			qc.Fields = append(qc.Fields, f)
			return statement.Column
		})
}
