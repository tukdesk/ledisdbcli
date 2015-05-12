package ledisdbcli

import (
	"github.com/tukdesk/goredis"
)

func (this *Client) Set(key interface{}, value []byte) error {
	reply, err := goredis.String(this.Do("SET", key, value))

	if err != nil {
		return err
	}

	if reply != "OK" {
		return newError(reply)
	}

	return nil
}

func (this *Client) Get(key interface{}) ([]byte, error) {
	return goredis.Bytes(this.Do("GET", key))
}

func (this *Client) Del(key ...interface{}) (int64, error) {
	if len(key) == 0 {
		return 0, nil
	}

	return goredis.Int64(this.Do("DEL", key...))
}
