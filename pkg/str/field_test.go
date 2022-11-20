package str

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yeungsean/ysq-db/pkg"
	provider "github.com/yeungsean/ysq-db/pkg/dbprovider"
	"github.com/yeungsean/ysq-db/pkg/dbprovider/mysql"
	"github.com/yeungsean/ysq-db/pkg/dbprovider/postgresql"
	"github.com/yeungsean/ysq-db/pkg/field"
)

func TestFieldMySQL(t *testing.T) {
	ctx := context.WithValue(context.TODO(), pkg.CtxKeyDBProvider, &mysql.Provider{})
	q := NewQuery(context.TODO(), mysql.Provider{}).Entity("user").Field("id", field.WithDefaultValue(0), field.WithAlias("id"))
	q.build()
	qCtx := q.ctxGetLambda()
	strs := provider.CtxGet(ctx).SelectFields(qCtx.Fields...)
	assert.Len(t, strs, 1)
	assert.Equal(t, []string{"IFNULL(id,0) AS id"}, strs)
}

func TestFieldPostgreSQL(t *testing.T) {
	ctx := context.WithValue(context.TODO(), pkg.CtxKeyDBProvider, &postgresql.Provider{})
	q := NewQuery(context.TODO(), mysql.Provider{}).Entity("user").Field("id", field.WithDefaultValue(11), field.WithAlias("id"))
	q.build()
	qCtx := q.ctxGetLambda()
	strs := provider.CtxGet(ctx).SelectFields(qCtx.Fields...)
	assert.Len(t, strs, 1)
	assert.Equal(t, []string{`COALESCE(id,11) AS id`}, strs)
}

func TestSelectPrefix(t *testing.T) {
	func() {
		ctx := context.WithValue(context.TODO(), pkg.CtxKeyDBProvider, &mysql.Provider{})
		q := NewQuery(ctx).Entity("user").SelectPrefix("user", "id", "name", "gender")
		q.build()
		qCtx := q.ctxGetLambda()
		strs := provider.CtxGet(ctx).SelectFields(qCtx.Fields...)
		assert.Len(t, strs, 3)
		assert.Equal(t, []string{"user.id", "user.name", "user.gender"}, strs)
	}()

	func() {
		ctx := context.WithValue(context.TODO(), pkg.CtxKeyDBProvider, &postgresql.Provider{})
		q := NewQuery(context.TODO(), mysql.Provider{}).Entity("user").SelectPrefix("user", "id", "name", "gender")
		q.build()
		qCtx := q.ctxGetLambda()
		strs := provider.CtxGet(ctx).SelectFields(qCtx.Fields...)
		assert.Len(t, strs, 3)
		assert.Equal(t, []string{"user.id", "user.name", "user.gender"}, strs)
	}()
}

func TestSelect(t *testing.T) {
	func() {
		ctx := context.WithValue(context.TODO(), pkg.CtxKeyDBProvider, &mysql.Provider{})
		q := NewQuery(context.TODO(), mysql.Provider{}).Entity("user").Select("id", "name", "gender")
		q.build()
		qCtx := q.ctxGetLambda()
		strs := provider.CtxGet(ctx).SelectFields(qCtx.Fields...)
		assert.Len(t, strs, 3)
		assert.Equal(t, []string{"id", "name", "gender"}, strs)
	}()

	func() {
		ctx := context.WithValue(context.TODO(), pkg.CtxKeyDBProvider, &postgresql.Provider{})
		q := NewQuery(context.TODO(), mysql.Provider{}).Entity("user").Select("id", "name", "gender")
		q.build()
		qCtx := q.ctxGetLambda()
		strs := provider.CtxGet(ctx).SelectFields(qCtx.Fields...)
		assert.Len(t, strs, 3)
		assert.Equal(t, []string{`id`, `name`, `gender`}, strs)
	}()
}
