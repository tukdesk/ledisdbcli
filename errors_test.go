package ledisdbcli

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotFound(t *testing.T) {
	c, app := setUp(t)
	defer tearDown(c, app)

	_, err := c.Get("not_found")
	assert.True(t, IsNotFound(err))
}
