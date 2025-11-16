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
		Auth *Auth
	}

	App struct {
		Address string
		Port string
		Debug bool
	}

	Postgres struct {
		User string
		Password string
		Name string
		Host string
		Port string
	}

	Auth struct {
		Secret string
		SigningMethod string `mapstructure:"signingMethod"`
		AccessExpirationSecs int `mapstructure:"accessExpirationSecs"`
		RefreshExpirationSecs int `mapstructure:"refreshExpirationSecs"`
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

		viper.SetDefault("app.address", "")
		viper.SetDefault("app.port", "8080")
		viper.SetDefault("app.debug", "true")

		viper.SetDefault("postgres.user", "postgres")
		viper.SetDefault("postgres.name", "postgres")
		viper.SetDefault("postgres.port", "5432")
		viper.SetDefault("postgres.host", "db")
		viper.SetDefault("postgres.password", "")

		viper.SetDefault("auth.signingMethod", "HS256")
		viper.SetDefault("auth.accessExpirationSecs", 60*60*24)
		viper.SetDefault("auth.refreshExpirationSecs", 60*60*24*7)

		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}

		if err := viper.Unmarshal(&configInstance); err != nil {
			panic(err)
		}
	})

	return configInstance
}
