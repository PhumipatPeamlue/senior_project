package main

import (
	"database/sql"
	"document_service/internal/adapter/handlers"
	"document_service/internal/adapter/repositories"
	"document_service/internal/conf"
	"document_service/internal/core/services"
	elasticsearchclient "document_service/internal/elasticsearch"
	"document_service/internal/middlewares"
	"document_service/internal/mysql"
	"document_service/internal/routes"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

var (
	appConfig *conf.AppConfig
	es        *elasticsearch.Client
	db        *sql.DB
)

func init() {
	appConfig = conf.LoadConfig()
	if _, err := os.Stat(appConfig.StoragePath); os.IsNotExist(err) {
		if err = os.Mkdir(appConfig.StoragePath, 0775); err != nil {
			log.Fatal(err)
		}
	}
	es = elasticsearchclient.ConnectEs(appConfig.EsURL)
	db = mysql.ConnectMySQL(appConfig.MySQLDsn)
}

func main() {
	imageInfoRepo := repositories.NewImageInfoMySQL(db)
	videoDocRepo := repositories.NewVideoDocES(es, appConfig.VideoDocIndex)
	drugDocRepo := repositories.NewDrugDocES(es, appConfig.DrugDocIndex)

	imageStorageService := services.NewImageStorageFileSystem(imageInfoRepo, appConfig.Host, appConfig.Port, appConfig.StoragePath)
	videoDocService := services.NewVideoDocService(videoDocRepo, imageStorageService)
	drugDocService := services.NewDrugDocService(drugDocRepo, imageStorageService)

	videoDocHandler := handlers.NewVideoDocHandler(videoDocService)
	drugDocHandler := handlers.NewDrugDocHandler(drugDocService)
	imageHandler := handlers.NewImageHandler(appConfig.StoragePath)

	r := gin.Default()

	r.Use(middlewares.Cors())
	r.Use(middlewares.HandleServiceError())

	routes.VideoDoc(r, videoDocHandler)
	routes.DrugDoc(r, drugDocHandler)
	routes.Image(r, imageHandler)

	if err := r.Run(fmt.Sprintf(":%s", appConfig.Port)); err != nil {
		log.Fatalln(err)
	}
}
