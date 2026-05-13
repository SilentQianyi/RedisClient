package redisClient

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

// ScriptCreateAndReturnValue 若键不存在则创建，返回当前值
var ScriptCreateAndReturnValue = redis.NewScript(`
	local val = redis.call("get", KEYS[1])
	if not val then
		redis.call("set", KEYS[1], ARGV[1])
		val = ARGV[1]
	end

	return val
`)

// CreateAndReturnValue 若键不存在则创建并返回新值，否则返回已有值。
func (r *RedisClient) CreateAndReturnValue(ctx context.Context, key string, value string) (string, error) {
	ret, e := ScriptCreateAndReturnValue.Run(ctx, r, []string{key}, value).Result()
	if e != nil {
		r.log.Error("CreateAndReturnValue ScriptCreateAndReturnValue.Run error! key[ %s ], err[ %s ]", key, e.Error())
		return "", ErrSystem
	}
	return ret.(string), nil
}

// ScriptSetCas CAS 比较并设置：仅当旧值匹配时才写入新值
var ScriptSetCas = redis.NewScript(`
	local expiration = tonumber(ARGV[3])

	local val = redis.call("get", KEYS[1])
	if not val then
		val = ""
	end

	if val == ARGV[1] then
		if expiration <= 0 then
			return redis.call("set", KEYS[1], ARGV[2])
		else
			return redis.call("setex", KEYS[1], tonumber(ARGV[3]), ARGV[2])
		end
	else
		return redis.error_reply("caserr")
	end
`)

// SetCas CAS 比较并设置：仅当 key 的当前值等于 oldval 时才写入 newval。
// 若 expiration > 0，设置过期时间。值不匹配返回 ErrSetCasConflict。
func (r *RedisClient) SetCas(ctx context.Context, key string, oldval, newval []byte, expiration time.Duration) error {
	e := ScriptSetCas.Run(ctx, r, []string{key}, oldval, newval, formatSec(expiration)).Err()
	if e != nil {
		r.log.Error("SetCas ScriptSetCas.Run error! key[ %s ], oldval[ %s ], newval[ %s ], expiration[ %s ], err[ %s ]", key, string(oldval), string(newval), expiration, e.Error())
		if e.Error() != "caserr" {
			return ErrSystem
		}
		return ErrSetCasConflict
	}
	return nil
}

func formatSec(dur time.Duration) int64 {
	if dur > 0 && dur < time.Second {
		return 1
	}
	return int64(dur / time.Second)
}

// ScriptGetAndDel 获取登录时间并删除键
var ScriptGetAndDel = redis.NewScript(`
	local val = redis.call("get", KEYS[1])
	if not val then
		return 0, redis.error_reply("notFind")
	end

	return val, redis.call("del", KEYS[1])
`)

// ScriptGetAndDel 原子获取键值并删除，用于登录时间的一次性读取。
func (r *RedisClient) GetAndDelValue(ctx context.Context, key string) (int64, error) {
	ret, e := ScriptGetAndDel.Run(ctx, r, []string{key}).Result()
	if e != nil {
		if e.Error() != "notFind" {
			r.log.Error("GetAndDelValue ScriptGetAndDel.Run error! key[ %s ], err[ %s ]", key, e.Error())
			return 0, ErrSetCasConflict
		}
		r.log.Info("GetAndDelValue key not found, key[ %s ]", key)
		return 0, nil
	}

	var val int64
	switch v := ret.(type) {
	case int64:
		val = v
	case string:
		val, e = strconv.ParseInt(v, 10, 64)
	default:
		e = errors.New("unknown type")
		r.log.Error("GetAndDelValue ScriptGetAndDel.Run error! unknown type, err[ %s ]", e.Error())
	}

	if e != nil {
		return 0, ErrSystem
	}
	return val, nil
}
