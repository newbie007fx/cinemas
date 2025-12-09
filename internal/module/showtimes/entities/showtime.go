package entities

import "time"

type Showtime struct {
	ID        uint      `db:"id"`
	TheaterID uint      `db:"theater_id"`
	MovieID   uint      `db:"movie_id"`
	ShowDate  time.Time `db:"show_date"`
	StartTime time.Time `db:"start_time"`
	EndTime   time.Time `db:"end_time"`
	Price     float64   `db:"price"`
}
