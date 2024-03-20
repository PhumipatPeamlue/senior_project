package http_gin

import (
	"document_service/internal/core"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DocHandler struct {
	service core.DocServiceInterface
}

func (h *DocHandler) handleError(c *gin.Context, err error) {
	c.Error(err)

	var errDocNotFound *core.ErrDocNotFound
	var errDocDuplicate *core.ErrDocDuplicate
	var errDocBadRequest *core.ErrDocBadRequest
	switch {
	case errors.As(err, &errDocNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": "document not found"})
	case errors.As(err, &errDocDuplicate):
		c.JSON(http.StatusConflict, gin.H{"error": "document already exists"})
	case errors.As(err, &errDocBadRequest):
		c.JSON(http.StatusBadRequest, gin.H{"error": "document bad request"})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
	}
}

func (h *DocHandler) SearchDoc(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page parameter. Please provide a valid integer value for 'page'."})
		return
	}
	pageSize, err := strconv.Atoi(c.Query("page_size"))
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page_size parameter. Please provide a valid integer value for 'page_size'"})
		return
	}
	keyword := c.Query("keyword")

	docs, total, err := h.service.SearchDoc(page, pageSize, keyword)
	if err != nil {
		h.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  docs,
		"total": total,
	})
}

func NewDocHandler(service core.DocServiceInterface) *DocHandler {
	return &DocHandler{
		service: service,
	}
}
