package ports

import (
	"context"

	"github.com/newbie007fx/cinemas/internal/errors"
	"github.com/newbie007fx/cinemas/internal/module/showtimes/entities"
)

type Repository interface {
	Create(ctx context.Context, showtime entities.Showtime) (*entities.Showtime, *errors.BaseError)
	GetByID(ctx context.Context, id uint) (*entities.Showtime, *errors.BaseError)
	List(ctx context.Context) ([]entities.Showtime, *errors.BaseError)
	Update(ctx context.Context, showtime entities.Showtime) (*entities.Showtime, *errors.BaseError)
	Delete(ctx context.Context, id uint) *errors.BaseError
	ValidateReferences(ctx context.Context, theaterID, movieID uint) *errors.BaseError
	ListSeatInventories(ctx context.Context, showtimeID uint) ([]entities.SeatInventory, *errors.BaseError)
}
