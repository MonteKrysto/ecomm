package handler

import (
	"log"

	"github.com/Montekrysto/ecomm/app/usecase"
	"github.com/Montekrysto/ecomm/domain/entity"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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
	var user entity.User

	// Parse JSON from request body
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Validate received data (simple example)
	if user.Email == "" || user.Password == "" || user.Name == "" {
		c.JSON(400, gin.H{"error": "Name, Email, and Password are required"})
		return
	}

	// Hashing the password before storing it to the database
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to encrypt password"})
		return
	}
	user.Password = string(hashedPassword)

	// Storing to the database
	err = h.userUseCase.Create(user)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"status": "User created successfully"})
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
