package ledisdbcli

import (
	"bytes"
	"os"
	"testing"

	"github.com/siddontang/ledisdb/config"
	"github.com/siddontang/ledisdb/server"
	"github.com/stretchr/testify/assert"
)

const (
	testDBAddr    = "127.0.0.1:16380"
	testDBDataDir = "./_testdata"
)

func setUp(t *testing.T) (*Client, *server.App) {
	ledisCfg := config.NewConfigDefault()
	ledisCfg.Addr = testDBAddr
	ledisCfg.DataDir = testDBDataDir
	app, err := server.NewApp(ledisCfg)

	go app.Run()

	if err != nil {
		t.Fatal(err)
		return nil, nil
	}

	cfg := Config{
		Addr:    testDBAddr,
		DBIndex: 15,
	}

	client, err := New(cfg)
	if err != nil {
		t.Fatal(err)
		return nil, nil
	}

	if err := client.FlushDB(); err != nil {
		t.Fatal(err)
		return nil, nil
	}

	return client, app
}

func tearDown(c *Client, app *server.App) {
	c.FlushDB()
	c.Close()
	app.Close()
	os.RemoveAll(testDBDataDir)
}

func TestKV(t *testing.T) {
	c, app := setUp(t)
	defer tearDown(c, app)

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
