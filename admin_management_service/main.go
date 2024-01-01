package main

import (
	"admin_management_service/config"
	"admin_management_service/handler/drug_doc_handler"
	"admin_management_service/handler/image_handler"
	"admin_management_service/handler/video_doc_handler"
	"admin_management_service/repositories/drug_doc_repository"
	"admin_management_service/repositories/image_info_repository"
	"admin_management_service/repositories/video_doc_repository"
	"admin_management_service/services/drug_doc_service"
	"admin_management_service/services/video_doc_service"
	"log"
	"path/filepath"

	"github.com/gin-contrib/cors"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gin-gonic/gin"
)

func main() {
	appConfig := config.LoadConfig()

	cfg := elasticsearch.Config{
		Addresses: []string{appConfig.ElasticURL},
	}
	es := config.ConnectES(cfg)

	db := config.ConnectMySQL(appConfig.MysqlDSN)
	defer db.Close()

	config.CreateIndex(es, "video_doc")
	config.CreateIndex(es, "drug_doc")
	config.CreateTable(filepath.Join("scripts", "database.sql"), db)

	imageInfoRepo := image_info_repository.New(db)
	imageHandler := image_handler.New(imageInfoRepo)

	videoDocRepo := video_doc_repository.New(es, "video_doc")
	videoDocService := video_doc_service.New()
	videoDocHandler := video_doc_handler.New(videoDocRepo, videoDocService)

	drugDocRepo := drug_doc_repository.New(es, "drug_doc")
	drugDocService := drug_doc_service.New()
	drugDocHandler := drug_doc_handler.New(drugDocRepo, drugDocService)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * 3600,
	}))

	videoDocGroup := r.Group("/video_doc")
	{
		videoDocGroup.GET("/:doc_id", videoDocHandler.GetVideoDoc)
		videoDocGroup.GET("/search", videoDocHandler.Search)
		videoDocGroup.POST("/", videoDocHandler.AddVideoDoc)
		videoDocGroup.PUT("/", videoDocHandler.UpdateVideoDoc)
		videoDocGroup.DELETE("/:doc_id", videoDocHandler.DeleteVideoDoc)
	}
	drugDocGroup := r.Group("/drug_doc")
	{
		drugDocGroup.GET("/:doc_id", drugDocHandler.GetDrugDoc)
		drugDocGroup.GET("/search", drugDocHandler.Search)
		drugDocGroup.POST("/", drugDocHandler.AddDrugDoc)
		drugDocGroup.PUT("/", drugDocHandler.UpdateDrugDoc)
		drugDocGroup.DELETE("/:doc_id", drugDocHandler.DeleteDrugDoc)
	}
	imageGroup := r.Group("/image")
	{
		imageGroup.GET("/:doc_id", imageHandler.GetImage)
		imageGroup.POST("/:doc_id", imageHandler.UploadImage)
		imageGroup.PUT("/:doc_id", imageHandler.ChangeImage)
		imageGroup.DELETE("/:doc_id", imageHandler.DeleteImage)
	}

	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
