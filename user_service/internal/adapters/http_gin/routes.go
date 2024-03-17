package http_gin

import "github.com/gin-gonic/gin"

func UserRoutes(r *gin.Engine, h *UserHandler) {
	userRouter := r.Group("/user")
	{
		userRouter.POST("/:line_user_id", h.AddNewUser)
		userRouter.PUT("/", h.ChangeTimeSetting)
		userRouter.GET("/:line_user_id", h.FindUserByLineUserID)
		userRouter.DELETE("/:line_user_id", h.RemoveUserByLineUserID)
	}
}
