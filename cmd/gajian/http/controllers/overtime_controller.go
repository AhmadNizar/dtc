package controllers

import (
	"net/http"

	"github.com/AhmadNizar/dtc/internal/application/usecase"
	"github.com/AhmadNizar/dtc/internal/dto"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type OvertimeHandler struct {
	UseCase usecase.Overtime
}

func NewOvertimeHandler(usecase usecase.Overtime) *OvertimeHandler {
	return &OvertimeHandler{UseCase: usecase}
}

func (h *OvertimeHandler) Create(c *gin.Context) {
	var req dto.CreateOvertimeDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.GeneralResponseDTO{OK: false, Message: "Invalid request payload"})
		return
	}

	result, err := h.UseCase.Create(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.GeneralResponseDTO{OK: false, Message: "Failed to create overtime"})
		return
	}

	c.JSON(http.StatusCreated, dto.GeneralResponseDTO{OK: true, Message: "Overtime created successfully", Data: result})
}

func (h *OvertimeHandler) Update(c *gin.Context) {
	var req dto.UpdateOvertimeDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.GeneralResponseDTO{OK: false, Message: "Invalid request payload"})
		return
	}

	result, err := h.UseCase.Update(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.GeneralResponseDTO{OK: false, Message: "Failed to update overtime"})
		return
	}

	c.JSON(http.StatusOK, dto.GeneralResponseDTO{OK: true, Message: "Overtime updated successfully", Data: result})
}

func (h *OvertimeHandler) GetByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.GeneralResponseDTO{OK: false, Message: "Invalid ID"})
		return
	}

	result, err := h.UseCase.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, dto.GeneralResponseDTO{OK: false, Message: "Overtime not found"})
		return
	}

	c.JSON(http.StatusOK, dto.GeneralResponseDTO{OK: true, Message: "Overtime fetched successfully", Data: result})
}

func (h *OvertimeHandler) ListAll(c *gin.Context) {
	list, err := h.UseCase.ListAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.GeneralResponseDTO{OK: false, Message: "Failed to list overtime"})
		return
	}

	c.JSON(http.StatusOK, dto.GeneralResponseDTO{OK: true, Message: "Overtime list retrieved", Data: list})
}
