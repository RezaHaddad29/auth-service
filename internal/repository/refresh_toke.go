package repository

import (
	"context"

	"github.com/RezaHaddad29/auth-service/internal/model"
	"github.com/jackc/pgx/v4/pgxpool"
)

type RefreshTokenRepository interface {
	Save(ctx context.Context, token model.RefreshToken) error
	FindByToken(ctx context.Context, token string) (*model.RefreshToken, error)
	DeleteByToken(ctx context.Context, token string) error
}

type refreshTokenRepo struct {
	db *pgxpool.Pool
}

func NewRefreshTokenRepository(db *pgxpool.Pool) RefreshTokenRepository {
	return &refreshTokenRepo{db}
}

func (r *refreshTokenRepo) Save(ctx context.Context, token model.RefreshToken) error {
	_, err := r.db.Exec(ctx, "INSERT INTO refresh_tokens (user_id, token, expires_at) VALUES ($1, $2, $3)", token.UserID, token.Token, token.ExpiresAt)
	return err
}

func (r *refreshTokenRepo) FindByToken(ctx context.Context, token string) (*model.RefreshToken, error) {
	row := r.db.QueryRow(ctx, "SELECT id, user_id, token, expires_at FROM refresh_tokens WHERE token = $1", token)
	var rt model.RefreshToken
	err := row.Scan(&rt.ID, &rt.UserID, &rt.Token, &rt.ExpiresAt)
	if err != nil {
		return nil, err
	}
	return &rt, nil
}

func (r *refreshTokenRepo) DeleteByToken(ctx context.Context, token string) error {
	_, err := r.db.Exec(ctx, "DELETE FROM refresh_tokens WHERE token = $1", token)
	return err
}
