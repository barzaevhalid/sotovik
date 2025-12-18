package utils

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userId int64, role string) (string, error) {
	claims := jwt.MapClaims{
		"sub":  strconv.FormatInt(userId, 10),
		"role": role,
		"exp":  time.Now().Add(2 * time.Second).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func GenerateRefreshToken(userId int64) (string, error) {
	claims := jwt.MapClaims{
		"sub": strconv.FormatInt(userId, 10),
		"exp": time.Now().Add(30 * 24 * time.Hour).Unix(), // 30 дней
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
}
