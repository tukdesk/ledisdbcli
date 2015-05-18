package ledisdbcli

import (
	"fmt"
	"time"

	"github.com/tukdesk/goredis"
)

func (this *Client) HDel(key interface{}, field ...interface{}) (int64, error) {
	if len(field) == 0 {
		return 0, nil
	}

	return goredis.Int64(this.Do("HDEL", this.appendKeyAndFields(key, field)...))
}

func (this *Client) HExists(key, field interface{}) (bool, error) {
	return goredis.Bool(this.Do("HEXISTS", key, field))
}

func (this *Client) HGet(key, field interface{}) ([]byte, error) {
	return goredis.Bytes(this.Do("HGET", key, field))
}

func (this *Client) HGetAll(key interface{}) (map[string][]byte, error) {
	values, err := goredis.Values(this.Do("HGETALL", key))
	if err != nil {
		return nil, err
	}
	return this.values2map(values)
}

func (this *Client) HKeys(key interface{}) ([]string, error) {
	return goredis.Strings(this.Do("HKEYS", key))
}

func (this *Client) HLen(key interface{}) (int64, error) {
	return goredis.Int64(this.Do("HLEN", key))
}

func (this *Client) HMGet(key interface{}, field ...interface{}) ([][]byte, error) {
	values, err := goredis.Values(this.Do("HMGET", this.appendKeyAndFields(key, field)...))
	if err != nil {
		return nil, err
	}
	return this.values2bytesslice(values)
}

func (this *Client) HMSet(key interface{}, fieldvalues ...interface{}) error {
	if len(fieldvalues) == 0 {
		return nil
	}
	if len(fieldvalues)%2 != 0 {
		return newError(fmt.Errorf("invalid args, expected field-value pairs, got %d values", len(fieldvalues)))
	}

	reply, err := this.Do("HMSET", this.appendKeyAndFields(key, fieldvalues)...)
	if err != nil {
		return nil
	}

	if reply != ReplyOK {
		return newError(reply)
	}

	return nil
}

func (this *Client) HSet(key, field, value interface{}) (bool, error) {
	return goredis.Bool(this.Do("HSET", key, field, value))
}

func (this *Client) HVals(key interface{}) ([][]byte, error) {
	values, err := goredis.Values(this.Do("HVALS", key))
	if err != nil {
		return nil, err
	}
	return this.values2bytesslice(values)
}

func (this *Client) HIncrBy(key, field interface{}, increment int64) (int64, error) {
	return goredis.Int64(this.Do("HINCRBY", key, field, increment))
}

func (this *Client) HClear(key interface{}) (int64, error) {
	return goredis.Int64(this.Do("HCLEAR", key))
}

func (this *Client) HMClear(keys ...interface{}) (int64, error) {
	if len(keys) == 0 {
		return 0, nil
	}
	return goredis.Int64(this.Do("HMCLEAR", keys...))
}

func (this *Client) HExpire(key interface{}, seconds time.Duration) (bool, error) {
	return goredis.Bool(this.Do("HEXPIRE", key, seconds/time.Second))
}

func (this *Client) HExpireAt(key interface{}, timestamp int64) (bool, error) {
	return goredis.Bool(this.Do("HEXPIREAT", key, timestamp))
}

func (this *Client) HTTL(key interface{}) (int64, error) {
	return goredis.Int64(this.Do("HTTL", key))
}

func (this *Client) HPersist(key interface{}) (bool, error) {
	return goredis.Bool(this.Do("HPERSIST", key))
}

func (this *Client) HKeyExists(key interface{}) (bool, error) {
	return goredis.Bool(this.Do("HKEYEXISTS", key))
}

func (this *Client) appendKeyAndFields(key interface{}, field []interface{}) []interface{} {
	return append([]interface{}{key}, field...)
}
