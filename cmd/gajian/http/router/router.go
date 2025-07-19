package router

import (
	"github.com/AhmadNizar/dtc/cmd/gajian/http/controllers"
	"github.com/AhmadNizar/dtc/internal/application/usecase"
	"github.com/gin-gonic/gin"
)

func NewRouter(
	userService usecase.User,
	attendanceService usecase.Attendance,
	attendancePeriodService usecase.AttendancePeriod,
	overtimeService usecase.Overtime,
	payrollService usecase.Payroll,
	payslipService usecase.Payslip,
	reimbursementService usecase.Reimbursement,
) *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1")

	// User Routes
	userHandler := controllers.NewUserHandler(userService)
	v1.GET("/users/:id", userHandler.GetUserByID)
	v1.GET("/users", userHandler.ListUsers)

	// Attendance Routes
	attendanceHandler := controllers.NewAttendanceHandler(attendanceService)
	v1.POST("/attendances", attendanceHandler.Create)

	// Attendance Period Routes
	attendancePeriodHandler := controllers.NewAttendancePeriodHandler(attendancePeriodService)
	v1.POST("/attendance-periods", attendancePeriodHandler.Create)
	v1.GET("/attendance-periods/:id", attendancePeriodHandler.GetByID)
	v1.GET("/attendance-periods", attendancePeriodHandler.List)

	// Overtime Routes
	overtimeHandler := controllers.NewOvertimeHandler(overtimeService)
	v1.POST("/overtimes", overtimeHandler.Create)
	v1.GET("/overtimes/:id", overtimeHandler.GetByID)
	v1.GET("/overtimes", overtimeHandler.ListAll)
	v1.PUT("/overtimes", overtimeHandler.Update)

	// Payroll Routes
	payrollHandler := controllers.NewPayrollHandler(payrollService)
	v1.POST("/payrolls", payrollHandler.Create)
	v1.GET("/payrolls/:id", payrollHandler.GetByID)
	v1.GET("/payrolls", payrollHandler.ListAll)
	v1.PUT("/payrolls", payrollHandler.Update)

	// Payslip Routes
	payslipHandler := controllers.NewPayslipHandler(payslipService)
	v1.POST("/payslips", payslipHandler.Create)
	v1.PUT("/payslips", payslipHandler.Update)

	// Reimbursement Routes
	reimbursementHandler := controllers.NewReimbursementHandler(reimbursementService)
	v1.POST("/reimbursements", reimbursementHandler.Create)
	v1.GET("/reimbursements/:id", reimbursementHandler.GetByID)
	v1.GET("/reimbursements", reimbursementHandler.ListAllInPeriod)
	v1.GET("/reimbursements/user", reimbursementHandler.ListByUserInPeriod)
	v1.PUT("/reimbursements", reimbursementHandler.Update)
	v1.DELETE("/reimbursements/:id", reimbursementHandler.Delete)

	return router
}
