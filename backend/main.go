package main

import (
	"log"

	"github.com/joho/godotenv"

	"DuyrepWebsiteBackend/internal/api"
	"DuyrepWebsiteBackend/internal/database"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, ", err)
	}

  database.Connect()
  api.Run()
}
