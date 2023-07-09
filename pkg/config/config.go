package config

import "gomodules.xyz/envconfig"

type DatabaseConfig struct {
	DSN string `envconfig:"DATABASE_DSN"`
}

type Config struct {
	DatabaseConfig DatabaseConfig
	Port string`default:"8080"`
}

func LoadConfig() (*Config, error) {
	var cfg Config
	err := envconfig.Process("", &cfg)
	return &cfg, err
}
