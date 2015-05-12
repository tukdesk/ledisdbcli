package ledisdbcli

import (
	"fmt"

	"github.com/tukdesk/goredis"
)

func IsNotFound(err error) bool {
	return err == goredis.ErrNil
}

func newError(reason interface{}) error {
	return fmt.Errorf("Ledis Client Error: %s", reason)
}
