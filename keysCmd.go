package redisClient

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

// ============================================================================
// 键操作
// ============================================================================

// Del 删除键
func (r *RedisClient) Del(ctx context.Context, keys ...string) *redis.IntCmd {
	if r.client != nil {
		return r.client.Del(ctx, keys...)
	}
	return r.clusterClient.Del(ctx, keys...)
}

// Unlink 异步删除键
func (r *RedisClient) Unlink(ctx context.Context, keys ...string) *redis.IntCmd {
	if r.client != nil {
		return r.client.Unlink(ctx, keys...)
	}
	return r.clusterClient.Unlink(ctx, keys...)
}

// Exists 检查键是否存在
func (r *RedisClient) Exists(ctx context.Context, keys ...string) *redis.IntCmd {
	if r.client != nil {
		return r.client.Exists(ctx, keys...)
	}
	return r.clusterClient.Exists(ctx, keys...)
}

// Expire 设置键过期时间
func (r *RedisClient) Expire(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd {
	if r.client != nil {
		return r.client.Expire(ctx, key, expiration)
	}
	return r.clusterClient.Expire(ctx, key, expiration)
}

// ExpireAt 设置键过期时间点
func (r *RedisClient) ExpireAt(ctx context.Context, key string, tm time.Time) *redis.BoolCmd {
	if r.client != nil {
		return r.client.ExpireAt(ctx, key, tm)
	}
	return r.clusterClient.ExpireAt(ctx, key, tm)
}

// ExpireTime 获取键过期时间点
func (r *RedisClient) ExpireTime(ctx context.Context, key string) *redis.DurationCmd {
	if r.client != nil {
		return r.client.ExpireTime(ctx, key)
	}
	return r.clusterClient.ExpireTime(ctx, key)
}

// ExpireNX 仅当键没有过期时间时设置过期
func (r *RedisClient) ExpireNX(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd {
	if r.client != nil {
		return r.client.ExpireNX(ctx, key, expiration)
	}
	return r.clusterClient.ExpireNX(ctx, key, expiration)
}

// ExpireXX 仅当键已有过期时间时设置过期
func (r *RedisClient) ExpireXX(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd {
	if r.client != nil {
		return r.client.ExpireXX(ctx, key, expiration)
	}
	return r.clusterClient.ExpireXX(ctx, key, expiration)
}

// ExpireGT 仅当新过期时间大于当前时设置
func (r *RedisClient) ExpireGT(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd {
	if r.client != nil {
		return r.client.ExpireGT(ctx, key, expiration)
	}
	return r.clusterClient.ExpireGT(ctx, key, expiration)
}

// ExpireLT 仅当新过期时间小于当前时设置
func (r *RedisClient) ExpireLT(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd {
	if r.client != nil {
		return r.client.ExpireLT(ctx, key, expiration)
	}
	return r.clusterClient.ExpireLT(ctx, key, expiration)
}

// PExpire 以毫秒设置键过期时间
func (r *RedisClient) PExpire(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd {
	if r.client != nil {
		return r.client.PExpire(ctx, key, expiration)
	}
	return r.clusterClient.PExpire(ctx, key, expiration)
}

// PExpireAt 以毫秒时间戳设置键过期时间点
func (r *RedisClient) PExpireAt(ctx context.Context, key string, tm time.Time) *redis.BoolCmd {
	if r.client != nil {
		return r.client.PExpireAt(ctx, key, tm)
	}
	return r.clusterClient.PExpireAt(ctx, key, tm)
}

// PExpireTime 获取键过期时间的毫秒时间戳
func (r *RedisClient) PExpireTime(ctx context.Context, key string) *redis.DurationCmd {
	if r.client != nil {
		return r.client.PExpireTime(ctx, key)
	}
	return r.clusterClient.PExpireTime(ctx, key)
}

// TTL 获取键剩余生存时间
func (r *RedisClient) TTL(ctx context.Context, key string) *redis.DurationCmd {
	if r.client != nil {
		return r.client.TTL(ctx, key)
	}
	return r.clusterClient.TTL(ctx, key)
}

// PTTL 获取键剩余生存时间（毫秒）
func (r *RedisClient) PTTL(ctx context.Context, key string) *redis.DurationCmd {
	if r.client != nil {
		return r.client.PTTL(ctx, key)
	}
	return r.clusterClient.PTTL(ctx, key)
}

// Persist 移除键的过期时间
func (r *RedisClient) Persist(ctx context.Context, key string) *redis.BoolCmd {
	if r.client != nil {
		return r.client.Persist(ctx, key)
	}
	return r.clusterClient.Persist(ctx, key)
}

// Keys 模糊查找键
func (r *RedisClient) Keys(ctx context.Context, pattern string) *redis.StringSliceCmd {
	if r.client != nil {
		return r.client.Keys(ctx, pattern)
	}
	return r.clusterClient.Keys(ctx, pattern)
}

// Type 获取键类型
func (r *RedisClient) Type(ctx context.Context, key string) *redis.StatusCmd {
	if r.client != nil {
		return r.client.Type(ctx, key)
	}
	return r.clusterClient.Type(ctx, key)
}

// RandomKey 随机返回一个键
func (r *RedisClient) RandomKey(ctx context.Context) *redis.StringCmd {
	if r.client != nil {
		return r.client.RandomKey(ctx)
	}
	return r.clusterClient.RandomKey(ctx)
}

// Rename 重命名键
func (r *RedisClient) Rename(ctx context.Context, key, newkey string) *redis.StatusCmd {
	if r.client != nil {
		return r.client.Rename(ctx, key, newkey)
	}
	return r.clusterClient.Rename(ctx, key, newkey)
}

// RenameNX 仅当新键名不存在时重命名
func (r *RedisClient) RenameNX(ctx context.Context, key, newkey string) *redis.BoolCmd {
	if r.client != nil {
		return r.client.RenameNX(ctx, key, newkey)
	}
	return r.clusterClient.RenameNX(ctx, key, newkey)
}

// Move 将键移动到另一个数据库
func (r *RedisClient) Move(ctx context.Context, key string, db int) *redis.BoolCmd {
	if r.client != nil {
		return r.client.Move(ctx, key, db)
	}
	return r.clusterClient.Move(ctx, key, db)
}

// Copy 复制键
func (r *RedisClient) Copy(ctx context.Context, sourceKey string, destKey string, db int, replace bool) *redis.IntCmd {
	if r.client != nil {
		return r.client.Copy(ctx, sourceKey, destKey, db, replace)
	}
	return r.clusterClient.Copy(ctx, sourceKey, destKey, db, replace)
}

// Touch 更新键的最后访问时间
func (r *RedisClient) Touch(ctx context.Context, keys ...string) *redis.IntCmd {
	if r.client != nil {
		return r.client.Touch(ctx, keys...)
	}
	return r.clusterClient.Touch(ctx, keys...)
}

// Sort 排序列表/集合/有序集合元素
func (r *RedisClient) Sort(ctx context.Context, key string, sort *Sort) *redis.StringSliceCmd {
	if r.client != nil {
		return r.client.Sort(ctx, key, sort)
	}
	return r.clusterClient.Sort(ctx, key, sort)
}

// SortStore 排序并存储结果
func (r *RedisClient) SortStore(ctx context.Context, key, store string, sort *Sort) *redis.IntCmd {
	if r.client != nil {
		return r.client.SortStore(ctx, key, store, sort)
	}
	return r.clusterClient.SortStore(ctx, key, store, sort)
}

// SortInterfaces 排序并返回 interface 切片
func (r *RedisClient) SortInterfaces(ctx context.Context, key string, sort *Sort) *redis.SliceCmd {
	if r.client != nil {
		return r.client.SortInterfaces(ctx, key, sort)
	}
	return r.clusterClient.SortInterfaces(ctx, key, sort)
}

// SortRO 只读排序
func (r *RedisClient) SortRO(ctx context.Context, key string, sort *Sort) *redis.StringSliceCmd {
	if r.client != nil {
		return r.client.SortRO(ctx, key, sort)
	}
	return r.clusterClient.SortRO(ctx, key, sort)
}

// Dump 序列化键值
func (r *RedisClient) Dump(ctx context.Context, key string) *redis.StringCmd {
	if r.client != nil {
		return r.client.Dump(ctx, key)
	}
	return r.clusterClient.Dump(ctx, key)
}

// Restore 反序列化键值
func (r *RedisClient) Restore(ctx context.Context, key string, ttl time.Duration, value string) *redis.StatusCmd {
	if r.client != nil {
		return r.client.Restore(ctx, key, ttl, value)
	}
	return r.clusterClient.Restore(ctx, key, ttl, value)
}

// RestoreReplace 反序列化键值（替换模式）
func (r *RedisClient) RestoreReplace(ctx context.Context, key string, ttl time.Duration, value string) *redis.StatusCmd {
	if r.client != nil {
		return r.client.RestoreReplace(ctx, key, ttl, value)
	}
	return r.clusterClient.RestoreReplace(ctx, key, ttl, value)
}

// ObjectFreq 获取键访问频率
func (r *RedisClient) ObjectFreq(ctx context.Context, key string) *redis.IntCmd {
	if r.client != nil {
		return r.client.ObjectFreq(ctx, key)
	}
	return r.clusterClient.ObjectFreq(ctx, key)
}

// ObjectRefCount 获取键引用计数
func (r *RedisClient) ObjectRefCount(ctx context.Context, key string) *redis.IntCmd {
	if r.client != nil {
		return r.client.ObjectRefCount(ctx, key)
	}
	return r.clusterClient.ObjectRefCount(ctx, key)
}

// ObjectEncoding 获取键内部编码
func (r *RedisClient) ObjectEncoding(ctx context.Context, key string) *redis.StringCmd {
	if r.client != nil {
		return r.client.ObjectEncoding(ctx, key)
	}
	return r.clusterClient.ObjectEncoding(ctx, key)
}

// ObjectIdleTime 获取键空闲时间
func (r *RedisClient) ObjectIdleTime(ctx context.Context, key string) *redis.DurationCmd {
	if r.client != nil {
		return r.client.ObjectIdleTime(ctx, key)
	}
	return r.clusterClient.ObjectIdleTime(ctx, key)
}

// Scan 增量迭代键空间
func (r *RedisClient) Scan(ctx context.Context, cursor uint64, match string, count int64) *redis.ScanCmd {
	if r.client != nil {
		return r.client.Scan(ctx, cursor, match, count)
	}
	return r.clusterClient.Scan(ctx, cursor, match, count)
}

// ScanType 按类型增量迭代键空间
func (r *RedisClient) ScanType(ctx context.Context, cursor uint64, match string, count int64, keyType string) *redis.ScanCmd {
	if r.client != nil {
		return r.client.ScanType(ctx, cursor, match, count, keyType)
	}
	return r.clusterClient.ScanType(ctx, cursor, match, count, keyType)
}

// Migrate 迁移键到另一 Redis 实例
func (r *RedisClient) Migrate(ctx context.Context, host, port, key string, db int, timeout time.Duration) *redis.StatusCmd {
	if r.client != nil {
		return r.client.Migrate(ctx, host, port, key, db, timeout)
	}
	return r.clusterClient.Migrate(ctx, host, port, key, db, timeout)
}
