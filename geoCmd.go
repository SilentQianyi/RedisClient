package redisClient

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// GeoAdd 添加地理位置
func (r *RedisClient) GeoAdd(ctx context.Context, key string, geoLocation ...*GeoLocation) *redis.IntCmd {
	if r.client != nil {
		return r.client.GeoAdd(ctx, key, geoLocation...)
	}
	return r.clusterClient.GeoAdd(ctx, key, geoLocation...)
}

// GeoPos 获取地理位置坐标
func (r *RedisClient) GeoPos(ctx context.Context, key string, members ...string) *redis.GeoPosCmd {
	if r.client != nil {
		return r.client.GeoPos(ctx, key, members...)
	}
	return r.clusterClient.GeoPos(ctx, key, members...)
}

// GeoRadius 按半径查询地理位置
func (r *RedisClient) GeoRadius(ctx context.Context, key string, longitude, latitude float64, query *GeoRadiusQuery) *redis.GeoLocationCmd {
	if r.client != nil {
		return r.client.GeoRadius(ctx, key, longitude, latitude, query)
	}
	return r.clusterClient.GeoRadius(ctx, key, longitude, latitude, query)
}

// GeoRadiusStore 按半径查询并存储结果
func (r *RedisClient) GeoRadiusStore(ctx context.Context, key string, longitude, latitude float64, query *GeoRadiusQuery) *redis.IntCmd {
	if r.client != nil {
		return r.client.GeoRadiusStore(ctx, key, longitude, latitude, query)
	}
	return r.clusterClient.GeoRadiusStore(ctx, key, longitude, latitude, query)
}

// GeoRadiusByMember 按成员半径查询地理位置
func (r *RedisClient) GeoRadiusByMember(ctx context.Context, key, member string, query *GeoRadiusQuery) *redis.GeoLocationCmd {
	if r.client != nil {
		return r.client.GeoRadiusByMember(ctx, key, member, query)
	}
	return r.clusterClient.GeoRadiusByMember(ctx, key, member, query)
}

// GeoRadiusByMemberStore 按成员半径查询并存储结果
func (r *RedisClient) GeoRadiusByMemberStore(ctx context.Context, key, member string, query *GeoRadiusQuery) *redis.IntCmd {
	if r.client != nil {
		return r.client.GeoRadiusByMemberStore(ctx, key, member, query)
	}
	return r.clusterClient.GeoRadiusByMemberStore(ctx, key, member, query)
}

// GeoSearch 搜索地理位置
func (r *RedisClient) GeoSearch(ctx context.Context, key string, q *GeoSearchQuery) *redis.StringSliceCmd {
	if r.client != nil {
		return r.client.GeoSearch(ctx, key, q)
	}
	return r.clusterClient.GeoSearch(ctx, key, q)
}

// GeoSearchLocation 搜索地理位置（含坐标）
func (r *RedisClient) GeoSearchLocation(ctx context.Context, key string, q *GeoSearchLocationQuery) *redis.GeoSearchLocationCmd {
	if r.client != nil {
		return r.client.GeoSearchLocation(ctx, key, q)
	}
	return r.clusterClient.GeoSearchLocation(ctx, key, q)
}

// GeoSearchStore 搜索地理位置并存储
func (r *RedisClient) GeoSearchStore(ctx context.Context, key, store string, q *GeoSearchStoreQuery) *redis.IntCmd {
	if r.client != nil {
		return r.client.GeoSearchStore(ctx, key, store, q)
	}
	return r.clusterClient.GeoSearchStore(ctx, key, store, q)
}

// GeoDist 计算两个地理位置的距离
func (r *RedisClient) GeoDist(ctx context.Context, key string, member1, member2, unit string) *redis.FloatCmd {
	if r.client != nil {
		return r.client.GeoDist(ctx, key, member1, member2, unit)
	}
	return r.clusterClient.GeoDist(ctx, key, member1, member2, unit)
}

// GeoHash 获取地理位置 geohash
func (r *RedisClient) GeoHash(ctx context.Context, key string, members ...string) *redis.StringSliceCmd {
	if r.client != nil {
		return r.client.GeoHash(ctx, key, members...)
	}
	return r.clusterClient.GeoHash(ctx, key, members...)
}
