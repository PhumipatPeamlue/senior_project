package main

import (
	"document_service/internal/adapters/handlers"
	"document_service/internal/adapters/repositories"
	"document_service/internal/conf"
	"document_service/internal/core/services"
	"document_service/internal/es"
	"document_service/internal/middlewares"
	"document_service/internal/routes"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var (
	appConfig *conf.AppConfig
	esClient  *elasticsearch.Client
)

func init() {
	appConfig = conf.LoadConfig()
	esClient = es.ConnectEs(appConfig.EsURL)
}

func main() {
	videoDocRepo := repositories.NewVideoDocRepositoryES(esClient, appConfig.VideoDocIndex)
	drugDocRepo := repositories.NewDrugDocRepositoryEs(esClient, appConfig.DrugDocIndex)

	imageStorageService := services.NewImageStorageService(http.Client{})
	videoDocService := services.NewVideoDocService(videoDocRepo, imageStorageService)
	drugDocService := services.NewDrugDocService(drugDocRepo, imageStorageService)

	videoDocHandler := handlers.NewVideoDocHandler(videoDocService)
	drugDocHandler := handlers.NewDrugDocHandler(drugDocService)

	r := gin.Default()

	r.Use(middlewares.Cors())
	r.Use(middlewares.ErrorHandler())

	routes.VideoDoc(r, videoDocHandler)
	routes.DrugDoc(r, drugDocHandler)

	if err := r.Run(fmt.Sprintf(":%s", appConfig.Port)); err != nil {
		log.Fatalln(err)
	}
}
