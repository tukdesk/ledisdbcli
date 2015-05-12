package ledisdbcli

import (
	"bytes"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKVScan(t *testing.T) {
	c := setUp(t)
	defer tearDown(c)

	count := 10
	for i := 0; i < count; i++ {
		if err := c.Set("key"+strconv.Itoa(i), []byte("value"+strconv.Itoa(i))); err != nil {
			assert.Nil(t, err)
		}
	}

	res, err := c.KVScan("", 100)
	assert.Nil(t, err)
	assert.Equal(t, count, res.Len())
	assert.False(t, res.HasMore())

	res1, err := c.KVScan("key"+strconv.Itoa(0), 5)
	assert.Nil(t, err)
	assert.Equal(t, 5, res1.Len())

	assert.True(t, bytes.Equal(res1.Keys()[0], []byte("key"+strconv.Itoa(1))))
	assert.True(t, bytes.Equal(res1.Keys()[res1.Len()-1], []byte("key"+strconv.Itoa(5))))

	assert.True(t, res1.HasMore())

	res2, err := c.KVScan(res1.Cursor(), 5)
	assert.Nil(t, err)
	assert.Equal(t, 4, res2.Len())
	assert.False(t, res2.HasMore())
}
