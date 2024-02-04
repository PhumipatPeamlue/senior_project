package conf

import "os"

type AppConfig struct {
	MySQLDsn    string
	Host        string
	Port        string
	StoragePath string
}

func LoadConfig() *AppConfig {
	return &AppConfig{
		MySQLDsn:    getEnv("MYSQL_DSN", "root:root@tcp(localhost:3306)/image_info_db?parseTime=true"),
		Host:        getEnv("HOST", "localhost"),
		Port:        getEnv("PORT", "8080"),
		StoragePath: getEnv("STORAGE_PATH", "uploads"),
	}
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
