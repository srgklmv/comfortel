package repository

import (
	"context"
	"database/sql"
	"errors"
)

type repository struct {
	conn *sql.DB
}

func New(conn *sql.DB) *repository {
	return &repository{conn: conn}
}

func getTxFromContext(ctx context.Context) (*sql.Tx, error) {
	tx, ok := ctx.Value("tx").(*sql.Tx)
	if !ok {
		return nil, errors.New("no tx in context")
	}

	return tx, nil
}
