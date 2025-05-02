package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/RezaHaddad29/auth-service/config"
	"github.com/RezaHaddad29/auth-service/internal/handler"
	"github.com/RezaHaddad29/auth-service/internal/middleware"
	"github.com/RezaHaddad29/auth-service/internal/repository"
	"github.com/RezaHaddad29/auth-service/internal/service"
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

	userRepo := repository.NewUserRepository(db.DB)
	refreshTokenRepo := repository.NewRefreshTokenRepository(db.DB)
	authSvc := service.NewAuthService(userRepo, refreshTokenRepo)
	authHandler := handler.NewAuthHandler(authSvc)

	http.HandleFunc("/api/register", authHandler.Register)
	http.HandleFunc("/api/login", authHandler.Login)
	http.HandleFunc("/api/profile", middleware.AuthMiddleware(authHandler.Profile))

	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
