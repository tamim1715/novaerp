package hr

import (
	"github.com/gin-gonic/gin"
	"github.com/tamim1715/novaerp/internal/app"
	"github.com/tamim1715/novaerp/internal/modules/employee"
	"github.com/tamim1715/novaerp/internal/modules/hr/allowance"
	"github.com/tamim1715/novaerp/internal/modules/hr/attendance"
	"github.com/tamim1715/novaerp/internal/modules/hr/leave"
	"github.com/tamim1715/novaerp/internal/modules/hr/leave/leaverequest"
	"github.com/tamim1715/novaerp/internal/modules/hr/overtime"
	"github.com/tamim1715/novaerp/internal/modules/hr/payroll"
	"github.com/tamim1715/novaerp/internal/modules/hr/shift"
)

// RegisterRoutes registers all HR submodules under /hr.
func RegisterRoutes(api *gin.RouterGroup, application *app.Application) {
	employeeRepo := employee.NewRepository(application.DB)
	hrGroup := api.Group("/hr")

	// Shift Submodule
	shiftRepo := shift.NewRepository(application.DB)
	shiftService := shift.NewService(shiftRepo, employeeRepo, application.Logger)
	shiftHandler := shift.NewHandler(shiftService)
	shift.RegisterRoutes(hrGroup.Group("/shifts"), shiftHandler)

	// Overtime Submodule
	otRepo := overtime.NewRepository(application.DB)
	otService := overtime.NewService(otRepo, application.Logger)
	otHandler := overtime.NewHandler(otService)
	overtime.RegisterRoutes(hrGroup.Group("/overtimes"), otHandler)

	// Allowance Submodule
	allowanceRepo := allowance.NewRepository(application.DB)
	allowanceService := allowance.NewService(allowanceRepo, employeeRepo, application.Logger)
	allowanceHandler := allowance.NewHandler(allowanceService)
	allowance.RegisterRoutes(hrGroup.Group("/allowances"), allowanceHandler)

	// Leave Submodule (delegates to leavetype and leaverequest)
	leave.RegisterRoutes(hrGroup.Group("/leaves"), application)
	leaveRequestRepo := leaverequest.NewRepository(application.DB)

	// Attendance Submodule (integrated with Shift, Overtime, Leave)
	attRepo := attendance.NewRepository(application.DB)
	attService := attendance.NewService(attRepo, employeeRepo, shiftService, otService, leaveRequestRepo, application.Logger)
	attHandler := attendance.NewHandler(attService)
	attendance.RegisterRoutes(hrGroup.Group("/attendances"), attHandler)

	// Payroll Submodule (integrated with Employee, Allowance, Attendance, Overtime, Leave)
	payrollRepo := payroll.NewRepository(application.DB)
	payrollService := payroll.NewService(payrollRepo, employeeRepo, allowanceService, attService, otService, leaveRequestRepo, application.Logger)
	payrollHandler := payroll.NewHandler(payrollService)
	payroll.RegisterRoutes(hrGroup.Group("/payrolls"), payrollHandler)
}
