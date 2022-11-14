package str

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yeungsean/ysq-db/pkg/field"
)

func TestQueryStringFields(t *testing.T) {
	func() {
		q := NewQuery().Entity("user").OrderAsc("id")
		str := q.String()
		assert.Equal(t, "SELECT * FROM `user` ORDER BY `id` ASC", str)
	}()

	func() {
		q := NewQuery().Entity("user").Select("id", "name").OrderAsc("id")
		str := q.String()
		assert.Equal(t, "SELECT `id`,`name` FROM `user` ORDER BY `id` ASC", str)
	}()

	func() {
		q := NewQuery("user").As("u").Select("id", "name")
		str := q.String()
		assert.Equal(t, "SELECT `id`,`name` FROM `user` AS u", str)
	}()

	func() {
		q := NewQuery("user").As("u").SelectPrefix("u", "id", "name")
		str := q.String()
		assert.Equal(t, "SELECT u.id,u.name FROM `user` AS u", str)
	}()

	func() {
		q := NewQuery("user").As("u").SelectPrefix("u", "id", "name").
			Field("age", field.WithDefaultValue(0))
		str := q.String()
		assert.Equal(t, "SELECT u.id,u.name,IFNULL(`age`,0) AS age FROM `user` AS u", str)
	}()
}

func TestQueryLazyLoad(t *testing.T) {
	q1 := NewQuery("user").As("u").InnerJoin("user_detail", "u.id=ud.id", "ud")
	q2 := q1.SelectPrefix("u", "id", "name")
	assert.Equal(t, "SELECT * FROM `user` AS u INNER JOIN `user_detail` AS ud ON u.id=ud.id", q1.String())
	assert.Equal(t, "SELECT u.id,u.name FROM `user` AS u INNER JOIN `user_detail` AS ud ON u.id=ud.id", q2.String())

	q3 := q1.Field("id", field.WithDefaultValue(999))
	assert.Equal(t, "SELECT IFNULL(`id`,999) AS id FROM `user` AS u INNER JOIN `user_detail` AS ud ON u.id=ud.id", q3.String())
}

func TestQueryStringJoin(t *testing.T) {
	func() {
		q := NewQuery("user").As("u").LeftJoin("user_detail", "u.id=ud.id", "ud")
		str := q.String()
		assert.Equal(t, "SELECT * FROM `user` AS u LEFT JOIN `user_detail` AS ud ON u.id=ud.id", str)
	}()

	func() {
		q := NewQuery("user").As("u").InnerJoin("user_detail", "u.id=ud.id", "ud")
		str := q.String()
		assert.Equal(t, "SELECT * FROM `user` AS u INNER JOIN `user_detail` AS ud ON u.id=ud.id", str)
	}()

	func() {
		q := NewQuery("user").As("u").RightJoin("user_detail", "u.id=ud.id", "ud")
		str := q.String()
		assert.Equal(t, "SELECT * FROM `user` AS u RIGHT JOIN `user_detail` AS ud ON u.id=ud.id", str)
	}()

	func() {
		q := NewQuery("user").As("u").RightJoin("user_detail", "u.id=user_detail.id")
		str := q.String()
		assert.Equal(t, "SELECT * FROM `user` AS u RIGHT JOIN `user_detail` ON u.id=user_detail.id", str)
	}()

	func() {
		q := NewQuery("user").As("u").
			SelectPrefix("u", "id", "name").
			InnerJoin("user_detail", "u.id=ud.id", "ud").
			SelectPrefix("ud", "addr").
			Field("age", field.WithDefaultValue(0))
		str := q.String()
		assert.Equal(t,
			"SELECT u.id,u.name,ud.addr,IFNULL(`age`,0) AS age FROM `user` AS u INNER JOIN `user_detail` AS ud ON u.id=ud.id",
			str)
	}()
}

func TestQueryStringWhere(t *testing.T) {
	func() {
		q := NewQuery("user").As("u").
			LeftJoin("user_detail", "u.id=ud.id", "ud").
			Equal("u.id", 1)
		str := q.String()
		assert.Equal(t,
			"SELECT * FROM `user` AS u LEFT JOIN `user_detail` AS ud ON u.id=ud.id WHERE u.id=?",
			str)
	}()
}
