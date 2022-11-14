package str

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yeungsean/ysq-db/internal"
	"github.com/yeungsean/ysq-db/internal/provider/mysql"
	"github.com/yeungsean/ysq-db/internal/provider/postgresql"
	"github.com/yeungsean/ysq-db/pkg/field"
)

func TestFieldMySQL(t *testing.T) {
	ctx := context.WithValue(context.TODO(), internal.CtxKeySourceProvider, &mysql.Provider{})
	q := NewQuery(`user`).Field("id", field.WithDefaultValue(0))
	q.build()
	qCtx := q.ctxGetLambda()
	strs := internal.CtxGetSourceProvider(ctx).SelectFieldsQuote(qCtx.Fields...)
	assert.Len(t, strs, 1)
	assert.Equal(t, []string{"IFNULL(`id`,0) AS id"}, strs)
}

func TestFieldPostgreSQL(t *testing.T) {
	ctx := context.WithValue(context.TODO(), internal.CtxKeySourceProvider, &postgresql.Provider{})
	q := NewQuery(`user`).Field("id", field.WithDefaultValue(11))
	q.build()
	qCtx := q.ctxGetLambda()
	strs := internal.CtxGetSourceProvider(ctx).SelectFieldsQuote(qCtx.Fields...)
	assert.Len(t, strs, 1)
	assert.Equal(t, []string{`COALESCE("id",11) AS id`}, strs)
}

func TestSelectPrefix(t *testing.T) {
	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeySourceProvider, &mysql.Provider{})
		q := NewQuery(`user`).SelectPrefix("user", "id", "name", "gender")
		q.build()
		qCtx := q.ctxGetLambda()
		strs := internal.CtxGetSourceProvider(ctx).SelectFieldsQuote(qCtx.Fields...)
		assert.Len(t, strs, 3)
		assert.Equal(t, []string{"user.id", "user.name", "user.gender"}, strs)
	}()

	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeySourceProvider, &postgresql.Provider{})
		q := NewQuery(`user`).SelectPrefix("user", "id", "name", "gender")
		q.build()
		qCtx := q.ctxGetLambda()
		strs := internal.CtxGetSourceProvider(ctx).SelectFieldsQuote(qCtx.Fields...)
		assert.Len(t, strs, 3)
		assert.Equal(t, []string{"user.id", "user.name", "user.gender"}, strs)
	}()
}

func TestSelect(t *testing.T) {
	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeySourceProvider, &mysql.Provider{})
		q := NewQuery(`user`).Select("id", "name", "gender")
		q.build()
		qCtx := q.ctxGetLambda()
		strs := internal.CtxGetSourceProvider(ctx).SelectFieldsQuote(qCtx.Fields...)
		assert.Len(t, strs, 3)
		assert.Equal(t, []string{"`id`", "`name`", "`gender`"}, strs)
	}()

	func() {
		ctx := context.WithValue(context.TODO(), internal.CtxKeySourceProvider, &postgresql.Provider{})
		q := NewQuery(`user`).Select("id", "name", "gender")
		q.build()
		qCtx := q.ctxGetLambda()
		strs := internal.CtxGetSourceProvider(ctx).SelectFieldsQuote(qCtx.Fields...)
		assert.Len(t, strs, 3)
		assert.Equal(t, []string{`"id"`, `"name"`, `"gender"`}, strs)
	}()
}
