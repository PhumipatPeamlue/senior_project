package main

import (
	"database/sql"
	"file_storage_service/internal/adapters/http_gin"
	"file_storage_service/internal/adapters/repositories"
	"file_storage_service/internal/core"
	"file_storage_service/internal/infrastructures"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	db *sql.DB
)

func init() {
	if err := os.Mkdir("bucket", 0755); err != nil {
		log.Fatal(err)
	}
	godotenv.Load()

	dsn := os.Getenv("FILE_INFO_DB")
	db = infrastructures.ConnectRDB("mysql", dsn)
}

func main() {
	defer db.Close()

	fileInfoRepository := repositories.NewFileInfoRepositorySQL(db)

	localFileStorageService := core.NewLocalFileStorageService(fileInfoRepository)

	localFileStorageHandler := http_gin.NewLocalFileStorageHandler(localFileStorageService)

	r := gin.Default()
	r.Use(http_gin.Cors())

	http_gin.ImageRoutes(r, localFileStorageHandler)

	err := r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
