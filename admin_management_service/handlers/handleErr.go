package handlers

import (
	"admin_management_service/errs"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleServiceErr(c *gin.Context, err error) {
	if errors.Is(err, errs.UnexpectedError) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else if errors.Is(err, errs.VideoDocNotFoundError) {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else if errors.Is(err, errs.DocImageNotFoundError) {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else if errors.Is(err, errs.DrugDocNotFoundError) {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}
	return
}
