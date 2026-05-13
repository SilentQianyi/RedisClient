package redisClient

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// PFAdd 向 HyperLogLog 添加元素
func (r *RedisClient) PFAdd(ctx context.Context, key string, els ...interface{}) *redis.IntCmd {
	if r.client != nil {
		return r.client.PFAdd(ctx, key, els...)
	}
	return r.clusterClient.PFAdd(ctx, key, els...)
}

// PFCount 获取 HyperLogLog 基数估算
func (r *RedisClient) PFCount(ctx context.Context, keys ...string) *redis.IntCmd {
	if r.client != nil {
		return r.client.PFCount(ctx, keys...)
	}
	return r.clusterClient.PFCount(ctx, keys...)
}

// PFMerge 合并多个 HyperLogLog
func (r *RedisClient) PFMerge(ctx context.Context, dest string, keys ...string) *redis.StatusCmd {
	if r.client != nil {
		return r.client.PFMerge(ctx, dest, keys...)
	}
	return r.clusterClient.PFMerge(ctx, dest, keys...)
}
