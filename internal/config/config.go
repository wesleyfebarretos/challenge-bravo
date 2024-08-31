package config

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

type DBConfig struct {
	User        string
	Port        string
	Password    string
	Address     string
	Host        string
	Name        string
	Driver      string
	PoolMaxConn int
}

type Config struct {
	ApiToken string
	AppEnv   string
	Port     string
	DB       DBConfig
}

var (
	Envs     Config
	initOnce sync.Once
)

func Init() {
	initOnce.Do(func() {
		Envs = Config{
			ApiToken: getEnv("API_TOKEN", "ToYaaRUiza7cYAMzD+Pk2ha9N2Xn3rwMpuhd2JVEQ/Usdbte6kFaIOoIWm6qXgOXt0qYZo3uHTvecySPo4p5zQ=="),
			AppEnv:   getEnv("APP_ENV", "development"),
			Port:     getEnv("PORT", "8080"),
			DB: DBConfig{
				Driver:      getEnv("DB_DRIVER", "postgres"),
				User:        getEnv("DB_USER", "root"),
				Password:    getEnv("DB_PASSWORD", "root"),
				Port:        getEnv("DB_PORT", "5432"),
				Name:        getEnv("DB_NAME", "challenge_bravo"),
				Host:        getEnv("DB_HOST", "challenge-bravo"),
				Address:     fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "5432")),
				PoolMaxConn: getEnvAsInt("DB_POOL_MAX_CONNECTION", 10),
			},
		}
	})
}

func getEnv(key string, callback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return callback
}

func getEnvAsInt64(key string, callback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return callback
		}

		return i
	}
	return callback
}

func getEnvAsInt(key string, callback int) int {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.Atoi(value)
		if err != nil {
			return callback
		}

		return i
	}
	return callback
}

func getEnvAsBool(key string, callback bool) bool {
	if value, ok := os.LookupEnv(key); ok {
		b, err := strconv.ParseBool(value)
		if err != nil {
			return callback
		}
		return b
	}
	return callback
}
