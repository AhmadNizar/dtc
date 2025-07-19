package controllers

import (
	"net/http"
	"time"

	"github.com/AhmadNizar/dtc/internal/application/usecase"
	"github.com/AhmadNizar/dtc/internal/dto"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ReimbursementHandler struct {
	service usecase.Reimbursement
}

func NewReimbursementHandler(service usecase.Reimbursement) *ReimbursementHandler {
	return &ReimbursementHandler{service: service}
}

func (h *ReimbursementHandler) Create(c *gin.Context) {
	var input dto.CreateReimbursementDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, dto.GeneralResponseDTO{
			OK:      false,
			Message: err.Error(),
		})
		return
	}

	reimbursement, err := h.service.Create(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.GeneralResponseDTO{
			OK:      false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, dto.GeneralResponseDTO{
		OK:      true,
		Message: "Successfully created reimbursement",
		Data:    reimbursement,
	})
}

func (h *ReimbursementHandler) Update(c *gin.Context) {
	var input dto.UpdateReimbursementDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, dto.GeneralResponseDTO{
			OK:      false,
			Message: err.Error(),
		})
		return
	}

	reimbursement, err := h.service.Update(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.GeneralResponseDTO{
			OK:      false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.GeneralResponseDTO{
		OK:      true,
		Message: "Successfully updated reimbursement",
		Data:    reimbursement,
	})
}

func (h *ReimbursementHandler) GetByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.GeneralResponseDTO{
			OK:      false,
			Message: "Invalid UUID",
		})
		return
	}

	reimbursement, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, dto.GeneralResponseDTO{
			OK:      false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.GeneralResponseDTO{
		OK:      true,
		Message: "Successfully fetched reimbursement",
		Data:    reimbursement,
	})
}

func (h *ReimbursementHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.GeneralResponseDTO{
			OK:      false,
			Message: "Invalid UUID",
		})
		return
	}

	if err := h.service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, dto.GeneralResponseDTO{
			OK:      false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.GeneralResponseDTO{
		OK:      true,
		Message: "Successfully deleted reimbursement",
	})
}

func (h *ReimbursementHandler) ListByUserInPeriod(c *gin.Context) {
	userIDParam := c.Query("user_id")
	startDateParam := c.Query("start_date")
	endDateParam := c.Query("end_date")

	userID, err := uuid.Parse(userIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.GeneralResponseDTO{
			OK:      false,
			Message: "Invalid user_id",
		})
		return
	}

	startDate, err := time.Parse("2006-01-02", startDateParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.GeneralResponseDTO{
			OK:      false,
			Message: "Invalid start_date",
		})
		return
	}

	endDate, err := time.Parse("2006-01-02", endDateParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.GeneralResponseDTO{
			OK:      false,
			Message: "Invalid end_date",
		})
		return
	}

	reimbursements, err := h.service.ListByUserInPeriod(userID, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.GeneralResponseDTO{
			OK:      false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.GeneralResponseDTO{
		OK:      true,
		Message: "Successfully fetched reimbursements by user",
		Data:    reimbursements,
	})
}

func (h *ReimbursementHandler) ListAllInPeriod(c *gin.Context) {
	startDateParam := c.Query("start_date")
	endDateParam := c.Query("end_date")

	startDate, err := time.Parse("2006-01-02", startDateParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.GeneralResponseDTO{
			OK:      false,
			Message: "Invalid start_date",
		})
		return
	}

	endDate, err := time.Parse("2006-01-02", endDateParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.GeneralResponseDTO{
			OK:      false,
			Message: "Invalid end_date",
		})
		return
	}

	reimbursements, err := h.service.ListAllInPeriod(startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.GeneralResponseDTO{
			OK:      false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.GeneralResponseDTO{
		OK:      true,
		Message: "Successfully fetched all reimbursements in period",
		Data:    reimbursements,
	})
}
