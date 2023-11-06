/*
 * Copyright icp-filing Author(https://houseme.github.io/bytedance/). All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * You can obtain one at https://github.com/houseme/bytedance.
 *
 */

package cache

import (
    "context"
    "time"
    
    "github.com/redis/go-redis/v9"
)

// Redis .redis cache
type Redis struct {
    conn redis.UniversalClient
}

// RedisOpts redis 连接属性
type RedisOpts struct {
    Host        string `yml:"host" json:"host"`
    Password    string `yml:"password" json:"password"`
    Database    int    `yml:"database" json:"database"`
    MaxIdle     int    `yml:"max_idle" json:"max_idle"`
    IdleTimeout int    `yml:"idle_timeout" json:"idle_timeout"`
}

// NewDefaultRedisOpts 实例化
func NewDefaultRedisOpts() *RedisOpts {
    return &RedisOpts{
        Host:     "127.0.0.1:6379",
        Password: "",
        Database: 0,
    }
}

// NewRedis 实例化
func NewRedis(_ context.Context, opts *RedisOpts) *Redis {
    return &Redis{conn: redis.NewUniversalClient(&redis.UniversalOptions{
        Addrs:           []string{opts.Host},
        DB:              opts.Database,
        Password:        opts.Password,
        ConnMaxIdleTime: time.Second * time.Duration(opts.IdleTimeout),
        MinIdleConns:    opts.MaxIdle,
    }),
    }
}

// SetConn 设置 conn
func (r *Redis) SetConn(conn redis.UniversalClient) {
    r.conn = conn
}

// Get 获取一个值
func (r *Redis) Get(ctx context.Context, key string) interface{} {
    result := r.conn.Get(ctx, key)
    if result.Err() != nil {
        return ""
    }
    return result.Val()
}

// Set 设置一个值
func (r *Redis) Set(ctx context.Context, key string, val interface{}, timeout time.Duration) error {
    return r.conn.SetEx(ctx, key, val, timeout).Err()
}

// IsExist 判断 key 是否存在
func (r *Redis) IsExist(ctx context.Context, key string) bool {
    return r.conn.Exists(ctx, key).Val() > 0
}

// Delete 删除
func (r *Redis) Delete(ctx context.Context, key string) error {
    return r.conn.Del(ctx, key).Err()
}
