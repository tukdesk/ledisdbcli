package ledisdbcli

import (
	"fmt"
)

func (this *Client) values2map(values []interface{}) (map[string][]byte, error) {
	if len(values)%2 != 0 {
		return nil, newError(fmt.Sprintf("invalid response, expected k-v pairs, got %d values", len(values)))
	}
	m := map[string][]byte{}
	for i := 0; i < len(values); i += 2 {
		if _, ok := values[i].([]byte); !ok {
			return nil, newError(fmt.Sprintf("invalid response, expected bytes, got %T", values[i]))
		}
		if _, ok := values[i+1].([]byte); !ok {
			return nil, newError(fmt.Sprintf("invalid response, expected bytes, got %T", values[i+1]))
		}
		m[string(values[i].([]byte))] = values[i+1].([]byte)
	}
	return m, nil
}

func (this *Client) values2bytesslice(values []interface{}) ([][]byte, error) {
	res := make([][]byte, len(values))
	for i, value := range values {
		if value == nil {
			res[i] = nil
			continue
		}
		if b, ok := value.([]byte); ok {
			res[i] = b
			continue
		}
		return nil, newError(fmt.Sprintf("invalid response, expected bytes or nil, got %T", value))
	}
	return res, nil
}
