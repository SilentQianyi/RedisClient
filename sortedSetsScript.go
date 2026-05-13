package redisClient

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// ScriptZAddValue 获取成员分数并累加
var ScriptZAddValue = redis.NewScript(`
	local val = redis.call("zscore", KEYS[1], ARGV[1])
	if not val then
		val = tonumber(ARGV[2])
	else
		val = math.floor(val) + tonumber(ARGV[2])
	end

	return redis.call("zadd", KEYS[1], val, ARGV[1])
`)

// ZAddValue 累加有序集合成员的分数（先获取当前分数再累加，避免覆盖）。
func (r *RedisClient) ZAddValue(ctx context.Context, key, member string, value float64) error {
	e := ScriptZAddValue.Run(ctx, r, []string{key}, member, value).Err()
	if e != nil {
		r.log.Error("ZAddValue ScriptZAddValue.Run error! key[ %s ], member[ %s ], value[ %f ], err[ %s ]", key, member, value, e.Error())
		return ErrSystem
	}
	return nil
}
