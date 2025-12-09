package database

import (
	"fmt"
	"log/slog"

	"github.com/newbie007fx/cinemas/platform/configuration"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DatabaseService struct {
	*sqlx.DB
	cs *configuration.ConfigService
}

func New(cs *configuration.ConfigService) *DatabaseService {
	return &DatabaseService{
		cs: cs,
	}
}

func (ds DatabaseService) getConectionString() string {
	config := ds.cs.GetConfig()
	configDb := config.Database

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", configDb.Host, configDb.Port, configDb.User, configDb.Password, configDb.Name)
}

func (ds *DatabaseService) Setup() (err error) {
	connstring := ds.getConectionString()

	ds.DB, err = sqlx.Connect("postgres", connstring)
	if err != nil {
		slog.Error(fmt.Sprint(err.Error()))
	}

	return
}

func (ds *DatabaseService) Shutdown() {
	ds.DB.Close()
}
