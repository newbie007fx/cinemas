package models

type Showtime struct {
	ID              uint            `json:"id"`
	TheaterID       uint            `json:"theater_id"`
	MovieID         uint            `json:"movie_id"`
	ShowDate        string          `json:"show_date"`
	StartTime       string          `json:"start_time"`
	EndTime         string          `json:"end_time"`
	Price           float64         `json:"price"`
	SeatInventories []SeatInventory `json:"seat_inventories,omitempty"`
}

type SeatInventory struct {
	ID         uint   `json:"id"`
	SeatID     uint   `json:"seat_id"`
	SeatRow    string `json:"seat_row"`
	SeatColumn int    `json:"seat_column"`
	Status     string `json:"status"`
}

type CreateShowtimeInput struct {
	TheaterID uint
	MovieID   uint
	ShowDate  string
	StartTime string
	EndTime   string
	Price     float64
}

type UpdateShowtimeInput struct {
	TheaterID uint
	MovieID   uint
	ShowDate  string
	StartTime string
	EndTime   string
	Price     float64
}
