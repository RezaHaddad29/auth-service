package service

import (
	"context"
	"fmt"
	"time"

	"github.com/RezaHaddad29/auth-service/internal/model"
	"github.com/RezaHaddad29/auth-service/internal/repository"
	"github.com/RezaHaddad29/auth-service/pkg/jwt"
	"github.com/google/uuid"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(ctx context.Context, user model.User) error
	AuthenticateUser(ctx context.Context, email, password string) (string, string, error)
}

type authService struct {
	userRepo         repository.UserRepository
	refreshTokenRepo repository.RefreshTokenRepository
}

func NewAuthService(userRepo repository.UserRepository, refreshTokenRepo repository.RefreshTokenRepository) AuthService {
	return &authService{userRepo: userRepo, refreshTokenRepo: refreshTokenRepo}
}

func (s *authService) Register(ctx context.Context, user model.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return s.userRepo.CreateUser(ctx, user)
}

func (s *authService) AuthenticateUser(ctx context.Context, userName, password string) (string, string, error) {
	user, err := s.userRepo.GetUserByUsername(ctx, userName)
	if err != nil {
		return "", "", fmt.Errorf("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", "", fmt.Errorf("invalid credentials")
	}

	accessToken, err := jwt.GenerateToken(user.ID, user.UserName)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate access token: %w", err)
	}

	refreshToken := uuid.NewString()
	expiresAt := time.Now().Add(7 * 24 * time.Hour)

	err = s.refreshTokenRepo.Save(ctx, model.RefreshToken{
		UserID:    user.ID,
		Token:     refreshToken,
		ExpiresAt: expiresAt,
	})
	if err != nil {
		return "", "", fmt.Errorf("failed to save refresh token: %w", err)
	}

	return accessToken, refreshToken, nil
}
