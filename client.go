package ledisdbcli

import (
	"github.com/tukdesk/goredis"
)

const (
	defaultIMaxdleConns = 4
)

type Client struct {
	c   *goredis.Client
	cfg Config
}

func New(cfg Config) (*Client, error) {
	c := goredis.NewClient(cfg.Addr, cfg.Password)
	if cfg.MaxIdleConns > 0 {
		c.SetMaxIdleConns(cfg.MaxIdleConns)
	}

	c.SetDBIndex(cfg.DBIndex)

	if _, err := c.Do("PING"); err != nil {
		return nil, err
	}

	client := &Client{c: c}
	return client, client.Ping()
}

func (this *Client) Do(cmd string, args ...interface{}) (interface{}, error) {
	return this.c.Do(cmd, args...)
}

func (this *Client) Close() {
	this.c.Close()
}
