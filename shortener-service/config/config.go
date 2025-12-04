package config

import (
	"os"
	"sync"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	Host        string
	Port        string
	JwtSecret   string
	DB          DBConfig
	GatewayPort string
}

type DBConfig struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     string
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
			Host:      os.Getenv("APP_HOST"),
			Port:      os.Getenv("APP_PORT"),
			JwtSecret: os.Getenv("JWT_SECRET"),
			DB: DBConfig{
				Host:     os.Getenv("DB_HOST"),
				User:     os.Getenv("DB_USER"),
				Password: os.Getenv("DB_PASSWORD"),
				Name:     os.Getenv("DB_NAME"),
				Port:     os.Getenv("DB_PORT"),
			},
			GatewayPort: os.Getenv("GATEWAY_PORT"),
		}
	})

	return instance
}
