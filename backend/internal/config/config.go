package config

import (
	"github.com/spf13/viper"

	"strings"
	"sync"
)

type (
	Config struct {
		App *App
		Postgres *Postgres
	}

	App struct {
		Address string
		Port string
	}

	Postgres struct {
		User string
		Password string
		Name string
		Host string
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

		viper.SetDefault("postgres.user", "postgres")
		viper.SetDefault("postgres.name", "postgres")
		viper.SetDefault("postgres.port", "5432")
		viper.SetDefault("postgres.host", "db")
		viper.SetDefault("postgres.password", "")

		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}

		if err := viper.Unmarshal(&configInstance); err != nil {
			panic(err)
		}
	})

	return configInstance
}
