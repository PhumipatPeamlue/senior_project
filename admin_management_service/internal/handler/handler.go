package handler

import (
	"admin_management_service/internal/elasticsearch/indices/drug_doc_index"
	"admin_management_service/internal/elasticsearch/indices/video_doc_index"
	"admin_management_service/internal/mysql/repositories/image_files"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var (
	matchAllQuery = `
	{
		"from": %d,
		"size": %d,
		"query": {
			"match_all": {} 
		} 
	}
	`
	searchQueryVideoDoc = `
	{
		"from": %d,
		"size": %d,
		"query": {
		  "multi_match": {
			"query": "%s",
			"fields": ["title", "description"]
		  }
		}
	}
	`
	searchQueryDrugDoc = `
	{
		"from": %d,
		"size": %d,
		"query": {
		  "multi_match": {
			"query": "%s",
			"fields": ["trade_name", "drug_name", "description"]
		  }
		}
	}
	`
)

type Handler struct {
	videoDocIndex video_doc_index.VideoDocIndexInterface
	drugDocIndex  drug_doc_index.DrugDocIndexInterface
	imageFileRepo image_files.ImageFilesRepoInterface
}

func New(video video_doc_index.VideoDocIndexInterface, drug drug_doc_index.DrugDocIndexInterface, imageFiles image_files.ImageFilesRepoInterface) *Handler {
	return &Handler{
		videoDocIndex: video,
		drugDocIndex:  drug,
		imageFileRepo: imageFiles,
	}
}

func (h *Handler) handleInternalServerError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": "Internal Server Error",
	})
}

func (h *Handler) handleNotFound(c *gin.Context, name string) {
	msg := fmt.Sprintf("%s is not found", name)
	c.JSON(http.StatusNotFound, gin.H{
		"message": msg,
	})
}

func (h *Handler) handleIntQuery(c *gin.Context, key string) (value int, err error) {
	value, err = strconv.Atoi(c.Query(key))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%s is not integer", key)})
	}
	return
}
