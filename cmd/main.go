package main

import (
	"log"

	"github.com/Shahrzad-Taherzadeh/cinemaTicket/internal/api"
)

func main() {
	app := api.NewServer()

	log.Println("Server is running on :8080")

	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}