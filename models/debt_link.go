package models

import "time"

// DebtLink stores a short-lived tokenized link that a debter can open to submit receipt
type DebtLink struct {
	ID        uint       `gorm:"primaryKey" json:"id"`             // primary key
	DebtLogID uint       `gorm:"index" json:"debt_log_id"`         // FK to DebtLog
	Token     string     `gorm:"uniqueIndex;size:64" json:"token"` // secure random token used in the shareable URL
	ExpiresAt *time.Time `json:"expires_at,omitempty"`             // optional expiry
	CreatedAt time.Time  `json:"created_at"`                       // created timestamp
}
