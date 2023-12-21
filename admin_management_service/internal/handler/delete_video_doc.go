package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) DeleteVideoDoc() func(c *gin.Context) {
	return func(c *gin.Context) {
		var err error
		defer func() {
			if err != nil {
				log.Println(err.Error())
			}
		}()

		id := c.Param("id")
		if err = h.videoDocIndex.Delete(id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "can't delete the video document"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "delete the video document successfully"})
	}
}
