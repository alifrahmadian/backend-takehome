package utils

import (
	"app/internal/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	UserID int64  `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateToken(user *models.User, secretKey string, ttl int) (string, error) {
	jwtSecret := []byte(secretKey)

	claims := Claims{
		Name:   user.Name,
		Email:  user.Email,
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(ttl))),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
