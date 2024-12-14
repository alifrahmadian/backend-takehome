package configs

import (
	"app/internal/handlers"
	"database/sql"
)

type Config struct {
	DB       *sql.DB
	Auth     *AuthConfig
	Handlers *Handlers
	Env      string
}

type Handlers struct {
	AuthHandler *handlers.AuthHandler
}

type AuthConfig struct {
	TTL       int
	SecretKey string
}
