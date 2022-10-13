package ysqdb

import (
	"context"

	"github.com/yeungsean/ysq-db/pkg"
	"github.com/yeungsean/ysq-db/pkg/column"
	"github.com/yeungsean/ysq-db/pkg/statments/join"
	"github.com/yeungsean/ysq-db/pkg/statments/ops"
	"github.com/yeungsean/ysq/pkg/delegate"
)

// Lambda ...
type Lambda[T any] struct {
	entity pkg.ITable
}

// From ...
func From[T pkg.ITable](entity pkg.ITable) *Lambda[T] {
	return &Lambda[T]{
		entity: entity,
	}
}

// Context ...
func (l *Lambda[T]) Context(ctx context.Context) *Lambda[T] {
	return l
}

// Cache ...
func (l *Lambda[T]) Cache() *Lambda[T] {
	return l
}

func varArgGet[T any](args ...T) T {
	var zero T
	if len(args) > 0 {
		zero = args[0]
	}
	return zero
}

// Where ...
func (l *Lambda[T]) Where(column column.Type, value any, opArgs ...ops.Type) *Lambda[T] {
	// op := varArgGet(opArgs)
	return l
}

// ColumnCompare ...
type ColumnCompare struct {
	Column column.Type
	Value  any
	Op     ops.Type
}

// And a=1 AND b=2
func (l *Lambda[T]) And(elements ...ColumnCompare) *Lambda[T] {
	return l
}

// Or a=1 OR b=2
func (l *Lambda[T]) Or(elements ...ColumnCompare) *Lambda[T] {
	return l
}

// Like a LIKE '%abc%'
func (l *Lambda[T]) Like(column column.Type, value string) *Lambda[T] {
	return l
}

// Join https://www.tektutorialshub.com/entity-framework/join-query-entity-framework/
func Join[TLeft, TRight, TResult pkg.ITable, TKey column.Type](
	left *Lambda[TLeft],
	right *Lambda[TRight],
	leftKeySelector delegate.Func1[TLeft, TKey],
	rightKeySelector delegate.Func1[TRight, TKey],
	jtArgs ...join.Type,
) *Lambda[TResult] {
	return nil
}

// Join ...
func (l *Lambda[T]) Join(entity pkg.ITable, leftColumn, rightColumn column.Type, jtArgs ...join.Type) *Lambda[T] {
	return l
}

// OrderAsc ORDER BY a
func (l *Lambda[T]) OrderAsc(columns ...column.Type) *Lambda[T] {
	return l
}

// OrderDesc ORDER BY a DESC
func (l *Lambda[T]) OrderDesc(columns ...column.Type) *Lambda[T] {
	return l
}

// GroupBy 分组
func (l *Lambda[T]) GroupBy(column column.Type) *Lambda[T] {
	return l
}

// Having 分组查询后的过滤
func (l *Lambda[T]) Having(column column.Type, value any, poArg ops.Type) *Lambda[T] {
	return l
}

// Limit limit
func (l *Lambda[T]) Limit(size int) *Lambda[T] {
	return l
}

// Skip skip
func (l *Lambda[T]) Skip(size int) *Lambda[T] {
	return l.Offset(size)
}

// Offset offset
func (l *Lambda[T]) Offset(size int) *Lambda[T] {
	return l
}

// Select 筛选字段
func (l *Lambda[T]) Select(columns ...column.Type) *Lambda[T] {
	return l
}

// Distinct 去重
func (l *Lambda[T]) Distinct(columns ...column.Type) *Lambda[T] {
	return l
}

/* 生成结果 */

// Count 统计总数
func (l *Lambda[T]) Count() (int64, error) {
	return 0, nil
}

// Avg 平均数
func (l *Lambda[T]) Avg() (float64, error) {
	return 0, nil
}

// Sum 总和
func (l *Lambda[T]) Sum() (float64, error) {
	return 0, nil
}

// Max 最大值
func (l *Lambda[T]) Max() (t T, err error) {
	return
}

// Min 最小值
func (l *Lambda[T]) Min() (t T, err error) {
	return
}

// Exists 记录是否存在
func (l *Lambda[T]) Exists() (bool, error) {
	return false, nil
}

// ToSlice ...
func (l *Lambda[T]) ToSlice(countArg ...int) ([]*T, error) {
	count := 2
	if len(countArg) > 0 {
		count = countArg[0]
	}
	res := make([]*T, 0, count)
	return res, nil
}

// One ...
func (l *Lambda[T]) One() (dest *T, err error) {
	return
}

// Scan ...
func (l *Lambda[T]) Scan(dest ...any) (err error) {
	return
}

// ScanStruct ...
func (l *Lambda[T]) ScanStruct(dest any) (err error) {
	return
}

// ScanMap ...
func (l *Lambda[T]) ScanMap(dest map[string]any) (err error) {
	return
}
