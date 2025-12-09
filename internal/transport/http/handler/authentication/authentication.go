package authentication

import (
	"github.com/newbie007fx/cinemas/internal/module/users/ports"
	"github.com/newbie007fx/cinemas/internal/transport/http/helpers/authentication"
)

type AuthHandlers struct {
	authToken   *authentication.TokenAuth
	userUsecase ports.Usecase
}

func New(userUsecase ports.Usecase, authToken *authentication.TokenAuth) *AuthHandlers {
	return &AuthHandlers{
		userUsecase: userUsecase,
		authToken:   authToken,
	}
}
