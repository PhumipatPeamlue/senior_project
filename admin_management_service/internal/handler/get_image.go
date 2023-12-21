package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"path/filepath"
)

func (h *Handler) GetImage() func(c *gin.Context) {
	return func(c *gin.Context) {
		var err error
		defer func() {
			if err != nil {
				log.Println(err.Error())
			}
		}()

		filename := c.Param("filename")
		path := filepath.Join("uploads", filename)

		c.File(path)
	}
}
