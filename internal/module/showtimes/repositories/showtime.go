package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/newbie007fx/cinemas/internal/errors"
	"github.com/newbie007fx/cinemas/internal/module/showtimes/entities"
	"github.com/newbie007fx/cinemas/internal/module/showtimes/ports"
	"github.com/newbie007fx/cinemas/platform/database"
)

const (
	seatInventoryStatusAvailable = "available"
)

type Repository struct {
	DB *database.DatabaseService
}

func New(db *database.DatabaseService) ports.Repository {
	return &Repository{
		DB: db,
	}
}

func (r Repository) Create(ctx context.Context, showtime entities.Showtime) (*entities.Showtime, *errors.BaseError) {
	tx, err := r.DB.BeginTxx(ctx, nil)
	if err != nil {
		return nil, errors.ErrorQueryDatabase.New(err.Error())
	}

	createQuery := `
		INSERT INTO showtimes ("theater_id", "movie_id", "show_date", "start_time", "end_time", "price")
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING "id", "theater_id", "movie_id", "show_date", "start_time", "end_time", "price"
	`

	var created entities.Showtime
	err = tx.QueryRowxContext(ctx, createQuery, showtime.TheaterID, showtime.MovieID, showtime.ShowDate, showtime.StartTime, showtime.EndTime, showtime.Price).StructScan(&created)
	if err != nil {
		tx.Rollback()
		return nil, errors.ErrorQueryDatabase.New(err.Error())
	}

	if err := r.createSeatInventories(ctx, tx, created.ID, created.TheaterID); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, errors.ErrorQueryDatabase.New(err.Error())
	}

	return &created, nil
}

func (r Repository) GetByID(ctx context.Context, id uint) (*entities.Showtime, *errors.BaseError) {
	query := `
		SELECT "id", "theater_id", "movie_id", "show_date", "start_time", "end_time", "price"
		FROM showtimes
		WHERE "id" = $1
	`

	var showtime entities.Showtime
	err := r.DB.GetContext(ctx, &showtime, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrorQueryNoRow.New("showtime not found")
		}
		return nil, errors.ErrorQueryDatabase.New(err.Error())
	}

	return &showtime, nil
}

func (r Repository) List(ctx context.Context) ([]entities.Showtime, *errors.BaseError) {
	query := `
		SELECT "id", "theater_id", "movie_id", "show_date", "start_time", "end_time", "price"
		FROM showtimes
		ORDER BY "show_date" ASC, "start_time" ASC
	`

	var showtimes []entities.Showtime
	err := r.DB.SelectContext(ctx, &showtimes, query)
	if err != nil {
		return nil, errors.ErrorQueryDatabase.New(err.Error())
	}

	return showtimes, nil
}

func (r Repository) Update(ctx context.Context, showtime entities.Showtime) (*entities.Showtime, *errors.BaseError) {
	query := `
		UPDATE showtimes
		SET "theater_id" = $1,
			"movie_id" = $2,
			"show_date" = $3,
			"start_time" = $4,
			"end_time" = $5,
			"price" = $6
		WHERE "id" = $7
		RETURNING "id", "theater_id", "movie_id", "show_date", "start_time", "end_time", "price"
	`

	var updated entities.Showtime
	err := r.DB.QueryRowxContext(ctx, query, showtime.TheaterID, showtime.MovieID, showtime.ShowDate, showtime.StartTime, showtime.EndTime, showtime.Price, showtime.ID).StructScan(&updated)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrorQueryNoRow.New("showtime not found")
		}

		return nil, errors.ErrorQueryDatabase.New(err.Error())
	}

	return &updated, nil
}

func (r Repository) Delete(ctx context.Context, id uint) *errors.BaseError {
	tx, err := r.DB.BeginTxx(ctx, nil)
	if err != nil {
		return errors.ErrorQueryDatabase.New(err.Error())
	}

	if _, err := tx.ExecContext(ctx, `DELETE FROM seat_inventories WHERE "showtime_id" = $1`, id); err != nil {
		tx.Rollback()
		return errors.ErrorQueryDatabase.New(err.Error())
	}

	res, err := tx.ExecContext(ctx, `DELETE FROM showtimes WHERE "id" = $1`, id)
	if err != nil {
		tx.Rollback()
		return errors.ErrorQueryDatabase.New(err.Error())
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		tx.Rollback()
		return errors.ErrorQueryDatabase.New(err.Error())
	}

	if rowsAffected == 0 {
		tx.Rollback()
		return errors.ErrorQueryNoRow.New("showtime not found")
	}

	if err := tx.Commit(); err != nil {
		return errors.ErrorQueryDatabase.New(err.Error())
	}

	return nil
}

func (r Repository) ValidateReferences(ctx context.Context, theaterID, movieID uint) *errors.BaseError {
	if err := r.ensureExists(ctx, "theaters", theaterID, "theater_id is invalid"); err != nil {
		return err
	}

	if err := r.ensureExists(ctx, "movies", movieID, "movie_id is invalid"); err != nil {
		return err
	}

	return nil
}

func (r Repository) ensureExists(ctx context.Context, table string, id uint, errMsg string) *errors.BaseError {
	query := fmt.Sprintf(`SELECT EXISTS(SELECT 1 FROM %s WHERE "id" = $1)`, table)

	var exists bool
	if err := r.DB.GetContext(ctx, &exists, query, id); err != nil {
		return errors.ErrorQueryDatabase.New(err.Error())
	}

	if !exists {
		return errors.ErrorValidation.New(errMsg)
	}

	return nil
}

func (r Repository) createSeatInventories(ctx context.Context, tx *sqlx.Tx, showtimeID, theaterID uint) *errors.BaseError {
	insertQuery := `
		INSERT INTO seat_inventories ("showtime_id", "seat_id", "status", "version")
		SELECT $1, "id", $2, 1
		FROM seats
		WHERE "theater_id" = $3
	`

	res, err := tx.ExecContext(ctx, insertQuery, showtimeID, seatInventoryStatusAvailable, theaterID)
	if err != nil {
		return errors.ErrorQueryDatabase.New(err.Error())
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return errors.ErrorQueryDatabase.New(err.Error())
	}

	if affected == 0 {
		return errors.ErrorValidation.New("selected theater has no seats configured")
	}

	return nil
}

func (r Repository) ListSeatInventories(ctx context.Context, showtimeID uint) ([]entities.SeatInventory, *errors.BaseError) {
	query := `
		SELECT si."id", si."seat_id", s."seat_row", s."seat_column", si."status"
		FROM seat_inventories si
		JOIN seats s ON s."id" = si."seat_id"
		WHERE si."showtime_id" = $1
		ORDER BY s."seat_row", s."seat_column", si."id"
	`

	var inventories []entities.SeatInventory
	if err := r.DB.SelectContext(ctx, &inventories, query, showtimeID); err != nil {
		return nil, errors.ErrorQueryDatabase.New(err.Error())
	}

	return inventories, nil
}
