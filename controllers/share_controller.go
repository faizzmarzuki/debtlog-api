package controllers

import (
	"net/http"
	"time"

	"github.com/faizzmarzuki/debtlog-api/config"
	"github.com/faizzmarzuki/debtlog-api/models"
	"github.com/gin-gonic/gin"
)

// HandleShareLink is the public handler for the /share/:token route
func HandleShareLink(c *gin.Context) {
	token := c.Param("token")
	var link models.DebtLink
	if err := config.DB.Where("token = ?", token).First(&link).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "invalid link"})
		return
	}
	if link.ExpiresAt != nil && link.ExpiresAt.Before(time.Now()) {
		c.JSON(http.StatusGone, gin.H{"error": "link expired"})
		return
	}
	// load debt log and return a minimal payload for the front-end
	var dl models.DebtLog
	config.DB.First(&dl, link.DebtLogID)
	var items []models.DebtLogDebter
	config.DB.Where("debt_log_id = ?", dl.ID).Find(&items)
	c.JSON(http.StatusOK, gin.H{"debt_log": dl, "items": items})
}
