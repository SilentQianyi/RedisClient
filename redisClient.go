package redisClient

import (
	"context"
	"crypto/tls"
	"errors"

	"github.com/redis/go-redis/v9"
)

// RedisClient Redis 客户端封装，同时支持单机和集群模式。
// 内部通过判断 client 是否为 nil 来分发到对应的底层客户端。
type RedisClient struct {
	client        *redis.Client
	clusterClient *redis.ClusterClient
	log           Logger
}

// NewRedisClient 根据配置创建 RedisClient 并验证连通性。
func NewRedisClient(cfg *Config, log Logger) (*RedisClient, error) {
	if cfg == nil {
		return nil, errors.New("newRedisClient cfg nil failed")
	}
	cfg = initConfig(cfg)

	if log == nil {
		log = &defaultLogger{}
	}

	c := &RedisClient{log: log}

	if len(cfg.Addrs) > 0 {
		options := &redis.ClusterOptions{
			Addrs:        cfg.Addrs,
			Password:     cfg.Password,
			DialTimeout:  cfg.dialTimeout(),
			ReadTimeout:  cfg.readTimeout(),
			WriteTimeout: cfg.writeTimeout(),
			MaxRetries:   cfg.MaxRetries,
			PoolSize:     cfg.PoolSize,
			MinIdleConns: cfg.MinIdle,
		}
		if cfg.IsTLS {
			options.TLSConfig = &tls.Config{
				MinVersion: tls.VersionTLS12,
			}
		}
		c.clusterClient = redis.NewClusterClient(options)
	} else {
		options := &redis.Options{
			Addr:         cfg.Addr,
			Password:     cfg.Password,
			DialTimeout:  cfg.dialTimeout(),
			ReadTimeout:  cfg.readTimeout(),
			WriteTimeout: cfg.writeTimeout(),
			MaxRetries:   cfg.MaxRetries,
			PoolSize:     cfg.PoolSize,
			MinIdleConns: cfg.MinIdle,
		}
		if cfg.IsTLS {
			options.TLSConfig = &tls.Config{
				MinVersion: tls.VersionTLS12,
			}
		}
		c.client = redis.NewClient(options)
	}

	log.Info("NewRedisClient init, cfg[ %+v ]", cfg)

	if err := c.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}
	log.Info("NewRedisClient ping success")

	return c, nil
}

// Close 关闭 Redis 连接
func (r *RedisClient) Close() error {
	if r.client != nil {
		return r.client.Close()
	}
	if r.clusterClient != nil {
		return r.clusterClient.Close()
	}
	return nil
}

// Client 返回底层 *redis.Client（单机模式），集群模式返回 nil
func (r *RedisClient) Client() *redis.Client {
	return r.client
}

// ClusterClient 返回底层 *redis.ClusterClient（集群模式），单机模式返回 nil
func (r *RedisClient) ClusterClient() *redis.ClusterClient {
	return r.clusterClient
}
