package str

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yeungsean/ysq-db/internal"
	"github.com/yeungsean/ysq-db/internal/expr/column"
	"github.com/yeungsean/ysq-db/internal/expr/cond"
	"github.com/yeungsean/ysq-db/internal/expr/ops"
	"github.com/yeungsean/ysq-db/internal/provider/mysql"
	"github.com/yeungsean/ysq-db/internal/provider/postgresql"
)

func TestIn(t *testing.T) {
	vals := []any{1, 2, 3}
	q := NewQuery(`user`).Field("name").In("id", vals)
	q.build()
	qCtx := q.ctxGetLambda()
	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeySourceProvider, &mysql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, "`id` IN(?,?,?)", qCtx.WhereClause.String(ctx))
		assert.Equal(t, vals, qCtx.Values)
	}()

	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeySourceProvider, &postgresql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, `"id" IN($1,$2,$3)`, qCtx.WhereClause.String(ctx))
		assert.Equal(t, vals, qCtx.Values)
	}()
}

func TestNotIn(t *testing.T) {
	vals := []any{1}
	q := NewQuery(`user`).Field("name").NotIn("id", vals)
	q.build()
	qCtx := q.ctxGetLambda()
	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeySourceProvider, &mysql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, "`id` NOT IN(?)", qCtx.WhereClause.String(ctx))
		assert.Equal(t, vals, qCtx.Values)
	}()

	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeySourceProvider, &postgresql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, `"id" NOT IN($1)`, qCtx.WhereClause.String(ctx))
		assert.Equal(t, vals, qCtx.Values)
	}()
}

func TestEqual(t *testing.T) {
	val := 1
	q := NewQuery(`user`).Field("name").Equal("id", val)
	q.build()
	qCtx := q.ctxGetLambda()
	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeySourceProvider, &mysql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, "`id`=?", qCtx.WhereClause.String(ctx))
		assert.Equal(t, []any{val}, qCtx.Values)
	}()

	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeySourceProvider, &postgresql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, `"id"=$1`, qCtx.WhereClause.String(ctx))
		assert.Equal(t, []any{val}, qCtx.Values)
	}()
}

func TestNotEqual(t *testing.T) {
	val := 1
	q := NewQuery(`user`).Field("name").NotEqual("id", val)
	q.build()
	qCtx := q.ctxGetLambda()
	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeySourceProvider, &mysql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, "`id`<>?", qCtx.WhereClause.String(ctx))
		assert.Equal(t, []any{val}, qCtx.Values)
	}()

	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeySourceProvider, &postgresql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, `"id"<>$1`, qCtx.WhereClause.String(ctx))
		assert.Equal(t, []any{val}, qCtx.Values)
	}()
}

func TestGreater(t *testing.T) {
	val := 1
	q := NewQuery(`user`).Field("name").Greater("id", val)
	q.build()
	qCtx := q.ctxGetLambda()
	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeySourceProvider, &mysql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, "`id`>?", qCtx.WhereClause.String(ctx))
		assert.Equal(t, []any{val}, qCtx.Values)
	}()

	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeySourceProvider, &postgresql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, `"id">$1`, qCtx.WhereClause.String(ctx))
		assert.Equal(t, []any{val}, qCtx.Values)
	}()
}

func TestGreaterOrEqual(t *testing.T) {
	val := 1
	q := NewQuery(`user`).Field("name").GreaterOrEqual("id", val)
	q.build()
	qCtx := q.ctxGetLambda()
	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeySourceProvider, &mysql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, "`id`>=?", qCtx.WhereClause.String(ctx))
		assert.Equal(t, []any{val}, qCtx.Values)
	}()

	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeySourceProvider, &postgresql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, `"id">=$1`, qCtx.WhereClause.String(ctx))
		assert.Equal(t, []any{val}, qCtx.Values)
	}()
}

func TestLess(t *testing.T) {
	val := 1
	q := NewQuery(`user`).Field("name").Less("id", val)
	q.build()
	qCtx := q.ctxGetLambda()
	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeySourceProvider, &mysql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, "`id`<?", qCtx.WhereClause.String(ctx))
		assert.Equal(t, []any{val}, qCtx.Values)
	}()

	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeySourceProvider, &postgresql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, `"id"<$1`, qCtx.WhereClause.String(ctx))
		assert.Equal(t, []any{val}, qCtx.Values)
	}()
}

func TestLessOrEqual(t *testing.T) {
	val := 1
	q := NewQuery(`user`).Field("name").LessOrEqual("id", val)
	q.build()
	qCtx := q.ctxGetLambda()
	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeySourceProvider, &mysql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, "`id`<=?", qCtx.WhereClause.String(ctx))
		assert.Equal(t, []any{val}, qCtx.Values)
	}()

	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeySourceProvider, &postgresql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, `"id"<=$1`, qCtx.WhereClause.String(ctx))
		assert.Equal(t, []any{val}, qCtx.Values)
	}()
}

func TestIsNull(t *testing.T) {
	q := NewQuery(`user`).Field("name").IsNull("id")
	q.build()
	qCtx := q.ctxGetLambda()
	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeySourceProvider, &mysql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, "`id` IS NULL", qCtx.WhereClause.String(ctx))
		assert.Len(t, qCtx.Values, 0)
	}()

	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeySourceProvider, &postgresql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, `"id" IS NULL`, qCtx.WhereClause.String(ctx))
		assert.Len(t, qCtx.Values, 0)
	}()
}

func TestIsNotNull(t *testing.T) {
	q := NewQuery(`user`).Field("name").IsNotNull("id")
	q.build()
	qCtx := q.ctxGetLambda()
	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeySourceProvider, &mysql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, "`id` IS NOT NULL", qCtx.WhereClause.String(ctx))
		assert.Len(t, qCtx.Values, 0)
	}()

	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeySourceProvider, &postgresql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, `"id" IS NOT NULL`, qCtx.WhereClause.String(ctx))
		assert.Len(t, qCtx.Values, 0)
	}()
}

func TestLike(t *testing.T) {
	q := NewQuery(`user`).Field("name").Like("name", "ysl")
	q.build()
	qCtx := q.ctxGetLambda()
	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeySourceProvider, &mysql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, "`name` LIKE ?", qCtx.WhereClause.String(ctx))
		assert.Len(t, qCtx.Values, 0)
	}()

	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeySourceProvider, &postgresql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, `"name" LIKE $1`, qCtx.WhereClause.String(ctx))
		assert.Len(t, qCtx.Values, 0)
	}()
}

func TestWhere(t *testing.T) {
	q := NewQuery(`user`).Field("name").Where(
		cond.All().AddChildren(
			cond.All().
				Add(column.New("id", column.WithValue(1))).
				Add(column.New("create_time", column.WithValue("2022-01-02T15:04:05"), column.WithOp(ops.GTE))),
			cond.Any().
				Add(column.New("id", column.WithValue(1), column.WithOp(ops.GTE))).
				Add(column.New("create_time", column.WithValue("2022-01-02T15:04:05"), column.WithOp(ops.LT))),
		),
	)
	q.build()
	qCtx := q.ctxGetLambda()

	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeySourceProvider, &mysql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, "((`id`=? AND `create_time`>=?) AND (`id`>=? OR `create_time`<?))",
			qCtx.WhereClause.String(ctx))
	}()

	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeySourceProvider, &postgresql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, `(("id"=$1 AND "create_time">=$2) AND ("id">=$3 OR "create_time"<$4))`,
			qCtx.WhereClause.String(ctx))
	}()
}

func TestBetween(t *testing.T) {
	q := NewQuery(`user`).Field("name").Between("id", 1, 10)
	q.build()
	qCtx := q.ctxGetLambda()

	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeySourceProvider, &mysql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, "`id` BETWEEN ? AND ?",
			qCtx.WhereClause.String(ctx))
	}()

	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeySourceProvider, &postgresql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, `"id" BETWEEN $1 AND $2`,
			qCtx.WhereClause.String(ctx))
	}()
}

func TestNotBetween(t *testing.T) {
	q := NewQuery(`user`).Field("name").NotBetween("id", 1, 10)
	q.build()
	qCtx := q.ctxGetLambda()

	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeySourceProvider, &mysql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, "`id` NOT BETWEEN ? AND ?",
			qCtx.WhereClause.String(ctx))
	}()

	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeySourceProvider, &postgresql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, `"id" NOT BETWEEN $1 AND $2`,
			qCtx.WhereClause.String(ctx))
	}()
}
