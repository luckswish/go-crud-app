package main

import (
	"github.com/joho/godotenv"
	"github.com/luckswish/go-crud-app/internal/config"
	"github.com/luckswish/go-crud-app/internal/db"
	"log"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cfg := config.LoadConfig()
	db.Connect(cfg)
}
