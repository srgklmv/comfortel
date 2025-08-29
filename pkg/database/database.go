package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/srgklmv/comfortel/pkg/logger"
)

func New(host, port, database, user, password string) (*sql.DB, error) {
	data := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		database,
	)

	db, err := sql.Open("postgres", data)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %w", err)
	}

	err = db.QueryRowContext(context.Background(), "select 1;").Err()
	if err != nil {
		return nil, fmt.Errorf("db.QueryRowContext: %w", err)
	}

	return db, nil
}

func Shutdown(db *sql.DB) error {
	return db.Close()
}

func Migrate(db *sql.DB, path string, version int) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		logger.Error("migrations driver set up error", slog.String("error", err.Error()))
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		path,
		"postgres", driver)
	if err != nil {
		logger.Error("migrate instance creation error", slog.String("error", err.Error()))
		return err
	}

	err = m.Migrate(uint(version))
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		logger.Error("migrations up error", slog.String("error", err.Error()))
		return err
	}

	return nil
}
