package str

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLimit(t *testing.T) {
	q := NewQuery().Entity("user", "u").Limit(10)
	q.build()
	assert.Equal(t, 10, q.ctxGetLambda().LimitCount)
	assert.Equal(t, 0, q.ctxGetLambda().LimitOffset)
}

func TestOffset(t *testing.T) {
	q := NewQuery().Entity("user", "u").Offset(10)
	q.build()
	assert.Equal(t, 10, q.ctxGetLambda().LimitOffset)
	assert.Equal(t, 0, q.ctxGetLambda().LimitCount)
}

func TestLimitOffset(t *testing.T) {
	q := NewQuery().Entity("user", "u").LimitOffset(10, 2)
	q.build()
	assert.Equal(t, 2, q.ctxGetLambda().LimitOffset)
	assert.Equal(t, 10, q.ctxGetLambda().LimitCount)
}
