package str

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yeungsean/ysq-db/pkg/join"
)

func TestLeftjoinExprString(t *testing.T) {
	je := joinExpr[string]{
		tableExpr: tableExpr[string]{
			table: "user_detail",
			alias: "ud",
		},
		jt:        join.Left,
		condition: "u.id = ud.user_id",
	}

	assert.Equal(t,
		`LEFT JOIN user_detail AS ud ON u.id = ud.user_id`,
		je.String())
}

func TestRightjoinExprString(t *testing.T) {
	je := joinExpr[string]{
		tableExpr: tableExpr[string]{
			table: "user_detail",
			alias: "ud",
		},
		jt:        join.Right,
		condition: "u.id = ud.user_id",
	}

	assert.Equal(t,
		`RIGHT JOIN user_detail AS ud ON u.id = ud.user_id`,
		je.String())
}

func TestInnerjoinExprString(t *testing.T) {
	je := joinExpr[string]{
		tableExpr: tableExpr[string]{
			table: "user_detail",
		},
		jt:        join.Inner,
		condition: "u.id = user_id",
	}

	assert.Equal(t,
		`INNER JOIN user_detail ON u.id = user_id`,
		je.String())
}

func TestLeftJoin(t *testing.T) {
	q := NewQuery().Entity("user", "u").
		LeftJoin("user_detail", "u.id = ud.user_id", "ud")
	q.build()
	js := q.ctxGetLambda().joins
	assert.Equal(t, js[0].jt, join.Left)
}

func TestRightJoin(t *testing.T) {
	q := NewQuery().Entity("user", "u").
		RightJoin("user_detail", "u.id = ud.user_id", "ud")
	q.build()
	js := q.ctxGetLambda().joins
	assert.Equal(t, js[0].jt, join.Right)
}

func TestInnerJoin(t *testing.T) {
	q := NewQuery().Entity("user", "u").
		InnerJoin("user_detail", "u.id = ud.user_id", "ud")
	q.build()
	js := q.ctxGetLambda().joins
	assert.Equal(t, js[0].jt, join.Inner)
}
