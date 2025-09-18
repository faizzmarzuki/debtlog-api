package config // config package holds shared app configuration like DB handle

import (
	"fmt"  // format strings
	"log"  // logging helper
	"os"   // read env vars
	"time" // connection pooling durations

	"gorm.io/driver/postgres" // Postgres driver for GORM
	"gorm.io/gorm"            // GORM ORM
)

var DB *gorm.DB // exported DB handle used across the app

func ConnectDatabase() {
	// prefer a full DSN if provided (useful for production envs)
	dsn := os.Getenv("DATABASE_DSN") // e.g. "host=... user=... password=... dbname=..."
	if dsn == "" {
		host := getenv("DB_HOST", "localhost")
		port := getenv("DB_PORT", "5432")
		user := getenv("DB_USER", "debtuser")
		password := getenv("DB_PASSWORD", "yourpassword")
		dbname := getenv("DB_NAME", "debtlogdb")
		tz := getenv("TIMEZONE", "UTC")
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
			host, user, password, dbname, port, tz) // build the DSN string
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}) // connect using GORM + Postgres driver
	if err != nil {
		log.Fatalf("failed to connect database: %v", err) // stop if DB unreachable
	}

	// configure underlying *sql.DB connection pool for production-friendly defaults
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get sql.DB: %v", err)
	}
	sqlDB.SetMaxOpenConns(25)                 // max open connections
	sqlDB.SetMaxIdleConns(25)                 // max idle connections
	sqlDB.SetConnMaxLifetime(5 * time.Minute) // connection max lifetime

	DB = db                           // assign global DB handle
	log.Println("database connected") // friendly startup log
}

// getenv returns the environment variable value or a fallback if not set
func getenv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
