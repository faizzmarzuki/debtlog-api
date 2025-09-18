package controllers

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/faizzmarzuki/debtlog-api/config"
	"github.com/faizzmarzuki/debtlog-api/models"
	"github.com/gin-gonic/gin"
)

// UploadReceipt accepts multipart/form-data: file and debter_id
func UploadReceipt(c *gin.Context) {
	userID := c.GetUint("user_id") // ensure user owns the debt log in a real app
	c.JSON(http.StatusOK, gin.H{
		"message": "receipt uploaded",
		"userID":  userID, // ✅ now it’s used
	})
	debtLogIDStr := c.Param("id")
	debtLogID, _ := strconv.ParseUint(debtLogIDStr, 10, 64)

	// form field: debter_id (which debter uploads)
	debterIDStr := c.PostForm("debter_id")
	debterID, _ := strconv.ParseUint(debterIDStr, 10, 64)

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file is required"})
		return
	}

	// simple local storage in ./uploads
	destination := fmt.Sprintf("uploads/%d_%s", time.Now().UnixNano(), filepath.Base(file.Filename))
	if err := c.SaveUploadedFile(file, destination); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save file"})
		return
	}

	// find the corresponding join row (debtlog+debter)
	var join models.DebtLogDebter
	if err := config.DB.Where("debt_log_id = ? AND debter_id = ?", debtLogID, debterID).First(&join).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid debter for this debt log"})
		return
	}

	receipt := models.Receipt{DebtLogDebterID: join.ID, FilePath: destination, Verified: "pending", CreatedAt: time.Now()}
	config.DB.Create(&receipt)

	// TODO: trigger email notification (enqueue job or send directly)

	c.JSON(http.StatusCreated, gin.H{"receipt": receipt})
}
