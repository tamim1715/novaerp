package leave

import (
	"github.com/gin-gonic/gin"
	"github.com/tamim1715/novaerp/internal/app"
	"github.com/tamim1715/novaerp/internal/modules/employee"
	"github.com/tamim1715/novaerp/internal/modules/hr/leave/leaverequest"
	"github.com/tamim1715/novaerp/internal/modules/hr/leave/leavetype"
)

// RegisterRoutes registers leave submodules (leavetype and leaverequest) under /hr/leaves.
func RegisterRoutes(leavesGroup *gin.RouterGroup, application *app.Application) {
	employeeRepo := employee.NewRepository(application.DB)

	// LeaveType Submodule
	leaveTypeRepo := leavetype.NewRepository(application.DB)
	leaveTypeService := leavetype.NewService(leaveTypeRepo, application.Logger)
	leaveTypeHandler := leavetype.NewHandler(leaveTypeService)

	// LeaveRequest Submodule
	leaveRequestRepo := leaverequest.NewRepository(application.DB)
	leaveRequestService := leaverequest.NewService(leaveRequestRepo, leaveTypeRepo, employeeRepo, application.Logger)
	leaveRequestHandler := leaverequest.NewHandler(leaveRequestService)

	// Register submodule routes
	leavetype.RegisterRoutes(leavesGroup.Group("/types"), leaveTypeHandler)
	leaverequest.RegisterRoutes(leavesGroup.Group("/requests"), leaveRequestHandler)
}
