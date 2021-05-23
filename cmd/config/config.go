package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBLogin string `mapstructure:"DB_LOGIN"`
	DBPass  string `mapstructure:"DB_PASS"`
	BaseURL string `mapstructure:"BASE_URL"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.SetConfigType("env")
	viper.SetConfigName(".env")
	viper.AddConfigPath(path)

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
