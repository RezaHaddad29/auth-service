package main

import (
	"github.com/RezaHaddad29/auth-service/db"
	"github.com/RezaHaddad29/auth-service/migrations"
	"log"
)

func main() {
	db.ConnectDB()

	err := migrations.RunMigrations()
	if err != nil {
		log.Fatal("Migration failed: ", err)
	}

	log.Println("Migrations applied successfully")
}
