package main

import (
	"github.com/luckswish/go-crud-app/internal/config"
	"github.com/luckswish/go-crud-app/internal/db"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	cfg := config.LoadConfig()
	db.Connect(cfg)
	db.RunMigrations(db.DB)
}
