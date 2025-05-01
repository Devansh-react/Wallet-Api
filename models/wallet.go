package models

import "time"

type Wallet struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	UserID    uint    `gorm:"unique;not null" json:"user_id"`
	Balance   float64 `gorm:"not null" json:"balance"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
