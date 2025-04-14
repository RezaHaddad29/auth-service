package main

import (
	"log"

	"github.com/RezaHaddad29/auth-service/config"
	"github.com/RezaHaddad29/auth-service/pkg/db"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("config load failed: %v", err)
	}

	err = db.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("Connect to DB failed: %v", err)
	}
	defer db.CloseDB()

	err = db.RunMigrations(cfg)
	if err != nil {
		log.Fatal("Migration failed: ", err)
	}

	log.Println("Migrations applied successfully")
}
