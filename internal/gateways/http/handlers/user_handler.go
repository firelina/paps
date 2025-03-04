package handlers

import (
	"github.com/gin-gonic/gin"
	"marine/internal/models"
	"marine/internal/usecase"
	"net/http"
	"strconv"
)

type UserHandler struct {
	userUsecase usecase.UserService
}

func NewUserHandler(u usecase.UserService) *UserHandler {
	return &UserHandler{u}
}

func (h *UserHandler) RegisterUser(c *gin.Context) {
	var newUser models.CreateUserRequest
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := c.Request.Context()
	userID, err := h.userUsecase.RegisterUser(ctx, &models.User{
		Username: newUser.Username,
		Password: newUser.Password,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, models.CreateUserResponse{ID: userID})
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	ctx := c.Request.Context()
	user, err := h.userUsecase.GetUser(ctx, id)
	if err != nil {
		if err == usecase.ErrUserNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "Пользователь не найден."})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, user)
}
