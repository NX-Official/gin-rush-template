package config

import (
	"fmt"
	"gin-rush-template/tools"
	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/viper"
)

var Path = "config/config.yaml"
var c = Config{}

func Read() {
	viper.SetConfigFile(Path)

	if tools.FileExist(Path) {
		tools.PanicIfErr(
			viper.ReadInConfig(),
			viper.Unmarshal(&c),
		)
	} else {
		fmt.Println("Config file not exist in ", Path, ". Using environment variables.")
		if err := envconfig.Process("", &c); err != nil {
			panic(err)
		}
	}

	//fmt.Printf("%+v\n", Config)
}

func Get() Config {
	return c
}
