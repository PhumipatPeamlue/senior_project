package handlers

import (
	"document_service/internal/core/ports"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
	service ports.DrugDocService
}

func (h *DrugDocHandler) AddNewDrugDocHandler(c *gin.Context) {
	var body AddNewDrugDocRequest
	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form format or missing required fields"})
		return
	}

	file, header, err := c.Request.FormFile("image")
	if err != nil {
		err = h.service.AddNewDrugDoc(body.TradeName, body.DrugName, body.Description, body.Preparation, body.Caution, nil, nil)
	} else {
		err = h.service.AddNewDrugDoc(body.TradeName, body.DrugName, body.Description, body.Preparation, body.Caution, &file, header)
	}

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "add new video document successfully"})
}

func (h *DrugDocHandler) GetDrugDocHandler(c *gin.Context) {
	docID := c.Param("doc_id")
	doc, imageURL, err := h.service.GetDrugDoc(docID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"doc":       doc,
		"image_url": imageURL,
	})
}

func (h *DrugDocHandler) SearchDrugDocHandler(c *gin.Context) {
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

	docs, total, err := h.service.SearchDrugDoc(page, pageSize, keyword)
	if err != nil {
		c.Error(err)
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form format or missing required fields"})
		return
	}

	file, header, err := c.Request.FormFile("image")
	if err != nil {
		err = h.service.ChangeDrugDocInfo(body.ID, body.TradeName, body.DrugName, body.Description, body.Preparation, body.Caution, body.DeleteImage, nil, nil)
	} else {
		err = h.service.ChangeDrugDocInfo(body.ID, body.TradeName, body.DrugName, body.Description, body.Preparation, body.Caution, body.DeleteImage, &file, header)
	}
	
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "change video document information successfully"})
}

func (h *DrugDocHandler) RemoveDrugDocHandler(c *gin.Context) {
	docID := c.Param("doc_id")
	if err := h.service.RemoveDrugDoc(docID); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "remove video document successfully"})
}

func NewDrugDocHandler(service ports.DrugDocService) *DrugDocHandler {
	return &DrugDocHandler{
		service: service,
	}
}
