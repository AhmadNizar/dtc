package controllers

import (
	"net/http"
	"time"

	"github.com/AhmadNizar/dtc/internal/application/usecase"
	"github.com/AhmadNizar/dtc/internal/dto"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AttendancePeriodHandler struct {
	UseCase usecase.AttendancePeriod
}

func NewAttendancePeriodHandler(usecase usecase.AttendancePeriod) *AttendancePeriodHandler {
	return &AttendancePeriodHandler{UseCase: usecase}
}

func (aph *AttendancePeriodHandler) Create(c *gin.Context) {
	var req dto.CreateAttendancePeriodDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.GeneralResponseDTO{
			OK:      false,
			Message: "Invalid request body",
		})
		return
	}

	period, err := aph.UseCase.Create(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.GeneralResponseDTO{
			OK:      false,
			Message: "Failed to create attendance period",
		})
		return
	}

	c.JSON(http.StatusCreated, dto.GeneralResponseDTO{
		OK:      true,
		Message: "Attendance period created",
		Data:    period,
	})
}

func (aph *AttendancePeriodHandler) List(c *gin.Context) {
	periods, err := aph.UseCase.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.GeneralResponseDTO{
			OK:      false,
			Message: "Failed to fetch attendance periods",
		})
		return
	}

	c.JSON(http.StatusOK, dto.GeneralResponseDTO{
		OK:      true,
		Message: "Successfully fetched attendance periods",
		Data:    periods,
	})
}

func (aph *AttendancePeriodHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.GeneralResponseDTO{
			OK:      false,
			Message: "Invalid period ID",
		})
		return
	}

	period, err := aph.UseCase.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, dto.GeneralResponseDTO{
			OK:      false,
			Message: "Attendance period not found",
		})
		return
	}

	c.JSON(http.StatusOK, dto.GeneralResponseDTO{
		OK:      true,
		Message: "Successfully fetched attendance period",
		Data:    period,
	})
}

func (aph *AttendancePeriodHandler) GetByDate(c *gin.Context) {
	dateStr := c.Query("date")
	if dateStr == "" {
		c.JSON(http.StatusBadRequest, dto.GeneralResponseDTO{
			OK:      false,
			Message: "Missing date query parameter",
		})
		return
	}

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.GeneralResponseDTO{
			OK:      false,
			Message: "Invalid date format. Use YYYY-MM-DD",
		})
		return
	}

	period, err := aph.UseCase.GetCurrentPeriod(date)
	if err != nil {
		c.JSON(http.StatusNotFound, dto.GeneralResponseDTO{
			OK:      false,
			Message: "Attendance period not found for given date",
		})
		return
	}

	c.JSON(http.StatusOK, dto.GeneralResponseDTO{
		OK:      true,
		Message: "Successfully fetched attendance period by date",
		Data:    period,
	})
}
