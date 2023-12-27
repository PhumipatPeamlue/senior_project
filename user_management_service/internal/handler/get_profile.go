package handler

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"user_management_service/internal/models"
)

func (h *Handler) GetUserProfile() func(c *gin.Context) {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		displayName, ok := session.Get("displayName").(string)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}
		pictureURL, ok := session.Get("pictureURL").(string)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}
		userID, ok := session.Get("userID").(string)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}

		profile := models.Profile{
			DisplayName: displayName,
			PictureURL:  pictureURL,
			UserID:      userID,
		}

		c.JSON(http.StatusOK, profile)
	}
}
