package http_gin

import (
	"document_service/internal/core"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AddNewDrugDocRequest struct {
	TradeName   string `form:"trade_name" binding:"required"`
	DrugName    string `form:"drug_name" binding:"required"`
	Description string `form:"description" binding:"required"`
	Preparation string `form:"preparation" binding:"required"`
	Caution     string `form:"caution" binding:"required"`
}

type ChangeDrugDocInfoRequest struct {
	ID          string `form:"id" binding:"required"`
	TradeName   string `form:"trade_name"`
	DrugName    string `form:"drug_name"`
	Description string `form:"description"`
	Preparation string `form:"preparation"`
	Caution     string `form:"caution"`
	DeleteImage bool   `form:"delete_image"`
}

type DrugDocHandler struct {
	service core.DrugDocServiceInterface
}

func (h *DrugDocHandler) handleError(c *gin.Context, err error) {
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

func (h *DrugDocHandler) AddNewDrugDocHandler(c *gin.Context) {
	var body AddNewDrugDocRequest
	if err := c.ShouldBind(&body); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form format or missing required fields"})
		return
	}

	docID, err := h.service.AddNewDrugDoc(body.TradeName, body.DrugName, body.Description, body.Preparation, body.Caution)
	if err != nil {
		h.handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "add new drug document successfully",
		"doc_id":  docID,
	})
}

func (h *DrugDocHandler) GetDrugDocHandler(c *gin.Context) {
	docID := c.Param("doc_id")
	doc, err := h.service.GetDrugDoc(docID)
	if err != nil {
		h.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, doc)
}

func (h *DrugDocHandler) SearchDrugDocHandler(c *gin.Context) {
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

	docs, total, err := h.service.SearchDrugDoc(page, pageSize, keyword)
	if err != nil {
		h.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  docs,
		"total": total,
	})
}

func (h *DrugDocHandler) ChangeDrugDocInfoHandler(c *gin.Context) {
	var body ChangeDrugDocInfoRequest
	if err := c.ShouldBind(&body); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form format or missing required fields"})
		return
	}

	err := h.service.ChangeDrugDocInfo(body.ID, body.TradeName, body.DrugName, body.Description, body.Preparation, body.Caution)
	if err != nil {
		h.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "change drug document information successfully"})
}

func (h *DrugDocHandler) RemoveDrugDocHandler(c *gin.Context) {
	docID := c.Param("doc_id")
	if err := h.service.RemoveDrugDoc(docID); err != nil {
		h.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "remove drug document successfully"})
}

func NewDrugDocHandler(service core.DrugDocServiceInterface) *DrugDocHandler {
	return &DrugDocHandler{
		service: service,
	}
}
