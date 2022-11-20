package str

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yeungsean/ysq-db/pkg/dbprovider/mysql"
	"github.com/yeungsean/ysq-db/pkg/field"
	"github.com/yeungsean/ysq-db/pkg/option"
)

func TestSum(t *testing.T) {
	baseQ := NewQuery(context.TODO(), &mysql.Provider{}).Entity("user")
	func() {
		q := baseQ.As(`u`).Sum("age", field.Option{DefaultValue: "0"})
		q.build()
		qCtx := q.ctxGetLambda()
		assert.Len(t, qCtx.Fields, 1)
		assert.Equal(t, "SELECT IFNULL(SUM(age),0) AS age FROM user AS u", q.String())
	}()

	func() {
		q := baseQ.Sum("age", field.Option{})
		q.build()
		qCtx := q.ctxGetLambda()
		assert.Len(t, qCtx.Fields, 1)
		assert.Equal(t, "SELECT SUM(age) AS age FROM user", q.String())
	}()
}

func TestAvg(t *testing.T) {
	baseQ := NewQuery(context.TODO(), &mysql.Provider{}).Entity("user")
	func() {
		q := baseQ.Avg("age", field.Option{
			DefaultValue: "0",
			Option: option.Option{
				Alias: "AvgAge",
			},
		})
		q.build()
		qCtx := q.ctxGetLambda()
		assert.Len(t, qCtx.Fields, 1)
		assert.Equal(t, "SELECT IFNULL(AVG(age),0) AS AvgAge FROM user", q.String())
	}()

	func() {
		q := baseQ.Avg("age", field.Option{})
		q.build()
		qCtx := q.ctxGetLambda()
		assert.Len(t, qCtx.Fields, 1)
		assert.Equal(t, "SELECT AVG(age) AS age FROM user", q.String())
	}()
}

func TestMax(t *testing.T) {
	baseQ := NewQuery(context.TODO(), &mysql.Provider{}).Entity("user")
	func() {
		q := baseQ.Max("age", field.Option{DefaultValue: "0"})
		q.build()
		qCtx := q.ctxGetLambda()
		assert.Len(t, qCtx.Fields, 1)
		assert.Equal(t, "SELECT IFNULL(MAX(age),0) AS age FROM user", q.String())
	}()

	func() {
		q := baseQ.Max("age", field.Option{})
		q.build()
		qCtx := q.ctxGetLambda()
		assert.Len(t, qCtx.Fields, 1)
		assert.Equal(t, "SELECT MAX(age) AS age FROM user", q.String())
	}()
}

func TestMin(t *testing.T) {
	baseQ := NewQuery(context.TODO(), &mysql.Provider{}).Entity("user")
	func() {
		q := baseQ.Min("age", field.Option{DefaultValue: "0"})
		q.build()
		qCtx := q.ctxGetLambda()
		assert.Len(t, qCtx.Fields, 1)
		assert.Equal(t, "SELECT IFNULL(MIN(age),0) AS age FROM user", q.String())
	}()

	func() {
		q := baseQ.Min("age", field.Option{})
		q.build()
		qCtx := q.ctxGetLambda()
		assert.Len(t, qCtx.Fields, 1)
		assert.Equal(t, "SELECT MIN(age) AS age FROM user", q.String())
	}()
}

func TestCount(t *testing.T) {
	baseQ := NewQuery(context.TODO(), &mysql.Provider{}).Entity("user")
	func() {
		q := baseQ.Count(field.Option{
			Option: option.Option{Alias: "cnt"},
		})
		q.build()
		qCtx := q.ctxGetLambda()
		assert.Len(t, qCtx.Fields, 1)
		assert.Equal(t, "SELECT COUNT(1) AS cnt FROM user", q.String())
	}()

	func() {
		q := baseQ.Count()
		q.build()
		qCtx := q.ctxGetLambda()
		assert.Len(t, qCtx.Fields, 1)
		assert.Equal(t, "SELECT COUNT(1) FROM user", q.String())
	}()
}
