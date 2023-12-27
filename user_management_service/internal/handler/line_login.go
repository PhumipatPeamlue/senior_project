package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) LineLogin() func(c *gin.Context) {
	return func(c *gin.Context) {
		clientID := h.lineClient.GetClientID()
		redirectURI := h.lineClient.GetRedirectURI()
		state := h.lineClient.GetState()
		format := "https://access.line.me/oauth2/v2.1/authorize?response_type=code&client_id=%s&redirect_uri=%s&state=%s&scope=profile"
		loginURL := fmt.Sprintf(format, clientID, redirectURI, state)
		c.Redirect(http.StatusTemporaryRedirect, loginURL)
	}
}
