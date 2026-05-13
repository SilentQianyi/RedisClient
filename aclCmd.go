package redisClient

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// ACLDryRun 模拟 ACL 权限检查
func (r *RedisClient) ACLDryRun(ctx context.Context, username string, command ...interface{}) *redis.StringCmd {
	if r.client != nil {
		return r.client.ACLDryRun(ctx, username, command...)
	}
	return r.clusterClient.ACLDryRun(ctx, username, command...)
}

// ACLLog 获取 ACL 日志
func (r *RedisClient) ACLLog(ctx context.Context, count int64) *redis.ACLLogCmd {
	if r.client != nil {
		return r.client.ACLLog(ctx, count)
	}
	return r.clusterClient.ACLLog(ctx, count)
}

// ACLLogReset 重置 ACL 日志
func (r *RedisClient) ACLLogReset(ctx context.Context) *redis.StatusCmd {
	if r.client != nil {
		return r.client.ACLLogReset(ctx)
	}
	return r.clusterClient.ACLLogReset(ctx)
}
