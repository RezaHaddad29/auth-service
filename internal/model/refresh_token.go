package model

import "time"

type RefreshToken struct {
	ID        int
	UserID    int64
	Token     string
	ExpiresAt time.Time
}
