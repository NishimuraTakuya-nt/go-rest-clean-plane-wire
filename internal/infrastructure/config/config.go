package config

import (
	"os"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	ServerAddress  string
	AllowedOrigins []string
	JWTSecretKey   string
	RequestTimeout time.Duration
	// その他の設定項目
}

func Load() *Config {
	return &Config{
		ServerAddress:  getEnv("SERVER_ADDRESS", ":8081"),
		AllowedOrigins: getEnvAsSlice("ALLOWED_ORIGINS", []string{"*"}),
		JWTSecretKey:   getEnv("JWT_SECRET_KEY", "jwt-secret"),
		RequestTimeout: getEnvAsDuration("REQUEST_TIMEOUT", 1*time.Second),
		// その他の設定項目の読み込み
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int) int {
	value := getEnv(key, "")
	if value == "" {
		return fallback
	}
	v, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}
	return v
}

func getEnvAsDuration(key string, fallback time.Duration) time.Duration {
	value := getEnv(key, "")
	if value == "" {
		return fallback
	}
	v, err := time.ParseDuration(value)
	if err != nil {
		return fallback
	}
	return v
}

func getEnvAsSlice(key string, fallback []string) []string {
	value := getEnv(key, "")
	if value == "" {
		return fallback
	}
	return strings.Split(value, ",")
}
