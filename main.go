package main

import (
	"log"
	"net/http"
	"wallet-api/config"
	"wallet-api/database"
	"wallet-api/routes"
)

func main() {
	config.LoadEnv()
	database.Connect()
	routes.RegisterRoutes()

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
