package models

import "time"

// Receipt represents an uploaded file (image/pdf) as proof of payment
type Receipt struct {
	ID              uint      `gorm:"primaryKey" json:"id"`                      // primary key
	DebtLogDebterID uint      `gorm:"index" json:"debt_log_debter_id"`           // FK to DebtLogDebter
	FilePath        string    `gorm:"size:1024" json:"file_path"`                // where file is stored (S3 path or local path)
	Verified        string    `gorm:"size:20;default:'pending'" json:"verified"` // pending|approved|rejected
	CreatedAt       time.Time `json:"uploaded_at"`                               // upload timestamp
}
