package controllers

import (
	"net/http"
	"time"

	"github.com/AhmadNizar/dtc/internal/application/usecase"
	"github.com/AhmadNizar/dtc/internal/dto"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AttendanceHandler struct {
	UseCase usecase.Attendance
}

func NewAttendanceHandler(usecase usecase.Attendance) *AttendanceHandler {
	return &AttendanceHandler{UseCase: usecase}
}

func (h *AttendanceHandler) Create(c *gin.Context) {
	var req dto.CreateAttendanceDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.GeneralResponseDTO{
			OK:      false,
			Message: "Invalid input",
		})
		return
	}

	result, err := h.UseCase.Create(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.GeneralResponseDTO{
			OK:      false,
			Message: "Failed to create attendance",
		})
		return
	}

	c.JSON(http.StatusOK, dto.GeneralResponseDTO{
		OK:      true,
		Message: "Attendance created successfully",
		Data:    result,
	})
}

func (h *AttendanceHandler) GetByUserAndDate(c *gin.Context) {
	userIDParam := c.Param("user_id")
	dateParam := c.Param("date")

	userID, err := uuid.Parse(userIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.GeneralResponseDTO{OK: false, Message: "Invalid user ID"})
		return
	}

	date, err := time.Parse("2006-01-02", dateParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.GeneralResponseDTO{OK: false, Message: "Invalid date format"})
		return
	}

	attendance, err := h.UseCase.GetByUserAndDate(userID, date)
	if err != nil {
		c.JSON(http.StatusNotFound, dto.GeneralResponseDTO{OK: false, Message: "Attendance not found"})
		return
	}

	c.JSON(http.StatusOK, dto.GeneralResponseDTO{
		OK:      true,
		Message: "Successfully retrieved attendance",
		Data:    attendance,
	})
}
