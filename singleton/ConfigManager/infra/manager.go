package infra

type Config struct{}

var envs = map[string]string{
	"DB_HOST": "localhost",
	"DB_PORT": "5432",
	"APP_ENV": "development",
}

func GetConfigManager() *Config {
	return &Config{}
}

func (c *Config) Get(key string) string {
	return envs[key]
}
