package redisClient

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// Publish 发布消息到频道
func (r *RedisClient) Publish(ctx context.Context, channel string, message interface{}) *redis.IntCmd {
	if r.client != nil {
		return r.client.Publish(ctx, channel, message)
	}
	return r.clusterClient.Publish(ctx, channel, message)
}

// SPublish 发布消息到分片频道
func (r *RedisClient) SPublish(ctx context.Context, channel string, message interface{}) *redis.IntCmd {
	if r.client != nil {
		return r.client.SPublish(ctx, channel, message)
	}
	return r.clusterClient.SPublish(ctx, channel, message)
}

// PubSubChannels 列出活跃频道
func (r *RedisClient) PubSubChannels(ctx context.Context, pattern string) *redis.StringSliceCmd {
	if r.client != nil {
		return r.client.PubSubChannels(ctx, pattern)
	}
	return r.clusterClient.PubSubChannels(ctx, pattern)
}

// PubSubNumSub 获取频道订阅数
func (r *RedisClient) PubSubNumSub(ctx context.Context, channels ...string) *redis.MapStringIntCmd {
	if r.client != nil {
		return r.client.PubSubNumSub(ctx, channels...)
	}
	return r.clusterClient.PubSubNumSub(ctx, channels...)
}

// PubSubNumPat 获取模式订阅数
func (r *RedisClient) PubSubNumPat(ctx context.Context) *redis.IntCmd {
	if r.client != nil {
		return r.client.PubSubNumPat(ctx)
	}
	return r.clusterClient.PubSubNumPat(ctx)
}

// PubSubShardChannels 列出分片频道
func (r *RedisClient) PubSubShardChannels(ctx context.Context, pattern string) *redis.StringSliceCmd {
	if r.client != nil {
		return r.client.PubSubShardChannels(ctx, pattern)
	}
	return r.clusterClient.PubSubShardChannels(ctx, pattern)
}

// PubSubShardNumSub 获取分片频道订阅数
func (r *RedisClient) PubSubShardNumSub(ctx context.Context, channels ...string) *redis.MapStringIntCmd {
	if r.client != nil {
		return r.client.PubSubShardNumSub(ctx, channels...)
	}
	return r.clusterClient.PubSubShardNumSub(ctx, channels...)
}
