package ledisdbcli

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotFound(t *testing.T) {
	c := setUp(t)
	defer tearDown(c)

	_, err := c.Get("not_found")
	assert.True(t, IsNotFound(err))
}
