package config

import (
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/rs/zerolog"
	"github.com/wesleyfebarretos/challenge-bravo/app/internal/enum"
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

type JWT struct {
	Secret           string
	ExpirationInHour int
}

type LogConfig struct {
	Output     *os.File
	Filename   string
	LogLevel   int
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
}

type RedisConfig struct {
	Password    string
	HostAndPort string
	Host        string
	Port        string
	User        string
}

type Config struct {
	Redis    RedisConfig
	ApiToken string
	AppEnv   string
	Port     string
	DB       DBConfig
	Jwt      JWT
	Log      LogConfig
}

var (
	Envs     Config
	initOnce sync.Once
)

func Init() {
	initOnce.Do(func() {
		Envs = Config{
			ApiToken: getEnv("API_TOKEN", "ToYaaRUiza7cYAMzD+Pk2ha9N2Xn3rwMpuhd2JVEQ/Usdbte6kFaIOoIWm6qXgOXt0qYZo3uHTvecySPo4p5zQ=="),
			AppEnv:   getEnv("APP_ENV", enum.DEVELOPMENT_ENVIROMENT),
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
			Jwt: JWT{
				Secret:           getEnv("API_TOKEN", "ToYaaRUiza7cYAMzD+Pk2ha9N2Xn3rwMpuhd2JVEQ/Usdbte6kFaIOoIWm6qXgOXt0qYZo3uHTvecySPo4p5zQ=="),
				ExpirationInHour: 48,
			},
			Log: LogConfig{
				LogLevel:   getEnvAsInt("LOG_LEVEL", int(zerolog.InfoLevel)),
				Filename:   getEnv("LOG_FILE_NAME", "logs/app.log"),
				MaxSize:    getEnvAsInt("LOG_MAX_SIZE", 5),
				MaxBackups: getEnvAsInt("LOG_MAX_BACKUPS", 5),
				MaxAge:     getEnvAsInt("LOG_MAX_AGE", 30),
				Compress:   getEnvAsBool("LOG_COMPRESS", true),
				Output:     os.Stdout,
			},
			Redis: RedisConfig{
				Password:    getEnv("REDIS_PASSWORD", "root"),
				HostAndPort: fmt.Sprintf("%s:%s", getEnv("REDIS_HOST", "redis"), getEnv("REDIS_PORT", "6379")),
				Host:        getEnv("REDIS_HOST", "localhost"),
				Port:        getEnv("REDIS_PORT", "6379"),
				User:        getEnv("REDIS_USER", ""),
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
