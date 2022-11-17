package cond

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yeungsean/ysq-db/internal"
	"github.com/yeungsean/ysq-db/internal/expr/column"
	"github.com/yeungsean/ysq-db/internal/provider/mysql"
)

func TestOr(t *testing.T) {
	cond := Any()
	assert.NotNil(t, cond)
}

func TestAnd(t *testing.T) {
	cond := All()
	assert.NotNil(t, cond)
	assert.Equal(t, cond.logic, And)
}

func TestAdd(t *testing.T) {
	ctx := context.TODO()
	ctx = context.WithValue(ctx, internal.CtxKeySourceProvider, &mysql.Provider{})
	func() {
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		cond := All().
			Add(column.New("column1")).
			Add(column.New("column2"))
		assert.Equal(t, "(column1=? AND column2=?)", cond.String(ctx))
	}()

	func() {
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		cond := Any().
			Add(column.New("column1")).
			Add(column.New("column2"))
		assert.Equal(t, "(column1=? OR column2=?)", cond.String(ctx))
	}()

	func() {
		ctx = internal.CtxResetFilterColumnIndex(ctx)
		condOr := Any().
			Add(column.New("or_col1").GreaterEqual(1)).
			Add(column.New("or_col2").LessEqual(2))
		condAnd := All().
			Add(column.New("and_col1").IsNull()).
			Add(column.New("and_col2").IsNotNull())
		cond := Any().AddChildren(condAnd).AddChildren(condOr)
		assert.Equal(t, "((and_col1 IS NULL AND and_col2 IS NOT NULL) OR (or_col1>=? OR or_col2<=?))",
			cond.String(ctx))
	}()
}
