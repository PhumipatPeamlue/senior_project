package drug_doc_handler

import (
	"errors"
	"log"
	"net/http"
	"senior_project/admin_management_service/services/drug_doc_service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GinDrugDocHandler struct {
	drugDocService drug_doc_service.DrugDocService
}

func NewGinDrugDocHandler(drugDocService drug_doc_service.DrugDocService) *GinDrugDocHandler {
	return &GinDrugDocHandler{
		drugDocService: drugDocService,
	}
}

func (h *GinDrugDocHandler) handleServiceError(c *gin.Context, err error) {
	log.Println(err)
	if errors.Is(err, drug_doc_service.ErrDrugDocNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "drug document not found"})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
	}
}

func (h *GinDrugDocHandler) HandleGetDrugDoc(c *gin.Context) {
	docID := c.Param("doc_id")
	res, err := h.drugDocService.GetDrugDoc(docID)
	if err != nil {
		h.handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, *res)
}

func (h *GinDrugDocHandler) HandleSearchDrugDoc(c *gin.Context) {
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

	res, err := h.drugDocService.SearchDrugDoc(page, pageSize, keyword)
	if err != nil {
		h.handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, *res)
}

func (h *GinDrugDocHandler) HandleCreateDrugDoc(c *gin.Context) {
	var newDocReq drug_doc_service.NewDrugDocRequest
	if err := c.ShouldBind(&newDocReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't bind JSON"})
		return
	}

	file, header, err := c.Request.FormFile("image")
	if err != nil {
		if err = h.drugDocService.CreateDrugDoc(&newDocReq); err != nil {
			h.handleServiceError(c, err)
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "create successfully"})
		return
	}

	if err = h.drugDocService.CreateDrugDocWithImage(&newDocReq, &file, header); err != nil {
		h.handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "create successfully"})
}

func (h *GinDrugDocHandler) HandleUpdateDrugDoc(c *gin.Context) {
	var updateDrugDocReq drug_doc_service.UpdateDrugDocRequest
	if err := c.ShouldBind(&updateDrugDocReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't bind JSON"})
		return
	}

	file, header, err := c.Request.FormFile("image")
	if err != nil {
		if err = h.drugDocService.UpdateDrugDoc(&updateDrugDocReq); err != nil {
			h.handleServiceError(c, err)
		}
		c.JSON(http.StatusOK, gin.H{"message": "update successfully"})
		return
	}

	if err = h.drugDocService.UpdateDrugDocWithImage(&updateDrugDocReq, &file, header); err != nil {
		h.handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "update successfully"})
}

func (h *GinDrugDocHandler) HandleDeleteDrugDoc(c *gin.Context) {
	docID := c.Param("doc_id")
	if err := h.drugDocService.DeleteDrugDoc(docID); err != nil {
		h.handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "delete successfully"})
}
