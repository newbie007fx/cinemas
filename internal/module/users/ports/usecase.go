package ports

import (
	"context"

	"github.com/newbie007fx/cinemas/internal/errors"
	"github.com/newbie007fx/cinemas/internal/module/users/models"
)

type Usecase interface {
	VerifyUsernamePassword(ctx context.Context, username, password string) (*models.User, *errors.BaseError)
}
