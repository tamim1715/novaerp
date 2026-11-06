package shift

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup, handler *Handler) {
	router.POST("", handler.CreateShift)
	router.GET("", handler.FindAllShifts)
	router.GET("/:id", handler.FindShiftByID)
	router.PUT("/:id", handler.UpdateShift)
	router.DELETE("/:id", handler.DeleteShift)

	// Employee Shift assignments
	router.POST("/assign", handler.AssignShift)
	router.GET("/assignments", handler.FindEmployeeShifts)
}
