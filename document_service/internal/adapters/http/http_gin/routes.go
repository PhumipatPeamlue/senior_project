package http_gin

import "github.com/gin-gonic/gin"

func DocRoutes(r *gin.Engine, h *DocHandler, vh *VideoDocHandler, dh *DrugDocHandler) {
	docRouter := r.Group("/document")
	{
		docRouter.GET("/search", h.SearchDoc)
	}

	videoDocRouter := docRouter.Group("/video_doc")
	{
		videoDocRouter.GET("/:doc_id", vh.GetVideoDocHandler)
		videoDocRouter.GET("/search", vh.SearchVideoDocHandler)
		videoDocRouter.POST("/", vh.AddNewVideoDocHandler)
		videoDocRouter.PUT("/", vh.ChangeVideoDocInfoHandler)
		videoDocRouter.DELETE("/:doc_id", vh.RemoveVideoDocHandler)
	}

	drugDocRouter := docRouter.Group("/drug_doc")
	{
		drugDocRouter.GET("/:doc_id", dh.GetDrugDocHandler)
		drugDocRouter.GET("/search", dh.SearchDrugDocHandler)
		drugDocRouter.POST("/", dh.AddNewDrugDocHandler)
		drugDocRouter.PUT("/", dh.ChangeDrugDocInfoHandler)
		drugDocRouter.DELETE("/:doc_id", dh.RemoveDrugDocHandler)
	}
}
