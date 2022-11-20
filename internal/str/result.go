package str

import "database/sql"

// Scan 查询并返回单条记录
func (q *Query[T]) Scan(data any) error {
	qc := q.ctxGetLambda()
	if tx := q.ctxGetTx(); tx == nil {
		return qc.DB.GetContext(q.ctx, data, q.String(), q.Values()...)
	} else {
		return tx.GetContext(q.ctx, data, q.String(), q.Values()...)
	}
}

// Slice 查询并返回多行记录
func (q *Query[T]) Slice(data any) error {
	qc := q.ctxGetLambda()
	if tx := q.ctxGetTx(); tx == nil {
		return qc.DB.SelectContext(q.ctx, data, q.String(), q.Values()...)
	} else {
		return tx.SelectContext(q.ctx, data, q.String(), q.Values()...)
	}
}

// ScanFields 查询并返回单条记录，指定变量接收
func (q *Query[T]) ScanFields(vals ...any) error {
	qc := q.ctxGetLambda()
	var row *sql.Row
	if tx := q.ctxGetTx(); tx == nil {
		row = qc.DB.QueryRowContext(q.ctx, q.String(), q.Values()...)
	} else {
		row = tx.QueryRowContext(q.ctx, q.String(), q.Values()...)
	}
	return row.Scan(vals...)
}
