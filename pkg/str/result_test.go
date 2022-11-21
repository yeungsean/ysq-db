package str

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/yeungsean/ysq-db/pkg/dbprovider/mysql"
)

func TestScan(t *testing.T) {
	type TmpUser struct {
		ID   int64
		Name string
	}
	func() {
		tx := &sqlx.Tx{Tx: &sql.Tx{}}
		p1 := gomonkey.ApplyMethod(reflect.TypeOf(tx), "GetContext",
			func(*sqlx.Tx, context.Context, any, string, ...any) error {
				return nil
			})
		defer p1.Reset()
		q := NewQuery(context.TODO(), mysql.Provider{}).WithTx(tx)
		q.build()
		tu := &TmpUser{}
		err := q.Scan(tu)
		assert.Nil(t, err)
	}()

	func() {
		db := &sqlx.DB{}
		p1 := gomonkey.ApplyMethod(reflect.TypeOf(db), "GetContext",
			func(*sqlx.DB, context.Context, any, string, ...any) error {
				return nil
			})
		defer p1.Reset()
		q := NewQuery(context.TODO(), mysql.Provider{}).WithDB(db)
		q.build()
		tu := &TmpUser{}
		err := q.Scan(tu)
		assert.Nil(t, err)
	}()
}

func TestSlice(t *testing.T) {
	type TmpUser struct {
		ID   int64
		Name string
	}
	func() {
		tx := &sqlx.Tx{Tx: &sql.Tx{}}
		p1 := gomonkey.ApplyMethod(reflect.TypeOf(tx), "SelectContext",
			func(*sqlx.Tx, context.Context, any, string, ...any) error {
				return nil
			})
		defer p1.Reset()
		q := NewQuery(context.TODO(), mysql.Provider{}).WithTx(tx)
		q.build()
		tus := make([]*TmpUser, 0, 1)
		err := q.Slice(tus)
		assert.Nil(t, err)
	}()
	func() {
		db := &sqlx.DB{}
		p1 := gomonkey.ApplyMethod(reflect.TypeOf(db), "SelectContext",
			func(*sqlx.DB, context.Context, any, string, ...any) error {
				return nil
			})
		defer p1.Reset()
		q := NewQuery(context.TODO(), mysql.Provider{}).WithDB(db)
		q.build()
		tus := make([]*TmpUser, 0, 1)
		err := q.Slice(tus)
		assert.Nil(t, err)
	}()
}

func TestScanFields(t *testing.T) {
	row := &sqlx.Row{}
	p2 := gomonkey.ApplyFunc((*sqlx.Row).Scan, func(*sqlx.Row, ...any) error {
		return nil
	})
	defer p2.Reset()
	func() {
		tx := &sqlx.Tx{Tx: &sql.Tx{}}
		p1 := gomonkey.ApplyMethod(reflect.TypeOf(tx), "QueryRowxContext",
			func(*sqlx.Tx, context.Context, string, ...any) *sqlx.Row {
				return row
			})
		defer p1.Reset()
		q := NewQuery(context.TODO(), mysql.Provider{}).WithTx(tx)
		q.build()
		var id int64
		var name string
		err := q.ScanFields(&id, &name)
		assert.Nil(t, err)
	}()
	func() {
		db := &sqlx.DB{}
		p1 := gomonkey.ApplyMethod(reflect.TypeOf(db), "QueryRowxContext",
			func(*sqlx.DB, context.Context, string, ...any) *sqlx.Row {
				return row
			})
		defer p1.Reset()
		q := NewQuery(context.TODO(), mysql.Provider{}).WithDB(db)
		q.build()
		var id int64
		var name string
		err := q.ScanFields(&id, &name)
		assert.Nil(t, err)
	}()
}
