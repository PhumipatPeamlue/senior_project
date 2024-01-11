package main

import (
	"admin_management_service/config"
	"admin_management_service/handlers/doc_image_handler"
	"admin_management_service/handlers/drug_doc_handler"
	"admin_management_service/handlers/video_doc_handler"
	"admin_management_service/repositories/doc_image_repository"
	"admin_management_service/repositories/drug_doc_repository"
	"admin_management_service/repositories/video_doc_repository"
	"admin_management_service/services/drug_doc_service"
	"admin_management_service/services/video_doc_service"
	"log"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	appConfig := config.LoadConfig()
	cfg := elasticsearch.Config{
		Addresses: []string{appConfig.ElasticURL},
	}
	es := config.InitElasticsearch(cfg)

	db := config.InitMysql(appConfig.MysqlDSN)

	videoDocRepo := video_doc_repository.NewVideoDocRepo(es, "video_doc")
	drugDocRepo := drug_doc_repository.NewDrugDocRepo(es, "drug_doc")
	docImageRepo := doc_image_repository.NewDocImageRepo(db)
	videoDocService := video_doc_service.NewVideoDocService(videoDocRepo, docImageRepo)
	drugDocService := drug_doc_service.NewDrugDocService(drugDocRepo, docImageRepo)
	videoDocHandler := video_doc_handler.NewVideoDocHandler(videoDocService)
	drugDocHandler := drug_doc_handler.NewDrugDocHandler(drugDocService)
	docImageHandler := doc_image_handler.NewDocImageHandler()

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
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
	docImageGroup := r.Group("/image")
	{
		docImageGroup.GET("/:filename", docImageHandler.GetImage)
	}

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
