package str

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yeungsean/ysq-db/pkg"
	"github.com/yeungsean/ysq-db/pkg/dbprovider/mysql"
	"github.com/yeungsean/ysq-db/pkg/dbprovider/postgresql"
)

func TestOrderAscMySQL(t *testing.T) {
	ctx := context.WithValue(context.TODO(), pkg.CtxKeyDBProvider, &mysql.Provider{})
	func() {
		q := NewQuery(context.TODO()).Entity("user").OrderAsc("id")
		q.build()
		assert.Len(t, q.ctxGetLambda().orders, 1)
		assert.Equal(t, "id ASC", q.ctxGetLambda().orders[0].String(ctx))
	}()

	func() {
		q := NewQuery(context.TODO()).Entity("user").OrderAsc("id").OrderDesc("create_time")
		q.build()
		assert.Len(t, q.ctxGetLambda().orders, 2)
		assert.Equal(t, "id ASC", q.ctxGetLambda().orders[0].String(ctx))
		assert.Equal(t, "create_time DESC", q.ctxGetLambda().orders[1].String(ctx))
	}()
}

func TestOrderDescMySQL(t *testing.T) {
	ctx := context.WithValue(context.TODO(), pkg.CtxKeyDBProvider, &mysql.Provider{})
	func() {
		q := NewQuery(context.TODO()).Entity("user").OrderDesc("id")
		q.build()
		assert.Len(t, q.ctxGetLambda().orders, 1)
		assert.Equal(t, "id DESC", q.ctxGetLambda().orders[0].String(ctx))
	}()

	func() {
		q := NewQuery(context.TODO()).Entity("user").OrderDesc("id").OrderAsc("create_time")
		q.build()
		assert.Len(t, q.ctxGetLambda().orders, 2)
		assert.Equal(t, "id DESC", q.ctxGetLambda().orders[0].String(ctx))
		assert.Equal(t, "create_time ASC", q.ctxGetLambda().orders[1].String(ctx))
	}()
}

func TestOrderAscPostgreSQL(t *testing.T) {
	ctx := context.WithValue(context.TODO(), pkg.CtxKeyDBProvider, &postgresql.Provider{})
	func() {
		q := NewQuery(context.TODO()).Entity("user").OrderAsc("id")
		q.build()
		assert.Len(t, q.ctxGetLambda().orders, 1)
		assert.Equal(t, `id ASC`, q.ctxGetLambda().orders[0].String(ctx))
	}()

	func() {
		q := NewQuery(context.TODO()).Entity("user").OrderAsc("id").OrderDesc("create_time")
		q.build()
		assert.Len(t, q.ctxGetLambda().orders, 2)
		assert.Equal(t, `id ASC`, q.ctxGetLambda().orders[0].String(ctx))
		assert.Equal(t, `create_time DESC`, q.ctxGetLambda().orders[1].String(ctx))
	}()
}

func TestOrderDescPostgreSQL(t *testing.T) {
	ctx := context.WithValue(context.TODO(), pkg.CtxKeyDBProvider, &postgresql.Provider{})
	func() {
		q := NewQuery(context.TODO()).Entity("user").OrderDesc("id")
		q.build()
		assert.Len(t, q.ctxGetLambda().orders, 1)
		assert.Equal(t, `id DESC`, q.ctxGetLambda().orders[0].String(ctx))
	}()

	func() {
		q := NewQuery(context.TODO()).Entity("user").OrderDesc("id").OrderAsc("create_time")
		q.build()
		assert.Len(t, q.ctxGetLambda().orders, 2)
		assert.Equal(t, `id DESC`, q.ctxGetLambda().orders[0].String(ctx))
		assert.Equal(t, `create_time ASC`, q.ctxGetLambda().orders[1].String(ctx))
	}()
}
