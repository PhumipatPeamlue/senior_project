package main

import (
	elasticsearch2 "admin_management_service/internal/elasticsearch"
	"admin_management_service/internal/elasticsearch/indices/drug_doc_index"
	"admin_management_service/internal/elasticsearch/indices/video_doc_index"
	"admin_management_service/internal/handler"
	"admin_management_service/internal/middlewares"
	"admin_management_service/internal/mysql"
	"admin_management_service/internal/mysql/repositories/image_files"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

func main() {
	cfg := elasticsearch.Config{
		Addresses: []string{fmt.Sprintf(os.Getenv("ES_URL"))},
	}
	es, err := elasticsearch2.Connect(cfg)
	if err != nil {
		log.Fatal(err.Error())
	}
	videoDocIndex := video_doc_index.New(es, "video_doc")
	if err = videoDocIndex.CreateIndex(); err != nil {
		log.Fatal(err.Error())
	}
	drugDocIndex := drug_doc_index.New(es, "drug_doc")
	if err = drugDocIndex.CreateIndex(); err != nil {
		log.Fatal(err.Error())
	}

	db, err := mysql.Connect(os.Getenv("MYSQL_DSN"))
	if err != nil {
		log.Fatal(err.Error())
	}
	imageFileRepo := image_files.New(db)
	if err = imageFileRepo.CreateTable(); err != nil {
		log.Fatal(err.Error())
	}

	h := handler.New(videoDocIndex, drugDocIndex, imageFileRepo)

	r := gin.Default()
	r.Use(middlewares.Cors())
	videoGroup := r.Group("/video")
	{
		videoGroup.GET("/:id", h.GetVideoDoc())
		videoGroup.GET("/search", h.SearchVideoDoc())
		videoGroup.POST("/", h.InsertVideoDoc())
		videoGroup.PUT("/", h.UpdateVideoDoc())
		videoGroup.DELETE("/:id", h.DeleteVideoDoc())
	}
	drugGroup := r.Group("/drug")
	{
		drugGroup.GET("/:id", h.GetDrugDoc())
		drugGroup.GET("/search", h.SearchDrugDoc())
		drugGroup.POST("/", h.InsertDrugDoc())
		drugGroup.PUT("/", h.UpdateDrugDoc())
		drugGroup.DELETE("/:id", h.DeleteDrugDoc())
	}
	imageGroup := r.Group("/image")
	{
		imageGroup.GET("/paths/:id", h.GetAllImagePaths())
		imageGroup.GET("/:filename", h.GetImage())
		imageGroup.POST("/:id", h.InsertImage())
		imageGroup.DELETE("/:id", h.DeleteImage())
	}

	if err = r.Run(":8080"); err != nil {
		log.Fatal(err.Error())
	}
}
