package str

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yeungsean/ysq-db/internal/expr/ops"
	"github.com/yeungsean/ysq-db/pkg"
	"github.com/yeungsean/ysq-db/pkg/dbprovider/mysql"
	"github.com/yeungsean/ysq-db/pkg/dbprovider/postgresql"
)

func TestGroupBy(t *testing.T) {
	func() {
		q := NewQuery(context.TODO(), mysql.Provider{}).Entity("user").GroupBy("gender,create_time")
		q.build()
		qCtx := q.ctxGetLambda()
		assert.Len(t, qCtx.Groups, 1)
	}()
}

func TestHavingOrMySQL(t *testing.T) {
	ctx := context.WithValue(context.TODO(), pkg.CtxKeyDBProvider, &mysql.Provider{})
	func() {
		ctx = pkg.CtxResetFilterColumnIndex(ctx)
		q := NewQuery(context.TODO(), &mysql.Provider{}).Entity("user").HavingOr("gender", 1)
		q.build()
		qCtx := q.ctxGetLambda()
		assert.Equal(t, "gender=?", qCtx.HavingClause.String(ctx))
	}()

	func() {
		ctx = pkg.CtxResetFilterColumnIndex(ctx)
		q := NewQuery(context.TODO(), &mysql.Provider{}).Entity("user").HavingOr("gender", 1).HavingOr("id", 100, ops.GTE)
		q.build()
		qCtx := q.ctxGetLambda()
		assert.Equal(t, "gender=? OR id>=?", qCtx.HavingClause.String(ctx))
	}()
}

func TestHavingAndMySQL(t *testing.T) {
	ctx := context.WithValue(context.TODO(), pkg.CtxKeyDBProvider, &mysql.Provider{})
	func() {
		ctx = pkg.CtxResetFilterColumnIndex(ctx)
		q := NewQuery(context.TODO()).Entity(`user`).HavingOr("gender", 1)
		q.build()
		qCtx := q.ctxGetLambda()
		assert.Equal(t, "gender=?", qCtx.HavingClause.String(ctx))
	}()

	func() {
		ctx = pkg.CtxResetFilterColumnIndex(ctx)
		q := NewQuery(context.TODO()).Entity(`user`).HavingAnd("gender", 1).HavingAnd("id", 100, ops.GTE)
		q.build()
		qCtx := q.ctxGetLambda()
		assert.Equal(t, "gender=? AND id>=?", qCtx.HavingClause.String(ctx))
	}()
}

func TestHavingOrPostgreSQL(t *testing.T) {
	ctx := context.WithValue(context.TODO(), pkg.CtxKeyDBProvider, &postgresql.Provider{})
	func() {
		ctx = pkg.CtxResetFilterColumnIndex(ctx)
		q := NewQuery(context.TODO()).Entity(`user`).HavingOr("gender", 1)
		q.build()
		qCtx := q.ctxGetLambda()
		assert.Equal(t, `gender=$1`, qCtx.HavingClause.String(ctx))
	}()

	func() {
		ctx = pkg.CtxResetFilterColumnIndex(ctx)
		q := NewQuery(context.TODO()).Entity(`user`).HavingOr("gender", 1).HavingOr("id", 100, ops.GTE)
		q.build()
		qCtx := q.ctxGetLambda()
		assert.Equal(t, `gender=$1 OR id>=$2`, qCtx.HavingClause.String(ctx))
	}()
}

func TestHavingAndPostgreSQL(t *testing.T) {
	ctx := context.WithValue(context.TODO(), pkg.CtxKeyDBProvider, &postgresql.Provider{})
	func() {
		ctx = pkg.CtxResetFilterColumnIndex(ctx)
		q := NewQuery(context.TODO()).Entity(`user`).HavingOr("gender", 1)
		q.build()
		qCtx := q.ctxGetLambda()
		assert.Equal(t, `gender=$1`, qCtx.HavingClause.String(ctx))
	}()

	func() {
		ctx = pkg.CtxResetFilterColumnIndex(ctx)
		q := NewQuery(context.TODO()).Entity(`user`).HavingAnd("gender", 1).HavingAnd("id", 100, ops.GTE)
		q.build()
		qCtx := q.ctxGetLambda()
		assert.Equal(t, `gender=$1 AND id>=$2`, qCtx.HavingClause.String(ctx))
	}()
}
