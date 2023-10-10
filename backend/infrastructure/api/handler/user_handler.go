package handler

import (
	"log"

	"github.com/Montekrysto/ecomm/app/usecase"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
	userUseCase usecase.UserUseCase
}

func NewUserHandler(u *usecase.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: *u,
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	// Your code to handle user creation here
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	log.Printf("Getting Users")
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))

	user, err := h.userUseCase.GetByID(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, user)
}
