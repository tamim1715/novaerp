package allowance

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup, handler *Handler) {
	router.POST("/types", handler.CreateType)
	router.GET("/types", handler.FindAllTypes)
	router.GET("/types/:id", handler.FindTypeByID)
	router.PUT("/types/:id", handler.UpdateType)
	router.DELETE("/types/:id", handler.DeleteType)

	router.POST("/assign", handler.AssignAllowance)
	router.GET("/assignments", handler.FindEmployeeAllowances)
	router.DELETE("/assignments/:id", handler.DeleteEmployeeAllowance)
}
