package controllers

import (
	"encoding/json"
	"net/http"
	"wallet-api/database"
	"wallet-api/models"
	"wallet-api/utils"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	if user.Name == "" || user.Email == "" {
		utils.RespondJSON(w, http.StatusBadRequest, "All feilds  are required")
		return
	}

	var existing models.User
	if err := database.DB.Where("email = ?", user.Email).First(&existing).Error; err == nil {
		utils.RespondJSON(w, http.StatusConflict, "User already exists")
		return
	}

	database.DB.Create(&user)
	utils.RespondJSON(w, http.StatusCreated, user)
}

func CreateWallet(w http.ResponseWriter, r *http.Request) {
	type WalletRequest struct {
		UserID        uint    `json:"user_id"`
		InitialAmount float64 `json:"initial_balance"`
	}

	var req WalletRequest
	json.NewDecoder(r.Body).Decode(&req)

	if req.UserID == 0 || req.InitialAmount < 0 {
		utils.RespondJSON(w, http.StatusBadRequest, "Invalid input")
		return
	}

	var user models.User
	if err := database.DB.First(&user, req.UserID).Error; err != nil {
		utils.RespondJSON(w, http.StatusNotFound, "User not found")
		return
	}

	var wallet models.Wallet
	if err := database.DB.Where("user_id = ?", req.UserID).First(&wallet).Error; err == nil {
		utils.RespondJSON(w, http.StatusBadRequest, "Wallet already exists")
		return
	}

	newWallet := models.Wallet{UserID: req.UserID, Balance: req.InitialAmount}
	database.DB.Create(&newWallet)
	utils.RespondJSON(w, http.StatusCreated, newWallet)
}

func GetBalance(w http.ResponseWriter, r *http.Request) {
	type Req struct {
		UserID uint `json:"user_id"`
	}
	var req Req
	json.NewDecoder(r.Body).Decode(&req)

	var wallet models.Wallet
	if err := database.DB.Where("user_id = ?", req.UserID).First(&wallet).Error; err != nil {
		utils.RespondJSON(w, http.StatusNotFound, "Wallet not found")
		return
	}

	utils.RespondJSON(w, http.StatusOK, map[string]float64{"balance": wallet.Balance})
}
