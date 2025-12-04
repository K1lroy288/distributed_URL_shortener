package config

import (
	"os"
	"sync"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	Host          string
	Port          string
	JwtSecret     string
	ShortenerHost string
	ShortenerPort string
	RedisHost     string
	RedisPort     string
}

var (
	instance *Config
	once     sync.Once
)

func GetConfig() *Config {
	if instance == nil {
		return loadConfig()
	}

	return instance
}

func loadConfig() *Config {
	once.Do(func() {
		instance = &Config{
			Host:          os.Getenv("APP_HOST"),
			Port:          os.Getenv("APP_PORT"),
			JwtSecret:     os.Getenv("JWT_SECRET"),
			ShortenerHost: os.Getenv("SHORTENER_HOST"),
			ShortenerPort: os.Getenv("SHORTENER_PORT"),
			RedisHost:     os.Getenv("REDIS_HOST"),
			RedisPort:     os.Getenv("REDIS_PORT"),
		}
	})

	return instance
}
