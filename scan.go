package ledisdbcli

import (
	"fmt"

	"github.com/tukdesk/goredis"
)

func (this *Client) KVScan(cursor interface{}, count int) (*KVScanResult, error) {
	return this.parseKVScanReply(this.Do("XSCAN", "KV", cursor, "count", count))
}

func (this *Client) parseScanReply(reply interface{}, replyErr error) ([]byte, [][]byte, error) {
	pieces, err := goredis.Values(reply, replyErr)
	if err != nil {
		return nil, nil, err
	}

	if len(pieces) != 2 {
		return nil, nil, newError(fmt.Sprintf("invalid response, expeted 2 parts, got %d", len(pieces)))
	}

	cursor, ok := pieces[0].([]byte)
	if !ok {
		return nil, nil, newError(fmt.Sprintf("invalid response, expected bytes, got %T", pieces[0]))
	}

	keysI, ok := pieces[1].([]interface{})
	if !ok {
		return nil, nil, newError(fmt.Sprintf("invalid response, expected slice of interface{}, got %T", pieces[1]))
	}

	keys := make([][]byte, len(keysI))
	for i, one := range keysI {
		key, ok := one.([]byte)
		if !ok {
			return nil, nil, newError(fmt.Sprintf("invalid key, expected bytes, got %T", one))
		}
		keys[i] = key
	}

	return cursor, keys, nil
}

func (this *Client) parseKVScanReply(reply interface{}, replyErr error) (*KVScanResult, error) {
	cursor, keys, err := this.parseScanReply(reply, replyErr)
	if err != nil {
		return nil, err
	}
	return &KVScanResult{
		cursor: cursor,
		keys:   keys,
	}, nil
}

type KVScanResult struct {
	cursor []byte
	keys   [][]byte
}

func (this *KVScanResult) HasMore() bool {
	return len(this.cursor) > 0
}

func (this *KVScanResult) Len() int {
	return len(this.keys)
}

func (this *KVScanResult) Cursor() []byte {
	return this.cursor
}

func (this *KVScanResult) Keys() [][]byte {
	return this.keys
}
