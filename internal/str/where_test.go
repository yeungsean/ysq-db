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

func TestAndIn(t *testing.T) {
	vals := []any{1, 2, 3}
	q := NewQuery(`user`).Field("name").AndIn("id", vals)
	q.build()
	qCtx := q.ctxGetLambda()
	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeyDBProvider, &mysql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, "id IN(?,?,?)", qCtx.WhereClause.String(ctx))
		assert.Equal(t, vals, qCtx.Values)
	}()

	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeyDBProvider, &postgresql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, `id IN($1,$2,$3)`, qCtx.WhereClause.String(ctx))
		assert.Equal(t, vals, qCtx.Values)
	}()
}

func TestOrIn(t *testing.T) {
	vals := []any{1, 2, 3}
	q := NewQuery(`user`).Field("name").AndIn("id", vals)
	q.build()
	qCtx := q.ctxGetLambda()
	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeyDBProvider, &mysql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, "id IN(?,?,?)", qCtx.WhereClause.String(ctx))
		assert.Equal(t, vals, qCtx.Values)
	}()

	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeyDBProvider, &postgresql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, `id IN($1,$2,$3)`, qCtx.WhereClause.String(ctx))
		assert.Equal(t, vals, qCtx.Values)
	}()
}

func TestAndNotIn(t *testing.T) {
	vals := []any{1}
	q := NewQuery(`user`).Field("name").AndNotIn("id", vals)
	q.build()
	qCtx := q.ctxGetLambda()
	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeyDBProvider, &mysql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, "id NOT IN(?)", qCtx.WhereClause.String(ctx))
		assert.Equal(t, vals, qCtx.Values)
	}()

	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeyDBProvider, &postgresql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, `id NOT IN($1)`, qCtx.WhereClause.String(ctx))
		assert.Equal(t, vals, qCtx.Values)
	}()
}

func TestOrNotIn(t *testing.T) {
	vals := []any{1}
	q := NewQuery(`user`).Field("name").OrNotIn("id", vals).OrGreater("age", 100)
	q.build()
	qCtx := q.ctxGetLambda()
	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeyDBProvider, &mysql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, "id NOT IN(?) OR age>?", qCtx.WhereClause.String(ctx))
		assert.Equal(t, []any{1, 100}, qCtx.Values)
	}()

	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeyDBProvider, &postgresql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, `id NOT IN($1) OR age>$2`, qCtx.WhereClause.String(ctx))
		assert.Equal(t, []any{1, 100}, qCtx.Values)
	}()
}

func TestAndEqual(t *testing.T) {
	val := 1
	q := NewQuery(`user`).Field("name").AndEqual("id", val)
	q.build()
	qCtx := q.ctxGetLambda()
	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeyDBProvider, &mysql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, "id=?", qCtx.WhereClause.String(ctx))
		assert.Equal(t, []any{val}, qCtx.Values)
	}()

	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeyDBProvider, &postgresql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, `id=$1`, qCtx.WhereClause.String(ctx))
		assert.Equal(t, []any{val}, qCtx.Values)
	}()
}

func TestOrEqual(t *testing.T) {
	val := 1
	q := NewQuery(`user`).Field("name").OrEqual("id", val).OrLess("age", 10)
	q.build()
	qCtx := q.ctxGetLambda()
	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeyDBProvider, &mysql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, "id=? OR age<?", qCtx.WhereClause.String(ctx))
		assert.Equal(t, []any{val, 10}, qCtx.Values)
	}()

	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeyDBProvider, &postgresql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, `id=$1 OR age<$2`, qCtx.WhereClause.String(ctx))
		assert.Equal(t, []any{val, 10}, qCtx.Values)
	}()
}

func TestAndNotEqual(t *testing.T) {
	val := 1
	q := NewQuery(`user`).Field("name").AndNotEqual("id", val)
	q.build()
	qCtx := q.ctxGetLambda()
	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeyDBProvider, &mysql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, "id<>?", qCtx.WhereClause.String(ctx))
		assert.Equal(t, []any{val}, qCtx.Values)
	}()

	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeyDBProvider, &postgresql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, `id<>$1`, qCtx.WhereClause.String(ctx))
		assert.Equal(t, []any{val}, qCtx.Values)
	}()
}

func TestOrNotEqual(t *testing.T) {
	val := 1
	q := NewQuery(`user`).Field("name").OrNotEqual("id", val).OrGreater("id", 100)
	q.build()
	qCtx := q.ctxGetLambda()
	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeyDBProvider, &mysql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, "id<>? OR id>?", qCtx.WhereClause.String(ctx))
		assert.Equal(t, []any{val, 100}, qCtx.Values)
	}()

	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeyDBProvider, &postgresql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, `id<>$1 OR id>$2`, qCtx.WhereClause.String(ctx))
		assert.Equal(t, []any{val, 100}, qCtx.Values)
	}()
}

func TestAndGreater(t *testing.T) {
	val := 1
	q := NewQuery(`user`).Field("name").AndGreater("id", val)
	q.build()
	qCtx := q.ctxGetLambda()
	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeyDBProvider, &mysql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, "id>?", qCtx.WhereClause.String(ctx))
		assert.Equal(t, []any{val}, qCtx.Values)
	}()

	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeyDBProvider, &postgresql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, `id>$1`, qCtx.WhereClause.String(ctx))
		assert.Equal(t, []any{val}, qCtx.Values)
	}()
}

func TestOrGreater(t *testing.T) {
	val := 1
	q := NewQuery(`user`).Field("name").OrGreater("id", val).OrLessOrEqual("age", 50)
	q.build()
	qCtx := q.ctxGetLambda()
	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeyDBProvider, &mysql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, "id>? OR age<=?", qCtx.WhereClause.String(ctx))
		assert.Equal(t, []any{val, 50}, qCtx.Values)
	}()

	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeyDBProvider, &postgresql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, `id>$1 OR age<=$2`, qCtx.WhereClause.String(ctx))
		assert.Equal(t, []any{val, 50}, qCtx.Values)
	}()
}

func TestAndGreaterOrEqual(t *testing.T) {
	val := 1
	q := NewQuery(`user`).Field("name").AndGreaterOrEqual("id", val)
	q.build()
	qCtx := q.ctxGetLambda()
	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeyDBProvider, &mysql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, "id>=?", qCtx.WhereClause.String(ctx))
		assert.Equal(t, []any{val}, qCtx.Values)
	}()

	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeyDBProvider, &postgresql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, `id>=$1`, qCtx.WhereClause.String(ctx))
		assert.Equal(t, []any{val}, qCtx.Values)
	}()
}

func TestOrGreaterOrEqual(t *testing.T) {
	val := 1
	q := NewQuery(`user`).Field("name").OrGreaterOrEqual("id", val).OrLess("age", 9)
	q.build()
	qCtx := q.ctxGetLambda()
	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeyDBProvider, &mysql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, "id>=? OR age<?", qCtx.WhereClause.String(ctx))
		assert.Equal(t, []any{val, 9}, qCtx.Values)
	}()

	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeyDBProvider, &postgresql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, `id>=$1 OR age<$2`, qCtx.WhereClause.String(ctx))
		assert.Equal(t, []any{val, 9}, qCtx.Values)
	}()
}

func TestAndLess(t *testing.T) {
	val := 1
	q := NewQuery(`user`).Field("name").AndLess("id", val)
	q.build()
	qCtx := q.ctxGetLambda()
	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeyDBProvider, &mysql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, "id<?", qCtx.WhereClause.String(ctx))
		assert.Equal(t, []any{val}, qCtx.Values)
	}()

	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeyDBProvider, &postgresql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, `id<$1`, qCtx.WhereClause.String(ctx))
		assert.Equal(t, []any{val}, qCtx.Values)
	}()
}

func TestOrLess(t *testing.T) {
	val := 1
	q := NewQuery(`user`).Field("name").OrLess("id", val).OrIn("age", []any{10, 25})
	q.build()
	qCtx := q.ctxGetLambda()
	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeyDBProvider, &mysql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, "id<? OR age IN(?,?)", qCtx.WhereClause.String(ctx))
		assert.Equal(t, []any{val, 10, 25}, qCtx.Values)
	}()

	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeyDBProvider, &postgresql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, `id<$1 OR age IN($2,$3)`, qCtx.WhereClause.String(ctx))
		assert.Equal(t, []any{val, 10, 25}, qCtx.Values)
	}()
}

func TestLessOrEqual(t *testing.T) {
	val := 1
	q := NewQuery(`user`).Field("name").AndLessOrEqual("id", val)
	q.build()
	qCtx := q.ctxGetLambda()
	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeyDBProvider, &mysql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, "id<=?", qCtx.WhereClause.String(ctx))
		assert.Equal(t, []any{val}, qCtx.Values)
	}()

	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeyDBProvider, &postgresql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, `id<=$1`, qCtx.WhereClause.String(ctx))
		assert.Equal(t, []any{val}, qCtx.Values)
	}()
}

func TestAndIsNull(t *testing.T) {
	q := NewQuery(`user`).Field("name").AndIsNull("id")
	q.build()
	qCtx := q.ctxGetLambda()
	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeyDBProvider, &mysql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, "id IS NULL", qCtx.WhereClause.String(ctx))
		assert.Len(t, qCtx.Values, 0)
	}()

	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeyDBProvider, &postgresql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, `id IS NULL`, qCtx.WhereClause.String(ctx))
		assert.Len(t, qCtx.Values, 0)
	}()
}

func TestOrIsNull(t *testing.T) {
	q := NewQuery(`user`).Field("name").AndEqual("age", 100).OrIsNull("id")
	q.build()
	qCtx := q.ctxGetLambda()
	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeyDBProvider, &mysql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, "age=? OR id IS NULL", qCtx.WhereClause.String(ctx))
		assert.Len(t, qCtx.Values, 1)
	}()

	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeyDBProvider, &postgresql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, `age=$1 OR id IS NULL`, qCtx.WhereClause.String(ctx))
		assert.Len(t, qCtx.Values, 1)
	}()
}

func TestAndIsNotNull(t *testing.T) {
	q := NewQuery(`user`).Field("name").AndIsNotNull("id")
	q.build()
	qCtx := q.ctxGetLambda()
	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeyDBProvider, &mysql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, "id IS NOT NULL", qCtx.WhereClause.String(ctx))
		assert.Len(t, qCtx.Values, 0)
	}()

	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeyDBProvider, &postgresql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, `id IS NOT NULL`, qCtx.WhereClause.String(ctx))
		assert.Len(t, qCtx.Values, 0)
	}()
}

func TestOrIsNotNull(t *testing.T) {
	q := NewQuery(`user`).Field("name").AndEqual("age", 10).OrIsNotNull("id")
	q.build()
	qCtx := q.ctxGetLambda()
	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeyDBProvider, &mysql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, "age=? OR id IS NOT NULL", qCtx.WhereClause.String(ctx))
		assert.Len(t, qCtx.Values, 1)
	}()

	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeyDBProvider, &postgresql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, `age=$1 OR id IS NOT NULL`, qCtx.WhereClause.String(ctx))
		assert.Len(t, qCtx.Values, 1)
	}()
}

func TestAndLike(t *testing.T) {
	q := NewQuery(`user`).Field("name").AndLike("name", "ysl")
	q.build()
	qCtx := q.ctxGetLambda()
	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeyDBProvider, &mysql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, "name LIKE ?", qCtx.WhereClause.String(ctx))
		assert.Len(t, qCtx.Values, 1)
	}()

	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeyDBProvider, &postgresql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, `name LIKE $1`, qCtx.WhereClause.String(ctx))
		assert.Len(t, qCtx.Values, 1)
	}()
}

func TestOrLike(t *testing.T) {
	q := NewQuery(`user`).Field("name").AndGreater("id", 10).OrLike("name", "ysl")
	q.build()
	qCtx := q.ctxGetLambda()
	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeyDBProvider, &mysql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, "id>? OR name LIKE ?", qCtx.WhereClause.String(ctx))
		assert.Len(t, qCtx.Values, 2)
	}()

	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeyDBProvider, &postgresql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, `id>$1 OR name LIKE $2`, qCtx.WhereClause.String(ctx))
		assert.Len(t, qCtx.Values, 2)
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
		ctx := context.WithValue(context.TODO(), internal.CtxKeyDBProvider, &mysql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, "((id=? AND create_time>=?) AND (id>=? OR create_time<?))",
			qCtx.WhereClause.String(ctx))
	}()

	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeyDBProvider, &postgresql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, `((id=$1 AND create_time>=$2) AND (id>=$3 OR create_time<$4))`,
			qCtx.WhereClause.String(ctx))
	}()
}

func TestAndBetween(t *testing.T) {
	q := NewQuery(`user`).Field("name").AndBetween("id", 1, 10)
	q.build()
	qCtx := q.ctxGetLambda()

	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeyDBProvider, &mysql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, "id BETWEEN ? AND ?",
			qCtx.WhereClause.String(ctx))
	}()

	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeyDBProvider, &postgresql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, `id BETWEEN $1 AND $2`,
			qCtx.WhereClause.String(ctx))
	}()
}

func TestOrBetween(t *testing.T) {
	q := NewQuery(`user`).Field("name").AndEqual("age", 10).OrBetween("id", 1, 10)
	q.build()
	qCtx := q.ctxGetLambda()

	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeyDBProvider, &mysql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, "age=? OR id BETWEEN ? AND ?",
			qCtx.WhereClause.String(ctx))
	}()

	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeyDBProvider, &postgresql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, `age=$1 OR id BETWEEN $2 AND $3`,
			qCtx.WhereClause.String(ctx))
	}()
}

func TestNotBetween(t *testing.T) {
	q := NewQuery(`user`).Field("name").AndNotBetween("id", 1, 10)
	q.build()
	qCtx := q.ctxGetLambda()

	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeyDBProvider, &mysql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, "id NOT BETWEEN ? AND ?",
			qCtx.WhereClause.String(ctx))
	}()

	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeyDBProvider, &postgresql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, `id NOT BETWEEN $1 AND $2`,
			qCtx.WhereClause.String(ctx))
	}()
}

func TestOrNotBetween(t *testing.T) {
	q := NewQuery(`user`).Field("name").AndEqual("age", 100).OrNotBetween("id", 1, 10)
	q.build()
	qCtx := q.ctxGetLambda()

	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeyDBProvider, &mysql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, "age=? OR id NOT BETWEEN ? AND ?",
			qCtx.WhereClause.String(ctx))
	}()

	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeyDBProvider, &postgresql.Provider{})
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		assert.Equal(t, `age=$1 OR id NOT BETWEEN $2 AND $3`,
			qCtx.WhereClause.String(ctx))
	}()
}
