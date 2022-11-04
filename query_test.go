package ysqdb

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yeungsean/ysq-db/pkg/cond"
	"github.com/yeungsean/ysq-db/pkg/statement"
)

type myEntity struct {
}

func (t myEntity) Table() string {
	return "entity_table"
}

func (t myEntity) Schema() string {
	return ""
}

func (t myEntity) String() string {
	if t.Schema() != "" {
		return fmt.Sprintf("%s.%s", t.Schema(), t.Table())
	}
	return t.Table()
}

func Test_QueryString(t *testing.T) {
	q := NewQuery(myEntity{}).
		Select("col1", "col2").
		Equal("col1", "good").
		GreaterOrEqual("col2", "very").
		Equal("col1", "boy", cond.Or)
	func() {
		ustr := q.
			IsNotNull("col1").
			OrderAsc("col1").
			OrderDesc("col2").
			String()
		assert.Equal(t,
			`SELECT col1,col2 FROM entity_table WHERE (col1=?) AND (col2>=?) OR (col1=?) AND (col1 IS NOT NULL) ORDER BY col1 ASC,col2 ASC`,
			ustr)
	}()

	//	func() {
	//		ustr := q.Equal("col3", "halou", cond.Or).Equal("col4", "louha", cond.Or).String()
	//		assert.Equal(t,
	//			`SELECT col1,col2 FROM entity_table WHERE (col1=?) AND (col2>=?) OR (col1=?) OR (col3=?) OR (col4=?)`,
	//			ustr)
	//	}()
}

func Test_String(t *testing.T) {
	q := NewQuery(myEntity{}).
		Select("col1", "col2").
		Equal("col1", "good").
		GreaterOrEqual("col2", "very").
		Equal("col1", "boy", cond.Or)
	q1 := q.
		IsNotNull("col1").
		OrderAsc("col1").
		OrderDesc("col2")
	q2 := q

	t.Error(q1.String(), q1.Values())
	t.Error(q2.String(), q2.Values())
}

func Test_build(t *testing.T) {
	q := NewQuery(myEntity{}).
		Select("col1", "col2").
		Equal("col1", "good").
		GreaterOrEqual("col2", "very").
		Equal("col1", "boy", cond.Or)

	hasExec := false
	q.Next = func() Iterator[myEntity] {
		hasExec = true
		return func() statement.Type {
			return statement.Nop
		}
	}

	assert.False(t, hasExec)
	assert.Equal(t, q.buildCnt, int32(0))
	q.build()
	q.build()
	q.build()
	assert.Equal(t, q.buildCnt, int32(1))
	assert.True(t, hasExec)
}
