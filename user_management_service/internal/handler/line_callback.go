package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) LineCallback() func(c *gin.Context) {
	return func(c *gin.Context) {
		var err error
		defer func() {
			if err != nil {
				log.Println(err.Error())
			}
		}()

		code := c.Query("code")
		state := c.Query("state")
		if state != h.lineClient.GetState() {
			c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
			return
		}

		resp, err := h.lineClient.GetAccessToken(code)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		}

		c.JSON(http.StatusOK, resp)
	}
}
