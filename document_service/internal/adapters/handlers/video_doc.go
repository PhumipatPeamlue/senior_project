package handlers

import (
	"document_service/internal/core/domains"
	"document_service/internal/core/ports"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AddNewVideoDocRequest struct {
	Title       string `json:"title" form:"title" binding:"required"`
	VideoURL    string `json:"video_url" form:"video_url" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
}

type ChangeVideoDocInfoRequest struct {
	ID          string `json:"id" form:"id" binding:"required"`
	Title       string `json:"title" form:"title"`
	VideoURL    string `json:"video_url" form:"video_url"`
	Description string `json:"description" form:"description"`
	DeleteImage bool   `form:"delete_image"`
}

type GetVideoDocResponse struct {
	Doc      domains.VideoDoc `json:"doc"`
	ImageURL string           `json:"image_url"`
}

type VideoDocHandler struct {
	service ports.VideoDocService
}

func (h *VideoDocHandler) AddNewVideoDocHandler(c *gin.Context) {
	var body AddNewVideoDocRequest
	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form format or missing required fields"})
		return
	}

	file, header, err := c.Request.FormFile("image")
	if err != nil {
		err = h.service.AddNewVideoDoc(body.Title, body.VideoURL, body.Description, nil, nil)
	} else {
		err = h.service.AddNewVideoDoc(body.Title, body.VideoURL, body.Description, &file, header)
	}

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "add new video document successfully"})
}

func (h *VideoDocHandler) GetVideoDocHandler(c *gin.Context) {
	docID := c.Param("doc_id")
	doc, imageURL, err := h.service.GetVideoDoc(docID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"doc":       doc,
		"image_url": imageURL,
	})
}

func (h *VideoDocHandler) SearchVideoDocHandler(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page parameter. Please provide a valid integer value for 'page'."})
		return
	}
	pageSize, err := strconv.Atoi(c.Query("page_size"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page_size parameter. Please provide a valid integer value for 'page_size'"})
		return
	}
	keyword := c.Query("keyword")

	docs, total, err := h.service.SearchVideoDoc(page, pageSize, keyword)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  docs,
		"total": total,
	})
}

func (h *VideoDocHandler) ChangeVideoDocInfoHandler(c *gin.Context) {
	var body ChangeVideoDocInfoRequest
	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form format or missing required fields"})
		return
	}

	file, header, err := c.Request.FormFile("image")
	if err != nil {
		err = h.service.ChangeVideoDocInfo(body.ID, body.Title, body.VideoURL, body.Description, body.DeleteImage, nil, nil)
	} else {
		err = h.service.ChangeVideoDocInfo(body.ID, body.Title, body.VideoURL, body.Description, body.DeleteImage, &file, header)
	}

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "change video document information successfully"})
}

func (h *VideoDocHandler) RemoveVideoDocHandler(c *gin.Context) {
	docID := c.Param("doc_id")
	if err := h.service.RemoveVideoDoc(docID); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "remove video document successfully"})
}

func NewVideoDocHandler(service ports.VideoDocService) *VideoDocHandler {
	return &VideoDocHandler{
		service: service,
	}
}
