package str

import (
	"github.com/yeungsean/ysq"
	"github.com/yeungsean/ysq-db/pkg/field"
	"github.com/yeungsean/ysq-db/pkg/statement"
)

// Select ...
func (q *Query[T]) Select(fields ...string) *Query[T] {
	return q.wrap(
		func(q *Query[T], qc *queryContext[T]) func() statement.Type {
			res := ysq.Select(
				ysq.FromSlice(fields),
				func(s string) *field.Field {
					return &field.Field{
						Name: field.Type(s),
					}
				}).ToSlice(len(fields))
			qc.Fields = append(qc.Fields, res...)
			return func() statement.Type {
				return statement.Column
			}
		})
}

// SelectPrefix ...
func (q *Query[T]) SelectPrefix(prefix string, fields ...string) *Query[T] {
	return q.wrap(func(q *Query[T], qc *queryContext[T]) func() statement.Type {
		res := ysq.Select(
			ysq.FromSlice(fields),
			func(s string) *field.Field {
				return &field.Field{
					Name:   field.Type(s),
					Prefix: prefix,
				}
			}).ToSlice(len(fields))
		qc.Fields = append(qc.Fields, res...)
		return func() statement.Type {
			return statement.Column
		}
	})
}

// Field ...
func (q *Query[T]) Field(prefix, fieldName, defaultValue string) *Query[T] {
	return q.wrap(func(q *Query[T], qc *queryContext[T]) func() statement.Type {
		qc.Fields = append(qc.Fields, &field.Field{
			Prefix:       prefix,
			Name:         field.Type(fieldName),
			DefaultValue: defaultValue,
		})
		return func() statement.Type {
			return statement.Column
		}
	})
}
