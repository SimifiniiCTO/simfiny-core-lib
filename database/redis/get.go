package redis

import (
	"context"
	"fmt"

	"github.com/gomodule/redigo/redis"
)

// Get reads a value from the cache
func (c *Client) Get(ctx context.Context, key string) ([]byte, error) {
	txn := c.telemetrySdk.GetTraceFromContext(ctx)
	span := c.telemetrySdk.StartRedisDatastoreSegment(txn, RedisReadFromCacheTxn.String())
	defer span.End()

	// validate the key
	if key == "" {
		return nil, fmt.Errorf("empty key")
	}

	conn := c.pool.Get()
	defer conn.Close()
	value, err := redis.String(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}

	return []byte(value), nil
}

// GetManyFromCache reads a value from the cache
func (c *Client) GetMany(ctx context.Context, keys []string) ([][]byte, error) {
	txn := c.telemetrySdk.GetTraceFromContext(ctx)
	span := c.telemetrySdk.StartRedisDatastoreSegment(txn, RedisReadManyFromCacheTxn.String())
	defer span.End()

	// validate the key
	if len(keys) == 0 {
		return nil, fmt.Errorf("empty key")
	}

	conn := c.pool.Get()
	defer conn.Close()

	var ifaceSlice []interface{}
	for _, s := range keys {
		ifaceSlice = append(ifaceSlice, s)
	}

	values, err := redis.ByteSlices(conn.Do("MGET", ifaceSlice...))
	if err != nil {
		return nil, err
	}

	return values, nil
}
