package config

import (
	"context"
	"fmt"

	"github.com/spf13/viper"
)

const ENV = ".env"

type Config struct{}

func NewConfig(ctx context.Context) (*Config, error) {
	viper.SetConfigFile(ENV)
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error loading configuration file: %v", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("error unmarshalling the config into a struct: %v", err)
	}

	return &cfg, nil
}
