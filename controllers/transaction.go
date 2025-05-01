package controllers

import (
	"encoding/json"
	"gorm.io/gorm"
	"net/http"
	"time"
	"wallet-api/database"
	"wallet-api/models"
	"wallet-api/utils"
)

func Transfer(w http.ResponseWriter, r *http.Request) {
	type TransferRequest struct {
		FromWalletID uint    `json:"from_wallet_id"`
		ToWalletID   uint    `json:"to_wallet_id"`
		Amount       float64 `json:"amount"`
	}

	var req TransferRequest
	json.NewDecoder(r.Body).Decode(&req)

	if req.Amount <= 0 || req.FromWalletID == req.ToWalletID {
		utils.RespondJSON(w, http.StatusBadRequest, "Invalid transfer request")
		return
	}

	var from models.Wallet
	var to models.Wallet

	if err := database.DB.First(&from, req.FromWalletID).Error; err != nil {
		utils.RespondJSON(w, http.StatusNotFound, "Sender wallet not found")
		return
	}
	if err := database.DB.First(&to, req.ToWalletID).Error; err != nil {
		utils.RespondJSON(w, http.StatusNotFound, "Receiver wallet not found")
		return
	}
	if from.Balance < req.Amount {
		utils.RespondJSON(w, http.StatusBadRequest, "Insufficient balance")
		return
	}

	database.DB.Transaction(func(tx *gorm.DB) error {
		from.Balance -= req.Amount
		to.Balance += req.Amount
		tx.Save(&from)
		tx.Save(&to)

		transaction := models.Transaction{
			FromWalletID:    from.ID,
			ToWalletID:      to.ID,
			Amount:          req.Amount,
			TransactionTime: time.Now(),
		}
		tx.Create(&transaction)
		return nil
	})

	utils.RespondJSON(w, http.StatusOK, "Transfer successful")
}

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	type Req struct {
		WalletID uint `json:"wallet_id"`
	}
	var req Req
	json.NewDecoder(r.Body).Decode(&req)

	var transactions []models.Transaction
	database.DB.Where("from_wallet_id = ? OR to_wallet_id = ?", req.WalletID, req.WalletID).
		Order("transaction_time desc").
		Find(&transactions)

	utils.RespondJSON(w, http.StatusOK, transactions)
}
