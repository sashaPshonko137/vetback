package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type Config struct {
	Env       string `yaml:"ENV"`
	SecretKey string `yaml:"SECRET_KEY" env-required:"true"`
	Port      string `yaml:"PORT" env-default:":3000"`
	DBUrl     string `yaml:"DB_URL" env-required:"true"`
	Cost      int    `yaml:"COST" env-required:"true"`
}

func NewConfig() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		panic("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("CONFIG_PATH does not exist")
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic(err)
	}

	return &cfg
}
