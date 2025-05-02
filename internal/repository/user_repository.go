package repository

import (
	"context"

	"github.com/RezaHaddad29/auth-service/internal/model"
	"github.com/jackc/pgx/v4/pgxpool"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user model.User) error
	GetUserByUsername(ctx context.Context, username string) (model.User, error)
}

type userRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(ctx context.Context, user model.User) error {
	query := `INSERT INTO users (username, password, email) VALUES ($1, $2, $3)`
	_, err := r.db.Exec(ctx, query, user.UserName, user.Password, user.Email)
	return err
}

func (r *userRepository) GetUserByUsername(ctx context.Context, username string) (model.User, error) {
	query := `SELECT id, username, password, email FROM users WHERE username = $1`
	row := r.db.QueryRow(ctx, query, username)

	var user model.User
	err := row.Scan(&user.ID, &user.UserName, &user.Password, &user.Email)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
