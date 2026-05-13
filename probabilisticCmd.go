package redisClient

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// ============================================================================
// Bloom Filter
// ============================================================================

// BFAdd 向 Bloom Filter 添加元素
func (r *RedisClient) BFAdd(ctx context.Context, key string, element interface{}) *redis.BoolCmd {
	if r.client != nil {
		return r.client.BFAdd(ctx, key, element)
	}
	return r.clusterClient.BFAdd(ctx, key, element)
}

// BFCard 获取 Bloom Filter 的基数
func (r *RedisClient) BFCard(ctx context.Context, key string) *redis.IntCmd {
	if r.client != nil {
		return r.client.BFCard(ctx, key)
	}
	return r.clusterClient.BFCard(ctx, key)
}

// BFExists 检查元素是否在 Bloom Filter 中
func (r *RedisClient) BFExists(ctx context.Context, key string, element interface{}) *redis.BoolCmd {
	if r.client != nil {
		return r.client.BFExists(ctx, key, element)
	}
	return r.clusterClient.BFExists(ctx, key, element)
}

// BFInfo 获取 Bloom Filter 信息
func (r *RedisClient) BFInfo(ctx context.Context, key string) *redis.BFInfoCmd {
	if r.client != nil {
		return r.client.BFInfo(ctx, key)
	}
	return r.clusterClient.BFInfo(ctx, key)
}

// BFInfoArg 获取 Bloom Filter 指定信息
func (r *RedisClient) BFInfoArg(ctx context.Context, key, option string) *redis.BFInfoCmd {
	if r.client != nil {
		return r.client.BFInfoArg(ctx, key, option)
	}
	return r.clusterClient.BFInfoArg(ctx, key, option)
}

// BFInfoCapacity 获取 Bloom Filter 容量信息
func (r *RedisClient) BFInfoCapacity(ctx context.Context, key string) *redis.BFInfoCmd {
	if r.client != nil {
		return r.client.BFInfoCapacity(ctx, key)
	}
	return r.clusterClient.BFInfoCapacity(ctx, key)
}

// BFInfoSize 获取 Bloom Filter 大小信息
func (r *RedisClient) BFInfoSize(ctx context.Context, key string) *redis.BFInfoCmd {
	if r.client != nil {
		return r.client.BFInfoSize(ctx, key)
	}
	return r.clusterClient.BFInfoSize(ctx, key)
}

// BFInfoFilters 获取 Bloom Filter 过滤器信息
func (r *RedisClient) BFInfoFilters(ctx context.Context, key string) *redis.BFInfoCmd {
	if r.client != nil {
		return r.client.BFInfoFilters(ctx, key)
	}
	return r.clusterClient.BFInfoFilters(ctx, key)
}

// BFInfoItems 获取 Bloom Filter 元素信息
func (r *RedisClient) BFInfoItems(ctx context.Context, key string) *redis.BFInfoCmd {
	if r.client != nil {
		return r.client.BFInfoItems(ctx, key)
	}
	return r.clusterClient.BFInfoItems(ctx, key)
}

// BFInfoExpansion 获取 Bloom Filter 扩展信息
func (r *RedisClient) BFInfoExpansion(ctx context.Context, key string) *redis.BFInfoCmd {
	if r.client != nil {
		return r.client.BFInfoExpansion(ctx, key)
	}
	return r.clusterClient.BFInfoExpansion(ctx, key)
}

// BFInsert 向 Bloom Filter 插入多个元素
func (r *RedisClient) BFInsert(ctx context.Context, key string, options *BFInsertOptions, elements ...interface{}) *redis.BoolSliceCmd {
	if r.client != nil {
		return r.client.BFInsert(ctx, key, options, elements...)
	}
	return r.clusterClient.BFInsert(ctx, key, options, elements...)
}

// BFMAdd 向 Bloom Filter 批量添加元素
func (r *RedisClient) BFMAdd(ctx context.Context, key string, elements ...interface{}) *redis.BoolSliceCmd {
	if r.client != nil {
		return r.client.BFMAdd(ctx, key, elements...)
	}
	return r.clusterClient.BFMAdd(ctx, key, elements...)
}

// BFMExists 批量检查元素是否在 Bloom Filter 中
func (r *RedisClient) BFMExists(ctx context.Context, key string, elements ...interface{}) *redis.BoolSliceCmd {
	if r.client != nil {
		return r.client.BFMExists(ctx, key, elements...)
	}
	return r.clusterClient.BFMExists(ctx, key, elements...)
}

// BFReserve 创建 Bloom Filter
func (r *RedisClient) BFReserve(ctx context.Context, key string, errorRate float64, capacity int64) *redis.StatusCmd {
	if r.client != nil {
		return r.client.BFReserve(ctx, key, errorRate, capacity)
	}
	return r.clusterClient.BFReserve(ctx, key, errorRate, capacity)
}

// BFReserveExpansion 创建 Bloom Filter（含扩展）
func (r *RedisClient) BFReserveExpansion(ctx context.Context, key string, errorRate float64, capacity, expansion int64) *redis.StatusCmd {
	if r.client != nil {
		return r.client.BFReserveExpansion(ctx, key, errorRate, capacity, expansion)
	}
	return r.clusterClient.BFReserveExpansion(ctx, key, errorRate, capacity, expansion)
}

// BFReserveNonScaling 创建非扩展 Bloom Filter
func (r *RedisClient) BFReserveNonScaling(ctx context.Context, key string, errorRate float64, capacity int64) *redis.StatusCmd {
	if r.client != nil {
		return r.client.BFReserveNonScaling(ctx, key, errorRate, capacity)
	}
	return r.clusterClient.BFReserveNonScaling(ctx, key, errorRate, capacity)
}

// BFReserveWithArgs 创建 Bloom Filter（带完整参数）
func (r *RedisClient) BFReserveWithArgs(ctx context.Context, key string, options *BFReserveOptions) *redis.StatusCmd {
	if r.client != nil {
		return r.client.BFReserveWithArgs(ctx, key, options)
	}
	return r.clusterClient.BFReserveWithArgs(ctx, key, options)
}

// BFScanDump 扫描并转储 Bloom Filter
func (r *RedisClient) BFScanDump(ctx context.Context, key string, iterator int64) *redis.ScanDumpCmd {
	if r.client != nil {
		return r.client.BFScanDump(ctx, key, iterator)
	}
	return r.clusterClient.BFScanDump(ctx, key, iterator)
}

// BFLoadChunk 加载 Bloom Filter 数据块
func (r *RedisClient) BFLoadChunk(ctx context.Context, key string, iterator int64, data interface{}) *redis.StatusCmd {
	if r.client != nil {
		return r.client.BFLoadChunk(ctx, key, iterator, data)
	}
	return r.clusterClient.BFLoadChunk(ctx, key, iterator, data)
}

// ============================================================================
// Cuckoo Filter
// ============================================================================

// CFAdd 向 Cuckoo Filter 添加元素
func (r *RedisClient) CFAdd(ctx context.Context, key string, element interface{}) *redis.BoolCmd {
	if r.client != nil {
		return r.client.CFAdd(ctx, key, element)
	}
	return r.clusterClient.CFAdd(ctx, key, element)
}

// CFAddNX 向 Cuckoo Filter 添加元素（不存在时）
func (r *RedisClient) CFAddNX(ctx context.Context, key string, element interface{}) *redis.BoolCmd {
	if r.client != nil {
		return r.client.CFAddNX(ctx, key, element)
	}
	return r.clusterClient.CFAddNX(ctx, key, element)
}

// CFCount 获取 Cuckoo Filter 中元素计数
func (r *RedisClient) CFCount(ctx context.Context, key string, element interface{}) *redis.IntCmd {
	if r.client != nil {
		return r.client.CFCount(ctx, key, element)
	}
	return r.clusterClient.CFCount(ctx, key, element)
}

// CFDel 从 Cuckoo Filter 删除元素
func (r *RedisClient) CFDel(ctx context.Context, key string, element interface{}) *redis.BoolCmd {
	if r.client != nil {
		return r.client.CFDel(ctx, key, element)
	}
	return r.clusterClient.CFDel(ctx, key, element)
}

// CFExists 检查元素是否在 Cuckoo Filter 中
func (r *RedisClient) CFExists(ctx context.Context, key string, element interface{}) *redis.BoolCmd {
	if r.client != nil {
		return r.client.CFExists(ctx, key, element)
	}
	return r.clusterClient.CFExists(ctx, key, element)
}

// CFInfo 获取 Cuckoo Filter 信息
func (r *RedisClient) CFInfo(ctx context.Context, key string) *redis.CFInfoCmd {
	if r.client != nil {
		return r.client.CFInfo(ctx, key)
	}
	return r.clusterClient.CFInfo(ctx, key)
}

// CFInsert 向 Cuckoo Filter 插入多个元素
func (r *RedisClient) CFInsert(ctx context.Context, key string, options *CFInsertOptions, elements ...interface{}) *redis.BoolSliceCmd {
	if r.client != nil {
		return r.client.CFInsert(ctx, key, options, elements...)
	}
	return r.clusterClient.CFInsert(ctx, key, options, elements...)
}

// CFInsertNX 向 Cuckoo Filter 插入多个元素（不存在时）
func (r *RedisClient) CFInsertNX(ctx context.Context, key string, options *CFInsertOptions, elements ...interface{}) *redis.IntSliceCmd {
	if r.client != nil {
		return r.client.CFInsertNX(ctx, key, options, elements...)
	}
	return r.clusterClient.CFInsertNX(ctx, key, options, elements...)
}

// CFMExists 批量检查元素是否在 Cuckoo Filter 中
func (r *RedisClient) CFMExists(ctx context.Context, key string, elements ...interface{}) *redis.BoolSliceCmd {
	if r.client != nil {
		return r.client.CFMExists(ctx, key, elements...)
	}
	return r.clusterClient.CFMExists(ctx, key, elements...)
}

// CFReserve 创建 Cuckoo Filter
func (r *RedisClient) CFReserve(ctx context.Context, key string, capacity int64) *redis.StatusCmd {
	if r.client != nil {
		return r.client.CFReserve(ctx, key, capacity)
	}
	return r.clusterClient.CFReserve(ctx, key, capacity)
}

// CFReserveWithArgs 创建 Cuckoo Filter（带完整参数）
func (r *RedisClient) CFReserveWithArgs(ctx context.Context, key string, options *CFReserveOptions) *redis.StatusCmd {
	if r.client != nil {
		return r.client.CFReserveWithArgs(ctx, key, options)
	}
	return r.clusterClient.CFReserveWithArgs(ctx, key, options)
}

// CFReserveExpansion 创建 Cuckoo Filter（含扩展）
func (r *RedisClient) CFReserveExpansion(ctx context.Context, key string, capacity int64, expansion int64) *redis.StatusCmd {
	if r.client != nil {
		return r.client.CFReserveExpansion(ctx, key, capacity, expansion)
	}
	return r.clusterClient.CFReserveExpansion(ctx, key, capacity, expansion)
}

// CFReserveBucketSize 创建 Cuckoo Filter（含桶大小）
func (r *RedisClient) CFReserveBucketSize(ctx context.Context, key string, capacity int64, bucketsize int64) *redis.StatusCmd {
	if r.client != nil {
		return r.client.CFReserveBucketSize(ctx, key, capacity, bucketsize)
	}
	return r.clusterClient.CFReserveBucketSize(ctx, key, capacity, bucketsize)
}

// CFReserveMaxIterations 创建 Cuckoo Filter（含最大迭代数）
func (r *RedisClient) CFReserveMaxIterations(ctx context.Context, key string, capacity int64, maxiterations int64) *redis.StatusCmd {
	if r.client != nil {
		return r.client.CFReserveMaxIterations(ctx, key, capacity, maxiterations)
	}
	return r.clusterClient.CFReserveMaxIterations(ctx, key, capacity, maxiterations)
}

// CFScanDump 扫描并转储 Cuckoo Filter
func (r *RedisClient) CFScanDump(ctx context.Context, key string, iterator int64) *redis.ScanDumpCmd {
	if r.client != nil {
		return r.client.CFScanDump(ctx, key, iterator)
	}
	return r.clusterClient.CFScanDump(ctx, key, iterator)
}

// CFLoadChunk 加载 Cuckoo Filter 数据块
func (r *RedisClient) CFLoadChunk(ctx context.Context, key string, iterator int64, data interface{}) *redis.StatusCmd {
	if r.client != nil {
		return r.client.CFLoadChunk(ctx, key, iterator, data)
	}
	return r.clusterClient.CFLoadChunk(ctx, key, iterator, data)
}

// ============================================================================
// Count-Min Sketch
// ============================================================================

// CMSIncrBy 增加 Count-Min Sketch 计数值
func (r *RedisClient) CMSIncrBy(ctx context.Context, key string, elements ...interface{}) *redis.IntSliceCmd {
	if r.client != nil {
		return r.client.CMSIncrBy(ctx, key, elements...)
	}
	return r.clusterClient.CMSIncrBy(ctx, key, elements...)
}

// CMSInfo 获取 Count-Min Sketch 信息
func (r *RedisClient) CMSInfo(ctx context.Context, key string) *redis.CMSInfoCmd {
	if r.client != nil {
		return r.client.CMSInfo(ctx, key)
	}
	return r.clusterClient.CMSInfo(ctx, key)
}

// CMSInitByDim 按维度初始化 Count-Min Sketch
func (r *RedisClient) CMSInitByDim(ctx context.Context, key string, width, height int64) *redis.StatusCmd {
	if r.client != nil {
		return r.client.CMSInitByDim(ctx, key, width, height)
	}
	return r.clusterClient.CMSInitByDim(ctx, key, width, height)
}

// CMSInitByProb 按概率初始化 Count-Min Sketch
func (r *RedisClient) CMSInitByProb(ctx context.Context, key string, errorRate, probability float64) *redis.StatusCmd {
	if r.client != nil {
		return r.client.CMSInitByProb(ctx, key, errorRate, probability)
	}
	return r.clusterClient.CMSInitByProb(ctx, key, errorRate, probability)
}

// CMSMerge 合并 Count-Min Sketch
func (r *RedisClient) CMSMerge(ctx context.Context, destKey string, sourceKeys ...string) *redis.StatusCmd {
	if r.client != nil {
		return r.client.CMSMerge(ctx, destKey, sourceKeys...)
	}
	return r.clusterClient.CMSMerge(ctx, destKey, sourceKeys...)
}

// CMSMergeWithWeight 加权合并 Count-Min Sketch
func (r *RedisClient) CMSMergeWithWeight(ctx context.Context, destKey string, sourceKeys map[string]int64) *redis.StatusCmd {
	if r.client != nil {
		return r.client.CMSMergeWithWeight(ctx, destKey, sourceKeys)
	}
	return r.clusterClient.CMSMergeWithWeight(ctx, destKey, sourceKeys)
}

// CMSQuery 查询 Count-Min Sketch 计数值
func (r *RedisClient) CMSQuery(ctx context.Context, key string, elements ...interface{}) *redis.IntSliceCmd {
	if r.client != nil {
		return r.client.CMSQuery(ctx, key, elements...)
	}
	return r.clusterClient.CMSQuery(ctx, key, elements...)
}

// ============================================================================
// Top-K
// ============================================================================

// TopKAdd 向 Top-K 添加元素
func (r *RedisClient) TopKAdd(ctx context.Context, key string, elements ...interface{}) *redis.StringSliceCmd {
	if r.client != nil {
		return r.client.TopKAdd(ctx, key, elements...)
	}
	return r.clusterClient.TopKAdd(ctx, key, elements...)
}

// TopKCount 获取 Top-K 元素计数
func (r *RedisClient) TopKCount(ctx context.Context, key string, elements ...interface{}) *redis.IntSliceCmd {
	if r.client != nil {
		return r.client.TopKCount(ctx, key, elements...)
	}
	return r.clusterClient.TopKCount(ctx, key, elements...)
}

// TopKIncrBy 增加 Top-K 元素计数
func (r *RedisClient) TopKIncrBy(ctx context.Context, key string, elements ...interface{}) *redis.StringSliceCmd {
	if r.client != nil {
		return r.client.TopKIncrBy(ctx, key, elements...)
	}
	return r.clusterClient.TopKIncrBy(ctx, key, elements...)
}

// TopKInfo 获取 Top-K 信息
func (r *RedisClient) TopKInfo(ctx context.Context, key string) *redis.TopKInfoCmd {
	if r.client != nil {
		return r.client.TopKInfo(ctx, key)
	}
	return r.clusterClient.TopKInfo(ctx, key)
}

// TopKList 列出 Top-K 元素
func (r *RedisClient) TopKList(ctx context.Context, key string) *redis.StringSliceCmd {
	if r.client != nil {
		return r.client.TopKList(ctx, key)
	}
	return r.clusterClient.TopKList(ctx, key)
}

// TopKListWithCount 列出 Top-K 元素及计数
func (r *RedisClient) TopKListWithCount(ctx context.Context, key string) *redis.MapStringIntCmd {
	if r.client != nil {
		return r.client.TopKListWithCount(ctx, key)
	}
	return r.clusterClient.TopKListWithCount(ctx, key)
}

// TopKQuery 查询元素是否在 Top-K 中
func (r *RedisClient) TopKQuery(ctx context.Context, key string, elements ...interface{}) *redis.BoolSliceCmd {
	if r.client != nil {
		return r.client.TopKQuery(ctx, key, elements...)
	}
	return r.clusterClient.TopKQuery(ctx, key, elements...)
}

// TopKReserve 创建 Top-K
func (r *RedisClient) TopKReserve(ctx context.Context, key string, k int64) *redis.StatusCmd {
	if r.client != nil {
		return r.client.TopKReserve(ctx, key, k)
	}
	return r.clusterClient.TopKReserve(ctx, key, k)
}

// TopKReserveWithOptions 创建 Top-K（带完整参数）
func (r *RedisClient) TopKReserveWithOptions(ctx context.Context, key string, k int64, width, depth int64, decay float64) *redis.StatusCmd {
	if r.client != nil {
		return r.client.TopKReserveWithOptions(ctx, key, k, width, depth, decay)
	}
	return r.clusterClient.TopKReserveWithOptions(ctx, key, k, width, depth, decay)
}

// ============================================================================
// T-Digest
// ============================================================================

// TDigestAdd 向 T-Digest 添加值
func (r *RedisClient) TDigestAdd(ctx context.Context, key string, elements ...float64) *redis.StatusCmd {
	if r.client != nil {
		return r.client.TDigestAdd(ctx, key, elements...)
	}
	return r.clusterClient.TDigestAdd(ctx, key, elements...)
}

// TDigestByRank 按排名获取值
func (r *RedisClient) TDigestByRank(ctx context.Context, key string, rank ...uint64) *redis.FloatSliceCmd {
	if r.client != nil {
		return r.client.TDigestByRank(ctx, key, rank...)
	}
	return r.clusterClient.TDigestByRank(ctx, key, rank...)
}

// TDigestByRevRank 按反向排名获取值
func (r *RedisClient) TDigestByRevRank(ctx context.Context, key string, rank ...uint64) *redis.FloatSliceCmd {
	if r.client != nil {
		return r.client.TDigestByRevRank(ctx, key, rank...)
	}
	return r.clusterClient.TDigestByRevRank(ctx, key, rank...)
}

// TDigestCDF 获取累积分布
func (r *RedisClient) TDigestCDF(ctx context.Context, key string, elements ...float64) *redis.FloatSliceCmd {
	if r.client != nil {
		return r.client.TDigestCDF(ctx, key, elements...)
	}
	return r.clusterClient.TDigestCDF(ctx, key, elements...)
}

// TDigestCreate 创建 T-Digest
func (r *RedisClient) TDigestCreate(ctx context.Context, key string) *redis.StatusCmd {
	if r.client != nil {
		return r.client.TDigestCreate(ctx, key)
	}
	return r.clusterClient.TDigestCreate(ctx, key)
}

// TDigestCreateWithCompression 创建 T-Digest（含压缩）
func (r *RedisClient) TDigestCreateWithCompression(ctx context.Context, key string, compression int64) *redis.StatusCmd {
	if r.client != nil {
		return r.client.TDigestCreateWithCompression(ctx, key, compression)
	}
	return r.clusterClient.TDigestCreateWithCompression(ctx, key, compression)
}

// TDigestInfo 获取 T-Digest 信息
func (r *RedisClient) TDigestInfo(ctx context.Context, key string) *redis.TDigestInfoCmd {
	if r.client != nil {
		return r.client.TDigestInfo(ctx, key)
	}
	return r.clusterClient.TDigestInfo(ctx, key)
}

// TDigestMax 获取 T-Digest 最大值
func (r *RedisClient) TDigestMax(ctx context.Context, key string) *redis.FloatCmd {
	if r.client != nil {
		return r.client.TDigestMax(ctx, key)
	}
	return r.clusterClient.TDigestMax(ctx, key)
}

// TDigestMin 获取 T-Digest 最小值
func (r *RedisClient) TDigestMin(ctx context.Context, key string) *redis.FloatCmd {
	if r.client != nil {
		return r.client.TDigestMin(ctx, key)
	}
	return r.clusterClient.TDigestMin(ctx, key)
}

// TDigestMerge 合并 T-Digest
func (r *RedisClient) TDigestMerge(ctx context.Context, destKey string, options *TDigestMergeOptions, sourceKeys ...string) *redis.StatusCmd {
	if r.client != nil {
		return r.client.TDigestMerge(ctx, destKey, options, sourceKeys...)
	}
	return r.clusterClient.TDigestMerge(ctx, destKey, options, sourceKeys...)
}

// TDigestQuantile 获取分位数
func (r *RedisClient) TDigestQuantile(ctx context.Context, key string, elements ...float64) *redis.FloatSliceCmd {
	if r.client != nil {
		return r.client.TDigestQuantile(ctx, key, elements...)
	}
	return r.clusterClient.TDigestQuantile(ctx, key, elements...)
}

// TDigestRank 获取值的排名
func (r *RedisClient) TDigestRank(ctx context.Context, key string, values ...float64) *redis.IntSliceCmd {
	if r.client != nil {
		return r.client.TDigestRank(ctx, key, values...)
	}
	return r.clusterClient.TDigestRank(ctx, key, values...)
}

// TDigestReset 重置 T-Digest
func (r *RedisClient) TDigestReset(ctx context.Context, key string) *redis.StatusCmd {
	if r.client != nil {
		return r.client.TDigestReset(ctx, key)
	}
	return r.clusterClient.TDigestReset(ctx, key)
}

// TDigestRevRank 获取值的反向排名
func (r *RedisClient) TDigestRevRank(ctx context.Context, key string, values ...float64) *redis.IntSliceCmd {
	if r.client != nil {
		return r.client.TDigestRevRank(ctx, key, values...)
	}
	return r.clusterClient.TDigestRevRank(ctx, key, values...)
}

// TDigestTrimmedMean 获取修剪平均值
func (r *RedisClient) TDigestTrimmedMean(ctx context.Context, key string, lowCutQuantile, highCutQuantile float64) *redis.FloatCmd {
	if r.client != nil {
		return r.client.TDigestTrimmedMean(ctx, key, lowCutQuantile, highCutQuantile)
	}
	return r.clusterClient.TDigestTrimmedMean(ctx, key, lowCutQuantile, highCutQuantile)
}
