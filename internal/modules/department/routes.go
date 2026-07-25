package department

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup, handler *Handler) {

	router.POST("", handler.Create)

	router.GET("", handler.FindAll)

	router.GET("/:id", handler.FindByID)

	router.PUT("/:id", handler.Update)

	router.DELETE("/:id", handler.Delete)
}
