package config

import (
	"os"
	"strings"

	"github.com/anjude/backend-beanflow/infrastructure/enum"
)

type Config struct {
	App          App          `mapstructure:"app"`
	DbConfig     DbConfig     `mapstructure:"db_config"`
	RedisConfig  RedisConfig  `mapstructure:"redis_config"`
	LoggerConfig LoggerConfig `mapstructure:"logger_config"`
}

type (
	App struct {
		Port      string `mapstructure:"port"`
		JwtKey    string `mapstructure:"jwt_key"`
		AppId     string `mapstructure:"app_id"`
		AppSecret string `mapstructure:"app_secret"`
	}

	DbConfig struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Address  string `mapstructure:"address"`
		Database string `mapstructure:"database"`
	}

	RedisConfig struct {
		Enable   bool   `mapstructure:"enable"`
		Addr     string `mapstructure:"addr"`
		Password string `mapstructure:"password"`
		DB       int    `mapstructure:"db"`
		PoolSize int    `mapstructure:"pool_size"`
	}
	LoggerConfig struct {
		Level string `mapstructure:"level"`
	}
)

func (c *Config) GetJwtKey() []byte {
	return []byte(c.App.JwtKey)
}

func GetEnv() enum.Env {
	env := os.Getenv(enum.EkEnv.ToString())
	if env == "" {
		env = os.Getenv(strings.ToLower(enum.EkEnv.ToString()))
	}
	if env == "" {
		env = enum.DEV.ToString()
	}
	return enum.Env(env)
}
