package controllers

import (
	"net/http"
	"strconv"

	"github.com/faizzmarzuki/debtlog-api/config"
	"github.com/faizzmarzuki/debtlog-api/models"
	"github.com/gin-gonic/gin"
)

// CreateDebter creates a new debter linked to the authenticated user
func CreateDebter(c *gin.Context) {
	userID := c.GetUint("user_id") // set by auth middleware
	var input models.Debter
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input.UserID = userID
	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create debter"})
		return
	}
	c.JSON(http.StatusCreated, input)
}

// ListDebters returns debters for the authenticated user
func ListDebters(c *gin.Context) {
	userID := c.GetUint("user_id")
	var debters []models.Debter
	config.DB.Where("user_id = ?", userID).Find(&debters)
	c.JSON(http.StatusOK, debters)
}

// UpdateDebter updates an existing debter (owner only)
func UpdateDebter(c *gin.Context) {
	userID := c.GetUint("user_id")
	idStr := c.Param("id")
	id, _ := strconv.ParseUint(idStr, 10, 64)
	var debter models.Debter
	if err := config.DB.First(&debter, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "debter not found"})
		return
	}
	if debter.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not allowed"})
		return
	}
	if err := c.ShouldBindJSON(&debter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&debter)
	c.JSON(http.StatusOK, debter)
}

// DeleteDebter soft-deletes a debter (hard delete for now)
func DeleteDebter(c *gin.Context) {
	userID := c.GetUint("user_id")
	idStr := c.Param("id")
	id, _ := strconv.ParseUint(idStr, 10, 64)
	var debter models.Debter
	if err := config.DB.First(&debter, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "debter not found"})
		return
	}
	if debter.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "not allowed"})
		return
	}
}
