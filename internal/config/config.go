package config

import (
	"strings"

	"github.com/spf13/viper"
)

const (
	CONFIG_NAME = "config"
	CONFIG_TYPE = "yaml"
	CONFIG_PATH = "./configs"
)

type Config struct {
	DB    *DB  `yaml:"db"`
	Debug bool `yaml:"debug"`
}

type DB struct {
	DSN    string `yaml:"dsn"`
	DRIVER string `yaml:"driver"`
}

func LoadConfig() (cfg *Config, err error) {
	viper.SetConfigName(CONFIG_NAME)
	viper.SetConfigType(CONFIG_TYPE)
	viper.AddConfigPath(CONFIG_PATH)

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err = viper.UnmarshalExact(&cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
