package controllers

import (
	"net/http"

	"github.com/AhmadNizar/dtc/internal/application/usecase"
	"github.com/AhmadNizar/dtc/internal/dto"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PayslipHandler struct {
	UseCase usecase.Payslip
}

func NewPayslipHandler(usecase usecase.Payslip) *PayslipHandler {
	return &PayslipHandler{UseCase: usecase}
}

func (plh *PayslipHandler) Create(c *gin.Context) {
	var payload dto.CreatePayslipDTO
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, dto.GeneralResponseDTO{
			OK:      false,
			Message: "Invalid request payload",
		})
		return
	}

	payslip, err := plh.UseCase.Create(payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.GeneralResponseDTO{
			OK:      false,
			Message: "Failed to create payslip",
		})
		return
	}

	c.JSON(http.StatusCreated, dto.GeneralResponseDTO{
		OK:      true,
		Message: "Payslip created successfully",
		Data:    payslip,
	})
}

func (plh *PayslipHandler) Update(c *gin.Context) {
	var payload dto.UpdatePayslipDTO
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, dto.GeneralResponseDTO{
			OK:      false,
			Message: "Invalid request payload",
		})
		return
	}

	updated, err := plh.UseCase.Update(payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.GeneralResponseDTO{
			OK:      false,
			Message: "Failed to update payslip",
		})
		return
	}

	c.JSON(http.StatusOK, dto.GeneralResponseDTO{
		OK:      true,
		Message: "Payslip updated successfully",
		Data:    updated,
	})
}

func (plh *PayslipHandler) GetByUserAndPayroll(c *gin.Context) {
	userIDStr := c.Query("user_id")
	payrollIDStr := c.Query("payroll_id")

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.GeneralResponseDTO{
			OK:      false,
			Message: "Invalid user_id",
		})
		return
	}

	payrollID, err := uuid.Parse(payrollIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.GeneralResponseDTO{
			OK:      false,
			Message: "Invalid payroll_id",
		})
		return
	}

	payslip, err := plh.UseCase.GetByUserAndPayroll(payrollID, userID)
	if err != nil {
		c.JSON(http.StatusNotFound, dto.GeneralResponseDTO{
			OK:      false,
			Message: "Payslip not found",
		})
		return
	}

	c.JSON(http.StatusOK, dto.GeneralResponseDTO{
		OK:      true,
		Message: "Payslip retrieved successfully",
		Data:    payslip,
	})
}

func (plh *PayslipHandler) ListByPayroll(c *gin.Context) {
	payrollIDStr := c.Param("payroll_id")

	payrollID, err := uuid.Parse(payrollIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.GeneralResponseDTO{
			OK:      false,
			Message: "Invalid payroll_id",
		})
		return
	}

	payslips, err := plh.UseCase.ListByPayroll(payrollID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.GeneralResponseDTO{
			OK:      false,
			Message: "Failed to list payslips",
		})
		return
	}

	c.JSON(http.StatusOK, dto.GeneralResponseDTO{
		OK:      true,
		Message: "Payslips retrieved successfully",
		Data:    payslips,
	})
}

func (plh *PayslipHandler) SummaryByPayroll(c *gin.Context) {
	payrollIDStr := c.Param("payroll_id")

	payrollID, err := uuid.Parse(payrollIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.GeneralResponseDTO{
			OK:      false,
			Message: "Invalid payroll_id",
		})
		return
	}

	summary, err := plh.UseCase.SummaryByPayroll(payrollID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.GeneralResponseDTO{
			OK:      false,
			Message: "Failed to get summary",
		})
		return
	}

	c.JSON(http.StatusOK, dto.GeneralResponseDTO{
		OK:      true,
		Message: "Payslip summary retrieved successfully",
		Data:    summary,
	})
}
