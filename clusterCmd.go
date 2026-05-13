package redisClient

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// ============================================================================
// 集群管理
// ============================================================================

// ClusterMyShardID 获取当前分片 ID
func (r *RedisClient) ClusterMyShardID(ctx context.Context) *redis.StringCmd {
	if r.client != nil {
		return r.client.ClusterMyShardID(ctx)
	}
	return r.clusterClient.ClusterMyShardID(ctx)
}

// ClusterSlots 获取槽位信息
func (r *RedisClient) ClusterSlots(ctx context.Context) *redis.ClusterSlotsCmd {
	if r.client != nil {
		return r.client.ClusterSlots(ctx)
	}
	return r.clusterClient.ClusterSlots(ctx)
}

// ClusterShards 获取分片信息
func (r *RedisClient) ClusterShards(ctx context.Context) *redis.ClusterShardsCmd {
	if r.client != nil {
		return r.client.ClusterShards(ctx)
	}
	return r.clusterClient.ClusterShards(ctx)
}

// ClusterLinks 获取集群链接信息
func (r *RedisClient) ClusterLinks(ctx context.Context) *redis.ClusterLinksCmd {
	if r.client != nil {
		return r.client.ClusterLinks(ctx)
	}
	return r.clusterClient.ClusterLinks(ctx)
}

// ClusterNodes 获取集群节点信息
func (r *RedisClient) ClusterNodes(ctx context.Context) *redis.StringCmd {
	if r.client != nil {
		return r.client.ClusterNodes(ctx)
	}
	return r.clusterClient.ClusterNodes(ctx)
}

// ClusterMeet 将节点加入集群
func (r *RedisClient) ClusterMeet(ctx context.Context, host, port string) *redis.StatusCmd {
	if r.client != nil {
		return r.client.ClusterMeet(ctx, host, port)
	}
	return r.clusterClient.ClusterMeet(ctx, host, port)
}

// ClusterForget 从集群中移除节点
func (r *RedisClient) ClusterForget(ctx context.Context, nodeID string) *redis.StatusCmd {
	if r.client != nil {
		return r.client.ClusterForget(ctx, nodeID)
	}
	return r.clusterClient.ClusterForget(ctx, nodeID)
}

// ClusterReplicate 将当前节点设为指定节点的从节点
func (r *RedisClient) ClusterReplicate(ctx context.Context, nodeID string) *redis.StatusCmd {
	if r.client != nil {
		return r.client.ClusterReplicate(ctx, nodeID)
	}
	return r.clusterClient.ClusterReplicate(ctx, nodeID)
}

// ClusterResetSoft 软重置集群
func (r *RedisClient) ClusterResetSoft(ctx context.Context) *redis.StatusCmd {
	if r.client != nil {
		return r.client.ClusterResetSoft(ctx)
	}
	return r.clusterClient.ClusterResetSoft(ctx)
}

// ClusterResetHard 硬重置集群
func (r *RedisClient) ClusterResetHard(ctx context.Context) *redis.StatusCmd {
	if r.client != nil {
		return r.client.ClusterResetHard(ctx)
	}
	return r.clusterClient.ClusterResetHard(ctx)
}

// ClusterInfo 获取集群信息
func (r *RedisClient) ClusterInfo(ctx context.Context) *redis.StringCmd {
	if r.client != nil {
		return r.client.ClusterInfo(ctx)
	}
	return r.clusterClient.ClusterInfo(ctx)
}

// ClusterKeySlot 获取键所属的槽位
func (r *RedisClient) ClusterKeySlot(ctx context.Context, key string) *redis.IntCmd {
	if r.client != nil {
		return r.client.ClusterKeySlot(ctx, key)
	}
	return r.clusterClient.ClusterKeySlot(ctx, key)
}

// ClusterGetKeysInSlot 获取指定槽位中的键
func (r *RedisClient) ClusterGetKeysInSlot(ctx context.Context, slot int, count int) *redis.StringSliceCmd {
	if r.client != nil {
		return r.client.ClusterGetKeysInSlot(ctx, slot, count)
	}
	return r.clusterClient.ClusterGetKeysInSlot(ctx, slot, count)
}

// ClusterCountFailureReports 获取节点的故障报告数
func (r *RedisClient) ClusterCountFailureReports(ctx context.Context, nodeID string) *redis.IntCmd {
	if r.client != nil {
		return r.client.ClusterCountFailureReports(ctx, nodeID)
	}
	return r.clusterClient.ClusterCountFailureReports(ctx, nodeID)
}

// ClusterCountKeysInSlot 获取槽位中的键数量
func (r *RedisClient) ClusterCountKeysInSlot(ctx context.Context, slot int) *redis.IntCmd {
	if r.client != nil {
		return r.client.ClusterCountKeysInSlot(ctx, slot)
	}
	return r.clusterClient.ClusterCountKeysInSlot(ctx, slot)
}

// ClusterDelSlots 删除槽位
func (r *RedisClient) ClusterDelSlots(ctx context.Context, slots ...int) *redis.StatusCmd {
	if r.client != nil {
		return r.client.ClusterDelSlots(ctx, slots...)
	}
	return r.clusterClient.ClusterDelSlots(ctx, slots...)
}

// ClusterDelSlotsRange 删除槽位范围
func (r *RedisClient) ClusterDelSlotsRange(ctx context.Context, min, max int) *redis.StatusCmd {
	if r.client != nil {
		return r.client.ClusterDelSlotsRange(ctx, min, max)
	}
	return r.clusterClient.ClusterDelSlotsRange(ctx, min, max)
}

// ClusterSaveConfig 保存集群配置
func (r *RedisClient) ClusterSaveConfig(ctx context.Context) *redis.StatusCmd {
	if r.client != nil {
		return r.client.ClusterSaveConfig(ctx)
	}
	return r.clusterClient.ClusterSaveConfig(ctx)
}

// ClusterSlaves 获取从节点列表
func (r *RedisClient) ClusterSlaves(ctx context.Context, nodeID string) *redis.StringSliceCmd {
	if r.client != nil {
		return r.client.ClusterSlaves(ctx, nodeID)
	}
	return r.clusterClient.ClusterSlaves(ctx, nodeID)
}

// ClusterFailover 触发集群故障转移
func (r *RedisClient) ClusterFailover(ctx context.Context) *redis.StatusCmd {
	if r.client != nil {
		return r.client.ClusterFailover(ctx)
	}
	return r.clusterClient.ClusterFailover(ctx)
}

// ClusterAddSlots 添加槽位
func (r *RedisClient) ClusterAddSlots(ctx context.Context, slots ...int) *redis.StatusCmd {
	if r.client != nil {
		return r.client.ClusterAddSlots(ctx, slots...)
	}
	return r.clusterClient.ClusterAddSlots(ctx, slots...)
}

// ClusterAddSlotsRange 添加槽位范围
func (r *RedisClient) ClusterAddSlotsRange(ctx context.Context, min, max int) *redis.StatusCmd {
	if r.client != nil {
		return r.client.ClusterAddSlotsRange(ctx, min, max)
	}
	return r.clusterClient.ClusterAddSlotsRange(ctx, min, max)
}

// ReadOnly 设置为只读模式（从节点）
func (r *RedisClient) ReadOnly(ctx context.Context) *redis.StatusCmd {
	if r.client != nil {
		return r.client.ReadOnly(ctx)
	}
	return r.clusterClient.ReadOnly(ctx)
}

// ReadWrite 设置为读写模式
func (r *RedisClient) ReadWrite(ctx context.Context) *redis.StatusCmd {
	if r.client != nil {
		return r.client.ReadWrite(ctx)
	}
	return r.clusterClient.ReadWrite(ctx)
}
