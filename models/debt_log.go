package models

import "time"

// DebtLog is the parent record representing a bill or shared expense
type DebtLog struct {
	ID          uint       `gorm:"primaryKey" json:"id"`                    // primary key
	UserID      uint       `gorm:"index" json:"user_id"`                    // owner of the debt log
	Title       string     `gorm:"size:255" json:"title"`                   // human friendly title
	TotalAmount float64    `json:"total_amount"`                            // total amount for the log
	Status      string     `gorm:"size:50;default:'pending'" json:"status"` // overall status
	DueDate     *time.Time `json:"due_date,omitempty"`                      // optional due date
	Description string     `gorm:"type:text" json:"description,omitempty"`  // optional description
	CreatedAt   time.Time  `json:"created_at"`                              // created timestamp
	UpdatedAt   time.Time  `json:"updated_at"`                              // updated timestamp
	// Note: we keep the many-to-many join table explicit (DebtLogDebter)
}
