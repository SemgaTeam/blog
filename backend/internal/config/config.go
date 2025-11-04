package config

import (
	"github.com/spf13/viper"

	"strings"
	"sync"
)

type (
	Config struct {
		App *App
	}

	App struct {
		Address string
		Port string
	}
)

var (
	once sync.Once
	configInstance *Config
)

func GetConfig() *Config {
	once.Do(func () {
		viper.AddConfigPath("./")
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")

		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		viper.SetDefault("app.address", "127.0.0.1")
		viper.SetDefault("app.port", "8080")

		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}

		if err := viper.Unmarshal(&configInstance); err != nil {
			panic(err)
		}
	})

	return configInstance
}
