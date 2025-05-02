package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("pBLxECVq7GSk;Z}!@dFauA~bn6/shr^z9X,#2")

func GenerateToken(userID int64, userName string) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  userID,
		"userName": userName,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
