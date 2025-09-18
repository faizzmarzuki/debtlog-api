package models

import "time"

// User model represents an application user (you) who manages debt logs
type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`              // primary key (auto increment)
	Name      string    `gorm:"size:255" json:"name"`              // display name
	Email     string    `gorm:"uniqueIndex;size:255" json:"email"` // unique login identifier
	Password  string    `gorm:"size:255" json:"-"`                 // hashed password, omit from JSON
	CreatedAt time.Time `json:"created_at"`                        // created timestamp
	UpdatedAt time.Time `json:"updated_at"`                        // updated timestamp
}
