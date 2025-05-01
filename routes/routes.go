package routes

import (
	"net/http"
	"wallet-api/controllers"
)

func RegisterRoutes() {
	http.HandleFunc("/user", controllers.CreateUser)
	http.HandleFunc("/wallet", controllers.CreateWallet)
	http.HandleFunc("/wallet/balance", controllers.GetBalance)
	http.HandleFunc("/wallet/transfer", controllers.Transfer)
	http.HandleFunc("/wallet/transactions", controllers.GetTransactions)
}
