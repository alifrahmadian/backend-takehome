package handlers

import (
	"app/internal/handlers/dtos"
	"app/internal/handlers/responses"
	"app/internal/models"
	"app/internal/services"
	"app/pkg/errors"
	"app/pkg/messages"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthHandler struct {
	AuthService services.AuthService
	SecretKey   string
	TTL         int
}

func NewAuthHandler(authService *services.AuthService, secretKey string, ttl int) *AuthHandler {
	return &AuthHandler{
		AuthService: *authService,
		SecretKey:   secretKey,
		TTL:         ttl,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req dtos.RegisterRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Field() {
			case "Name":
				responses.ErrorResponse(c, http.StatusBadRequest, errors.ErrNameRequired.Error())
				return
			case "Email":
				responses.ErrorResponse(c, http.StatusBadRequest, errors.ErrEmailRequired.Error())
				return
			case "Password":
				responses.ErrorResponse(c, http.StatusBadRequest, errors.ErrPasswordRequired.Error())
				return
			}
		}
	}

	user := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	err = h.AuthService.Register(user)
	if err != nil {
		if err == errors.ErrEmailExist {
			responses.ErrorResponse(c, http.StatusConflict, err.Error())
			return
		}

		responses.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	responses.SuccessResponse(c, messages.MsgRegisterSuccess, nil)
}
