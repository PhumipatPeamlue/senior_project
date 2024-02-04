package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"image_storage_service/internal/adapters/handlers"
	"image_storage_service/internal/adapters/repositories"
	"image_storage_service/internal/conf"
	"image_storage_service/internal/core/services"
	"image_storage_service/internal/middlewares"
	"image_storage_service/internal/mysql"
	"image_storage_service/internal/routes"
	"log"
	"os"
)

var (
	appConfig *conf.AppConfig
	db        *sql.DB
)

func init() {
	appConfig = conf.LoadConfig()
	if _, err := os.Stat(appConfig.StoragePath); os.IsNotExist(err) {
		if err = os.Mkdir(appConfig.StoragePath, 0775); err != nil {
			log.Fatal(err)
		}
	}
	db = mysql.ConnectMySQL(appConfig.MySQLDsn)
}

func main() {
	imageInfoRepo := repositories.NewImageInfoRepositoryMySQL(db)

	imageStorageService := services.NewImageStorageService(imageInfoRepo, appConfig.Host, appConfig.Port, appConfig.StoragePath)

	imageHandler := handlers.NewImageHandler(imageStorageService, appConfig.StoragePath)

	r := gin.Default()

	r.Use(middlewares.Cors())
	r.Use(middlewares.ErrorHandler())

	routes.Image(r, imageHandler)

	if err := r.Run(fmt.Sprintf(":%s", appConfig.Port)); err != nil {
		log.Fatalln(err)
	}
}
