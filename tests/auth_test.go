package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/faizzmarzuki/debtlog-api/config"
	"github.com/faizzmarzuki/debtlog-api/models"
	"github.com/faizzmarzuki/debtlog-api/routes"
	"github.com/gin-gonic/gin"
)

// helper: build a test router with real DB connection
func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)

	// Connect to your real DB
	config.ConnectDatabase()

	// Auto-migrate to ensure tables exist
	config.DB.AutoMigrate(&models.User{}, &models.Debter{}, &models.DebtLog{}, &models.DebtLogDebter{}, &models.Receipt{}, &models.DebtLink{})

	r := gin.Default()
	routes.SetupRouter(r)
	return r
}

// Test user registration
func TestRegister(t *testing.T) {
	router := setupRouter()

	// Sample request payload
	body := map[string]string{
		"name":     "Test User",
		"email":    "testuser@example.com",
		"password": "secret123",
	}
	jsonBody, _ := json.Marshal(body)

	// Build HTTP request
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	// Record response
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Check response
	if w.Code != http.StatusOK && w.Code != http.StatusCreated {
		t.Errorf("expected status 200/201, got %d", w.Code)
	}
}

// Test login (depends on TestRegister having run first)
func TestLogin(t *testing.T) {
	router := setupRouter()

	body := map[string]string{
		"email":    "testuser@example.com",
		"password": "secret123",
	}
	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
}
