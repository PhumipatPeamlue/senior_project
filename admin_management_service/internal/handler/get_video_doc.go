package handler

import (
	"admin_management_service/internal/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) GetVideoDoc() func(c *gin.Context) {
	return func(c *gin.Context) {
		var err error
		defer func() {
			if err != nil {
				log.Println(err.Error())
			}
		}()

		id := c.Param("id")
		getResult, err := h.videoDocIndex.Get(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("can't get the video document with id = %s", id),
			})
			return
		}

		res := models.VideoDocDto{
			ID:       getResult.ID,
			VideoDoc: getResult.Source,
		}
		c.JSON(http.StatusOK, res)
	}
}
