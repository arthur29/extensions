/*
 * Copyright (c) 2016 TFG Co <backend@tfgco.com>
 * Author: TFG Co <backend@tfgco.com>
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of
 * this software and associated documentation files (the "Software"), to deal in
 * the Software without restriction, including without limitation the rights to
 * use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
 * the Software, and to permit persons to whom the Software is furnished to do so,
 * subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
 * FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
 * COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
 * IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
 * CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package interfaces

import (
	"time"

	"github.com/go-redis/redis"
)

//RedisClient represents the contract for a redis client
type RedisClient interface {
	BLPop(timeout time.Duration, keys ...string) *redis.StringSliceCmd
	Close() error
	Eval(script string, keys []string, args ...interface{}) *redis.Cmd
	EvalSha(sha1 string, keys []string, args ...interface{}) *redis.Cmd
	Get(key string) *redis.StringCmd
	HGetAll(string) *redis.StringStringMapCmd
	HMSet(string, map[string]interface{}) *redis.StatusCmd
	Ping() *redis.StatusCmd
	RPush(key string, values ...interface{}) *redis.IntCmd
	ScriptExists(scripts ...string) *redis.BoolSliceCmd
	ScriptLoad(script string) *redis.StringCmd
	Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	SetNX(key string, value interface{}, expiration time.Duration) *redis.BoolCmd
	SPopN(key string, count int64) *redis.StringSliceCmd
	TxPipeline() redis.Pipeliner
	ZRangeByScore(key string, opt redis.ZRangeBy) *redis.StringSliceCmd
	RPopLPush(source string, destination string) *redis.StringCmd
	SAdd(key string, members ...interface{}) *redis.IntCmd
	SIsMember(key string, member interface{}) *redis.BoolCmd
}
