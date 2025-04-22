package infra

import "fmt"

type Config struct{}

var configInstance *Config

var envs = map[string]string{
	"DB_HOST": "localhost",
	"DB_PORT": "5432",
	"APP_ENV": "development",
}

func GetConfigManager() *Config {
	if configInstance == nil {
		fmt.Println("Creating new instance")
		configInstance = &Config{}
	}

	return configInstance
}

func (c *Config) Get(key string) string {
	return envs[key]
}
