package hr

import (
	"github.com/gin-gonic/gin"
	"github.com/tamim1715/novaerp/internal/app"
	"github.com/tamim1715/novaerp/internal/modules/employee"
	"github.com/tamim1715/novaerp/internal/modules/hr/attendance"
	"github.com/tamim1715/novaerp/internal/modules/hr/leave"
	"github.com/tamim1715/novaerp/internal/modules/hr/payroll"
)

// RegisterRoutes registers all HR submodules (leave, attendance, payroll) under /hr.
func RegisterRoutes(api *gin.RouterGroup, application *app.Application) {
	employeeRepo := employee.NewRepository(application.DB)
	hrGroup := api.Group("/hr")

	// Leave Submodule (delegates to leavetype and leaverequest)
	leave.RegisterRoutes(hrGroup.Group("/leaves"), application)

	// Attendance Submodule
	attRepo := attendance.NewRepository(application.DB)
	attService := attendance.NewService(attRepo, employeeRepo, application.Logger)
	attHandler := attendance.NewHandler(attService)
	attendance.RegisterRoutes(hrGroup.Group("/attendances"), attHandler)

	// Payroll Submodule
	payrollRepo := payroll.NewRepository(application.DB)
	payrollService := payroll.NewService(payrollRepo, employeeRepo, application.Logger)
	payrollHandler := payroll.NewHandler(payrollService)
	payroll.RegisterRoutes(hrGroup.Group("/payrolls"), payrollHandler)
}
