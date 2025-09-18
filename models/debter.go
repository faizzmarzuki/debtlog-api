package models

import "time"

// Debter is a person who owes money in a debt log (e.g., Ali, Siti)
type Debter struct {
	ID        uint      `gorm:"primaryKey" json:"id"`                       // primary key
	UserID    uint      `gorm:"index" json:"user_id"`                       // owner (the app user who created this debter)
	Name      string    `gorm:"size:255" json:"name"`                       // debtor's name
	Phone     string    `gorm:"size:100;default:''" json:"phone,omitempty"` // optional phone
	Email     string    `gorm:"size:255;default:''" json:"email,omitempty"` // optional email
	Notes     string    `gorm:"type:text" json:"notes,omitempty"`           // free-form notes
	CreatedAt time.Time `json:"created_at"`                                 // created timestamp
	UpdatedAt time.Time `json:"updated_at"`                                 // updated timestamp
}
