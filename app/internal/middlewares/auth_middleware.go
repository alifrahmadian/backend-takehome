package middlewares

import (
	"app/internal/handlers/responses"
	e "app/pkg/errors"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	UserID int64  `json:"user_id"`
	jwt.RegisteredClaims
}

func AuthMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			responses.ErrorResponse(c, http.StatusUnauthorized, e.ErrNoAuthorizationHeader.Error())
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			responses.ErrorResponse(c, http.StatusUnauthorized, e.ErrInvalidTokenFormat.Error())
			return
		}

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}

			return []byte(secretKey), nil
		})

		if err != nil || !token.Valid {
			responses.ErrorResponse(c, http.StatusUnauthorized, e.ErrTokenInvalid.Error())
			c.Abort()
			return
		}

		if claims.ExpiresAt.Time.Before(time.Now()) {
			responses.ErrorResponse(c, http.StatusUnauthorized, e.ErrTokenExpired.Error())
			c.Abort()
			return
		}

		c.Set("name", claims.Name)
		c.Set("user_id", claims.UserID)
		c.Set("email", claims.Email)

		c.Next()
	}
}
