package redisClient

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

// LPush 向列表左侧推入元素
func (r *RedisClient) LPush(ctx context.Context, key string, values ...interface{}) *redis.IntCmd {
	if r.client != nil {
		return r.client.LPush(ctx, key, values...)
	}
	return r.clusterClient.LPush(ctx, key, values...)
}

// LPushX 仅当列表存在时左侧推入
func (r *RedisClient) LPushX(ctx context.Context, key string, values ...interface{}) *redis.IntCmd {
	if r.client != nil {
		return r.client.LPushX(ctx, key, values...)
	}
	return r.clusterClient.LPushX(ctx, key, values...)
}

// RPush 向列表右侧推入元素
func (r *RedisClient) RPush(ctx context.Context, key string, values ...interface{}) *redis.IntCmd {
	if r.client != nil {
		return r.client.RPush(ctx, key, values...)
	}
	return r.clusterClient.RPush(ctx, key, values...)
}

// RPushX 仅当列表存在时右侧推入
func (r *RedisClient) RPushX(ctx context.Context, key string, values ...interface{}) *redis.IntCmd {
	if r.client != nil {
		return r.client.RPushX(ctx, key, values...)
	}
	return r.clusterClient.RPushX(ctx, key, values...)
}

// LPop 从列表左侧弹出元素
func (r *RedisClient) LPop(ctx context.Context, key string) *redis.StringCmd {
	if r.client != nil {
		return r.client.LPop(ctx, key)
	}
	return r.clusterClient.LPop(ctx, key)
}

// LPopCount 从列表左侧弹出多个元素
func (r *RedisClient) LPopCount(ctx context.Context, key string, count int) *redis.StringSliceCmd {
	if r.client != nil {
		return r.client.LPopCount(ctx, key, count)
	}
	return r.clusterClient.LPopCount(ctx, key, count)
}

// RPop 从列表右侧弹出元素
func (r *RedisClient) RPop(ctx context.Context, key string) *redis.StringCmd {
	if r.client != nil {
		return r.client.RPop(ctx, key)
	}
	return r.clusterClient.RPop(ctx, key)
}

// RPopCount 从列表右侧弹出多个元素
func (r *RedisClient) RPopCount(ctx context.Context, key string, count int) *redis.StringSliceCmd {
	if r.client != nil {
		return r.client.RPopCount(ctx, key, count)
	}
	return r.clusterClient.RPopCount(ctx, key, count)
}

// RPopLPush 将元素从源列表弹出并推入目标列表
func (r *RedisClient) RPopLPush(ctx context.Context, source, destination string) *redis.StringCmd {
	if r.client != nil {
		return r.client.RPopLPush(ctx, source, destination)
	}
	return r.clusterClient.RPopLPush(ctx, source, destination)
}

// LMove 原子移动列表元素
func (r *RedisClient) LMove(ctx context.Context, source, destination, srcpos, destpos string) *redis.StringCmd {
	if r.client != nil {
		return r.client.LMove(ctx, source, destination, srcpos, destpos)
	}
	return r.clusterClient.LMove(ctx, source, destination, srcpos, destpos)
}

// BLMove 阻塞式原子移动列表元素
func (r *RedisClient) BLMove(ctx context.Context, source, destination, srcpos, destpos string, timeout time.Duration) *redis.StringCmd {
	if r.client != nil {
		return r.client.BLMove(ctx, source, destination, srcpos, destpos, timeout)
	}
	return r.clusterClient.BLMove(ctx, source, destination, srcpos, destpos, timeout)
}

// LMPop 从多个列表弹出元素
func (r *RedisClient) LMPop(ctx context.Context, direction string, count int64, keys ...string) *redis.KeyValuesCmd {
	if r.client != nil {
		return r.client.LMPop(ctx, direction, count, keys...)
	}
	return r.clusterClient.LMPop(ctx, direction, count, keys...)
}

// BLMPop 阻塞式从多个列表弹出元素
func (r *RedisClient) BLMPop(ctx context.Context, timeout time.Duration, direction string, count int64, keys ...string) *redis.KeyValuesCmd {
	if r.client != nil {
		return r.client.BLMPop(ctx, timeout, direction, count, keys...)
	}
	return r.clusterClient.BLMPop(ctx, timeout, direction, count, keys...)
}

// BRPopLPush 阻塞式弹出并推入
func (r *RedisClient) BRPopLPush(ctx context.Context, source, destination string, timeout time.Duration) *redis.StringCmd {
	if r.client != nil {
		return r.client.BRPopLPush(ctx, source, destination, timeout)
	}
	return r.clusterClient.BRPopLPush(ctx, source, destination, timeout)
}

// BLPop 阻塞式左侧弹出
func (r *RedisClient) BLPop(ctx context.Context, timeout time.Duration, keys ...string) *redis.StringSliceCmd {
	if r.client != nil {
		return r.client.BLPop(ctx, timeout, keys...)
	}
	return r.clusterClient.BLPop(ctx, timeout, keys...)
}

// BRPop 阻塞式右侧弹出
func (r *RedisClient) BRPop(ctx context.Context, timeout time.Duration, keys ...string) *redis.StringSliceCmd {
	if r.client != nil {
		return r.client.BRPop(ctx, timeout, keys...)
	}
	return r.clusterClient.BRPop(ctx, timeout, keys...)
}

// LLen 获取列表长度
func (r *RedisClient) LLen(ctx context.Context, key string) *redis.IntCmd {
	if r.client != nil {
		return r.client.LLen(ctx, key)
	}
	return r.clusterClient.LLen(ctx, key)
}

// LRange 获取列表范围元素
func (r *RedisClient) LRange(ctx context.Context, key string, start, stop int64) *redis.StringSliceCmd {
	if r.client != nil {
		return r.client.LRange(ctx, key, start, stop)
	}
	return r.clusterClient.LRange(ctx, key, start, stop)
}

// LIndex 获取列表指定索引元素
func (r *RedisClient) LIndex(ctx context.Context, key string, index int64) *redis.StringCmd {
	if r.client != nil {
		return r.client.LIndex(ctx, key, index)
	}
	return r.clusterClient.LIndex(ctx, key, index)
}

// LInsert 在列表元素前后插入
func (r *RedisClient) LInsert(ctx context.Context, key, op string, pivot, value interface{}) *redis.IntCmd {
	if r.client != nil {
		return r.client.LInsert(ctx, key, op, pivot, value)
	}
	return r.clusterClient.LInsert(ctx, key, op, pivot, value)
}

// LInsertBefore 在列表元素前插入
func (r *RedisClient) LInsertBefore(ctx context.Context, key string, pivot, value interface{}) *redis.IntCmd {
	if r.client != nil {
		return r.client.LInsertBefore(ctx, key, pivot, value)
	}
	return r.clusterClient.LInsertBefore(ctx, key, pivot, value)
}

// LInsertAfter 在列表元素后插入
func (r *RedisClient) LInsertAfter(ctx context.Context, key string, pivot, value interface{}) *redis.IntCmd {
	if r.client != nil {
		return r.client.LInsertAfter(ctx, key, pivot, value)
	}
	return r.clusterClient.LInsertAfter(ctx, key, pivot, value)
}

// LSet 设置列表指定索引元素
func (r *RedisClient) LSet(ctx context.Context, key string, index int64, value interface{}) *redis.StatusCmd {
	if r.client != nil {
		return r.client.LSet(ctx, key, index, value)
	}
	return r.clusterClient.LSet(ctx, key, index, value)
}

// LRem 删除列表元素
func (r *RedisClient) LRem(ctx context.Context, key string, count int64, value interface{}) *redis.IntCmd {
	if r.client != nil {
		return r.client.LRem(ctx, key, count, value)
	}
	return r.clusterClient.LRem(ctx, key, count, value)
}

// LTrim 修剪列表
func (r *RedisClient) LTrim(ctx context.Context, key string, start, stop int64) *redis.StatusCmd {
	if r.client != nil {
		return r.client.LTrim(ctx, key, start, stop)
	}
	return r.clusterClient.LTrim(ctx, key, start, stop)
}

// LPos 查找元素位置
func (r *RedisClient) LPos(ctx context.Context, key string, value string, args redis.LPosArgs) *redis.IntCmd {
	if r.client != nil {
		return r.client.LPos(ctx, key, value, args)
	}
	return r.clusterClient.LPos(ctx, key, value, args)
}

// LPosCount 查找元素所有位置
func (r *RedisClient) LPosCount(ctx context.Context, key string, value string, count int64, args redis.LPosArgs) *redis.IntSliceCmd {
	if r.client != nil {
		return r.client.LPosCount(ctx, key, value, count, args)
	}
	return r.clusterClient.LPosCount(ctx, key, value, count, args)
}
