package str

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yeungsean/ysq-db/internal/expr/join"
	"github.com/yeungsean/ysq-db/internal/expr/table"
	"github.com/yeungsean/ysq-db/pkg"
	"github.com/yeungsean/ysq-db/pkg/dbprovider/mysql"
	"github.com/yeungsean/ysq-db/pkg/dbprovider/postgresql"
	"github.com/yeungsean/ysq-db/pkg/option"
)

func TestLeftjoinExprString(t *testing.T) {
	je := join.Expr[string]{
		Expr: table.Expr[string]{
			Table:  "user_detail",
			Option: option.Option{Alias: "ud"},
		},
		Type:      join.Left,
		Condition: "u.id = ud.user_id",
	}

	func() {
		ctx := context.WithValue(context.TODO(), pkg.CtxKeyDBProvider, &mysql.Provider{})
		assert.Equal(t,
			" LEFT JOIN user_detail AS ud ON u.id = ud.user_id",
			je.String(ctx))
	}()

	func() {
		ctx := context.WithValue(context.TODO(), pkg.CtxKeyDBProvider, &postgresql.Provider{})
		assert.Equal(t,
			` LEFT JOIN user_detail AS ud ON u.id = ud.user_id`,
			je.String(ctx))
	}()
}

func TestRightjoinExprString(t *testing.T) {
	je := join.Expr[string]{
		Expr: table.Expr[string]{
			Table:  "user_detail",
			Option: option.Option{Alias: "ud"},
		},
		Type:      join.Right,
		Condition: "u.id = ud.user_id",
	}

	ctx := context.WithValue(context.TODO(), pkg.CtxKeyDBProvider, &mysql.Provider{})
	assert.Equal(t,
		" RIGHT JOIN user_detail AS ud ON u.id = ud.user_id",
		je.String(ctx))
}

func TestInnerjoinExprString(t *testing.T) {
	je := join.Expr[string]{
		Expr: table.Expr[string]{
			Table: "user_detail",
		},
		Type:      join.Inner,
		Condition: "u.id = user_id",
	}

	ctx := context.WithValue(context.TODO(), pkg.CtxKeyDBProvider, &mysql.Provider{})
	assert.Equal(t,
		" INNER JOIN user_detail ON u.id = user_id",
		je.String(ctx))
}

func TestLeftJoin(t *testing.T) {
	q := NewQuery(context.TODO()).Entity("user", option.WithAlias("u")).
		LeftJoin("user_detail", "u.id = ud.user_id", option.WithAlias("ud"))
	q.build()
	js := q.ctxGetLambda().joins
	assert.Equal(t, js[0].Type, join.Left)
}

func TestRightJoin(t *testing.T) {
	q := NewQuery(context.TODO()).Entity("user", option.WithAlias("u")).
		RightJoin("user_detail", "u.id = ud.user_id", option.WithAlias("ud"))
	q.build()
	js := q.ctxGetLambda().joins
	assert.Equal(t, js[0].Type, join.Right)
}

func TestInnerJoin(t *testing.T) {
	q := NewQuery(context.TODO()).Entity("user", option.WithAlias("u")).
		InnerJoin("user_detail", "u.id = ud.user_id", option.WithAlias("ud"))
	q.build()
	js := q.ctxGetLambda().joins
	assert.Equal(t, js[0].Type, join.Inner)
}
