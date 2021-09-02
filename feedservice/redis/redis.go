package redis

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

const (
	NVT_FILENAME_POS           = 0
	NVT_REQUIRED_KEYS_POS      = 1
	NVT_MANDATORY_KEYS_POS     = 2
	NVT_EXCLUDED_KEYS_POS      = 3
	NVT_REQUIRED_UDP_PORTS_POS = 4
	NVT_REQUIRED_PORTS_POS     = 5
	NVT_DEPENDENCIES_POS       = 6
	NVT_TAGS_POS               = 7
	NVT_CVES_POS               = 8
	NVT_BIDS_POS               = 9
	NVT_XREFS_POS              = 10
	NVT_CATEGORY_POS           = 11
	NVT_FAMILY_POS             = 13
	NVT_NAME_POS               = 14
)

type RedisConnection struct {
	pool *redis.Pool
}

func (rc RedisConnection) GetList(db int, key string, start int, end int) ([]string, error) {
	conn := rc.pool.Get()
	defer conn.Close()

	conn.Do("SELECT", db)

	data, err := redis.Strings(conn.Do("LRANGE", key, start, end))
	if err != nil {
		return nil, fmt.Errorf("unable to get list: %s: %s", key, err)
	}
	return data, nil
}

func (rc RedisConnection) Close() error {
	return rc.pool.Close()
}

func NewRedisConnection(network string, address string) *RedisConnection {
	return &RedisConnection{
		pool: &redis.Pool{
			MaxIdle:     3,
			IdleTimeout: 240 * time.Second,

			Dial: func() (redis.Conn, error) {
				c, err := redis.Dial(network, address)
				if err != nil {
					return nil, err
				}
				return c, err
			},

			TestOnBorrow: func(c redis.Conn, t time.Time) error {
				_, err := c.Do("PING")
				return err
			},
		},
	}
}
