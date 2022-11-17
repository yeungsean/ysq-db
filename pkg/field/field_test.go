package field

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	func() {
		f := New(Type("id"))
		assert.Equal(t, Type("id"), f.Name)
		assert.Empty(t, f.Prefix)
	}()
	func() {
		f := New(Type("u.id"))
		assert.Equal(t, Type("id"), f.Name)
		assert.Equal(t, "u", f.Prefix)
	}()
}
