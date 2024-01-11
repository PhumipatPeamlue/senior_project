package config

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/indices/create"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

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

func InitElasticsearch(cfg elasticsearch.Config) (es *elasticsearch.TypedClient) {
	es, err := elasticsearch.NewTypedClient(cfg)
	if err != nil {
		log.Fatal(err)
	}

	es.Indices.Create("video_doc").Request(&create.Request{
		Mappings: &types.TypeMapping{
			Properties: map[string]types.Property{
				"title":       types.NewTextProperty(),
				"video_url":   types.NewTextProperty(),
				"description": types.NewTextProperty(),
				"create_at":   types.NewDateProperty(),
				"update_at":   types.NewDateProperty(),
			},
		},
	}).Do(nil)

	es.Indices.Create("drug_doc").Request(&create.Request{
		Mappings: &types.TypeMapping{
			Properties: map[string]types.Property{
				"trade_name":  types.NewTextProperty(),
				"drug_name":   types.NewTextProperty(),
				"description": types.NewTextProperty(),
				"preparation": types.NewTextProperty(),
				"caution":     types.NewTextProperty(),
				"create_at":   types.NewDateProperty(),
				"update_at":   types.NewDateProperty(),
			},
		},
	}).Do(nil)

	return
}

func InitMysql(dsn string) (db *sql.DB) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	script, err := os.ReadFile(filepath.Join("scripts", "database.sql"))
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(string(script))
	if err != nil {
		log.Fatal(err)
	}
	return
}
