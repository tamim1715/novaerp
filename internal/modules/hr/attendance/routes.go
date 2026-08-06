package attendance

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup, handler *Handler) {
	router.POST("/check-in", handler.CheckIn)
	router.POST("/check-out", handler.CheckOut)
	router.POST("", handler.Create)
	router.GET("", handler.FindAll)
	router.GET("/:id", handler.FindByID)
}
