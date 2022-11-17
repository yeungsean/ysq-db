package str

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yeungsean/ysq-db/internal/expr/column"
	"github.com/yeungsean/ysq-db/internal/expr/ops"
	"github.com/yeungsean/ysq-db/pkg"
	"github.com/yeungsean/ysq-db/pkg/field"
)

func TestQueryStringFields(t *testing.T) {
	func() {
		q := NewQuery().Entity("user").OrderAsc("id")
		str := q.String()
		assert.Equal(t, "SELECT * FROM user ORDER BY id ASC", str)
	}()

	func() {
		q := NewQuery().Entity("user").Select("id", "name").OrderAsc("id")
		str := q.String()
		assert.Equal(t, "SELECT id,name FROM user ORDER BY id ASC", str)
	}()

	func() {
		q := NewQuery("user").As("u").Select("id", "name")
		str := q.String()
		assert.Equal(t, "SELECT id,name FROM user AS u", str)
	}()

	func() {
		q := NewQuery("user").As("u").SelectPrefix("u", "id", "name")
		str := q.String()
		assert.Equal(t, "SELECT u.id,u.name FROM user AS u", str)
	}()

	func() {
		q := NewQuery("user").As("u").SelectPrefix("u", "id", "name").
			Field("age", field.WithDefaultValue(0))
		str := q.String()
		assert.Equal(t, "SELECT u.id,u.name,IFNULL(age,0) AS age FROM user AS u", str)
	}()

	func() {
		q := NewQuery("user").As("u").SelectPrefix("u", "id", "name").
			Field("age", field.WithDefaultValue(0), field.WithQuote(), field.WithAlias("age1"))
		str := q.String()
		assert.Equal(t, "SELECT u.id,u.name,IFNULL(`age`,0) AS age1 FROM user AS u", str)
	}()
}

func TestQueryLazyLoad(t *testing.T) {
	q1 := NewQuery("user").As("u").InnerJoin("user_detail", "u.id=ud.id", pkg.WithAlias("ud"))
	q2 := q1.SelectPrefix("u", "id", "name")
	assert.Equal(t, "SELECT * FROM user AS u INNER JOIN user_detail AS ud ON u.id=ud.id", q1.String())
	assert.Equal(t, "SELECT u.id,u.name FROM user AS u INNER JOIN user_detail AS ud ON u.id=ud.id", q2.String())

	q3 := q1.Field("id", field.WithDefaultValue(999))
	assert.Equal(t, "SELECT IFNULL(id,999) AS id FROM user AS u INNER JOIN user_detail AS ud ON u.id=ud.id", q3.String())
}

func TestQueryStringJoin(t *testing.T) {
	func() {
		q := NewQuery("user").As("u").LeftJoin("user_detail", "u.id=ud.id", pkg.WithAlias("ud"))
		str := q.String()
		assert.Equal(t, "SELECT * FROM user AS u LEFT JOIN user_detail AS ud ON u.id=ud.id", str)
	}()

	func() {
		q := NewQuery("user").As("u").InnerJoin("user_detail", "u.id=ud.id", pkg.WithAlias("ud"))
		str := q.String()
		assert.Equal(t, "SELECT * FROM user AS u INNER JOIN user_detail AS ud ON u.id=ud.id", str)
	}()

	func() {
		q := NewQuery("user").As("u").RightJoin("user_detail", "u.id=ud.id", pkg.WithAlias("ud"))
		str := q.String()
		assert.Equal(t, "SELECT * FROM user AS u RIGHT JOIN user_detail AS ud ON u.id=ud.id", str)
	}()

	func() {
		q := NewQuery("user").As("u").RightJoin("user_detail", "u.id=user_detail.id")
		str := q.String()
		assert.Equal(t, "SELECT * FROM user AS u RIGHT JOIN user_detail ON u.id=user_detail.id", str)
	}()

	func() {
		q := NewQuery("user").As("u").
			SelectPrefix("u", "id", "name").
			InnerJoin("user_detail", "u.id=ud.id", pkg.WithAlias("ud")).
			SelectPrefix("ud", "addr").
			Field("age", field.WithDefaultValue(0))
		str := q.String()
		assert.Equal(t,
			"SELECT u.id,u.name,ud.addr,IFNULL(age,0) AS age FROM user AS u INNER JOIN user_detail AS ud ON u.id=ud.id",
			str)
	}()
}

func TestQueryStringWhere(t *testing.T) {
	func() {
		q := NewQuery("user").As("u").
			LeftJoin("user_detail", "u.id=ud.id", pkg.WithAlias("ud")).
			AndEqual("u.id", 1)
		str := q.String()
		assert.Equal(t,
			"SELECT * FROM user AS u LEFT JOIN user_detail AS ud ON u.id=ud.id WHERE u.id=?",
			str)
		assert.Equal(t, []any{1}, q.Values())
	}()

	func() {
		baseQ := NewQuery(`user`).As(`u`).LeftJoin("user_detail", "u.id=ud.id", pkg.WithAlias("ud")).AndEqual("u.id", 1)
		q1 := baseQ.AndGreaterOrEqual("ud.age", 10)
		q2 := baseQ.AndLess("ud.age", 10).AndIsNotNull("ud.addr")
		q3 := baseQ.AndBetween("ud.age", 1, 10)
		q4 := q3.AndIn("job", []any{"it", "teacher"})

		assert.Equal(t,
			"SELECT * FROM user AS u LEFT JOIN user_detail AS ud ON u.id=ud.id WHERE u.id=?",
			baseQ.String())
		assert.Equal(t, []any{1}, baseQ.Values())

		assert.Equal(t,
			"SELECT * FROM user AS u LEFT JOIN user_detail AS ud ON u.id=ud.id WHERE u.id=? AND ud.age>=?",
			q1.String())
		assert.Equal(t, []any{1, 10}, q1.Values())

		assert.Equal(t,
			"SELECT * FROM user AS u LEFT JOIN user_detail AS ud ON u.id=ud.id WHERE u.id=? AND ud.age<? AND ud.addr IS NOT NULL",
			q2.String())
		assert.Equal(t, []any{1, 10}, q2.Values())

		assert.Equal(t,
			"SELECT * FROM user AS u LEFT JOIN user_detail AS ud ON u.id=ud.id WHERE u.id=? AND ud.age BETWEEN ? AND ?",
			q3.String())
		assert.Equal(t, []any{1, 1, 10}, q3.Values())

		assert.Equal(t,
			"SELECT * FROM user AS u LEFT JOIN user_detail AS ud ON u.id=ud.id WHERE u.id=? AND ud.age BETWEEN ? AND ? AND job IN(?,?)",
			q4.String())
		assert.Equal(t, []any{1, 1, 10, "it", "teacher"}, q4.Values())
	}()
}

func TestQueryGroups(t *testing.T) {
	func() {
		q1 := NewQuery(`user`).As(`u`).GroupBy("age").Select("age").Count(field.Option{
			Option: pkg.Option{Alias: "cnt"},
		})
		assert.Equal(t,
			"SELECT age,COUNT(1) AS cnt FROM user AS u GROUP BY age",
			q1.String())
	}()

	func() {
		q1 := NewQuery(`user`).As(`u`).GroupBy("age").Count(field.Option{
			Option: pkg.Option{Alias: "cnt"},
		})
		assert.Equal(t,
			"SELECT COUNT(1) AS cnt FROM user AS u GROUP BY age",
			q1.String())
	}()

	func() {
		q1 := NewQuery(`user`).As(`u`).GroupBy("age").Max("age")
		assert.Equal(t,
			"SELECT MAX(age) AS age FROM user AS u GROUP BY age",
			q1.String())
	}()

	func() {
		q1 := NewQuery(`user`).As(`u`).GroupBy("age").Min("age")
		assert.Equal(t,
			"SELECT MIN(age) AS age FROM user AS u GROUP BY age",
			q1.String())
	}()

	func() {
		q1 := NewQuery(`user`).As(`u`).GroupBy("age").Avg("age")
		assert.Equal(t,
			"SELECT AVG(age) AS age FROM user AS u GROUP BY age",
			q1.String())
	}()

	func() {
		q1 := NewQuery(`user`).As(`u`).GroupBy("age").Sum("age")
		assert.Equal(t,
			"SELECT SUM(age) AS age FROM user AS u GROUP BY age",
			q1.String())

		q2 := NewQuery(`user`).As(`u`).GroupBy("age").Sum("age", field.Option{
			Option: pkg.Option{
				Quote: true,
			},
		})
		assert.Equal(t,
			"SELECT SUM(`age`) AS age FROM user AS u GROUP BY age",
			q2.String())
	}()
}

func TestQueryGroupByHaving(t *testing.T) {
	func() {
		q1 := NewQuery(`user`).As(`u`).GroupBy("age").Sum("age").HavingAnd("age", 10, ops.GT)
		q2 := q1.HavingOr("SUM(age)", 100, ops.GTE)
		assert.Equal(t,
			"SELECT SUM(age) AS age FROM user AS u GROUP BY age HAVING age>?",
			q1.String())
		assert.Equal(t,
			"SELECT SUM(age) AS age FROM user AS u GROUP BY age HAVING age>? OR SUM(age)>=?",
			q2.String())
	}()
}

func TestQueryQuote(t *testing.T) {
	func() {
		q := NewQuery().Entity("user", pkg.WithQuote()).
			AndEqual("gender", "male", column.WithQuote()).
			AndGreater("id", 10).
			LimitOffset(10, 0).
			Select("id", "gender", "name", "age").
			Field("height", field.WithAlias("h"), field.WithQuote()).
			Field("weight", field.WithQuote())
		assert.Equal(t,
			"SELECT id,gender,name,age,`height` AS h,`weight` FROM `user` WHERE `gender`=? AND id>? LIMIT 10",
			q.String())
	}()

	func() {
		q := NewQuery("user").As("u").
			LeftJoin("user_role", "u.id=ur.user_id", pkg.WithAlias("ur")).
			AndGreaterOrEqual("u.id", 100).
			SelectPrefix("u", "id", "name", "age").
			Select("ur.role_id").
			LimitOffset(10, 10)
		assert.Equal(t,
			"SELECT u.id,u.name,u.age,ur.role_id FROM user AS u LEFT JOIN user_role AS ur ON u.id=ur.user_id WHERE u.id>=? LIMIT 10 OFFSET 10",
			q.String())
	}()
}
