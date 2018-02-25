package mysql

import (
	"testing"

	asst "github.com/stretchr/testify/assert"
)

func TestAdapter_Placeholders(t *testing.T) {
	assert := asst.New(t)
	a := New()
	assert.Equal("", a.Placeholders(0))
	assert.Equal("?", a.Placeholders(1))
	assert.Equal("?,?,?", a.Placeholders(3))
}
