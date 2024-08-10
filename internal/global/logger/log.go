package logger

import (
	"gin-rush-template/config"
	"log/slog"
	"os"
	"sync"
)

var (
	instance *slog.Logger
	once     sync.Once
)

// Get 获取全局 Logger 实例
func Get() *slog.Logger {
	once.Do(func() {
		var handler slog.Handler
		switch config.Get().Mode {
		case config.ModeDebug:
			handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{AddSource: false})
		case config.ModeRelease:
			handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true})
		}
		instance = slog.New(handler)
	})
	return instance
}

// New 创建一个新的 Logger 实例
func New(module string) *slog.Logger {
	return Get().With("module", module)
}
