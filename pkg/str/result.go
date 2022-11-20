package str

import "database/sql"

// Scan 查询并返回单条记录
func (q *Query[T]) Scan(data any) error {
	if tx := q.ctxGetTx(); tx != nil {
		return tx.GetContext(q.ctx, data, q.String(), q.Args()...)
	} else {
		return q.ctxGetDB().GetContext(q.ctx, data, q.String(), q.Args()...)
	}
}

// Slice 查询并返回多行记录
func (q *Query[T]) Slice(data any) error {
	if tx := q.ctxGetTx(); tx != nil {
		return tx.SelectContext(q.ctx, data, q.String(), q.Args()...)
	}
	return q.ctxGetDB().SelectContext(q.ctx, data, q.String(), q.Args()...)
}

// ScanFields 查询并返回单条记录，指定变量接收
func (q *Query[T]) ScanFields(vals ...any) error {
	var row *sql.Row
	if tx := q.ctxGetTx(); tx != nil {
		row = tx.QueryRowContext(q.ctx, q.String(), q.Args()...)
	} else {
		row = q.ctxGetDB().QueryRowContext(q.ctx, q.String(), q.Args()...)
	}
	return row.Scan(vals...)
}
