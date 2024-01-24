package handlers

import (
	"document_service/internal/core/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DrugDocHandler struct {
	service services.DrugDocService
}

func NewDrugDocHandler(service services.DrugDocService) *DrugDocHandler {
	return &DrugDocHandler{service: service}
}

func (h *DrugDocHandler) AddNewDrugDoc(c *gin.Context) {
	var err error
	var req services.AddNewDrugDocRequest
	if err = c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form format or missing required fields"})
		return
	}

	file, header, err := c.Request.FormFile("image")
	if err != nil {
		err = h.service.AddNewDrugDoc(req, nil, nil)
	} else {
		err = h.service.AddNewDrugDoc(req, &file, header)
	}

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "drug document added successfully"})
}

func (h *DrugDocHandler) GetDrugDoc(c *gin.Context) {
	docID := c.Param("doc_id")
	res, err := h.service.GetDrugDoc(docID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *DrugDocHandler) SearchDrugDoc(c *gin.Context) {
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

	res, err := h.service.SearchDrugDoc(page, pageSize, keyword)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *DrugDocHandler) ChangeDrugDocInfo(c *gin.Context) {
	var err error
	var req services.ChangeDrugDocInfoRequest
	if err = c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid form format or missing required fields"})
		return
	}

	file, header, err := c.Request.FormFile("image")
	if err != nil {
		err = h.service.ChangeDrugDocInfo(req, nil, nil)
	} else {
		err = h.service.ChangeDrugDocInfo(req, &file, header)
	}

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "drug document's information changed successfully"})
}

func (h *DrugDocHandler) RemoveDrugDoc(c *gin.Context) {
	docID := c.Param("doc_id")
	if err := h.service.RemoveDrugDoc(docID); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "drug document removed successfully"})
}
