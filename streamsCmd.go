package redisClient

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// XAdd 向流中添加消息
func (r *RedisClient) XAdd(ctx context.Context, a *XAddArgs) *redis.StringCmd {
	if r.client != nil {
		return r.client.XAdd(ctx, a)
	}
	return r.clusterClient.XAdd(ctx, a)
}

// XDel 删除流消息
func (r *RedisClient) XDel(ctx context.Context, stream string, ids ...string) *redis.IntCmd {
	if r.client != nil {
		return r.client.XDel(ctx, stream, ids...)
	}
	return r.clusterClient.XDel(ctx, stream, ids...)
}

// XLen 获取流长度
func (r *RedisClient) XLen(ctx context.Context, stream string) *redis.IntCmd {
	if r.client != nil {
		return r.client.XLen(ctx, stream)
	}
	return r.clusterClient.XLen(ctx, stream)
}

// XRange 按范围查询流消息（正序）
func (r *RedisClient) XRange(ctx context.Context, stream, start, stop string) *redis.XMessageSliceCmd {
	if r.client != nil {
		return r.client.XRange(ctx, stream, start, stop)
	}
	return r.clusterClient.XRange(ctx, stream, start, stop)
}

// XRangeN 按范围查询流消息（正序，限制数量）
func (r *RedisClient) XRangeN(ctx context.Context, stream, start, stop string, count int64) *redis.XMessageSliceCmd {
	if r.client != nil {
		return r.client.XRangeN(ctx, stream, start, stop, count)
	}
	return r.clusterClient.XRangeN(ctx, stream, start, stop, count)
}

// XRevRange 按范围查询流消息（逆序）
func (r *RedisClient) XRevRange(ctx context.Context, stream, start, stop string) *redis.XMessageSliceCmd {
	if r.client != nil {
		return r.client.XRevRange(ctx, stream, start, stop)
	}
	return r.clusterClient.XRevRange(ctx, stream, start, stop)
}

// XRevRangeN 按范围查询流消息（逆序，限制数量）
func (r *RedisClient) XRevRangeN(ctx context.Context, stream, start, stop string, count int64) *redis.XMessageSliceCmd {
	if r.client != nil {
		return r.client.XRevRangeN(ctx, stream, start, stop, count)
	}
	return r.clusterClient.XRevRangeN(ctx, stream, start, stop, count)
}

// XRead 读取流消息
func (r *RedisClient) XRead(ctx context.Context, a *XReadArgs) *redis.XStreamSliceCmd {
	if r.client != nil {
		return r.client.XRead(ctx, a)
	}
	return r.clusterClient.XRead(ctx, a)
}

// XReadStreams 读取流消息（简化版）
func (r *RedisClient) XReadStreams(ctx context.Context, streams ...string) *redis.XStreamSliceCmd {
	if r.client != nil {
		return r.client.XReadStreams(ctx, streams...)
	}
	return r.clusterClient.XReadStreams(ctx, streams...)
}

// XGroupCreate 创建消费者组
func (r *RedisClient) XGroupCreate(ctx context.Context, stream, group, start string) *redis.StatusCmd {
	if r.client != nil {
		return r.client.XGroupCreate(ctx, stream, group, start)
	}
	return r.clusterClient.XGroupCreate(ctx, stream, group, start)
}

// XGroupCreateMkStream 创建消费者组（自动创建流）
func (r *RedisClient) XGroupCreateMkStream(ctx context.Context, stream, group, start string) *redis.StatusCmd {
	if r.client != nil {
		return r.client.XGroupCreateMkStream(ctx, stream, group, start)
	}
	return r.clusterClient.XGroupCreateMkStream(ctx, stream, group, start)
}

// XGroupSetID 设置消费者组 ID
func (r *RedisClient) XGroupSetID(ctx context.Context, stream, group, start string) *redis.StatusCmd {
	if r.client != nil {
		return r.client.XGroupSetID(ctx, stream, group, start)
	}
	return r.clusterClient.XGroupSetID(ctx, stream, group, start)
}

// XGroupDestroy 销毁消费者组
func (r *RedisClient) XGroupDestroy(ctx context.Context, stream, group string) *redis.IntCmd {
	if r.client != nil {
		return r.client.XGroupDestroy(ctx, stream, group)
	}
	return r.clusterClient.XGroupDestroy(ctx, stream, group)
}

// XGroupCreateConsumer 创建消费者
func (r *RedisClient) XGroupCreateConsumer(ctx context.Context, stream, group, consumer string) *redis.IntCmd {
	if r.client != nil {
		return r.client.XGroupCreateConsumer(ctx, stream, group, consumer)
	}
	return r.clusterClient.XGroupCreateConsumer(ctx, stream, group, consumer)
}

// XGroupDelConsumer 删除消费者
func (r *RedisClient) XGroupDelConsumer(ctx context.Context, stream, group, consumer string) *redis.IntCmd {
	if r.client != nil {
		return r.client.XGroupDelConsumer(ctx, stream, group, consumer)
	}
	return r.clusterClient.XGroupDelConsumer(ctx, stream, group, consumer)
}

// XReadGroup 消费者组读取消息
func (r *RedisClient) XReadGroup(ctx context.Context, a *XReadGroupArgs) *redis.XStreamSliceCmd {
	if r.client != nil {
		return r.client.XReadGroup(ctx, a)
	}
	return r.clusterClient.XReadGroup(ctx, a)
}

// XAck 确认消息
func (r *RedisClient) XAck(ctx context.Context, stream, group string, ids ...string) *redis.IntCmd {
	if r.client != nil {
		return r.client.XAck(ctx, stream, group, ids...)
	}
	return r.clusterClient.XAck(ctx, stream, group, ids...)
}

// XClaim 认领消息
func (r *RedisClient) XClaim(ctx context.Context, a *XClaimArgs) *redis.XMessageSliceCmd {
	if r.client != nil {
		return r.client.XClaim(ctx, a)
	}
	return r.clusterClient.XClaim(ctx, a)
}

// XClaimJustID 认领消息（仅返回 ID）
func (r *RedisClient) XClaimJustID(ctx context.Context, a *XClaimArgs) *redis.StringSliceCmd {
	if r.client != nil {
		return r.client.XClaimJustID(ctx, a)
	}
	return r.clusterClient.XClaimJustID(ctx, a)
}

// XAutoClaim 自动认领消息
func (r *RedisClient) XAutoClaim(ctx context.Context, a *XAutoClaimArgs) *redis.XAutoClaimCmd {
	if r.client != nil {
		return r.client.XAutoClaim(ctx, a)
	}
	return r.clusterClient.XAutoClaim(ctx, a)
}

// XAutoClaimJustID 自动认领消息（仅返回 ID）
func (r *RedisClient) XAutoClaimJustID(ctx context.Context, a *XAutoClaimArgs) *redis.XAutoClaimJustIDCmd {
	if r.client != nil {
		return r.client.XAutoClaimJustID(ctx, a)
	}
	return r.clusterClient.XAutoClaimJustID(ctx, a)
}

// XPending 查询待处理消息
func (r *RedisClient) XPending(ctx context.Context, stream, group string) *redis.XPendingCmd {
	if r.client != nil {
		return r.client.XPending(ctx, stream, group)
	}
	return r.clusterClient.XPending(ctx, stream, group)
}

// XPendingExt 查询待处理消息（扩展）
func (r *RedisClient) XPendingExt(ctx context.Context, a *XPendingExtArgs) *redis.XPendingExtCmd {
	if r.client != nil {
		return r.client.XPendingExt(ctx, a)
	}
	return r.clusterClient.XPendingExt(ctx, a)
}

// XInfoStream 查询流信息
func (r *RedisClient) XInfoStream(ctx context.Context, key string) *redis.XInfoStreamCmd {
	if r.client != nil {
		return r.client.XInfoStream(ctx, key)
	}
	return r.clusterClient.XInfoStream(ctx, key)
}

// XInfoStreamFull 查询流完整信息
func (r *RedisClient) XInfoStreamFull(ctx context.Context, key string, count int) *redis.XInfoStreamFullCmd {
	if r.client != nil {
		return r.client.XInfoStreamFull(ctx, key, count)
	}
	return r.clusterClient.XInfoStreamFull(ctx, key, count)
}

// XInfoGroups 查询消费者组信息
func (r *RedisClient) XInfoGroups(ctx context.Context, key string) *redis.XInfoGroupsCmd {
	if r.client != nil {
		return r.client.XInfoGroups(ctx, key)
	}
	return r.clusterClient.XInfoGroups(ctx, key)
}

// XInfoConsumers 查询消费者信息
func (r *RedisClient) XInfoConsumers(ctx context.Context, key string, group string) *redis.XInfoConsumersCmd {
	if r.client != nil {
		return r.client.XInfoConsumers(ctx, key, group)
	}
	return r.clusterClient.XInfoConsumers(ctx, key, group)
}

// XTrimMaxLen 按最大长度修剪流
func (r *RedisClient) XTrimMaxLen(ctx context.Context, key string, maxLen int64) *redis.IntCmd {
	if r.client != nil {
		return r.client.XTrimMaxLen(ctx, key, maxLen)
	}
	return r.clusterClient.XTrimMaxLen(ctx, key, maxLen)
}

// XTrimMaxLenApprox 按最大长度近似修剪流
func (r *RedisClient) XTrimMaxLenApprox(ctx context.Context, key string, maxLen, limit int64) *redis.IntCmd {
	if r.client != nil {
		return r.client.XTrimMaxLenApprox(ctx, key, maxLen, limit)
	}
	return r.clusterClient.XTrimMaxLenApprox(ctx, key, maxLen, limit)
}

// XTrimMinID 按最小 ID 修剪流
func (r *RedisClient) XTrimMinID(ctx context.Context, key string, minID string) *redis.IntCmd {
	if r.client != nil {
		return r.client.XTrimMinID(ctx, key, minID)
	}
	return r.clusterClient.XTrimMinID(ctx, key, minID)
}

// XTrimMinIDApprox 按最小 ID 近似修剪流
func (r *RedisClient) XTrimMinIDApprox(ctx context.Context, key string, minID string, limit int64) *redis.IntCmd {
	if r.client != nil {
		return r.client.XTrimMinIDApprox(ctx, key, minID, limit)
	}
	return r.clusterClient.XTrimMinIDApprox(ctx, key, minID, limit)
}
