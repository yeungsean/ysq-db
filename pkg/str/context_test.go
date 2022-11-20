package str

import (
	"context"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/yeungsean/ysq-db/pkg/dbprovider/mysql"
)

func TestCtxGetTx(t *testing.T) {
	q := NewQuery(context.TODO()).WithDBProvider(&mysql.Provider{})
	func() {
		q1 := q.WithTx(&sqlx.Tx{})
		q1.build()
		tx := q1.ctxGetTx()
		assert.NotNil(t, tx)
	}()

	func() {
		q.build()
		tx := q.ctxGetTx()
		assert.Nil(t, tx)
	}()
}

func TestCtxGetDB(t *testing.T) {
	q := NewQuery(context.TODO()).WithDBProvider(&mysql.Provider{})
	func() {
		q1 := q.WithDB(&sqlx.DB{})
		q1.build()
		db := q1.ctxGetDB()
		assert.NotNil(t, db)
	}()

	func() {
		q.build()
		db := q.ctxGetDB()
		assert.Nil(t, db)
	}()
}
