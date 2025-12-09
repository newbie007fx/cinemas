package migration

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

const migrationLocation = "platform/database/migration/files"

func New(db *sql.DB) *MigrationService {
	return &MigrationService{
		DB: db,
	}
}

type MigrationService struct {
	migrator *migrate.Migrate
	DB       *sql.DB
}

func (ds *MigrationService) Setup() (err error) {
	driver, err := postgres.WithInstance(ds.DB, &postgres.Config{})
	if err != nil {
		return err
	}

	ds.migrator, err = migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", migrationLocation),
		"postgres", driver)

	return
}

func (ds *MigrationService) Migrate() error {
	if err := ds.migrator.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			version, _, _ := ds.migrator.Version()
			fmt.Printf("No new migration found. version %d is the latest version\n", version)

			return nil
		}

		return err
	}

	fmt.Println("Migration has running succesfully")

	return nil
}

func (ds *MigrationService) ForceMigrate(version int) error {
	fmt.Printf("Will running force migration for version %d\n", version)
	if err := ds.migrator.Force(version); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			version, _, _ := ds.migrator.Version()
			fmt.Printf("No new migration found. version %d is the latest version\n", version)

			return nil
		}

		return err
	}

	fmt.Println("Migration has running succesfully")

	return nil
}

func (ds *MigrationService) Rollback() error {
	version, _, err := ds.migrator.Version()
	if err != nil {
		if errors.Is(err, migrate.ErrNilVersion) {
			fmt.Println("No active migration found.")
			return nil
		}
		return err
	}

	fmt.Printf("Rollingback migration with version %d will be execute\n", version)

	if err := ds.migrator.Steps(-1); err != nil {
		return err
	}

	fmt.Println("Rollback has running succesfully")

	return nil
}

func (ds *MigrationService) CreateMirationFiles(name string) (err error) {
	unixTime := time.Now().Unix()

	fileUp := fmt.Sprintf("%s/%d_%s.up.sql", migrationLocation, unixTime, name)
	fileDown := fmt.Sprintf("%s/%d_%s.down.sql", migrationLocation, unixTime, name)

	if _, err := os.Create(fileUp); err != nil {
		return err
	}

	if _, err := os.Create(fileDown); err != nil {
		os.Remove(fileUp)
		return err
	}

	fmt.Println("Migration files created successfully.")

	return nil
}

func (ds *MigrationService) Shutdown() {}
