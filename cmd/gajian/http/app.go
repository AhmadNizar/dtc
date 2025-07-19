package gajian

import (
	"log"

	"github.com/AhmadNizar/dtc/cmd/gajian/http/router"
	"github.com/AhmadNizar/dtc/internal/application/service"
	"github.com/AhmadNizar/dtc/internal/config"
	dbinfrastructure "github.com/AhmadNizar/dtc/internal/infrastructure/db"
	persistence "github.com/AhmadNizar/dtc/internal/infrastructure/persistence/pg"
)

func Start(config *config.Config) {
	dsn := config.GetDSN()

	db, err := dbinfrastructure.SetupDB(dsn)

	if err != nil {
		panic("Failed to connect to database")
	}

	userRepo := persistence.NewUserRepository(db)
	attendanceRepo := persistence.NewAttendanceRepository(db)
	attendancePeridoRepo := persistence.NewAttendancePeriodRepository(db)
	overtimeRepo := persistence.NewOvertimeRepository(db)
	payrollRepo := persistence.NewPayrollRepository(db)
	payslipRepo := persistence.NewPayslipRepository(db)
	reimbursementRepo := persistence.NewReimbursementRepository(db)

	userService := service.NewUserService(userRepo)
	attendanceService := service.NewAttendanceService(attendanceRepo)
	attendancePeriodService := service.NewAttendancePeriodService(attendancePeridoRepo)
	overtimeService := service.NewOvertimeService(overtimeRepo)
	payrollService := service.NewPayrollService(payrollRepo)
	payslipService := service.NewPayslipService(payslipRepo)
	reimbursementService := service.NewReimbursementService(reimbursementRepo)

	router := router.NewRouter(userService, attendanceService, attendancePeriodService, overtimeService, payrollService, payslipService, reimbursementService)

	log.Println("Starting server on :", config.AppPort)
	if err := router.Run("0.0.0.0:" + config.AppPort); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
