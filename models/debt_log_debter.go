package models

import "time"

// DebtLogDebter is the join table that links a DebtLog with multiple Debters
type DebtLogDebter struct {
	ID         uint      `gorm:"primaryKey" json:"id"`                                 // primary key
	DebtLogID  uint      `gorm:"index" json:"debt_log_id"`                             // FK to DebtLog
	DebterID   uint      `gorm:"index" json:"debter_id"`                               // FK to Debter
	AmountDue  float64   `json:"amount_due"`                                           // how much this debter owes
	AmountPaid float64   `json:"amount_paid"`                                          // how much has been paid
	Status     string    `gorm:"size:50;default:'unpaid'" json:"status"`               // unpaid|partial|paid
	CreatedAt  time.Time `json:"created_at"`                                           // created timestamp
	UpdatedAt  time.Time `json:"updated_at"`                                           // updated timestamp
	Receipts   []Receipt `gorm:"foreignKey:DebtLogDebterID" json:"receipts,omitempty"` // receipts uploaded by this debter
}
