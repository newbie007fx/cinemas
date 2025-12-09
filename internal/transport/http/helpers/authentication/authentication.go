package authentication

import (
	"github.com/newbie007fx/cinemas/internal/module/users/models"
	"github.com/newbie007fx/cinemas/platform/validation"

	"github.com/golang-jwt/jwt/v4"
)

type LoginRequst struct {
	Username string `validate:"required" json:"username"`
	Password string `validate:"required" json:"password"`
}

func (lq *LoginRequst) Validate() error {
	return validation.Validate(lq)
}

type LoginResponse struct {
	AccessToken  string      `json:"access_token"`
	RefreshToken string      `json:"refresh_token"`
	User         models.User `json:"user"`
}

type TokenType string

type JwtCustomClaims struct {
	TokenType TokenType         `json:"token_type,omitempty"`
	Data      map[string]string `json:"data"`
	jwt.RegisteredClaims
}
