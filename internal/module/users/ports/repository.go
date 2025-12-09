package ports

import (
	"context"

	"github.com/newbie007fx/cinemas/internal/errors"
	"github.com/newbie007fx/cinemas/internal/module/users/entities"
)

type Repository interface {
	FindByUsername(ctx context.Context, username string) (*entities.User, *errors.BaseError)
}
