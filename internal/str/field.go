package str

import (
	"github.com/yeungsean/ysq"
	"github.com/yeungsean/ysq-db/internal/expr/field"
	"github.com/yeungsean/ysq-db/internal/expr/statement"
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
func (q *Query[T]) Field(prefix, fieldName, defaultValue string) *Query[T] {
	return q.wrap(
		func(q *Query[T], qc *queryContext[T]) statement.Type {
			qc.Fields = append(qc.Fields, &field.Field{
				Name: field.Type(fieldName),
				FieldOption: field.FieldOption{
					DefaultValue: defaultValue,
					Prefix:       prefix,
				},
			})
			return statement.Column
		})
}

// FieldWihtoutPrefix ...
func (q *Query[T]) FieldWihtoutPrefix(fieldName, defaultValue string) *Query[T] {
	return q.Field("", fieldName, defaultValue)
}
