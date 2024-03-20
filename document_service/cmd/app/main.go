package main

import (
	"document_service/internal/adapters/http/http_gin"
	"document_service/internal/adapters/repositories"
	"document_service/internal/core"
	"document_service/internal/infrastructures"
	"log"
	"os"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	esClient *elasticsearch.Client
)

func init() {
	godotenv.Load()
	esURL := os.Getenv("ES_URL")
	esClient = infrastructures.ConnectES(esURL)
}

func main() {
	videoDocRepository := repositories.NewVideoDocRepositoryES(esClient, os.Getenv("VIDEO_DOC_INDEX"))
	drugDocRepository := repositories.NewDrugDocRepositoryEs(esClient, os.Getenv("DRUG_DOC_INDEX"))
	docRepository := repositories.NewDocRepositoryES(esClient, os.Getenv("VIDEO_DOC_INDEX"), os.Getenv("DRUG_DOC_INDEX"))

	videoDocService := core.NewVideoDocService(videoDocRepository)
	drugDocService := core.NewDrugDocService(drugDocRepository)
	docService := core.NewDocService(docRepository)

	videoDocHandler := http_gin.NewVideoDocHandler(videoDocService)
	drugDocHandler := http_gin.NewDrugDocHandler(drugDocService)
	docHandler := http_gin.NewDocHandler(docService)

	r := gin.Default()
	r.Use(http_gin.Cors())

	http_gin.DocRoutes(r, docHandler, videoDocHandler, drugDocHandler)

	err := r.Run()
	if err != nil {
		log.Fatal()
	}
}
