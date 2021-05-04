package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ServiceName string `envconfig:"SERVICE_NAME" default:"book-rest-api"`
	Environment string `envconfig:"ENV" default:"dev"`
	Port        int    `envconfig:"PORT" default:"8080" required:"true"`

	DBHost         string `envconfig:"DB_HOST" default:"localhost"`
	DBPort         string `envconfig:"DB_PORT" default:"33062"`
	DBUserName     string `envconfig:"DB_USERNAME" default:"root"`
	DBPassword     string `envconfig:"DB_PASSWORD" default:"password"`
	DBDatabaseName string `envconfig:"DB_DBNAME" default:"gca"`
	DBLogMode      int    `envconfig:"DB_LOG_MODE" default:"3"`

	RedisHost string `envconfig:"REDIS_HOST" default:"localhost"`
	RedisPort string `envconfig:"REDIS_PORT" default:"33792"`
}

func New() Config {
	cfg := Config{}
	envconfig.MustProcess("", &cfg)
	cfg.DBPassword = "password"
	return cfg
}
