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
		VideoDocIndex: getEnv("VIDEO_DOC_INDEX", "video_doc"),
		DrugDocIndex:  getEnv("DRUG_DOC_INDEX", "drug_doc"),
		Port:          getEnv("PORT", "8081"),
	}
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
