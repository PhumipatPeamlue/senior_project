package conf

import "os"

type AppConfig struct {
	MySQLDsn string
	Port     string
}

func LoadConfig() *AppConfig {
	return &AppConfig{
		MySQLDsn: getEnv("MYSQL_DSN", "root:root@tcp(localhost:3306)/reminder_service_db?parseTime=true"),
		Port:     getEnv("PORT", "8080"),
	}
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
