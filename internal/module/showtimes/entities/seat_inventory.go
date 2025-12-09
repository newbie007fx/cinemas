package entities

type SeatInventory struct {
	ID         uint   `db:"id"`
	SeatID     uint   `db:"seat_id"`
	SeatRow    string `db:"seat_row"`
	SeatColumn int    `db:"seat_column"`
	Status     string `db:"status"`
}
