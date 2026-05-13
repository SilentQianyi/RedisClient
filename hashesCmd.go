package redisClient

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

// HSet 设置哈希字段
func (r *RedisClient) HSet(ctx context.Context, key string, values ...interface{}) *redis.IntCmd {
	if r.client != nil {
		return r.client.HSet(ctx, key, values...)
	}
	return r.clusterClient.HSet(ctx, key, values...)
}

// HSetNX 仅当字段不存在时设置哈希字段
func (r *RedisClient) HSetNX(ctx context.Context, key, field string, value interface{}) *redis.BoolCmd {
	if r.client != nil {
		return r.client.HSetNX(ctx, key, field, value)
	}
	return r.clusterClient.HSetNX(ctx, key, field, value)
}

// HMSet 批量设置哈希字段
func (r *RedisClient) HMSet(ctx context.Context, key string, values ...interface{}) *redis.BoolCmd {
	if r.client != nil {
		return r.client.HMSet(ctx, key, values...)
	}
	return r.clusterClient.HMSet(ctx, key, values...)
}

// HGet 获取哈希字段值
func (r *RedisClient) HGet(ctx context.Context, key, field string) *redis.StringCmd {
	if r.client != nil {
		return r.client.HGet(ctx, key, field)
	}
	return r.clusterClient.HGet(ctx, key, field)
}

// HMGet 批量获取哈希字段值
func (r *RedisClient) HMGet(ctx context.Context, key string, fields ...string) *redis.SliceCmd {
	if r.client != nil {
		return r.client.HMGet(ctx, key, fields...)
	}
	return r.clusterClient.HMGet(ctx, key, fields...)
}

// HGetAll 获取所有哈希字段
func (r *RedisClient) HGetAll(ctx context.Context, key string) *redis.MapStringStringCmd {
	if r.client != nil {
		return r.client.HGetAll(ctx, key)
	}
	return r.clusterClient.HGetAll(ctx, key)
}

// HDel 删除哈希字段
func (r *RedisClient) HDel(ctx context.Context, key string, fields ...string) *redis.IntCmd {
	if r.client != nil {
		return r.client.HDel(ctx, key, fields...)
	}
	return r.clusterClient.HDel(ctx, key, fields...)
}

// HExists 检查哈希字段是否存在
func (r *RedisClient) HExists(ctx context.Context, key, field string) *redis.BoolCmd {
	if r.client != nil {
		return r.client.HExists(ctx, key, field)
	}
	return r.clusterClient.HExists(ctx, key, field)
}

// HKeys 获取所有哈希字段名
func (r *RedisClient) HKeys(ctx context.Context, key string) *redis.StringSliceCmd {
	if r.client != nil {
		return r.client.HKeys(ctx, key)
	}
	return r.clusterClient.HKeys(ctx, key)
}

// HVals 获取所有哈希字段值
func (r *RedisClient) HVals(ctx context.Context, key string) *redis.StringSliceCmd {
	if r.client != nil {
		return r.client.HVals(ctx, key)
	}
	return r.clusterClient.HVals(ctx, key)
}

// HLen 获取哈希字段数量
func (r *RedisClient) HLen(ctx context.Context, key string) *redis.IntCmd {
	if r.client != nil {
		return r.client.HLen(ctx, key)
	}
	return r.clusterClient.HLen(ctx, key)
}

// HIncrBy 哈希字段整数值自增
func (r *RedisClient) HIncrBy(ctx context.Context, key, field string, incr int64) *redis.IntCmd {
	if r.client != nil {
		return r.client.HIncrBy(ctx, key, field, incr)
	}
	return r.clusterClient.HIncrBy(ctx, key, field, incr)
}

// HIncrByFloat 哈希字段浮点值自增
func (r *RedisClient) HIncrByFloat(ctx context.Context, key, field string, incr float64) *redis.FloatCmd {
	if r.client != nil {
		return r.client.HIncrByFloat(ctx, key, field, incr)
	}
	return r.clusterClient.HIncrByFloat(ctx, key, field, incr)
}

// HScan 增量迭代哈希字段
func (r *RedisClient) HScan(ctx context.Context, key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	if r.client != nil {
		return r.client.HScan(ctx, key, cursor, match, count)
	}
	return r.clusterClient.HScan(ctx, key, cursor, match, count)
}

// HRandField 随机返回哈希字段
func (r *RedisClient) HRandField(ctx context.Context, key string, count int) *redis.StringSliceCmd {
	if r.client != nil {
		return r.client.HRandField(ctx, key, count)
	}
	return r.clusterClient.HRandField(ctx, key, count)
}

// HRandFieldWithValues 随机返回哈希字段及值
func (r *RedisClient) HRandFieldWithValues(ctx context.Context, key string, count int) *redis.KeyValueSliceCmd {
	if r.client != nil {
		return r.client.HRandFieldWithValues(ctx, key, count)
	}
	return r.clusterClient.HRandFieldWithValues(ctx, key, count)
}

// HExpire 设置哈希字段过期时间
func (r *RedisClient) HExpire(ctx context.Context, key string, expiration time.Duration, fields ...string) *redis.IntSliceCmd {
	if r.client != nil {
		return r.client.HExpire(ctx, key, expiration, fields...)
	}
	return r.clusterClient.HExpire(ctx, key, expiration, fields...)
}

// HPTTL 获取哈希字段剩余生存时间（毫秒）
func (r *RedisClient) HPTTL(ctx context.Context, key string, fields ...string) *redis.IntSliceCmd {
	if r.client != nil {
		return r.client.HPTTL(ctx, key, fields...)
	}
	return r.clusterClient.HPTTL(ctx, key, fields...)
}

// HTTL 获取哈希字段剩余生存时间（秒）
func (r *RedisClient) HTTL(ctx context.Context, key string, fields ...string) *redis.IntSliceCmd {
	if r.client != nil {
		return r.client.HTTL(ctx, key, fields...)
	}
	return r.clusterClient.HTTL(ctx, key, fields...)
}

// HScanNoValues 增量迭代哈希字段（不返回值）
func (r *RedisClient) HScanNoValues(ctx context.Context, key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	if r.client != nil {
		return r.client.HScanNoValues(ctx, key, cursor, match, count)
	}
	return r.clusterClient.HScanNoValues(ctx, key, cursor, match, count)
}

// HExpireWithArgs 设置哈希字段过期（带参数）
func (r *RedisClient) HExpireWithArgs(ctx context.Context, key string, expiration time.Duration, expirationArgs redis.HExpireArgs, fields ...string) *redis.IntSliceCmd {
	if r.client != nil {
		return r.client.HExpireWithArgs(ctx, key, expiration, expirationArgs, fields...)
	}
	return r.clusterClient.HExpireWithArgs(ctx, key, expiration, expirationArgs, fields...)
}

// HPExpire 以毫秒设置哈希字段过期
func (r *RedisClient) HPExpire(ctx context.Context, key string, expiration time.Duration, fields ...string) *redis.IntSliceCmd {
	if r.client != nil {
		return r.client.HPExpire(ctx, key, expiration, fields...)
	}
	return r.clusterClient.HPExpire(ctx, key, expiration, fields...)
}

// HPExpireWithArgs 以毫秒设置哈希字段过期（带参数）
func (r *RedisClient) HPExpireWithArgs(ctx context.Context, key string, expiration time.Duration, expirationArgs redis.HExpireArgs, fields ...string) *redis.IntSliceCmd {
	if r.client != nil {
		return r.client.HPExpireWithArgs(ctx, key, expiration, expirationArgs, fields...)
	}
	return r.clusterClient.HPExpireWithArgs(ctx, key, expiration, expirationArgs, fields...)
}

// HExpireAt 设置哈希字段过期时间点
func (r *RedisClient) HExpireAt(ctx context.Context, key string, tm time.Time, fields ...string) *redis.IntSliceCmd {
	if r.client != nil {
		return r.client.HExpireAt(ctx, key, tm, fields...)
	}
	return r.clusterClient.HExpireAt(ctx, key, tm, fields...)
}

// HExpireAtWithArgs 设置哈希字段过期时间点（带参数）
func (r *RedisClient) HExpireAtWithArgs(ctx context.Context, key string, tm time.Time, expirationArgs redis.HExpireArgs, fields ...string) *redis.IntSliceCmd {
	if r.client != nil {
		return r.client.HExpireAtWithArgs(ctx, key, tm, expirationArgs, fields...)
	}
	return r.clusterClient.HExpireAtWithArgs(ctx, key, tm, expirationArgs, fields...)
}

// HPExpireAt 以毫秒时间戳设置哈希字段过期时间点
func (r *RedisClient) HPExpireAt(ctx context.Context, key string, tm time.Time, fields ...string) *redis.IntSliceCmd {
	if r.client != nil {
		return r.client.HPExpireAt(ctx, key, tm, fields...)
	}
	return r.clusterClient.HPExpireAt(ctx, key, tm, fields...)
}

// HPExpireAtWithArgs 以毫秒时间戳设置哈希字段过期时间点（带参数）
func (r *RedisClient) HPExpireAtWithArgs(ctx context.Context, key string, tm time.Time, expirationArgs redis.HExpireArgs, fields ...string) *redis.IntSliceCmd {
	if r.client != nil {
		return r.client.HPExpireAtWithArgs(ctx, key, tm, expirationArgs, fields...)
	}
	return r.clusterClient.HPExpireAtWithArgs(ctx, key, tm, expirationArgs, fields...)
}

// HPersist 移除哈希字段过期
func (r *RedisClient) HPersist(ctx context.Context, key string, fields ...string) *redis.IntSliceCmd {
	if r.client != nil {
		return r.client.HPersist(ctx, key, fields...)
	}
	return r.clusterClient.HPersist(ctx, key, fields...)
}

// HExpireTime 获取哈希字段过期时间点
func (r *RedisClient) HExpireTime(ctx context.Context, key string, fields ...string) *redis.IntSliceCmd {
	if r.client != nil {
		return r.client.HExpireTime(ctx, key, fields...)
	}
	return r.clusterClient.HExpireTime(ctx, key, fields...)
}

// HPExpireTime 获取哈希字段过期时间的毫秒时间戳
func (r *RedisClient) HPExpireTime(ctx context.Context, key string, fields ...string) *redis.IntSliceCmd {
	if r.client != nil {
		return r.client.HPExpireTime(ctx, key, fields...)
	}
	return r.clusterClient.HPExpireTime(ctx, key, fields...)
}
