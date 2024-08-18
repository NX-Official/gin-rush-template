package config

import (
	"github.com/cristalhq/aconfig"
)

const defaultFilePath = "config.yaml"

var c Config

func Init(path ...string) {
	filePath := defaultFilePath
	if len(path) == 1 {
		filePath = path[0]
	}
	loader := aconfig.LoaderFor(&c, aconfig.Config{
		SkipFlags: true,               // 跳过命令行解析
		SkipEnv:   false,              // 不跳过环境变量解析
		SkipFiles: false,              // 不跳过文件解析
		Files:     []string{filePath}, // 如果想要支持json,toml格式的文件，可以在这里添加
		FileDecoders: map[string]aconfig.FileDecoder{
			".yaml": New(),
		},
		AllowUnknownFields: true, // 允许未知字段
	})

	if err := loader.Load(); err != nil {
		panic(err)
	}

	return
}

func Set(config Config) {
	c = config
}

func Get() Config {
	return c
}

func IsRelease() bool {
	return c.Mode == ModeRelease
}

func IsDebug() bool {
	return c.Mode == ModeDebug
}
