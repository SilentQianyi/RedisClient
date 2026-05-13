package redisClient

import "errors"

var (
	// ErrSetCasConflict CAS 比较失败，值已被其他客户端修改
	ErrSetCasConflict = errors.New("redisclient: cas conflict")
	// ErrKeyNotFound 键不存在
	ErrKeyNotFound = errors.New("redisclient: key not found")
	// ErrSystem 系统错误
	ErrSystem = errors.New("redisclient: system error")
	// ErrNotInitialized 客户端未初始化
	ErrNotInitialized = errors.New("redisclient: client not initialized")
)
