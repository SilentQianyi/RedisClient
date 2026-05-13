package redisClient

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// Eval 执行 Lua 脚本
func (r *RedisClient) Eval(ctx context.Context, script string, keys []string, args ...interface{}) *redis.Cmd {
	if r.client != nil {
		return r.client.Eval(ctx, script, keys, args...)
	}
	return r.clusterClient.Eval(ctx, script, keys, args...)
}

// EvalSha 按 SHA1 执行已缓存的 Lua 脚本
func (r *RedisClient) EvalSha(ctx context.Context, sha1 string, keys []string, args ...interface{}) *redis.Cmd {
	if r.client != nil {
		return r.client.EvalSha(ctx, sha1, keys, args...)
	}
	return r.clusterClient.EvalSha(ctx, sha1, keys, args...)
}

// EvalRO 只读执行 Lua 脚本
func (r *RedisClient) EvalRO(ctx context.Context, script string, keys []string, args ...interface{}) *redis.Cmd {
	if r.client != nil {
		return r.client.EvalRO(ctx, script, keys, args...)
	}
	return r.clusterClient.EvalRO(ctx, script, keys, args...)
}

// EvalShaRO 按 SHA1 只读执行已缓存的 Lua 脚本
func (r *RedisClient) EvalShaRO(ctx context.Context, sha1 string, keys []string, args ...interface{}) *redis.Cmd {
	if r.client != nil {
		return r.client.EvalShaRO(ctx, sha1, keys, args...)
	}
	return r.clusterClient.EvalShaRO(ctx, sha1, keys, args...)
}

// ScriptExists 检查脚本是否已缓存
func (r *RedisClient) ScriptExists(ctx context.Context, hashes ...string) *redis.BoolSliceCmd {
	if r.client != nil {
		return r.client.ScriptExists(ctx, hashes...)
	}
	return r.clusterClient.ScriptExists(ctx, hashes...)
}

// ScriptFlush 清空脚本缓存
func (r *RedisClient) ScriptFlush(ctx context.Context) *redis.StatusCmd {
	if r.client != nil {
		return r.client.ScriptFlush(ctx)
	}
	return r.clusterClient.ScriptFlush(ctx)
}

// ScriptKill 终止当前运行的脚本
func (r *RedisClient) ScriptKill(ctx context.Context) *redis.StatusCmd {
	if r.client != nil {
		return r.client.ScriptKill(ctx)
	}
	return r.clusterClient.ScriptKill(ctx)
}

// ScriptLoad 加载脚本到缓存
func (r *RedisClient) ScriptLoad(ctx context.Context, script string) *redis.StringCmd {
	if r.client != nil {
		return r.client.ScriptLoad(ctx, script)
	}
	return r.clusterClient.ScriptLoad(ctx, script)
}

// FunctionLoad 加载函数库
func (r *RedisClient) FunctionLoad(ctx context.Context, code string) *redis.StringCmd {
	if r.client != nil {
		return r.client.FunctionLoad(ctx, code)
	}
	return r.clusterClient.FunctionLoad(ctx, code)
}

// FunctionLoadReplace 替换函数库
func (r *RedisClient) FunctionLoadReplace(ctx context.Context, code string) *redis.StringCmd {
	if r.client != nil {
		return r.client.FunctionLoadReplace(ctx, code)
	}
	return r.clusterClient.FunctionLoadReplace(ctx, code)
}

// FunctionDelete 删除函数库
func (r *RedisClient) FunctionDelete(ctx context.Context, libName string) *redis.StringCmd {
	if r.client != nil {
		return r.client.FunctionDelete(ctx, libName)
	}
	return r.clusterClient.FunctionDelete(ctx, libName)
}

// FunctionFlush 清空所有函数库
func (r *RedisClient) FunctionFlush(ctx context.Context) *redis.StringCmd {
	if r.client != nil {
		return r.client.FunctionFlush(ctx)
	}
	return r.clusterClient.FunctionFlush(ctx)
}

// FunctionKill 终止正在运行的函数
func (r *RedisClient) FunctionKill(ctx context.Context) *redis.StringCmd {
	if r.client != nil {
		return r.client.FunctionKill(ctx)
	}
	return r.clusterClient.FunctionKill(ctx)
}

// FunctionFlushAsync 异步清空所有函数库
func (r *RedisClient) FunctionFlushAsync(ctx context.Context) *redis.StringCmd {
	if r.client != nil {
		return r.client.FunctionFlushAsync(ctx)
	}
	return r.clusterClient.FunctionFlushAsync(ctx)
}

// FunctionList 列出函数库
func (r *RedisClient) FunctionList(ctx context.Context, q FunctionListQuery) *redis.FunctionListCmd {
	if r.client != nil {
		return r.client.FunctionList(ctx, q)
	}
	return r.clusterClient.FunctionList(ctx, q)
}

// FunctionDump 导出函数库
func (r *RedisClient) FunctionDump(ctx context.Context) *redis.StringCmd {
	if r.client != nil {
		return r.client.FunctionDump(ctx)
	}
	return r.clusterClient.FunctionDump(ctx)
}

// FunctionRestore 恢复函数库
func (r *RedisClient) FunctionRestore(ctx context.Context, libDump string) *redis.StringCmd {
	if r.client != nil {
		return r.client.FunctionRestore(ctx, libDump)
	}
	return r.clusterClient.FunctionRestore(ctx, libDump)
}

// FunctionStats 获取函数统计信息
func (r *RedisClient) FunctionStats(ctx context.Context) *redis.FunctionStatsCmd {
	if r.client != nil {
		return r.client.FunctionStats(ctx)
	}
	return r.clusterClient.FunctionStats(ctx)
}

// FCall 调用函数
func (r *RedisClient) FCall(ctx context.Context, function string, keys []string, args ...interface{}) *redis.Cmd {
	if r.client != nil {
		return r.client.FCall(ctx, function, keys, args...)
	}
	return r.clusterClient.FCall(ctx, function, keys, args...)
}

// FCallRo 只读调用函数
func (r *RedisClient) FCallRo(ctx context.Context, function string, keys []string, args ...interface{}) *redis.Cmd {
	if r.client != nil {
		return r.client.FCallRo(ctx, function, keys, args...)
	}
	return r.clusterClient.FCallRo(ctx, function, keys, args...)
}

// FCallRO 只读调用函数 (别名)
func (r *RedisClient) FCallRO(ctx context.Context, function string, keys []string, args ...interface{}) *redis.Cmd {
	if r.client != nil {
		return r.client.FCallRO(ctx, function, keys, args...)
	}
	return r.clusterClient.FCallRO(ctx, function, keys, args...)
}
