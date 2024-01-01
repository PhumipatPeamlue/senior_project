package config

import "os"

type AppConfig struct {
	ElasticURL string
	MysqlDSN   string
}

func LoadConfig() *AppConfig {
	return &AppConfig{
		ElasticURL: getEnv("ES_URL", "http://localhost:9200"),
		MysqlDSN:   getEnv("MYSQL_DSN", "root:root@tcp(localhost:3306)/senior_project"),
	}
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
