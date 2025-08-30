package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Database Database
	Server   Server
	JWT      JWT
	Env      string
}

type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

type Server struct {
	Host string
	Port string
}

type JWT struct {
	Secret      string
	ExpireHours int
}

func LoadConfig() (*Config, error) {
	// 加载.env文件
	if err := godotenv.Load(); err != nil {
		// 如果没有.env文件，继续使用环境变量
	}

	expireHours, _ := strconv.Atoi(getEnv("JWT_EXPIRE_HOURS", "24"))

	config := &Config{
		Database: Database{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "3306"),
			User:     getEnv("DB_USER", "root"),
			Password: getEnv("DB_PASSWORD", ""),
			DBName:   getEnv("DB_NAME", "littlecode"),
		},
		Server: Server{
			Host: getEnv("SERVER_HOST", "localhost"),
			Port: getEnv("SERVER_PORT", "8080"),
		},
		JWT: JWT{
			Secret:      getEnv("JWT_SECRET", "default_secret"),
			ExpireHours: expireHours,
		},
		Env: getEnv("ENV", "development"),
	}

	return config, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
