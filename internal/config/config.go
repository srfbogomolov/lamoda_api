package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

const (
	CONFIG_FILE = "config.yaml"
	CONFIG_PATH = "."
)

type Config struct {
	Postgres *PostgresConfig `yaml:"postgres"`
	Debug    bool            `yaml:"debug"`
}

type PostgresConfig struct {
	URL string `yaml:"url"`
}

func LoadConfig() (cfg *Config, err error) {
	viper.SetConfigFile(CONFIG_FILE)
	viper.AddConfigPath(CONFIG_PATH)

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error loading config: %w", err)
	}

	if err = viper.UnmarshalExact(&cfg); err != nil {
		return nil, fmt.Errorf("error unmarshalling config into the structure: %w", err)
	}

	return cfg, nil
}
