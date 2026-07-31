package department

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup, handler *Handler) {

	router.POST("", handler.CreateDoc)

	router.GET("", handler.FindAllDoc)

	router.GET("/:id", handler.FindByIDDoc)

	router.PUT("/:id", handler.UpdateDoc)

	router.DELETE("/:id", handler.DeleteDoc)
}
