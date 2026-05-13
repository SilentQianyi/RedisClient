package redisClient

// Logger 日志接口，调用方可注入自己的日志实现
type Logger interface {
	Info(format string, args ...interface{})
	Error(format string, args ...interface{})
	Fatal(format string, args ...interface{})
}

// defaultLogger 默认日志实现（空操作）
type defaultLogger struct{}

func (d *defaultLogger) Info(format string, args ...interface{})  {}
func (d *defaultLogger) Error(format string, args ...interface{}) {}
func (d *defaultLogger) Fatal(format string, args ...interface{}) {}
