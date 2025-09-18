package controllers

import (
	"net/http"
	"time"

	"github.com/faizzmarzuki/debtlog-api/config" // DB handle
	"github.com/faizzmarzuki/debtlog-api/models" // models
	"github.com/faizzmarzuki/debtlog-api/utils"  // jwt helper (replace path)
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt" // for password hashing
)

// Register request payload
type registerPayload struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// Register creates a new user with hashed password
func Register(c *gin.Context) {
	var payload registerPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}

	user := models.User{Name: payload.Name, Email: payload.Email, Password: string(hash)} // create model
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	// return created user (without password)
	c.JSON(http.StatusCreated, gin.H{"user": gin.H{"id": user.ID, "name": user.Name, "email": user.Email, "created_at": time.Now()}})
}

// Login payload
type loginPayload struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// Login authenticates a user and returns a JWT
func Login(c *gin.Context) {
	var payload loginPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := config.DB.Where("email = ?", payload.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	// compare password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	// create JWT
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "login successful",
		"token":   token, // ✅ use token so it's not unused
	})
}
