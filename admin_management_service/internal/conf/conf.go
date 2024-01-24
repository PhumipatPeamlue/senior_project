package conf

import "os"

type AppConfig struct {
	EsURL         string
	MySQLDsn      string
	VideoDocIndex string
	DrugDocIndex  string
	Host          string
	Port          string
	StoragePath   string
}

func LoadConfig() *AppConfig {
	return &AppConfig{
		EsURL:         getEnv("ES_URL", "http://localhost:9200"),
		MySQLDsn:      getEnv("MYSQL_DSN", "root:root@tcp(localhost:3306)/image_info_db?parseTime=true"),
		VideoDocIndex: getEnv("VIDEO_DOC_INDEX", "video_doc"),
		DrugDocIndex:  getEnv("DRUG_DOC_INDEX", "drug_doc"),
		Host:          getEnv("HOST", "localhost"),
		Port:          getEnv("PORT", "8080"),
		StoragePath:   getEnv("STORAGE_PATH", "uploads"),
	}
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
