package controllers

import (
	"net/http"

	"github.com/AhmadNizar/dtc/internal/application/usecase"
	"github.com/AhmadNizar/dtc/internal/dto"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// UserHandler handles user-related HTTP requests
type UserHandler struct {
	UseCase usecase.User
}

// NewUserHandler returns a new UserHandler
func NewUserHandler(usecase usecase.User) *UserHandler {
	return &UserHandler{UseCase: usecase}
}

func (uh *UserHandler) ListUsers(c *gin.Context) {
	users, err := uh.UseCase.ListAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.GeneralResponseDTO{
			OK:      false,
			Message: "failed to fetch user data",
		})
		return
	}
	c.JSON(http.StatusOK, dto.GeneralResponseDTO{
		OK:      true,
		Message: "Successfully get user data",
		Data:    users,
	})
}

func (uh *UserHandler) GetUserByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, dto.GeneralResponseDTO{
			OK:      false,
			Message: "Invalid user ID",
		})

		return
	}

	user, err := uh.UseCase.GetByID(id)

	if err != nil {
		c.JSON(http.StatusNotFound, dto.GeneralResponseDTO{
			OK:      false,
			Message: "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, dto.GeneralResponseDTO{
		OK:      true,
		Message: "Successfully get user data by id",
		Data:    user,
	})
}

func (uh *UserHandler) GetUserByUsername(c *gin.Context) {
	username := c.Param("username")
	user, err := uh.UseCase.GetByUsername(username)
	if err != nil {
		c.JSON(http.StatusNotFound, dto.GeneralResponseDTO{
			OK:      false,
			Message: "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, dto.GeneralResponseDTO{
		OK:      true,
		Message: "Successfully get user data by username",
		Data:    user,
	})
}
