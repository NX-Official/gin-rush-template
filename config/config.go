package config

type Config struct {
	Host   string `envconfig:"HOST"`
	Port   string `envconfig:"PORT"`
	Prefix string `envconfig:"PREFIX"`
	Mysql  Mysql
	JWT    JWT
}

type Mysql struct {
	Host     string `envconfig:"HOST"`
	Port     string `envconfig:"PORT"`
	Username string `envconfig:"USERNAME"`
	Password string `envconfig:"PASSWORD"`
	DBName   string `envconfig:"DB_NAME"`
}

type JWT struct {
	AccessSecret string `envconfig:"ACCESS_SECRET"`
	AccessExpire int64  `envconfig:"ACCESS_EXPIRE"`
}
