package controllers

import (
	"net/http"

	"github.com/AhmadNizar/dtc/internal/application/usecase"
	"github.com/AhmadNizar/dtc/internal/dto"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PayrollHandler interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	GetByID(c *gin.Context)
	ListAll(c *gin.Context)
	ListByPeriod(c *gin.Context)
}

type payrollHandler struct {
	UseCase usecase.Payroll
}

func NewPayrollHandler(usecase usecase.Payroll) PayrollHandler {
	return &payrollHandler{UseCase: usecase}
}

func (ph *payrollHandler) Create(c *gin.Context) {
	var input dto.CreatePayrollDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := ph.UseCase.Create(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, result)
}

func (ph *payrollHandler) Update(c *gin.Context) {
	var input dto.UpdatePayrollDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := ph.UseCase.Update(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (ph *payrollHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	result, err := ph.UseCase.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (ph *payrollHandler) ListAll(c *gin.Context) {
	result, err := ph.UseCase.ListAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (ph *payrollHandler) ListByPeriod(c *gin.Context) {
	periodIDStr := c.Param("period_id")
	periodID, err := uuid.Parse(periodIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	result, err := ph.UseCase.ListByPeriod(periodID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
