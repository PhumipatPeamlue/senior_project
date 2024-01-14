package main

import (
	"log"
	"os"
	"path/filepath"
	"senior_project/admin_management_service/config"
	"senior_project/admin_management_service/handlers/drug_doc_handler"
	"senior_project/admin_management_service/handlers/image_handler"
	"senior_project/admin_management_service/handlers/video_doc_handler"
	"senior_project/admin_management_service/repositories/drug_doc_repository"
	"senior_project/admin_management_service/repositories/image_storage_repository"
	"senior_project/admin_management_service/repositories/video_doc_repository"
	"senior_project/admin_management_service/services/drug_doc_service"
	"senior_project/admin_management_service/services/video_doc_service"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	es *elasticsearch.Client
)

func init() {
	dirPath := filepath.Join(".", "uploads")
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			log.Println("Error creating directory:", err)
			return
		}
		log.Println("Directory created:", dirPath)
	} else {
		log.Println("Directory already exists:", dirPath)
	}

	appConfig := config.LoadConfig()
	cfg := elasticsearch.Config{
		Addresses: []string{appConfig.EsURL},
	}
	es = config.InitElasticsearch(cfg)
}

func main() {
	imageStorageRepo := image_storage_repository.NewLocalImageStorageRepository("uploads")
	videoDocRepo := video_doc_repository.NewEsVideoDocRepository(es, "video_doc")
	drugDocRepo := drug_doc_repository.NewEsDrugDocRepository(es, "drug_doc")

	videoDocService := video_doc_service.New(videoDocRepo, imageStorageRepo)
	drugDocService := drug_doc_service.New(drugDocRepo, imageStorageRepo)

	videoDocHandler := video_doc_handler.NewGinVideoDocHandler(videoDocService)
	drugDocHandler := drug_doc_handler.NewGinDrugDocHandler(drugDocService)
	imageHandler := image_handler.NewGinImageHandler()

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	videoDocGroup := r.Group("/video_doc")
	{
		videoDocGroup.GET("/:doc_id", videoDocHandler.HandleGetVideoDoc)
		videoDocGroup.GET("/search", videoDocHandler.HandleSearchVideoDoc)
		videoDocGroup.POST("/", videoDocHandler.HandleCreateVideoDoc)
		videoDocGroup.PUT("/", videoDocHandler.HandleUpdateVideoDoc)
		videoDocGroup.DELETE("/:doc_id", videoDocHandler.HandleDeleteVideoDoc)
	}
	drugDocGroup := r.Group("/drug_doc")
	{
		drugDocGroup.GET("/:doc_id", drugDocHandler.HandleGetDrugDoc)
		drugDocGroup.GET("/search", drugDocHandler.HandleSearchDrugDoc)
		drugDocGroup.POST("/", drugDocHandler.HandleCreateDrugDoc)
		drugDocGroup.PUT("/", drugDocHandler.HandleUpdateDrugDoc)
		drugDocGroup.DELETE("/:doc_id", drugDocHandler.HandleDeleteDrugDoc)
	}
	imageGroup := r.Group("/image")
	{
		imageGroup.GET("/:image_name", imageHandler.HandlerGetImage)
	}

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
