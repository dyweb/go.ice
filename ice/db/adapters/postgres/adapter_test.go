package postgres

import (
	"testing"

	asst "github.com/stretchr/testify/assert"
)

func TestAdapter_Placeholders(t *testing.T) {
	assert := asst.New(t)
	a := New()
	assert.Equal("", a.Placeholders(0))
	assert.Equal("$1", a.Placeholders(1))
	assert.Equal("$1,$2,$3", a.Placeholders(3))
}
