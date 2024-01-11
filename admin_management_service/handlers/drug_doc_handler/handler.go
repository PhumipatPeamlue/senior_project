package drug_doc_handler

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

type DrugDocHandler struct {
	drugDocService ports.DrugDocService
}

func NewDrugDocHandler(drugDocService ports.DrugDocService) *DrugDocHandler {
	return &DrugDocHandler{
		drugDocService: drugDocService,
	}
}

func (h *DrugDocHandler) HandleCreateDrugDoc(c *gin.Context) {
	var doc models.DrugDoc
	if err := c.ShouldBind(&doc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	image, err := c.FormFile("image")
	if err != nil {
		if err = h.drugDocService.CreateDrugDoc(doc, "", ""); err != nil {
			handlers.HandleServiceErr(c, err)
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "create successfully"})
		return
	}

	imageName := fmt.Sprintf("%s_%s", time.Now().String(), image.Filename)
	imagePath := filepath.Join("uploads", imageName)
	if err = h.drugDocService.CreateDrugDoc(doc, imageName, imagePath); err != nil {
		handlers.HandleServiceErr(c, err)
		return
	}

	if err = c.SaveUploadedFile(image, imagePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "create successfully"})
}

func (h *DrugDocHandler) HandleGetDrugDoc(c *gin.Context) {
	doc, imageURL, err := h.drugDocService.GetDrugDoc(c.Param("doc_id"))
	if err != nil {
		handlers.HandleServiceErr(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"doc":      doc,
		"imageUrl": imageURL,
	})
}

func (h *DrugDocHandler) HandleSearchDrugDoc(c *gin.Context) {
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

	docs, total, err := h.drugDocService.SearchDrugDoc((page-1)*pageSize, pageSize, keyword)
	if err != nil {
		handlers.HandleServiceErr(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  docs,
		"total": total,
	})
}

func (h *DrugDocHandler) HandleUpdateDrugDoc(c *gin.Context) {
	var doc models.DrugDocWithID
	if err := c.ShouldBind(&doc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	image, err := c.FormFile("image")
	if err != nil {
		if err = h.drugDocService.UpdateDrugDoc(doc, "", ""); err != nil {
			handlers.HandleServiceErr(c, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "update successfully"})
		return
	}

	imageName := fmt.Sprintf("%s_%s", time.Now().String(), image.Filename)
	imagePath := filepath.Join("uploads", imageName)
	if err = h.drugDocService.UpdateDrugDoc(doc, imageName, imagePath); err != nil {
		handlers.HandleServiceErr(c, err)
		return
	}

	if err = c.SaveUploadedFile(image, imagePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "update successfully"})
}

func (h *DrugDocHandler) HandleDeleteDrugDoc(c *gin.Context) {
	if err := h.drugDocService.DeleteDrugDoc(c.Param("doc_id")); err != nil {
		handlers.HandleServiceErr(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "delete successfully"})
}
