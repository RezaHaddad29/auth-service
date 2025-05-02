package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/RezaHaddad29/auth-service/internal/middleware"
	"github.com/RezaHaddad29/auth-service/internal/model"
	"github.com/RezaHaddad29/auth-service/internal/service"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req model.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	user := model.User{
		UserName: req.UserName,
		Password: req.Password,
		Email:    req.Email,
	}

	if err := h.authService.Register(r.Context(), user); err != nil {
		fmt.Println("Error creating user:", err)
		if err.Error() == "user already exists" {
			http.Error(w, "User already exists", http.StatusConflict)
			return
		}
		http.Error(w, "Could not create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User registered successfully"))
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req model.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	accessToken, refreshToken, err := h.authService.AuthenticateUser(r.Context(), req.UserName, req.Password)
	if err != nil {
		fmt.Println("Error authenticating user:", err)
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	resp := map[string]string{"access_token": accessToken, "refresh_token": refreshToken}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *AuthHandler) Profile(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserIDKey)
	userName := r.Context().Value(middleware.UserNameKey)

	resp := map[string]interface{}{
		"user_id":  userID,
		"userName": userName,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
