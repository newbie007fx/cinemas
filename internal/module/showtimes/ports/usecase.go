package ports

import (
	"context"

	"github.com/newbie007fx/cinemas/internal/errors"
	"github.com/newbie007fx/cinemas/internal/module/showtimes/models"
)

type Usecase interface {
	Create(ctx context.Context, input models.CreateShowtimeInput) (*models.Showtime, *errors.BaseError)
	List(ctx context.Context) ([]models.Showtime, *errors.BaseError)
	GetByID(ctx context.Context, id uint) (*models.Showtime, *errors.BaseError)
	Update(ctx context.Context, id uint, input models.UpdateShowtimeInput) (*models.Showtime, *errors.BaseError)
	Delete(ctx context.Context, id uint) *errors.BaseError
}
