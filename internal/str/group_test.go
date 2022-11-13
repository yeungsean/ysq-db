package str

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yeungsean/ysq-db/internal"
	"github.com/yeungsean/ysq-db/internal/expr/ops"
	"github.com/yeungsean/ysq-db/internal/provider/mysql"
	"github.com/yeungsean/ysq-db/internal/provider/postgresql"
)

func TestGroupBy(t *testing.T) {
	func() {
		q := NewQuery("user").GroupBy("gender,create_time")
		q.build()
		qCtx := q.ctxGetLambda()
		assert.Len(t, qCtx.Groups, 1)
	}()
}

func TestHavingOrMySQL(t *testing.T) {
	ctx := context.WithValue(context.TODO(), internal.CtxKeySourceProvider, &mysql.Provider{})
	func() {
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		q := NewQuery(`user`).HavingOr("gender", 1)
		q.build()
		qCtx := q.ctxGetLambda()
		assert.Equal(t, "`gender`=?", qCtx.HavingClause.String(ctx))
	}()

	func() {
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		q := NewQuery(`user`).HavingOr("gender", 1).HavingOr("id", 100, ops.GTE)
		q.build()
		qCtx := q.ctxGetLambda()
		assert.Equal(t, "`gender`=? OR `id`>=?", qCtx.HavingClause.String(ctx))
	}()
}

func TestHavingAndMySQL(t *testing.T) {
	ctx := context.WithValue(context.TODO(), internal.CtxKeySourceProvider, &mysql.Provider{})
	func() {
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		q := NewQuery(`user`).HavingOr("gender", 1)
		q.build()
		qCtx := q.ctxGetLambda()
		assert.Equal(t, "`gender`=?", qCtx.HavingClause.String(ctx))
	}()

	func() {
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		q := NewQuery(`user`).HavingAnd("gender", 1).HavingAnd("id", 100, ops.GTE)
		q.build()
		qCtx := q.ctxGetLambda()
		assert.Equal(t, "`gender`=? AND `id`>=?", qCtx.HavingClause.String(ctx))
	}()
}

func TestHavingOrPostgreSQL(t *testing.T) {
	ctx := context.WithValue(context.TODO(), internal.CtxKeySourceProvider, &postgresql.Provider{})
	func() {
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		q := NewQuery(`user`).HavingOr("gender", 1)
		q.build()
		qCtx := q.ctxGetLambda()
		assert.Equal(t, `"gender"=$1`, qCtx.HavingClause.String(ctx))
	}()

	func() {
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		q := NewQuery(`user`).HavingOr("gender", 1).HavingOr("id", 100, ops.GTE)
		q.build()
		qCtx := q.ctxGetLambda()
		assert.Equal(t, `"gender"=$1 OR "id">=$2`, qCtx.HavingClause.String(ctx))
	}()
}

func TestHavingAndPostgreSQL(t *testing.T) {
	ctx := context.WithValue(context.TODO(), internal.CtxKeySourceProvider, &postgresql.Provider{})
	func() {
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		q := NewQuery(`user`).HavingOr("gender", 1)
		q.build()
		qCtx := q.ctxGetLambda()
		assert.Equal(t, `"gender"=$1`, qCtx.HavingClause.String(ctx))
	}()

	func() {
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		q := NewQuery(`user`).HavingAnd("gender", 1).HavingAnd("id", 100, ops.GTE)
		q.build()
		qCtx := q.ctxGetLambda()
		assert.Equal(t, `"gender"=$1 AND "id">=$2`, qCtx.HavingClause.String(ctx))
	}()
}
