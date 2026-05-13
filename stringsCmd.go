package redisClient

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

// Set 设置键值
func (r *RedisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	if r.client != nil {
		return r.client.Set(ctx, key, value, expiration)
	}
	return r.clusterClient.Set(ctx, key, value, expiration)
}

// SetArgs 设置键值（带额外参数）
func (r *RedisClient) SetArgs(ctx context.Context, key string, value interface{}, a SetArgs) *redis.StatusCmd {
	if r.client != nil {
		return r.client.SetArgs(ctx, key, value, a)
	}
	return r.clusterClient.SetArgs(ctx, key, value, a)
}

// SetEx 设置键值并指定过期时间
func (r *RedisClient) SetEx(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	if r.client != nil {
		return r.client.SetEx(ctx, key, value, expiration)
	}
	return r.clusterClient.SetEx(ctx, key, value, expiration)
}

// SetNX 键不存在时设置
func (r *RedisClient) SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.BoolCmd {
	if r.client != nil {
		return r.client.SetNX(ctx, key, value, expiration)
	}
	return r.clusterClient.SetNX(ctx, key, value, expiration)
}

// SetXX 键存在时设置
func (r *RedisClient) SetXX(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.BoolCmd {
	if r.client != nil {
		return r.client.SetXX(ctx, key, value, expiration)
	}
	return r.clusterClient.SetXX(ctx, key, value, expiration)
}

// Get 获取键值
func (r *RedisClient) Get(ctx context.Context, key string) *redis.StringCmd {
	if r.client != nil {
		return r.client.Get(ctx, key)
	}
	return r.clusterClient.Get(ctx, key)
}

// GetSet 设置新值并返回旧值
func (r *RedisClient) GetSet(ctx context.Context, key string, value interface{}) *redis.StringCmd {
	if r.client != nil {
		return r.client.GetSet(ctx, key, value)
	}
	return r.clusterClient.GetSet(ctx, key, value)
}

// GetEx 获取键值并设置过期时间
func (r *RedisClient) GetEx(ctx context.Context, key string, expiration time.Duration) *redis.StringCmd {
	if r.client != nil {
		return r.client.GetEx(ctx, key, expiration)
	}
	return r.clusterClient.GetEx(ctx, key, expiration)
}

// GetDel 获取键值并删除键
func (r *RedisClient) GetDel(ctx context.Context, key string) *redis.StringCmd {
	if r.client != nil {
		return r.client.GetDel(ctx, key)
	}
	return r.clusterClient.GetDel(ctx, key)
}

// GetRange 获取子字符串
func (r *RedisClient) GetRange(ctx context.Context, key string, start, end int64) *redis.StringCmd {
	if r.client != nil {
		return r.client.GetRange(ctx, key, start, end)
	}
	return r.clusterClient.GetRange(ctx, key, start, end)
}

// SetRange 覆写指定偏移处的字符串
func (r *RedisClient) SetRange(ctx context.Context, key string, offset int64, value string) *redis.IntCmd {
	if r.client != nil {
		return r.client.SetRange(ctx, key, offset, value)
	}
	return r.clusterClient.SetRange(ctx, key, offset, value)
}

// StrLen 获取字符串长度
func (r *RedisClient) StrLen(ctx context.Context, key string) *redis.IntCmd {
	if r.client != nil {
		return r.client.StrLen(ctx, key)
	}
	return r.clusterClient.StrLen(ctx, key)
}

// Append 追加字符串
func (r *RedisClient) Append(ctx context.Context, key, value string) *redis.IntCmd {
	if r.client != nil {
		return r.client.Append(ctx, key, value)
	}
	return r.clusterClient.Append(ctx, key, value)
}

// Incr 键值自增 1
func (r *RedisClient) Incr(ctx context.Context, key string) *redis.IntCmd {
	if r.client != nil {
		return r.client.Incr(ctx, key)
	}
	return r.clusterClient.Incr(ctx, key)
}

// IncrBy 键值自增指定值
func (r *RedisClient) IncrBy(ctx context.Context, key string, value int64) *redis.IntCmd {
	if r.client != nil {
		return r.client.IncrBy(ctx, key, value)
	}
	return r.clusterClient.IncrBy(ctx, key, value)
}

// IncrByFloat 键值自增浮点数
func (r *RedisClient) IncrByFloat(ctx context.Context, key string, value float64) *redis.FloatCmd {
	if r.client != nil {
		return r.client.IncrByFloat(ctx, key, value)
	}
	return r.clusterClient.IncrByFloat(ctx, key, value)
}

// Decr 键值自减 1
func (r *RedisClient) Decr(ctx context.Context, key string) *redis.IntCmd {
	if r.client != nil {
		return r.client.Decr(ctx, key)
	}
	return r.clusterClient.Decr(ctx, key)
}

// DecrBy 键值自减指定值
func (r *RedisClient) DecrBy(ctx context.Context, key string, decrement int64) *redis.IntCmd {
	if r.client != nil {
		return r.client.DecrBy(ctx, key, decrement)
	}
	return r.clusterClient.DecrBy(ctx, key, decrement)
}

// MSet 批量设置键值
func (r *RedisClient) MSet(ctx context.Context, values ...interface{}) *redis.StatusCmd {
	if r.client != nil {
		return r.client.MSet(ctx, values...)
	}
	return r.clusterClient.MSet(ctx, values...)
}

// MSetNX 批量设置键值（仅当所有键都不存在时）
func (r *RedisClient) MSetNX(ctx context.Context, values ...interface{}) *redis.BoolCmd {
	if r.client != nil {
		return r.client.MSetNX(ctx, values...)
	}
	return r.clusterClient.MSetNX(ctx, values...)
}

// MGet 批量获取键值
func (r *RedisClient) MGet(ctx context.Context, keys ...string) *redis.SliceCmd {
	if r.client != nil {
		return r.client.MGet(ctx, keys...)
	}
	return r.clusterClient.MGet(ctx, keys...)
}

// LCS 最长公共子序列
func (r *RedisClient) LCS(ctx context.Context, q *LCSQuery) *redis.LCSCmd {
	if r.client != nil {
		return r.client.LCS(ctx, q)
	}
	return r.clusterClient.LCS(ctx, q)
}
