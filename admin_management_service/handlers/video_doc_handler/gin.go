package video_doc_handler

import (
	"errors"
	"log"
	"net/http"
	"senior_project/admin_management_service/services/video_doc_service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GinVideoDocHandler struct {
	videoDocService video_doc_service.VideoDocService
}

func NewGinVideoDocHandler(videoDocService video_doc_service.VideoDocService) *GinVideoDocHandler {
	return &GinVideoDocHandler{
		videoDocService: videoDocService,
	}
}

func (h *GinVideoDocHandler) handleServiceError(c *gin.Context, err error) {
	log.Println(err)
	if errors.Is(err, video_doc_service.ErrVideoDocNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "video document not found"})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
	}
}

func (h *GinVideoDocHandler) HandleGetVideoDoc(c *gin.Context) {
	docID := c.Param("doc_id")
	res, err := h.videoDocService.GetVideoDoc(docID)
	if err != nil {
		h.handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, *res)
}

func (h *GinVideoDocHandler) HandleSearchVideoDoc(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "page isn't integer"})
		return
	}
	pageSize, err := strconv.Atoi(c.Query("page_size"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "page_size isn't integer"})
		return
	}
	keyword := c.Query("keyword")

	res, err := h.videoDocService.SearchVideoDoc(page, pageSize, keyword)
	if err != nil {
		h.handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, *res)
}

func (h *GinVideoDocHandler) HandleCreateVideoDoc(c *gin.Context) {
	var newDocReq video_doc_service.NewVideoDocRequest
	if err := c.ShouldBind(&newDocReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't bind JSON"})
		return
	}

	file, header, err := c.Request.FormFile("image")
	if err != nil {
		if err = h.videoDocService.CreateVideoDoc(&newDocReq); err != nil {
			h.handleServiceError(c, err)
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "create successfully"})
		return
	}

	if err = h.videoDocService.CreateVideoDocWithImage(&newDocReq, &file, header); err != nil {
		h.handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "create successfully"})
}

func (h *GinVideoDocHandler) HandleUpdateVideoDoc(c *gin.Context) {
	var updateVideoDocReq video_doc_service.UpdateVideoDocRequest
	if err := c.ShouldBind(&updateVideoDocReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't bind JSON"})
		return
	}

	file, header, err := c.Request.FormFile("image")
	if err != nil {
		if err = h.videoDocService.UpdateVideoDoc(&updateVideoDocReq); err != nil {
			h.handleServiceError(c, err)
		}
		c.JSON(http.StatusOK, gin.H{"message": "update successfully"})
		return
	}

	if err = h.videoDocService.UpdateVideoDocWithImage(&updateVideoDocReq, &file, header); err != nil {
		h.handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "update successfully"})
}

func (h *GinVideoDocHandler) HandleDeleteVideoDoc(c *gin.Context) {
	docID := c.Param("doc_id")
	if err := h.videoDocService.DeleteVideoDoc(docID); err != nil {
		h.handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "delete successfully"})
}
