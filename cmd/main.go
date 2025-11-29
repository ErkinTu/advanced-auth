package main

import (
	"AdvAuthGo/internal/database"
	"log"
)

func main() {
	database.Connect()

	log.Println("Server started")
}
