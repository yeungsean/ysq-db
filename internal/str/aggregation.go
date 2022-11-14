package str

import (
	"github.com/yeungsean/ysq-db/internal/expr/aggregation"
	"github.com/yeungsean/ysq-db/internal/expr/statement"
	"github.com/yeungsean/ysq-db/pkg/field"
)

func (q *Query[T]) agg(fn aggregation.Type, f string, opts ...field.FieldOption) *Query[T] {
	return q.wrap(func(q *Query[T], qc *queryContext[T]) statement.Type {
		var opt field.FieldOption
		if len(opts) > 0 {
			opt = opts[0]
		}
		if opt.Alias == "" {
			opt.Alias = f
		}
		field.WithAggregation(fn)(&opt)
		qc.Fields = append(qc.Fields, &field.Field{
			Name:        field.Type(f),
			FieldOption: opt,
		})
		return statement.Column
	})
}

// Sum 总和
func (q *Query[T]) Sum(f string, opts ...field.FieldOption) *Query[T] {
	return q.agg(aggregation.Sum, f, opts...)
}

// Avg 平均数
func (q *Query[T]) Avg(f string, opts ...field.FieldOption) *Query[T] {
	return q.agg(aggregation.Avg, f, opts...)
}

// Min 最小值
func (q *Query[T]) Min(f string, opts ...field.FieldOption) *Query[T] {
	return q.agg(aggregation.Min, f, opts...)
}

// Max 最大值
func (q *Query[T]) Max(f string, opts ...field.FieldOption) *Query[T] {
	return q.agg(aggregation.Max, f, opts...)
}

// Count 总数
func (q *Query[T]) Count(opts ...field.FieldOption) *Query[T] {
	return q.wrap(func(q *Query[T], qc *queryContext[T]) statement.Type {
		var opt field.FieldOption
		if len(opts) > 0 {
			opt = opts[0]
		}
		field.WithAggregation(aggregation.Count)(&opt)
		qc.Fields = append(qc.Fields, &field.Field{
			Name:        field.Type(`1`),
			FieldOption: opt,
		})
		return statement.Column
	})
}
