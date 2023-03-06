package config

import (
	"strings"

	"github.com/spf13/viper"
)

const (
	configName = "config"
	configType = "yaml"
	configPath = "./configs"
)

type Config struct {
	DB     *DB     `yaml:"db"`
	Server *Server `yaml:"server"`
	Debug  bool    `yaml:"debug"`
}

type DB struct {
	DSN    string `yaml:"dsn"`
	Driver string `yaml:"driver"`
}

type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func LoadConfig() (cfg *Config, err error) {
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	viper.AddConfigPath(configPath)

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
