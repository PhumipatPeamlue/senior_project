package config

import (
	"os"
)

type AppConfig struct {
	EsURL string
}

func LoadConfig() *AppConfig {
	return &AppConfig{
		EsURL: getEnv("ES_URL", "http://localhost:9200"),
	}
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
