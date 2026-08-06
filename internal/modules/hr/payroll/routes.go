package payroll

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup, handler *Handler) {
	router.POST("", handler.CreatePeriod)
	router.GET("", handler.FindAllPeriods)
	router.GET("/:id", handler.FindPeriodByID)
	router.POST("/:id/process", handler.ProcessPayroll)
	router.GET("/:id/payslips", handler.GetPayslips)
	router.POST("/:id/pay", handler.MarkPaid)
}
