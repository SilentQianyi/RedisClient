package redisClient

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// SetBit 设置位值
func (r *RedisClient) SetBit(ctx context.Context, key string, offset int64, value int) *redis.IntCmd {
	if r.client != nil {
		return r.client.SetBit(ctx, key, offset, value)
	}
	return r.clusterClient.SetBit(ctx, key, offset, value)
}

// GetBit 获取位值
func (r *RedisClient) GetBit(ctx context.Context, key string, offset int64) *redis.IntCmd {
	if r.client != nil {
		return r.client.GetBit(ctx, key, offset)
	}
	return r.clusterClient.GetBit(ctx, key, offset)
}

// BitCount 统计位数
func (r *RedisClient) BitCount(ctx context.Context, key string, bitCount *BitCount) *redis.IntCmd {
	if r.client != nil {
		return r.client.BitCount(ctx, key, bitCount)
	}
	return r.clusterClient.BitCount(ctx, key, bitCount)
}

// BitOpAnd 位与操作
func (r *RedisClient) BitOpAnd(ctx context.Context, destKey string, keys ...string) *redis.IntCmd {
	if r.client != nil {
		return r.client.BitOpAnd(ctx, destKey, keys...)
	}
	return r.clusterClient.BitOpAnd(ctx, destKey, keys...)
}

// BitOpOr 位或操作
func (r *RedisClient) BitOpOr(ctx context.Context, destKey string, keys ...string) *redis.IntCmd {
	if r.client != nil {
		return r.client.BitOpOr(ctx, destKey, keys...)
	}
	return r.clusterClient.BitOpOr(ctx, destKey, keys...)
}

// BitOpXor 位异或操作
func (r *RedisClient) BitOpXor(ctx context.Context, destKey string, keys ...string) *redis.IntCmd {
	if r.client != nil {
		return r.client.BitOpXor(ctx, destKey, keys...)
	}
	return r.clusterClient.BitOpXor(ctx, destKey, keys...)
}

// BitOpNot 位非操作
func (r *RedisClient) BitOpNot(ctx context.Context, destKey string, key string) *redis.IntCmd {
	if r.client != nil {
		return r.client.BitOpNot(ctx, destKey, key)
	}
	return r.clusterClient.BitOpNot(ctx, destKey, key)
}

// BitPos 查找第一个为 0 或 1 的位
func (r *RedisClient) BitPos(ctx context.Context, key string, bit int64, pos ...int64) *redis.IntCmd {
	if r.client != nil {
		return r.client.BitPos(ctx, key, bit, pos...)
	}
	return r.clusterClient.BitPos(ctx, key, bit, pos...)
}

// BitField 位域操作
func (r *RedisClient) BitField(ctx context.Context, key string, values ...interface{}) *redis.IntSliceCmd {
	if r.client != nil {
		return r.client.BitField(ctx, key, values...)
	}
	return r.clusterClient.BitField(ctx, key, values...)
}

// BitFieldRO 只读位域操作
func (r *RedisClient) BitFieldRO(ctx context.Context, key string, values ...interface{}) *redis.IntSliceCmd {
	if r.client != nil {
		return r.client.BitFieldRO(ctx, key, values...)
	}
	return r.clusterClient.BitFieldRO(ctx, key, values...)
}

// BitPosSpan 按 span 查找位位置
func (r *RedisClient) BitPosSpan(ctx context.Context, key string, bit int8, start, end int64, span string) *redis.IntCmd {
	if r.client != nil {
		return r.client.BitPosSpan(ctx, key, bit, start, end, span)
	}
	return r.clusterClient.BitPosSpan(ctx, key, bit, start, end, span)
}
