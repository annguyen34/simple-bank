package util

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver             string        `mapstructure:"DB_DRIVER" env:"DB_DRIVER"`
	DBSource             string        `mapstructure:"DB_SOURCE" env:"DB_SOURCE"`
	HTTPServerAddress    string        `mapstructure:"HTTP_SERVER_ADDRESS" env:"HTTP_SERVER_ADDRESS"`
	GRPCServerAddress    string        `mapstructure:"GRPC_SERVER_ADDRESS" env:"GRPC_SERVER_ADDRESS"`
	TokenSymmetricKey    string        `mapstructure:"TOKEN_SYMMETRIC_KEY" env:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION" env:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION" env:"REFRESH_TOKEN_DURATION"`
	RedisServerAddress   string        `mapstructure:"REDIS_SERVER_ADDRESS" env:"REDIS_SERVER_ADDRESS"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
