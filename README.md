# RedisClient

Go Redis 客户端封装库，基于 [go-redis/v9](https://github.com/redis/go-redis) v9，同时支持单机和集群模式。

## 版本信息

版本号格式: `v<Go版本>.<go-redis大版本>.<本项目迭代版本>`

| 项目 | 版本                  | 说明                      |
|------|---------------------|-------------------------|
| Go | **1.19**            | 编译运行环境                  |
| go-redis/v9 | **v9.7.3** (v9 大版本) | Redis 客户端底层库            |
| RedisClient | **v0.19.9-0**       | v0.go版本.go-redis版本-项目版本 |

## 环境依赖

| 依赖 | 版本 | 说明 |
|------|------|------|
| Go | **1.19+** | 编译运行环境 |
| go-redis/v9 | **v9.7.3** | Redis 客户端底层库 |

## 安装

```bash
go get github.com/SilentQianyi/RedisClient
```

## 快速开始

```go
package main

import (
    "context"
    "fmt"
    "log"
    "time"

    RedisClient "RedisClient"
)

func main() {
    // 单机模式
    client, err := RedisClient.NewRedisClient(&RedisClient.Config{
        Addr:    "127.0.0.1:6379",
        Password: "password",
    }, log)
    if err != nil {
        panic(err)
    }
    defer client.Close()

    ctx := context.Background()

    // 字符串操作
    client.Set(ctx, "key", "value", 10*time.Second)
    val, _ := client.Get(ctx, "key").Result()
    fmt.Println(val)

    // 集群模式
    clusterClient, _ := RedisClient.NewRedisClient(&RedisClient.Config{
        Addrs: []string{"127.0.0.1:7000", "127.0.0.1:7001", "127.0.0.1:7002"},
    }, nil)
    defer clusterClient.Close()
}
```

## 架构

`RedisClient` 内部持有 `*redis.Client`（单机）或 `*redis.ClusterClient`（集群），所有方法通过 `client != nil` 判断分发到对应底层客户端：

```
                   +----------------+
                   |  RedisClient   |
                   |                |
                   | client *Client |
                   | cluster *Clust |
                   +-------+--------+
                           |
              +------------+------------+
              |                         |
     r.client != nil            r.clusterClient
      (单机模式)                (集群模式)
```

## 配置

```go
type Config struct {
    Addr     string   // 单机地址 (host:port)
    Addrs    []string // 集群地址列表，非空时启用集群模式
    Password string   // 认证密码
    IsTLS    bool     // 是否启用 TLS

    DialTimeout  uint32 // 连接超时 (ms)，默认 5000
    ReadTimeout  uint32 // 读超时 (ms)，默认 5000
    WriteTimeout uint32 // 写超时 (ms)，默认 5000
    MaxRetries   int    // 最大重试次数，默认 1
    PoolSize     int    // 连接池大小，默认 50
    MinIdle      int    // 最小空闲连接数，默认 20
}
```

- `Config` 为 nil 时使用默认值（127.0.0.1:6379）
- `len(Addrs) > 0` 启用集群模式，否则使用 `Addr` 单机模式

## 日志

库内所有日志通过 `Logger` 接口输出，调用方注入自己的日志实现：

```go
type Logger interface {
    Info(format string, args ...interface{})
    Error(format string, args ...interface{})
    Fatal(format string, args ...interface{})
}
```

注入 nil 使用默认空操作实现（静默丢弃日志）。

### 日志格式规范

**Error 日志** — `方法名 包名.函数名 error! param[ %v ]`

```go
r.log.Error("SetCas ScriptSetCas.Run error! key[ %s ], oldval[ %s ], err[ %s ]", key, string(oldval), e.Error())
```

**Info 日志** — `方法名 说明, param[ %v ]`

```go
r.log.Info("GetAndDelValue key not found, key[ %s ]", key)
r.log.Info("NewRedisClient init, cfg[ %+v ]", cfg)
```

参数统一使用 `[ %v ]` 包裹（两端空格）。

## 文件结构

| 文件 | 说明 | 方法数 |
|------|------|--------|
| `redisClient.go` | RedisClient 结构体、NewRedisClient、Close、Client/ClusterClient | 4 |
| `config.go` | Config 配置 | - |
| `errors.go` | 哨兵错误 | - |
| `logger.go` | Logger 接口 + defaultLogger | - |
| `connectionCmd.go` | Ping, Echo, Client*, Config*, Info, DBSize, Flush*, Time, Memory* | ~50 |
| `keysCmd.go` | Del, Exists, Expire*, TTL, Keys, Scan, Rename, Type, Sort, Touch, Unlink, Copy | ~40 |
| `stringsCmd.go` | Set*, Get*, MSet, MGet, Incr*, Decr*, Append, GetRange, SetRange, StrLen, LCS | ~30 |
| `stringsScript.go` | CreateAndReturnValue, SetCas, GetAndDelValue (Lua scripts) | 3 |
| `hashesCmd.go` | HSet*, HGet*, HMGet, HGetAll, HDel, HLen, HIncrBy*, HScan, HExpire*, HPersist | ~40 |
| `listsCmd.go` | LPush*, RPush*, LPop*, RPop*, LLen, LRange, LMove*, BLMove*, BLPop, BRPop, LMPop | ~35 |
| `setsCmd.go` | SAdd, SRem, SCard, SMembers, SIsMember, SInter*, SUnion*, SDiff*, SMove, SScan, SPop* | ~30 |
| `sortedSetsCmd.go` | ZAdd*, ZRem*, ZCard, ZScore, ZIncrBy, ZRange*, ZRevRange*, ZPop*, ZInter*, ZUnion*, ZDiff*, ZMPop, BZMPop | ~40 |
| `sortedSetsScript.go` | ZAddValue (Lua script: 累加分数) | 1 |
| `geoCmd.go` | GeoAdd, GeoPos, GeoSearch*, GeoDist, GeoHash | ~10 |
| `streamsCmd.go` | XAdd, XRead*, XGroup*, XAck, XClaim*, XPending*, XInfo*, XTrim*, XAutoClaim* | ~35 |
| `hyperloglogCmd.go` | PFAdd, PFCount, PFMerge | 3 |
| `bitmapCmd.go` | SetBit, GetBit, BitCount, BitOpAnd/Or/Xor/Not, BitPos, BitField* | ~12 |
| `scriptingCmd.go` | Eval, EvalSha, ScriptLoad, ScriptExists, ScriptFlush, Function*, FCall* | ~20 |
| `pubsubCmd.go` | Publish, SPublish, PubSub* | ~10 |
| `aclCmd.go` | ACLDryRun, ACLLog, ACLLogReset | 3 |
| `clusterCmd.go` | Cluster*, ReadOnly, ReadWrite | ~24 |
| `probabilisticCmd.go` | BF*, CF*, CMS*, TopK*, TDigest* (Bloom/Cuckoo/CountMinSketch/TopK/TDigest) | ~66 |

## 高级方法

### Lua 脚本

| 方法 | 用途 |
|------|------|
| `CreateAndReturnValue` | 键不存在则创建并返回，否则返回已有值 |
| `SetCas` | CAS 比较并设置：仅当旧值匹配才写入新值 |
| `GetAndDelValue` | 原子获取键值并删除 |
| `ZAddValue` | 累加有序集合成员的分数（先获取当前值再累加） |

### 逃逸出口

对于 Pipeline/Transaction/Watch/Subscribe 以及扩展模块（RedisJSON、Search、TimeSeries），通过底层客户端直接操作：

```go
// 管道
pipe := client.Client().Pipeline()

// 扩展模块
jsonVal, _ := client.Client().JSONGet(ctx, "key", "$").Result()
```

## 错误

```go
var (
    ErrSetCasConflict  = errors.New("redisclient: cas conflict")      // CAS 值冲突
    ErrKeyNotFound     = errors.New("redisclient: key not found")     // 键不存在
    ErrSystem          = errors.New("redisclient: system error")      // 系统错误
    ErrNotInitialized  = errors.New("redisclient: client not initialized") // 未初始化
)
```

## 扩展模块

以下 Redis 模块通过 `Client()` / `ClusterClient()` 获取底层客户端直接访问：

- **RedisJSON** — `client.Client().JSONSet(ctx, key, path, value)`
- **RediSearch** — `client.Client().FTSearch(ctx, index, query)`
- **RedisTimeSeries** — `client.Client().TSAdd(ctx, key, value)`
- **RedisBloom** — 概率数据结构已封装在 `probabilisticCmd.go` 中
- **RedisGears** / **RedisGraph** — 通过底层客户端直接调用

## License

MIT
