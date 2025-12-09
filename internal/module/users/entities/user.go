package entities

import (
	"time"
)

type User struct {
	ID        uint      `db:"id"`
	Name      string    `db:"name"`
	Username  string    `db:"username"`
	Password  string    `db:"password"`
	Email     string    `db:"email"`
	CreatedAt time.Time `db:"created_at"`
}
