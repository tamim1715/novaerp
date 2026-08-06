package leaverequest

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup, handler *Handler) {
	router.POST("", handler.CreateLeaveRequest)
	router.GET("", handler.FindAllLeaveRequests)
	router.GET("/:id", handler.FindLeaveRequestByID)
	router.PUT("/:id/status", handler.UpdateLeaveStatus)
}
