package redisClient

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// SAdd 向集合添加成员
func (r *RedisClient) SAdd(ctx context.Context, key string, members ...interface{}) *redis.IntCmd {
	if r.client != nil {
		return r.client.SAdd(ctx, key, members...)
	}
	return r.clusterClient.SAdd(ctx, key, members...)
}

// SRem 删除集合成员
func (r *RedisClient) SRem(ctx context.Context, key string, members ...interface{}) *redis.IntCmd {
	if r.client != nil {
		return r.client.SRem(ctx, key, members...)
	}
	return r.clusterClient.SRem(ctx, key, members...)
}

// SCard 获取集合成员数
func (r *RedisClient) SCard(ctx context.Context, key string) *redis.IntCmd {
	if r.client != nil {
		return r.client.SCard(ctx, key)
	}
	return r.clusterClient.SCard(ctx, key)
}

// SMembers 获取集合所有成员
func (r *RedisClient) SMembers(ctx context.Context, key string) *redis.StringSliceCmd {
	if r.client != nil {
		return r.client.SMembers(ctx, key)
	}
	return r.clusterClient.SMembers(ctx, key)
}

// SMembersMap 获取集合所有成员（map 格式）
func (r *RedisClient) SMembersMap(ctx context.Context, key string) *redis.StringStructMapCmd {
	if r.client != nil {
		return r.client.SMembersMap(ctx, key)
	}
	return r.clusterClient.SMembersMap(ctx, key)
}

// SIsMember 判断是否为集合成员
func (r *RedisClient) SIsMember(ctx context.Context, key string, member interface{}) *redis.BoolCmd {
	if r.client != nil {
		return r.client.SIsMember(ctx, key, member)
	}
	return r.clusterClient.SIsMember(ctx, key, member)
}

// SMIsMember 批量判断是否为集合成员
func (r *RedisClient) SMIsMember(ctx context.Context, key string, members ...interface{}) *redis.BoolSliceCmd {
	if r.client != nil {
		return r.client.SMIsMember(ctx, key, members...)
	}
	return r.clusterClient.SMIsMember(ctx, key, members...)
}

// SPop 随机弹出集合成员
func (r *RedisClient) SPop(ctx context.Context, key string) *redis.StringCmd {
	if r.client != nil {
		return r.client.SPop(ctx, key)
	}
	return r.clusterClient.SPop(ctx, key)
}

// SPopN 随机弹出多个集合成员
func (r *RedisClient) SPopN(ctx context.Context, key string, count int64) *redis.StringSliceCmd {
	if r.client != nil {
		return r.client.SPopN(ctx, key, count)
	}
	return r.clusterClient.SPopN(ctx, key, count)
}

// SRandMember 随机获取集合成员
func (r *RedisClient) SRandMember(ctx context.Context, key string) *redis.StringCmd {
	if r.client != nil {
		return r.client.SRandMember(ctx, key)
	}
	return r.clusterClient.SRandMember(ctx, key)
}

// SRandMemberN 随机获取多个集合成员
func (r *RedisClient) SRandMemberN(ctx context.Context, key string, count int64) *redis.StringSliceCmd {
	if r.client != nil {
		return r.client.SRandMemberN(ctx, key, count)
	}
	return r.clusterClient.SRandMemberN(ctx, key, count)
}

// SMove 移动成员到另一集合
func (r *RedisClient) SMove(ctx context.Context, source, destination string, member interface{}) *redis.BoolCmd {
	if r.client != nil {
		return r.client.SMove(ctx, source, destination, member)
	}
	return r.clusterClient.SMove(ctx, source, destination, member)
}

// SInter 集合交集
func (r *RedisClient) SInter(ctx context.Context, keys ...string) *redis.StringSliceCmd {
	if r.client != nil {
		return r.client.SInter(ctx, keys...)
	}
	return r.clusterClient.SInter(ctx, keys...)
}

// SInterCard 集合交集基数
func (r *RedisClient) SInterCard(ctx context.Context, limit int64, keys ...string) *redis.IntCmd {
	if r.client != nil {
		return r.client.SInterCard(ctx, limit, keys...)
	}
	return r.clusterClient.SInterCard(ctx, limit, keys...)
}

// SInterStore 集合并集并存储
func (r *RedisClient) SInterStore(ctx context.Context, destination string, keys ...string) *redis.IntCmd {
	if r.client != nil {
		return r.client.SInterStore(ctx, destination, keys...)
	}
	return r.clusterClient.SInterStore(ctx, destination, keys...)
}

// SUnion 集合并集
func (r *RedisClient) SUnion(ctx context.Context, keys ...string) *redis.StringSliceCmd {
	if r.client != nil {
		return r.client.SUnion(ctx, keys...)
	}
	return r.clusterClient.SUnion(ctx, keys...)
}

// SUnionStore 集合并集并存储
func (r *RedisClient) SUnionStore(ctx context.Context, destination string, keys ...string) *redis.IntCmd {
	if r.client != nil {
		return r.client.SUnionStore(ctx, destination, keys...)
	}
	return r.clusterClient.SUnionStore(ctx, destination, keys...)
}

// SDiff 集合差集
func (r *RedisClient) SDiff(ctx context.Context, keys ...string) *redis.StringSliceCmd {
	if r.client != nil {
		return r.client.SDiff(ctx, keys...)
	}
	return r.clusterClient.SDiff(ctx, keys...)
}

// SDiffStore 集合差集并存储
func (r *RedisClient) SDiffStore(ctx context.Context, destination string, keys ...string) *redis.IntCmd {
	if r.client != nil {
		return r.client.SDiffStore(ctx, destination, keys...)
	}
	return r.clusterClient.SDiffStore(ctx, destination, keys...)
}

// SScan 增量迭代集合成员
func (r *RedisClient) SScan(ctx context.Context, key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	if r.client != nil {
		return r.client.SScan(ctx, key, cursor, match, count)
	}
	return r.clusterClient.SScan(ctx, key, cursor, match, count)
}
