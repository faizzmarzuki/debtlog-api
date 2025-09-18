package controllers

import (
	"net/http"
	"time"

	"github.com/faizzmarzuki/debtlog-api/config"
	"github.com/faizzmarzuki/debtlog-api/models"
	"github.com/faizzmarzuki/debtlog-api/utils"
	"github.com/gin-gonic/gin"
)

// CreateDebtLog payload simplified
type createDebtLogPayload struct {
	Title       string  `json:"title" binding:"required"`
	TotalAmount float64 `json:"total_amount" binding:"required"`
	DebterIDs   []uint  `json:"debter_ids" binding:"required"` // list of debter IDs to attach
}

// CreateDebtLog creates a debt log and associates debters in the join table
func CreateDebtLog(c *gin.Context) {
	userID := c.GetUint("user_id")
	var payload createDebtLogPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dl := models.DebtLog{UserID: userID, Title: payload.Title, TotalAmount: payload.TotalAmount, Status: "pending", CreatedAt: time.Now(), UpdatedAt: time.Now()}
	if err := config.DB.Create(&dl).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create debt log"})
		return
	}

	// Distribute equally for now
	share := payload.TotalAmount / float64(len(payload.DebterIDs))
	for _, did := range payload.DebterIDs {
		d := models.DebtLogDebter{DebtLogID: dl.ID, DebterID: did, AmountDue: share, AmountPaid: 0, Status: "unpaid", CreatedAt: time.Now(), UpdatedAt: time.Now()}
		config.DB.Create(&d)
	}

	// generate a share token (simple random string) - in production use crypto/rand
	token := utils.GenerateTokenString()
	expires := time.Now().Add(7 * 24 * time.Hour)
	dlLink := models.DebtLink{DebtLogID: dl.ID, Token: token, ExpiresAt: &expires, CreatedAt: time.Now()}
	config.DB.Create(&dlLink)

	c.JSON(http.StatusCreated, gin.H{"debt_log": dl, "share_token": token})
}

// GetDebtLog returns debt log with its debters and statuses
func GetDebtLog(c *gin.Context) {
	id := c.Param("id")
	var dl models.DebtLog
	if err := config.DB.First(&dl, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "debt log not found"})
		return
	}
	// load join rows and receipts
	var items []models.DebtLogDebter
	config.DB.Where("debt_log_id = ?", dl.ID).Preload("Receipts").Find(&items)
	c.JSON(http.StatusOK, gin.H{"debt_log": dl, "items": items})
}
