package redisClient

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

// ZAdd 向有序集合添加成员
func (r *RedisClient) ZAdd(ctx context.Context, key string, members ...redis.Z) *redis.IntCmd {
	if r.client != nil {
		return r.client.ZAdd(ctx, key, members...)
	}
	return r.clusterClient.ZAdd(ctx, key, members...)
}

// ZAddLT 仅当新分数小于当前分数时更新
func (r *RedisClient) ZAddLT(ctx context.Context, key string, members ...redis.Z) *redis.IntCmd {
	if r.client != nil {
		return r.client.ZAddLT(ctx, key, members...)
	}
	return r.clusterClient.ZAddLT(ctx, key, members...)
}

// ZAddGT 仅当新分数大于当前分数时更新
func (r *RedisClient) ZAddGT(ctx context.Context, key string, members ...redis.Z) *redis.IntCmd {
	if r.client != nil {
		return r.client.ZAddGT(ctx, key, members...)
	}
	return r.clusterClient.ZAddGT(ctx, key, members...)
}

// ZAddNX 仅当成员不存在时添加
func (r *RedisClient) ZAddNX(ctx context.Context, key string, members ...redis.Z) *redis.IntCmd {
	if r.client != nil {
		return r.client.ZAddNX(ctx, key, members...)
	}
	return r.clusterClient.ZAddNX(ctx, key, members...)
}

// ZAddXX 仅当成员已存在时更新
func (r *RedisClient) ZAddXX(ctx context.Context, key string, members ...redis.Z) *redis.IntCmd {
	if r.client != nil {
		return r.client.ZAddXX(ctx, key, members...)
	}
	return r.clusterClient.ZAddXX(ctx, key, members...)
}

// ZRem 删除有序集合成员
func (r *RedisClient) ZRem(ctx context.Context, key string, members ...interface{}) *redis.IntCmd {
	if r.client != nil {
		return r.client.ZRem(ctx, key, members...)
	}
	return r.clusterClient.ZRem(ctx, key, members...)
}

// ZRemRangeByScore 按分数范围删除成员
func (r *RedisClient) ZRemRangeByScore(ctx context.Context, key, min, max string) *redis.IntCmd {
	if r.client != nil {
		return r.client.ZRemRangeByScore(ctx, key, min, max)
	}
	return r.clusterClient.ZRemRangeByScore(ctx, key, min, max)
}

// ZRemRangeByRank 按排名范围删除成员
func (r *RedisClient) ZRemRangeByRank(ctx context.Context, key string, start, stop int64) *redis.IntCmd {
	if r.client != nil {
		return r.client.ZRemRangeByRank(ctx, key, start, stop)
	}
	return r.clusterClient.ZRemRangeByRank(ctx, key, start, stop)
}

// ZRemRangeByLex 按字典序范围删除成员
func (r *RedisClient) ZRemRangeByLex(ctx context.Context, key, min, max string) *redis.IntCmd {
	if r.client != nil {
		return r.client.ZRemRangeByLex(ctx, key, min, max)
	}
	return r.clusterClient.ZRemRangeByLex(ctx, key, min, max)
}

// ZCard 获取有序集合成员数
func (r *RedisClient) ZCard(ctx context.Context, key string) *redis.IntCmd {
	if r.client != nil {
		return r.client.ZCard(ctx, key)
	}
	return r.clusterClient.ZCard(ctx, key)
}

// ZCount 获取分数范围内的成员数
func (r *RedisClient) ZCount(ctx context.Context, key, min, max string) *redis.IntCmd {
	if r.client != nil {
		return r.client.ZCount(ctx, key, min, max)
	}
	return r.clusterClient.ZCount(ctx, key, min, max)
}

// ZScore 获取有序集合成员分数
func (r *RedisClient) ZScore(ctx context.Context, key, member string) *redis.FloatCmd {
	if r.client != nil {
		return r.client.ZScore(ctx, key, member)
	}
	return r.clusterClient.ZScore(ctx, key, member)
}

// ZMScore 批量获取有序集合成员分数
func (r *RedisClient) ZMScore(ctx context.Context, key string, members ...string) *redis.FloatSliceCmd {
	if r.client != nil {
		return r.client.ZMScore(ctx, key, members...)
	}
	return r.clusterClient.ZMScore(ctx, key, members...)
}

// ZIncrBy 增加成员分数
func (r *RedisClient) ZIncrBy(ctx context.Context, key string, increment float64, member string) *redis.FloatCmd {
	if r.client != nil {
		return r.client.ZIncrBy(ctx, key, increment, member)
	}
	return r.clusterClient.ZIncrBy(ctx, key, increment, member)
}

// ZRank 获取成员正序排名
func (r *RedisClient) ZRank(ctx context.Context, key, member string) *redis.IntCmd {
	if r.client != nil {
		return r.client.ZRank(ctx, key, member)
	}
	return r.clusterClient.ZRank(ctx, key, member)
}

// ZRevRank 获取成员逆序排名
func (r *RedisClient) ZRevRank(ctx context.Context, key, member string) *redis.IntCmd {
	if r.client != nil {
		return r.client.ZRevRank(ctx, key, member)
	}
	return r.clusterClient.ZRevRank(ctx, key, member)
}

// ZRange 按排名正序查询成员
func (r *RedisClient) ZRange(ctx context.Context, key string, start, stop int64) *redis.StringSliceCmd {
	if r.client != nil {
		return r.client.ZRange(ctx, key, start, stop)
	}
	return r.clusterClient.ZRange(ctx, key, start, stop)
}

// ZRangeWithScores 按排名正序查询成员（含分数）
func (r *RedisClient) ZRangeWithScores(ctx context.Context, key string, start, stop int64) *redis.ZSliceCmd {
	if r.client != nil {
		return r.client.ZRangeWithScores(ctx, key, start, stop)
	}
	return r.clusterClient.ZRangeWithScores(ctx, key, start, stop)
}

// ZRevRange 按排名逆序查询成员
func (r *RedisClient) ZRevRange(ctx context.Context, key string, start, stop int64) *redis.StringSliceCmd {
	if r.client != nil {
		return r.client.ZRevRange(ctx, key, start, stop)
	}
	return r.clusterClient.ZRevRange(ctx, key, start, stop)
}

// ZRevRangeWithScores 按排名逆序查询成员（含分数）
func (r *RedisClient) ZRevRangeWithScores(ctx context.Context, key string, start, stop int64) *redis.ZSliceCmd {
	if r.client != nil {
		return r.client.ZRevRangeWithScores(ctx, key, start, stop)
	}
	return r.clusterClient.ZRevRangeWithScores(ctx, key, start, stop)
}

// ZRangeByScore 按分数范围查询成员
func (r *RedisClient) ZRangeByScore(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.StringSliceCmd {
	if r.client != nil {
		return r.client.ZRangeByScore(ctx, key, opt)
	}
	return r.clusterClient.ZRangeByScore(ctx, key, opt)
}

// ZRevRangeByScoreWithScores 按分数范围逆序查询成员（含分数）
func (r *RedisClient) ZRevRangeByScoreWithScores(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.ZSliceCmd {
	if r.client != nil {
		return r.client.ZRevRangeByScoreWithScores(ctx, key, opt)
	}
	return r.clusterClient.ZRevRangeByScoreWithScores(ctx, key, opt)
}

// ZPopMax 弹出分数最高的成员
func (r *RedisClient) ZPopMax(ctx context.Context, key string, count ...int64) *redis.ZSliceCmd {
	if r.client != nil {
		return r.client.ZPopMax(ctx, key, count...)
	}
	return r.clusterClient.ZPopMax(ctx, key, count...)
}

// ZPopMin 弹出分数最低的成员
func (r *RedisClient) ZPopMin(ctx context.Context, key string, count ...int64) *redis.ZSliceCmd {
	if r.client != nil {
		return r.client.ZPopMin(ctx, key, count...)
	}
	return r.clusterClient.ZPopMin(ctx, key, count...)
}

// BZPopMax 阻塞式弹出分数最高的成员
func (r *RedisClient) BZPopMax(ctx context.Context, timeout time.Duration, keys ...string) *redis.ZWithKeyCmd {
	if r.client != nil {
		return r.client.BZPopMax(ctx, timeout, keys...)
	}
	return r.clusterClient.BZPopMax(ctx, timeout, keys...)
}

// BZPopMin 阻塞式弹出分数最低的成员
func (r *RedisClient) BZPopMin(ctx context.Context, timeout time.Duration, keys ...string) *redis.ZWithKeyCmd {
	if r.client != nil {
		return r.client.BZPopMin(ctx, timeout, keys...)
	}
	return r.clusterClient.BZPopMin(ctx, timeout, keys...)
}

// ZInter 有序集合交集
func (r *RedisClient) ZInter(ctx context.Context, store *redis.ZStore) *redis.StringSliceCmd {
	if r.client != nil {
		return r.client.ZInter(ctx, store)
	}
	return r.clusterClient.ZInter(ctx, store)
}

// ZInterWithScores 有序集合交集（含分数）
func (r *RedisClient) ZInterWithScores(ctx context.Context, store *redis.ZStore) *redis.ZSliceCmd {
	if r.client != nil {
		return r.client.ZInterWithScores(ctx, store)
	}
	return r.clusterClient.ZInterWithScores(ctx, store)
}

// ZInterStore 有序集合交集并存储
func (r *RedisClient) ZInterStore(ctx context.Context, destination string, store *redis.ZStore) *redis.IntCmd {
	if r.client != nil {
		return r.client.ZInterStore(ctx, destination, store)
	}
	return r.clusterClient.ZInterStore(ctx, destination, store)
}

// ZUnion 有序集合并集
func (r *RedisClient) ZUnion(ctx context.Context, store redis.ZStore) *redis.StringSliceCmd {
	if r.client != nil {
		return r.client.ZUnion(ctx, store)
	}
	return r.clusterClient.ZUnion(ctx, store)
}

// ZUnionWithScores 有序集合并集（含分数）
func (r *RedisClient) ZUnionWithScores(ctx context.Context, store redis.ZStore) *redis.ZSliceCmd {
	if r.client != nil {
		return r.client.ZUnionWithScores(ctx, store)
	}
	return r.clusterClient.ZUnionWithScores(ctx, store)
}

// ZUnionStore 有序集合并集并存储
func (r *RedisClient) ZUnionStore(ctx context.Context, dest string, store *redis.ZStore) *redis.IntCmd {
	if r.client != nil {
		return r.client.ZUnionStore(ctx, dest, store)
	}
	return r.clusterClient.ZUnionStore(ctx, dest, store)
}

// ZDiff 有序集合差集
func (r *RedisClient) ZDiff(ctx context.Context, keys ...string) *redis.StringSliceCmd {
	if r.client != nil {
		return r.client.ZDiff(ctx, keys...)
	}
	return r.clusterClient.ZDiff(ctx, keys...)
}

// ZDiffWithScores 有序集合差集（含分数）
func (r *RedisClient) ZDiffWithScores(ctx context.Context, keys ...string) *redis.ZSliceCmd {
	if r.client != nil {
		return r.client.ZDiffWithScores(ctx, keys...)
	}
	return r.clusterClient.ZDiffWithScores(ctx, keys...)
}

// ZDiffStore 有序集合差集并存储
func (r *RedisClient) ZDiffStore(ctx context.Context, destination string, keys ...string) *redis.IntCmd {
	if r.client != nil {
		return r.client.ZDiffStore(ctx, destination, keys...)
	}
	return r.clusterClient.ZDiffStore(ctx, destination, keys...)
}

// ZRandMember 随机获取有序集合成员
func (r *RedisClient) ZRandMember(ctx context.Context, key string, count int) *redis.StringSliceCmd {
	if r.client != nil {
		return r.client.ZRandMember(ctx, key, count)
	}
	return r.clusterClient.ZRandMember(ctx, key, count)
}

// ZRandMemberWithScores 随机获取有序集合成员（含分数）
func (r *RedisClient) ZRandMemberWithScores(ctx context.Context, key string, count int) *redis.ZSliceCmd {
	if r.client != nil {
		return r.client.ZRandMemberWithScores(ctx, key, count)
	}
	return r.clusterClient.ZRandMemberWithScores(ctx, key, count)
}

// ZScan 增量迭代有序集合成员
func (r *RedisClient) ZScan(ctx context.Context, key string, cursor uint64, match string, count int64) *redis.ScanCmd {
	if r.client != nil {
		return r.client.ZScan(ctx, key, cursor, match, count)
	}
	return r.clusterClient.ZScan(ctx, key, cursor, match, count)
}

// BZMPop 阻塞式从多个有序集合弹出元素
func (r *RedisClient) BZMPop(ctx context.Context, timeout time.Duration, order string, count int64, keys ...string) *redis.ZSliceWithKeyCmd {
	if r.client != nil {
		return r.client.BZMPop(ctx, timeout, order, count, keys...)
	}
	return r.clusterClient.BZMPop(ctx, timeout, order, count, keys...)
}

// ZAddArgs 添加有序集合成员（带完整参数）
func (r *RedisClient) ZAddArgs(ctx context.Context, key string, args redis.ZAddArgs) *redis.IntCmd {
	if r.client != nil {
		return r.client.ZAddArgs(ctx, key, args)
	}
	return r.clusterClient.ZAddArgs(ctx, key, args)
}

// ZAddArgsIncr 添加有序集合成员（带参数和增量）
func (r *RedisClient) ZAddArgsIncr(ctx context.Context, key string, args redis.ZAddArgs) *redis.FloatCmd {
	if r.client != nil {
		return r.client.ZAddArgsIncr(ctx, key, args)
	}
	return r.clusterClient.ZAddArgsIncr(ctx, key, args)
}

// ZInterCard 计算有序集合交集的基数
func (r *RedisClient) ZInterCard(ctx context.Context, limit int64, keys ...string) *redis.IntCmd {
	if r.client != nil {
		return r.client.ZInterCard(ctx, limit, keys...)
	}
	return r.clusterClient.ZInterCard(ctx, limit, keys...)
}

// ZLexCount 按字典序统计有序集合成员数量
func (r *RedisClient) ZLexCount(ctx context.Context, key, min, max string) *redis.IntCmd {
	if r.client != nil {
		return r.client.ZLexCount(ctx, key, min, max)
	}
	return r.clusterClient.ZLexCount(ctx, key, min, max)
}

// ZMPop 从多个有序集合弹出元素
func (r *RedisClient) ZMPop(ctx context.Context, order string, count int64, keys ...string) *redis.ZSliceWithKeyCmd {
	if r.client != nil {
		return r.client.ZMPop(ctx, order, count, keys...)
	}
	return r.clusterClient.ZMPop(ctx, order, count, keys...)
}

// ZRangeArgs 按复杂参数范围获取成员
func (r *RedisClient) ZRangeArgs(ctx context.Context, z redis.ZRangeArgs) *redis.StringSliceCmd {
	if r.client != nil {
		return r.client.ZRangeArgs(ctx, z)
	}
	return r.clusterClient.ZRangeArgs(ctx, z)
}

// ZRangeArgsWithScores 按复杂参数范围获取成员（带分数）
func (r *RedisClient) ZRangeArgsWithScores(ctx context.Context, z redis.ZRangeArgs) *redis.ZSliceCmd {
	if r.client != nil {
		return r.client.ZRangeArgsWithScores(ctx, z)
	}
	return r.clusterClient.ZRangeArgsWithScores(ctx, z)
}

// ZRangeByLex 按字典序范围获取成员
func (r *RedisClient) ZRangeByLex(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.StringSliceCmd {
	if r.client != nil {
		return r.client.ZRangeByLex(ctx, key, opt)
	}
	return r.clusterClient.ZRangeByLex(ctx, key, opt)
}

// ZRangeByScoreWithScores 按分数范围获取成员（带分数）
func (r *RedisClient) ZRangeByScoreWithScores(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.ZSliceCmd {
	if r.client != nil {
		return r.client.ZRangeByScoreWithScores(ctx, key, opt)
	}
	return r.clusterClient.ZRangeByScoreWithScores(ctx, key, opt)
}

// ZRangeStore 按复杂参数范围获取成员并存储
func (r *RedisClient) ZRangeStore(ctx context.Context, dst string, z redis.ZRangeArgs) *redis.IntCmd {
	if r.client != nil {
		return r.client.ZRangeStore(ctx, dst, z)
	}
	return r.clusterClient.ZRangeStore(ctx, dst, z)
}

// ZRankWithScore 获取成员排名（带分数）
func (r *RedisClient) ZRankWithScore(ctx context.Context, key, member string) *redis.RankWithScoreCmd {
	if r.client != nil {
		return r.client.ZRankWithScore(ctx, key, member)
	}
	return r.clusterClient.ZRankWithScore(ctx, key, member)
}

// ZRevRangeByLex 按字典序反向获取成员
func (r *RedisClient) ZRevRangeByLex(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.StringSliceCmd {
	if r.client != nil {
		return r.client.ZRevRangeByLex(ctx, key, opt)
	}
	return r.clusterClient.ZRevRangeByLex(ctx, key, opt)
}

// ZRevRangeByScore 按分数反向获取成员
func (r *RedisClient) ZRevRangeByScore(ctx context.Context, key string, opt *redis.ZRangeBy) *redis.StringSliceCmd {
	if r.client != nil {
		return r.client.ZRevRangeByScore(ctx, key, opt)
	}
	return r.clusterClient.ZRevRangeByScore(ctx, key, opt)
}

// ZRevRankWithScore 获取成员反向排名（带分数）
func (r *RedisClient) ZRevRankWithScore(ctx context.Context, key, member string) *redis.RankWithScoreCmd {
	if r.client != nil {
		return r.client.ZRevRankWithScore(ctx, key, member)
	}
	return r.clusterClient.ZRevRankWithScore(ctx, key, member)
}
