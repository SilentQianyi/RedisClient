# redisclient

Go Redis 客户端封装库，基于 go-redis/v9，同时支持单机和集群模式。

## Build & Test

```bash
go build ./...    # 编译
go test ./...     # 测试
go vet ./...      # 静态分析
```

## Project Structure

```
redisclient/
├── redisclient.go     # RedisClient struct, NewRedisClient, Close, Client(), ClusterClient()
├── config.go          # Config + DefaultConfig()
├── errors.go          # 哨兵错误
├── logger.go          # Logger 接口
├── connectionCmd.go   # Ping, Echo, Client*, Config*, Info, DBSize, Flush*, Time 等
├── keysCmd.go         # Del, Exists, Expire*, TTL, Keys, Scan, Rename, Type 等
├── stringsCmd.go      # Set*, Get*, MSet, MGet, Incr*, Decr*, Append, StrLen 等
├── hashesCmd.go       # HSet*, HGet*, HMGet, HGetAll, HDel, HLen, HScan 等
├── listsCmd.go        # LPush*, RPush*, LPop*, RPop*, LMove*, BLPop, BRPop 等
├── setsCmd.go         # SAdd, SRem, SCard, SMembers, SIsMember, SInter*, SUnion*, SDiff* 等
├── sortedSetsCmd.go   # ZAdd*, ZRem*, ZRange*, ZRevRange*, ZScore, ZIncrBy, ZPop* 等
├── geoCmd.go          # GeoAdd, GeoPos, GeoSearch*, GeoDist, GeoHash
├── streamsCmd.go      # XAdd, XRead*, XGroup*, XAck, XClaim*, XPending*, XInfo*, XTrim* 等
├── hyperloglogCmd.go  # PFAdd, PFCount, PFMerge
├── bitmapCmd.go       # SetBit, GetBit, BitCount, BitOp*, BitPos, BitField
├── scriptingCmd.go    # Eval, EvalSha, ScriptLoad, ScriptExists, ScriptFlush, Function*, FCall* 等
├── stringsScript.go   # ScriptSetCas/SetCas, ScriptCreateAndReturnValue/CreateAndReturnValue, ScriptGetAndDel/GetAndDelValue
├── sortedSetsScript.go # ScriptZAddValue, ZAddValue
├── pubsubCmd.go       # Publish, SPublish, PubSub*
├── aclCmd.go          # ACLDryRun, ACLLog, ACLLogReset, ACLSetUser, ACLDelUser, ACLList 等
├── clusterCmd.go      # ClusterMyID, ClusterSlots, ClusterNodes, ClusterInfo, ReadOnly, ReadWrite 等
├── probabilisticCmd.go # BF*, CF*, CMS*, TopK*, TDigest* (Bloom/Cuckoo/CountMinSketch/TopK/TDigest)
├── go.mod / go.sum
└── CLAUDE.md
```

## Architecture

`RedisClient` 内部持有 `*redis.Client`（单机）或 `*redis.ClusterClient`（集群），每个方法根据 `client != nil` 判断并分发到对应底层客户端。所有方法返回 go-redis 原始 Cmd 类型。

- **构造函数**: `NewRedisClient(cfg *Config, log Logger)` — nil config 使用默认值，nil logger 使用空操作实现
- **模式判断**: `len(cfg.Addrs) > 0` 启用集群模式，否则使用 `cfg.Addr` 单机模式
- **逃逸出口**: `Client()` / `ClusterClient()` 可获取底层客户端，用于 Pipeline/Transaction 等场景

## Logging

调用方实现 `Logger` 接口（Info/Error/Fatal）注入。未注入时使用 `defaultLogger`（静默丢弃日志）。

### 日志格式规范

**Error 日志**：`方法名 包名.函数名 error! param1[ %v ], param2[ %v ]`

```go
r.log.Error("SetCas ScriptSetCas.Run error! key[ %s ], oldval[ %s ], err[ %s ]", key, string(oldval), e.Error())
```

**Info 日志**：`方法名 说明, param1[ %v ], param2[ %v ]`

```go
r.log.Info("GetAndDelValue key not found, key[ %s ]", key)
```

**规则**：
- 参数使用 `[ %v ]` 包裹（两端空格）
- Error 日志必须包含报错的完整函数名（含包名）
- 尽量包含所有可用的函数参数和错误信息

## Key Design Decisions

- 方法返回 go-redis 原始 Cmd 类型，调用方直接调用 `.Err()` / `.Result()` / `.Val()` 获取结果
- Lua 脚本使用 `redis.NewScript().Run(ctx, client, keys, args...)`，自动处理 SCRIPT LOAD/EVALSHA 回退
- `StatefulCmdable` 方法（Select, SwapDB, ClientSetName 等）不纳入封装，因其仅在单边可用
- 扩展模块（RedisJSON, Search, TimeSeries 等）通过 `Client()`/`ClusterClient()` 直接访问
