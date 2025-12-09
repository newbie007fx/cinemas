package showtimes

import (
	"github.com/newbie007fx/cinemas/internal/module/showtimes/models"
	"github.com/newbie007fx/cinemas/platform/validation"
)

type showtimePayload struct {
	TheaterID uint    `json:"theater_id" validate:"required"`
	MovieID   uint    `json:"movie_id" validate:"required"`
	ShowDate  string  `json:"show_date" validate:"required"`
	StartTime string  `json:"start_time" validate:"required"`
	EndTime   string  `json:"end_time" validate:"required"`
	Price     float64 `json:"price" validate:"required,gt=0"`
}

func (sp *showtimePayload) Validate() error {
	return validation.Validate(sp)
}

func (sp showtimePayload) toCreateInput() models.CreateShowtimeInput {
	return models.CreateShowtimeInput{
		TheaterID: sp.TheaterID,
		MovieID:   sp.MovieID,
		ShowDate:  sp.ShowDate,
		StartTime: sp.StartTime,
		EndTime:   sp.EndTime,
		Price:     sp.Price,
	}
}

func (sp showtimePayload) toUpdateInput() models.UpdateShowtimeInput {
	return models.UpdateShowtimeInput{
		TheaterID: sp.TheaterID,
		MovieID:   sp.MovieID,
		ShowDate:  sp.ShowDate,
		StartTime: sp.StartTime,
		EndTime:   sp.EndTime,
		Price:     sp.Price,
	}
}
