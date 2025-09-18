package main // executable entry point for the application

import (
	"log" // simple logging
	"os"  // read environment variables

	"github.com/gin-gonic/gin" // web framework
	"github.com/joho/godotenv" // load .env into env vars (dev convenience)

	// Replace the module path below with your module from go.mod, e.g. github.com/mohammadfaiz/debtlog-api
	"github.com/faizzmarzuki/debtlog-api/config" // package that holds ConnectDatabase() and DB handle
	"github.com/faizzmarzuki/debtlog-api/models" // package containing GORM model definitions
	"github.com/faizzmarzuki/debtlog-api/routes" //package that contains all the endpoints
)

func main() {
	_ = godotenv.Load() // try loading .env silently (no error if missing)

	config.ConnectDatabase() // connect to Postgres and set config.DB

	// Auto-migrate all models: creates/updates DB tables to match structs
	if err := config.DB.AutoMigrate(
		&models.User{},
		&models.Debter{},
		&models.DebtLog{},
		&models.DebtLogDebter{},
		&models.Receipt{},
		&models.DebtLink{},
	); err != nil {
		log.Fatalf("AutoMigrate failed: %v", err) // fatal and exit if migrations fail
	}

	// Init Gin
	r := gin.Default()

	// Register routes
	routes.SetupRouter(r)

	port := os.Getenv("PORT") // allow overriding port via env
	if port == "" {
		port = "8080"
	} // default to 8080 in dev
	log.Printf("listening on :%s", port) // log startup port
	log.Fatal(r.Run(":" + port))         // start HTTP server and log fatal on error
}
