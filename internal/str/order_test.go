package str

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrderAsc(t *testing.T) {
	func() {
		q := NewQuery().Entity("user").OrderAsc("id")
		q.build()
		assert.Len(t, q.ctxGetLambda().orders, 1)
		assert.Equal(t, "id ASC", q.ctxGetLambda().orders[0].String())
	}()

	func() {
		q := NewQuery().Entity("user").OrderAsc("id").OrderDesc("create_time")
		q.build()
		assert.Len(t, q.ctxGetLambda().orders, 2)
		assert.Equal(t, "id ASC", q.ctxGetLambda().orders[0].String())
		assert.Equal(t, "create_time DESC", q.ctxGetLambda().orders[1].String())
	}()
}

func TestOrderDesc(t *testing.T) {
	func() {
		q := NewQuery().Entity("user").OrderDesc("id")
		q.build()
		assert.Len(t, q.ctxGetLambda().orders, 1)
		assert.Equal(t, "id DESC", q.ctxGetLambda().orders[0].String())
	}()

	func() {
		q := NewQuery().Entity("user").OrderDesc("id").OrderAsc("create_time")
		q.build()
		assert.Len(t, q.ctxGetLambda().orders, 2)
		assert.Equal(t, "id DESC", q.ctxGetLambda().orders[0].String())
		assert.Equal(t, "create_time ASC", q.ctxGetLambda().orders[1].String())
	}()
}
