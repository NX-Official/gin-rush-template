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

// New 创建一个新的 Logger 实例
func New(module string) *slog.Logger {
	once.Do(func() {
		var handler slog.Handler
		switch config.Get().Mode {
		case config.DebugMode:
			handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{AddSource: true})
		case config.ReleaseMode:
			handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true})
		}
		instance = slog.New(handler)
	})
	return instance.With("module", module)
}
