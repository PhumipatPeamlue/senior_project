package video_doc_handler

import (
	"admin_management_service/handlers"
	"admin_management_service/models"
	"admin_management_service/ports"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"strconv"
	"time"
)

type VideoDocHandler struct {
	videoDocService ports.VideoDocService
}

func NewVideoDocHandler(videoDocService ports.VideoDocService) *VideoDocHandler {
	return &VideoDocHandler{
		videoDocService: videoDocService,
	}
}

func (h *VideoDocHandler) HandleCreateVideoDoc(c *gin.Context) {
	var doc models.VideoDoc
	if err := c.ShouldBind(&doc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	image, err := c.FormFile("image")
	if err != nil {
		if err = h.videoDocService.CreateVideoDoc(doc, "", ""); err != nil {
			handlers.HandleServiceErr(c, err)
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "create successfully"})
		return
	}

	imageName := fmt.Sprintf("%s_%s", time.Now().String(), image.Filename)
	imagePath := filepath.Join("uploads", imageName)
	if err = h.videoDocService.CreateVideoDoc(doc, imageName, imagePath); err != nil {
		handlers.HandleServiceErr(c, err)
		return
	}

	if err = c.SaveUploadedFile(image, imagePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "create successfully"})
}

func (h *VideoDocHandler) HandleGetVideoDoc(c *gin.Context) {
	doc, imageURL, err := h.videoDocService.GetVideoDoc(c.Param("doc_id"))
	if err != nil {
		handlers.HandleServiceErr(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"doc":      doc,
		"imageUrl": imageURL,
	})
}

func (h *VideoDocHandler) HandleSearchVideoDoc(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}
	pageSize, err := strconv.Atoi(c.Query("page_size"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}
	keyword := c.Query("keyword")

	docs, total, err := h.videoDocService.SearchVideoDoc((page-1)*pageSize, pageSize, keyword)
	if err != nil {
		handlers.HandleServiceErr(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  docs,
		"total": total,
	})
}

func (h *VideoDocHandler) HandleUpdateVideoDoc(c *gin.Context) {
	var doc models.VideoDocWithId
	if err := c.ShouldBind(&doc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	image, err := c.FormFile("image")
	if err != nil {
		if err = h.videoDocService.UpdateVideoDoc(doc, "", ""); err != nil {
			handlers.HandleServiceErr(c, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "update successfully"})
		return
	}

	imageName := fmt.Sprintf("%s_%s", time.Now().String(), image.Filename)
	imagePath := filepath.Join("uploads", imageName)
	if err = h.videoDocService.UpdateVideoDoc(doc, imageName, imagePath); err != nil {
		handlers.HandleServiceErr(c, err)
		return
	}

	if err = c.SaveUploadedFile(image, imagePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "update successfully"})
}

func (h *VideoDocHandler) HandleDeleteVideoDoc(c *gin.Context) {
	if err := h.videoDocService.DeleteVideoDoc(c.Param("doc_id")); err != nil {
		handlers.HandleServiceErr(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "delete successfully"})
}
