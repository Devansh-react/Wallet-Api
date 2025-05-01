package models

import "time"

type Transaction struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	FromWalletID    uint      `json:"from_wallet_id"`
	ToWalletID      uint      `json:"to_wallet_id"`
	Amount          float64   `gorm:"not null" json:"amount"`
	TransactionTime time.Time `json:"transaction_time"`
}
