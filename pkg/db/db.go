package db

import (
	"context"
	"fmt"

	"github.com/RezaHaddad29/auth-service/config"
	"github.com/jackc/pgx/v4/pgxpool"
)

var DB *pgxpool.Pool

func ConnectDB(cfg *config.Config) error {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DbName)

	var err error
	DB, err = pgxpool.Connect(context.Background(), connString)
	if err != nil {
		return err
	}

	return nil
}

func CloseDB() {
	DB.Close()
}
