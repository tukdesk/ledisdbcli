package ledisdbcli

import (
	"github.com/tukdesk/goredis"
)

func (this *Client) Ping() error {
	reply, err := goredis.String(this.Do("PING"))

	if err != nil {
		return err
	}

	if reply != "PONG" {
		return newError(reply)
	}

	return nil
}

func (this *Client) FlushDB() error {
	_, err := this.Do("FLUSHDB")
	return err
}
