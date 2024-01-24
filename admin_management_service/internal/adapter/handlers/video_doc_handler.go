package handlers

import (
	"document_service/internal/core/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type VideoDocHandler struct {
	service services.VideoDocService
}

func NewVideoDocHandler(service services.VideoDocService) *VideoDocHandler {
	return &VideoDocHandler{
		service: service,
	}
}

func (h *VideoDocHandler) AddNewVideoDoc(c *gin.Context) {
	var err error
	var req services.AddNewVideoDocRequest
	if err = c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form format or missing required fields"})
		return
	}

	file, header, err := c.Request.FormFile("image")
	if err != nil {
		err = h.service.AddNewVideoDoc(req, nil, nil)
	} else {
		err = h.service.AddNewVideoDoc(req, &file, header)
	}

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "video document added successfully"})
}

func (h *VideoDocHandler) GetVideoDoc(c *gin.Context) {
	docID := c.Param("doc_id")
	res, err := h.service.GetVideoDoc(docID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *VideoDocHandler) SearchVideoDoc(c *gin.Context) {
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

	res, err := h.service.SearchVideoDoc(page, pageSize, keyword)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *VideoDocHandler) ChangeVideoDocInfo(c *gin.Context) {
	var err error
	var req services.ChangeVideoDocInfoRequest
	if err = c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form format or missing required fields"})
		return
	}

	file, header, err := c.Request.FormFile("image")
	if err != nil {
		err = h.service.ChangeVideoDocInfo(req, nil, nil)
	} else {
		err = h.service.ChangeVideoDocInfo(req, &file, header)
	}

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "video document's information changed successfully"})
}

func (h *VideoDocHandler) RemoveVideoDoc(c *gin.Context) {
	docID := c.Param("doc_id")
	if err := h.service.RemoveVideoDoc(docID); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "video document removed successfully"})
}
