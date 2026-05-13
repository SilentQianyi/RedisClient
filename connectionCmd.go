package redisClient

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

// ============================================================================
// 连接
// ============================================================================

// Ping 测试连接
func (r *RedisClient) Ping(ctx context.Context) *redis.StatusCmd {
	if r.client != nil {
		return r.client.Ping(ctx)
	}
	return r.clusterClient.Ping(ctx)
}

// Echo 回显消息
func (r *RedisClient) Echo(ctx context.Context, message interface{}) *redis.StringCmd {
	if r.client != nil {
		return r.client.Echo(ctx, message)
	}
	return r.clusterClient.Echo(ctx, message)
}

// ============================================================================
// 客户端管理
// ============================================================================

// ClientGetName 获取客户端名称
func (r *RedisClient) ClientGetName(ctx context.Context) *redis.StringCmd {
	if r.client != nil {
		return r.client.ClientGetName(ctx)
	}
	return r.clusterClient.ClientGetName(ctx)
}

// ClientList 列出客户端
func (r *RedisClient) ClientList(ctx context.Context) *redis.StringCmd {
	if r.client != nil {
		return r.client.ClientList(ctx)
	}
	return r.clusterClient.ClientList(ctx)
}

// ClientID 获取客户端 ID
func (r *RedisClient) ClientID(ctx context.Context) *redis.IntCmd {
	if r.client != nil {
		return r.client.ClientID(ctx)
	}
	return r.clusterClient.ClientID(ctx)
}

// ClientKill 关闭指定客户端
func (r *RedisClient) ClientKill(ctx context.Context, ipPort string) *redis.StatusCmd {
	if r.client != nil {
		return r.client.ClientKill(ctx, ipPort)
	}
	return r.clusterClient.ClientKill(ctx, ipPort)
}

// ClientKillByFilter 按条件关闭客户端
func (r *RedisClient) ClientKillByFilter(ctx context.Context, keys ...string) *redis.IntCmd {
	if r.client != nil {
		return r.client.ClientKillByFilter(ctx, keys...)
	}
	return r.clusterClient.ClientKillByFilter(ctx, keys...)
}

// ClientPause 暂停客户端
func (r *RedisClient) ClientPause(ctx context.Context, dur time.Duration) *redis.BoolCmd {
	if r.client != nil {
		return r.client.ClientPause(ctx, dur)
	}
	return r.clusterClient.ClientPause(ctx, dur)
}

// ClientUnpause 恢复客户端
func (r *RedisClient) ClientUnpause(ctx context.Context) *redis.BoolCmd {
	if r.client != nil {
		return r.client.ClientUnpause(ctx)
	}
	return r.clusterClient.ClientUnpause(ctx)
}

// ClientUnblock 取消客户端阻塞
func (r *RedisClient) ClientUnblock(ctx context.Context, id int64) *redis.IntCmd {
	if r.client != nil {
		return r.client.ClientUnblock(ctx, id)
	}
	return r.clusterClient.ClientUnblock(ctx, id)
}

// ClientUnblockWithError 取消客户端阻塞并返回错误
func (r *RedisClient) ClientUnblockWithError(ctx context.Context, id int64) *redis.IntCmd {
	if r.client != nil {
		return r.client.ClientUnblockWithError(ctx, id)
	}
	return r.clusterClient.ClientUnblockWithError(ctx, id)
}

// ClientInfo 获取客户端信息
func (r *RedisClient) ClientInfo(ctx context.Context) *redis.ClientInfoCmd {
	if r.client != nil {
		return r.client.ClientInfo(ctx)
	}
	return r.clusterClient.ClientInfo(ctx)
}

// ============================================================================
// 服务器管理
// ============================================================================

// ConfigGet 获取配置
func (r *RedisClient) ConfigGet(ctx context.Context, parameter string) *redis.MapStringStringCmd {
	if r.client != nil {
		return r.client.ConfigGet(ctx, parameter)
	}
	return r.clusterClient.ConfigGet(ctx, parameter)
}

// ConfigSet 设置配置
func (r *RedisClient) ConfigSet(ctx context.Context, parameter, value string) *redis.StatusCmd {
	if r.client != nil {
		return r.client.ConfigSet(ctx, parameter, value)
	}
	return r.clusterClient.ConfigSet(ctx, parameter, value)
}

// ConfigResetStat 重置统计
func (r *RedisClient) ConfigResetStat(ctx context.Context) *redis.StatusCmd {
	if r.client != nil {
		return r.client.ConfigResetStat(ctx)
	}
	return r.clusterClient.ConfigResetStat(ctx)
}

// ConfigRewrite 重写配置文件
func (r *RedisClient) ConfigRewrite(ctx context.Context) *redis.StatusCmd {
	if r.client != nil {
		return r.client.ConfigRewrite(ctx)
	}
	return r.clusterClient.ConfigRewrite(ctx)
}

// Info 获取服务器信息
func (r *RedisClient) Info(ctx context.Context, sections ...string) *redis.StringCmd {
	if r.client != nil {
		return r.client.Info(ctx, sections...)
	}
	return r.clusterClient.Info(ctx, sections...)
}

// InfoMap 获取服务器信息（map 格式）
func (r *RedisClient) InfoMap(ctx context.Context, sections ...string) *redis.InfoCmd {
	if r.client != nil {
		return r.client.InfoMap(ctx, sections...)
	}
	return r.clusterClient.InfoMap(ctx, sections...)
}

// DBSize 获取数据库键数量
func (r *RedisClient) DBSize(ctx context.Context) *redis.IntCmd {
	if r.client != nil {
		return r.client.DBSize(ctx)
	}
	return r.clusterClient.DBSize(ctx)
}

// FlushAll 清空所有数据库
func (r *RedisClient) FlushAll(ctx context.Context) *redis.StatusCmd {
	if r.client != nil {
		return r.client.FlushAll(ctx)
	}
	return r.clusterClient.FlushAll(ctx)
}

// FlushAllAsync 异步清空所有数据库
func (r *RedisClient) FlushAllAsync(ctx context.Context) *redis.StatusCmd {
	if r.client != nil {
		return r.client.FlushAllAsync(ctx)
	}
	return r.clusterClient.FlushAllAsync(ctx)
}

// FlushDB 清空当前数据库
func (r *RedisClient) FlushDB(ctx context.Context) *redis.StatusCmd {
	if r.client != nil {
		return r.client.FlushDB(ctx)
	}
	return r.clusterClient.FlushDB(ctx)
}

// FlushDBAsync 异步清空当前数据库
func (r *RedisClient) FlushDBAsync(ctx context.Context) *redis.StatusCmd {
	if r.client != nil {
		return r.client.FlushDBAsync(ctx)
	}
	return r.clusterClient.FlushDBAsync(ctx)
}

// Time 获取服务器时间
func (r *RedisClient) Time(ctx context.Context) *redis.TimeCmd {
	if r.client != nil {
		return r.client.Time(ctx)
	}
	return r.clusterClient.Time(ctx)
}

// LastSave 获取最后一次保存时间
func (r *RedisClient) LastSave(ctx context.Context) *redis.IntCmd {
	if r.client != nil {
		return r.client.LastSave(ctx)
	}
	return r.clusterClient.LastSave(ctx)
}

// Save 同步保存数据
func (r *RedisClient) Save(ctx context.Context) *redis.StatusCmd {
	if r.client != nil {
		return r.client.Save(ctx)
	}
	return r.clusterClient.Save(ctx)
}

// BgSave 异步保存数据
func (r *RedisClient) BgSave(ctx context.Context) *redis.StatusCmd {
	if r.client != nil {
		return r.client.BgSave(ctx)
	}
	return r.clusterClient.BgSave(ctx)
}

// BgRewriteAOF 异步重写 AOF
func (r *RedisClient) BgRewriteAOF(ctx context.Context) *redis.StatusCmd {
	if r.client != nil {
		return r.client.BgRewriteAOF(ctx)
	}
	return r.clusterClient.BgRewriteAOF(ctx)
}

// Shutdown 关闭服务器
func (r *RedisClient) Shutdown(ctx context.Context) *redis.StatusCmd {
	if r.client != nil {
		return r.client.Shutdown(ctx)
	}
	return r.clusterClient.Shutdown(ctx)
}

// ShutdownSave 保存并关闭服务器
func (r *RedisClient) ShutdownSave(ctx context.Context) *redis.StatusCmd {
	if r.client != nil {
		return r.client.ShutdownSave(ctx)
	}
	return r.clusterClient.ShutdownSave(ctx)
}

// ShutdownNoSave 不保存关闭服务器
func (r *RedisClient) ShutdownNoSave(ctx context.Context) *redis.StatusCmd {
	if r.client != nil {
		return r.client.ShutdownNoSave(ctx)
	}
	return r.clusterClient.ShutdownNoSave(ctx)
}

// SlowLogGet 获取慢日志
func (r *RedisClient) SlowLogGet(ctx context.Context, num int64) *redis.SlowLogCmd {
	if r.client != nil {
		return r.client.SlowLogGet(ctx, num)
	}
	return r.clusterClient.SlowLogGet(ctx, num)
}

// MemoryUsage 获取键内存占用
func (r *RedisClient) MemoryUsage(ctx context.Context, key string, samples ...int) *redis.IntCmd {
	if r.client != nil {
		return r.client.MemoryUsage(ctx, key, samples...)
	}
	return r.clusterClient.MemoryUsage(ctx, key, samples...)
}

// CommandList 获取命令列表
func (r *RedisClient) CommandList(ctx context.Context, filter *redis.FilterBy) *redis.StringSliceCmd {
	if r.client != nil {
		return r.client.CommandList(ctx, filter)
	}
	return r.clusterClient.CommandList(ctx, filter)
}

// CommandGetKeys 从命令中提取键名
func (r *RedisClient) CommandGetKeys(ctx context.Context, commands ...interface{}) *redis.StringSliceCmd {
	if r.client != nil {
		return r.client.CommandGetKeys(ctx, commands...)
	}
	return r.clusterClient.CommandGetKeys(ctx, commands...)
}

// CommandGetKeysAndFlags 从命令中提取键名及标志
func (r *RedisClient) CommandGetKeysAndFlags(ctx context.Context, commands ...interface{}) *redis.KeyFlagsCmd {
	if r.client != nil {
		return r.client.CommandGetKeysAndFlags(ctx, commands...)
	}
	return r.clusterClient.CommandGetKeysAndFlags(ctx, commands...)
}

// Quit 断开连接
func (r *RedisClient) Quit(ctx context.Context) *redis.StatusCmd {
	if r.client != nil {
		return r.client.Quit(ctx)
	}
	return r.clusterClient.Quit(ctx)
}

// SlaveOf 将当前服务器设为指定服务器的从服务器
func (r *RedisClient) SlaveOf(ctx context.Context, host, port string) *redis.StatusCmd {
	if r.client != nil {
		return r.client.SlaveOf(ctx, host, port)
	}
	return r.clusterClient.SlaveOf(ctx, host, port)
}

// DebugObject 调试键对象
func (r *RedisClient) DebugObject(ctx context.Context, key string) *redis.StringCmd {
	if r.client != nil {
		return r.client.DebugObject(ctx, key)
	}
	return r.clusterClient.DebugObject(ctx, key)
}

// ModuleLoadex 加载模块
func (r *RedisClient) ModuleLoadex(ctx context.Context, conf *redis.ModuleLoadexConfig) *redis.StringCmd {
	if r.client != nil {
		return r.client.ModuleLoadex(ctx, conf)
	}
	return r.clusterClient.ModuleLoadex(ctx, conf)
}

// Command 获取所有命令信息
func (r *RedisClient) Command(ctx context.Context) *redis.CommandsInfoCmd {
	if r.client != nil {
		return r.client.Command(ctx)
	}
	return r.clusterClient.Command(ctx)
}
