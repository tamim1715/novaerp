package leavetype

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup, handler *Handler) {
	router.POST("", handler.CreateLeaveType)
	router.GET("", handler.FindAllLeaveTypes)
	router.GET("/:id", handler.FindLeaveTypeByID)
	router.PUT("/:id", handler.UpdateLeaveType)
	router.DELETE("/:id", handler.DeleteLeaveType)
}
