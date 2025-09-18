package routes

import (
	"github.com/faizzmarzuki/debtlog-api/controllers" // controllers package
	"github.com/faizzmarzuki/debtlog-api/middleware"  // JWT auth middleware
	"github.com/gin-gonic/gin"                        // Gin framework
)

// RegisterRoutes takes an existing Gin router and attaches all routes.
// This is cleaner because main.go owns the Gin engine, not the routes package.
func SetupRouter(r *gin.Engine) {

	// --- Public routes (no auth required) ---

	// User auth
	r.POST("/register", controllers.Register) // Register a new user
	r.POST("/login", controllers.Login)       // Login with username/password

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Shareable link for debtors (public)
	r.GET("/share/:token", controllers.HandleShareLink)

	// --- Protected routes (require JWT auth) ---
	auth := r.Group("/")                  // Create a group for protected routes
	auth.Use(middleware.AuthMiddleware()) // Attach JWT middleware

	// Debter CRUD
	auth.POST("/debters", controllers.CreateDebter)       // Create new debter
	auth.GET("/debters", controllers.ListDebters)         // List all debters
	auth.PUT("/debters/:id", controllers.UpdateDebter)    // Update debter by ID
	auth.DELETE("/debters/:id", controllers.DeleteDebter) // Delete debter by ID

	// Debt logs
	auth.POST("/debtlogs", controllers.CreateDebtLog) // Create debt log
	auth.GET("/debtlogs/:id", controllers.GetDebtLog) // Get debt log details

	// Receipts (nested under debtlog)
	auth.POST("/debtlogs/:id/receipts", controllers.UploadReceipt) // Upload receipt
}
