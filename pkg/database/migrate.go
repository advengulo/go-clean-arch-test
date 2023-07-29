package database

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) error {

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	driver, err := postgres.WithInstance(sqlDB, &postgres.Config{})
	if err != nil {
		return err
	}

	// File source
	fsrc, err := (&file.File{}).Open("file://../go-clean-arch-test/pkg/database/migrations")
	if err != nil {
		return err
	}

	m, err := migrate.NewWithInstance("file", fsrc, "postgres", driver)
	if err != nil {
		return err
	}

	// Migrate to the latest version
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
