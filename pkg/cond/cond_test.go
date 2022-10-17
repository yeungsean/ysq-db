package cond

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yeungsean/ysq-db/pkg/column"
)

func TestOr(t *testing.T) {
	cond := Or()
	assert.NotNil(t, cond)
}

func TestAnd(t *testing.T) {
	cond := And()
	assert.NotNil(t, cond)
	assert.True(t, cond.isAll)
}

func TestAdd(t *testing.T) {
	func() {
		cond := And().Add(column.New("column1")).Add(column.New("column2"))
		assert.Equal(t, "(column1=? AND column2=?)", cond.String())
	}()

	func() {
		cond := Or().Add(column.New("column1")).Add(column.New("column2"))
		assert.Equal(t, "(column1=? OR column2=?)", cond.String())
	}()

	func() {
		condOr := Or().Add(column.New("or_col1").GreaterEqual(1)).Add(column.New("or_col2").LessEqual(2))
		condAnd := And().Add(column.New("and_col1").IsNull()).Add(column.New("and_col2").IsNotNull())
		cond := Any().AddChildren(condAnd).AddChildren(condOr)
		assert.Equal(t, `((and_col1 IS NULL AND and_col2 IS NOT NULL) OR (or_col1>=? OR or_col2<=?))`, cond.String())
	}()
}
