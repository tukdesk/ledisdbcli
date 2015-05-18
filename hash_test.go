package ledisdbcli

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	c, app := setUp(t)
	defer tearDown(c, app)

	HKey := "hash_key"
	HSetField := "hset_field"
	HSetFieldNonExists := "hset_field_non_exists"
	HSetVal := "hset_val"

	affected, err := c.HSet(HKey, HSetField, HSetVal)
	assert.Nil(t, err)
	assert.True(t, affected)

	affected, err = c.HSet(HKey, HSetField, HSetVal)
	assert.Nil(t, err)
	assert.False(t, affected)

	val, err := c.HGet(HKey, HSetField)
	assert.Nil(t, err)
	assert.Equal(t, []byte(HSetVal), val)

	exists, err := c.HExists(HKey, HSetField)
	assert.Nil(t, err)
	assert.True(t, exists)

	_, err = c.HGet(HKey, HSetFieldNonExists)
	assert.True(t, IsNotFound(err))

	cleared, err := c.HClear(HKey)
	assert.Nil(t, err)
	assert.Equal(t, 1, cleared)

	exists, err = c.HExists(HKey, HSetField)
	assert.Nil(t, err)
	assert.False(t, exists)

	// hincrby
	HIncrByField := "hincrby_field"
	var incr int64 = 5
	after, err := c.HIncrBy(HKey, HIncrByField, incr)
	assert.Nil(t, err)
	assert.Equal(t, incr, after)
	c.HClear(HKey)

	// hget all
	HMF1 := "hgetall_field1"
	HMV1 := "hgetall_value1"
	HMF2 := "hgetall_field2"
	HMV2 := "hgetall_value2"
	HMF3 := "hgetall_field3"
	HMV3 := "hgetall_value3"

	err = c.HMSet(HKey, HMF1, HMV1, HMF2, HMV2, HMF3, HMV3)
	assert.Nil(t, err)

	l, err := c.HLen(HKey)
	assert.Nil(t, err)
	assert.Equal(t, 3, l)

	keys, err := c.HKeys(HKey)
	assert.Nil(t, err)
	assert.Len(t, keys, 3)

	vals, err := c.HVals(HKey)
	assert.Nil(t, err)
	assert.Len(t, vals, 3)

	mgetvals, err := c.HMGet(HKey, HMF1, HMF2, HMF3, "non_exists")
	assert.Nil(t, err)
	assert.Len(t, mgetvals, 4)
	assert.Equal(t, []byte(HMV1), mgetvals[0])
	assert.Equal(t, []byte(HMV2), mgetvals[1])
	assert.Equal(t, []byte(HMV3), mgetvals[2])
	assert.Nil(t, mgetvals[3])

	m, err := c.HGetAll(HKey)
	assert.Nil(t, err)
	assert.Len(t, m, 3)
	assert.Equal(t, []byte(HMV1), m[HMF1])
	assert.Equal(t, []byte(HMV2), m[HMF2])
	assert.Equal(t, []byte(HMV3), m[HMF3])

	num, err := c.HClear(HKey)
	assert.Nil(t, err)
	assert.Equal(t, 3, num)

	// hmclear
	num, err = c.HMClear("a", "b", "c", "d")
	assert.Nil(t, err)
	assert.Equal(t, 4, num)
}
