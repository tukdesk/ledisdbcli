package ledisdbcli

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setUp(t *testing.T) *Client {
	cfg := Config{
		Addr:    "127.0.0.1:16380",
		DBIndex: 15,
	}

	client, err := New(cfg)
	if err != nil {
		t.Fatal(err)
		return nil
	}

	if err := client.FlushDB(); err != nil {
		t.Fatal(err)
		return nil
	}

	return client
}

func tearDown(c *Client) {
	c.FlushDB()
	c.Close()
}

func TestKV(t *testing.T) {
	c := setUp(t)
	defer tearDown(c)

	key := "test_key"
	value := []byte("test_value")

	_, err := c.Get(key)
	assert.True(t, IsNotFound(err))

	err = c.Set(key, value)
	assert.Nil(t, err)

	got, err := c.Get(key)
	assert.Nil(t, err)
	assert.True(t, bytes.Equal(value, got))

	count, err := c.Del(key)
	assert.Nil(t, err)
	assert.Equal(t, count, 1)

	_, err = c.Get(key)
	assert.True(t, IsNotFound(err))
}
