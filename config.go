package redisClient

import "time"

// Config Redis 客户端配置
type Config struct {
	// Addr 单机模式地址（host:port）
	Addr string
	// Addrs 集群模式地址列表，多个地址自动启用集群模式
	Addrs []string
	// Password 认证密码，无密码留空
	Password string
	// DialTimeout 连接超时（毫秒），默认 5000
	DialTimeout uint32
	// ReadTimeout 读超时（毫秒），默认 3000
	ReadTimeout uint32
	// WriteTimeout 写超时（毫秒），默认 3000
	WriteTimeout uint32
	MaxRetries   int
	// PoolSize 连接池大小，默认使用 go-redis 默认值
	PoolSize int
	// MinIdle 最小空闲连接数
	MinIdle int
	// IsTLS 是否启用 TLS
	IsTLS bool
}

// initConfig 返回默认配置
func initConfig(cfg *Config) *Config {

	dialTimeout := uint32(5000)
	if cfg.DialTimeout > 0 {
		dialTimeout = cfg.DialTimeout
	}
	readTimeout := uint32(5000)
	if cfg.ReadTimeout > 0 {
		readTimeout = cfg.ReadTimeout
	}
	writeTimeout := uint32(5000)
	if cfg.WriteTimeout > 0 {
		writeTimeout = cfg.WriteTimeout
	}
	maxRetries := 1
	if cfg.MaxRetries > 0 {
		maxRetries = cfg.MaxRetries
	}
	poolSize := 50
	if cfg.PoolSize > 0 {
		poolSize = cfg.PoolSize
	}
	minIdle := 20
	if cfg.MinIdle > 0 {
		minIdle = cfg.MinIdle
	}

	return &Config{
		Addr:         cfg.Addr,
		Addrs:        cfg.Addrs,
		Password:     cfg.Password,
		DialTimeout:  dialTimeout,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		MaxRetries:   maxRetries,
		PoolSize:     poolSize,
		MinIdle:      minIdle,
		IsTLS:        cfg.IsTLS,
	}
}

func (c *Config) dialTimeout() time.Duration {
	if c.DialTimeout == 0 {
		return 5 * time.Second
	}
	return time.Duration(c.DialTimeout) * time.Millisecond
}

func (c *Config) readTimeout() time.Duration {
	if c.ReadTimeout == 0 {
		return 3 * time.Second
	}
	return time.Duration(c.ReadTimeout) * time.Millisecond
}

func (c *Config) writeTimeout() time.Duration {
	if c.WriteTimeout == 0 {
		return 3 * time.Second
	}
	return time.Duration(c.WriteTimeout) * time.Millisecond
}
