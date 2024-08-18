package config

type Mode string

const (
	ModeDebug   Mode = "debug"
	ModeRelease Mode = "release"
)

// 记得在config.go文件中添加yaml:"Mode"，否则无法解析

type Config struct {
	Host   string `env:"HOST" yaml:"Host"`
	Port   string `env:"PORT" yaml:"Port"`
	Prefix string `env:"PREFIX" yaml:"Prefix"`
	Mode   Mode   `env:"MODE" yaml:"Mode"`
	OTel   OTel   `yaml:"OTel"`
	Mysql  Mysql  `yaml:"Mysql"`
	JWT    JWT    `yaml:"JWT"`
}

type Mysql struct {
	Host     string `env:"MYSQL_HOST" yaml:"Host"`
	Port     string `env:"MYSQL_PORT" yaml:"Port"`
	Username string `env:"USERNAME" yaml:"Username"`
	Password string `env:"PASSWORD" yaml:"Password"`
	DBName   string `env:"DB_NAME" yaml:"DBName"`
}

type JWT struct {
	AccessSecret string `env:"ACCESS_SECRET" yaml:"AccessSecret"`
	AccessExpire int64  `env:"ACCESS_EXPIRE" yaml:"AccessExpire"`
}

type OTel struct {
	Enable      bool   `env:"ENABLE" yaml:"Enable"`
	ServiceName string `env:"SERVICE_NAME" yaml:"ServiceName"`
	Endpoint    string `env:"ENDPOINT" yaml:"Endpoint"`
	AgentHost   string `env:"AGENT_HOST" yaml:"AgentHost"`
	AgentPort   string `env:"AGENT_PORT" yaml:"AgentPort"`
}
