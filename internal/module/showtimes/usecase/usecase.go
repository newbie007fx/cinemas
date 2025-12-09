package usecase

import (
	"context"
	"time"

	"github.com/newbie007fx/cinemas/internal/errors"
	"github.com/newbie007fx/cinemas/internal/module/showtimes/entities"
	"github.com/newbie007fx/cinemas/internal/module/showtimes/models"
	"github.com/newbie007fx/cinemas/internal/module/showtimes/ports"
)

const (
	dateLayout = "2006-01-02"
	timeLayout = "15:04"
)

type Usecase struct {
	repo ports.Repository
}

func New(repo ports.Repository) ports.Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (u Usecase) Create(ctx context.Context, input models.CreateShowtimeInput) (*models.Showtime, *errors.BaseError) {
	entity, err := u.buildEntityFromInput(input, 0)
	if err != nil {
		return nil, err
	}

	if err := u.repo.ValidateReferences(ctx, entity.TheaterID, entity.MovieID); err != nil {
		return nil, err
	}

	created, err := u.repo.Create(ctx, *entity)
	if err != nil {
		return nil, err
	}

	return u.mapEntityToModel(created, nil), nil
}

func (u Usecase) List(ctx context.Context) ([]models.Showtime, *errors.BaseError) {
	showtimes, err := u.repo.List(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]models.Showtime, 0, len(showtimes))
	for i := range showtimes {
		result = append(result, *u.mapEntityToModel(&showtimes[i], nil))
	}

	return result, nil
}

func (u Usecase) GetByID(ctx context.Context, id uint) (*models.Showtime, *errors.BaseError) {
	showtime, err := u.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	inventories, err := u.repo.ListSeatInventories(ctx, id)
	if err != nil {
		return nil, err
	}

	return u.mapEntityToModel(showtime, inventories), nil
}

func (u Usecase) Update(ctx context.Context, id uint, input models.UpdateShowtimeInput) (*models.Showtime, *errors.BaseError) {
	entity, err := u.buildEntityFromInput(models.CreateShowtimeInput{
		TheaterID: input.TheaterID,
		MovieID:   input.MovieID,
		ShowDate:  input.ShowDate,
		StartTime: input.StartTime,
		EndTime:   input.EndTime,
		Price:     input.Price,
	}, id)
	if err != nil {
		return nil, err
	}

	if err := u.repo.ValidateReferences(ctx, entity.TheaterID, entity.MovieID); err != nil {
		return nil, err
	}

	updated, err := u.repo.Update(ctx, *entity)
	if err != nil {
		return nil, err
	}

	return u.mapEntityToModel(updated, nil), nil
}

func (u Usecase) Delete(ctx context.Context, id uint) *errors.BaseError {
	return u.repo.Delete(ctx, id)
}

func (Usecase) buildEntityFromInput(input models.CreateShowtimeInput, id uint) (*entities.Showtime, *errors.BaseError) {
	showDate, err := time.Parse(dateLayout, input.ShowDate)
	if err != nil {
		return nil, errors.ErrorValidation.New("show_date must use format YYYY-MM-DD")
	}

	startTime, err := time.Parse(timeLayout, input.StartTime)
	if err != nil {
		return nil, errors.ErrorValidation.New("start_time must use format HH:MM")
	}

	endTime, err := time.Parse(timeLayout, input.EndTime)
	if err != nil {
		return nil, errors.ErrorValidation.New("end_time must use format HH:MM")
	}

	if !endTime.After(startTime) {
		return nil, errors.ErrorValidation.New("end_time must be after start_time")
	}

	return &entities.Showtime{
		ID:        id,
		TheaterID: input.TheaterID,
		MovieID:   input.MovieID,
		ShowDate:  showDate,
		StartTime: startTime,
		EndTime:   endTime,
		Price:     input.Price,
	}, nil
}

func (Usecase) mapEntityToModel(entity *entities.Showtime, inventories []entities.SeatInventory) *models.Showtime {
	showtime := &models.Showtime{
		ID:        entity.ID,
		TheaterID: entity.TheaterID,
		MovieID:   entity.MovieID,
		ShowDate:  entity.ShowDate.Format(dateLayout),
		StartTime: entity.StartTime.Format(timeLayout),
		EndTime:   entity.EndTime.Format(timeLayout),
		Price:     entity.Price,
	}

	if len(inventories) > 0 {
		showtime.SeatInventories = make([]models.SeatInventory, 0, len(inventories))
		for _, inv := range inventories {
			showtime.SeatInventories = append(showtime.SeatInventories, models.SeatInventory{
				ID:         inv.ID,
				SeatID:     inv.SeatID,
				SeatRow:    inv.SeatRow,
				SeatColumn: inv.SeatColumn,
				Status:     inv.Status,
			})
		}
	}

	return showtime
}
