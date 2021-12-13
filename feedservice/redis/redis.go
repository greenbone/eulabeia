// Copyright (C) 2021 Greenbone Networks GmbH
//
// SPDX-License-Identifier: AGPL-3.0-or-later
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

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
	NVT_FAMILY_POS             = 12
	NVT_NAME_POS               = 13
)

type RedisConnection struct {
	pool *redis.Pool
}

func (rc RedisConnection) GetList(
	db int,
	key string,
	start int,
	end int,
) ([]string, error) {
	conn := rc.pool.Get()
	defer conn.Close()

	conn.Do("SELECT", db)

	data, err := redis.Strings(conn.Do("LRANGE", key, start, end))
	if err != nil {
		return nil, fmt.Errorf("unable to get list: %s: %s", key, err)
	}
	return data, nil
}

func (rc RedisConnection) GetKeys(db int, filter string) ([]string, error) {
	conn := rc.pool.Get()
	defer conn.Close()

	conn.Do("SELECT", db)

	data, err := redis.Strings(conn.Do("KEYS", filter))
	if err != nil {
		return nil, fmt.Errorf(
			"unable to get keys with filter %s: %s",
			filter,
			err,
		)
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
