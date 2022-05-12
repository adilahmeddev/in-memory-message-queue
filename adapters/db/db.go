package db

import (
	"database/sql"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
)

type db struct {
	*sql.DB
}

func NewDB() (*db, error) {
	pgxConfig, err := pgx.ParseConfig("postgres://postgres:password@localhost:5432/event-notif")
	if err != nil {
		return nil, err
	}

	return &db{DB: stdlib.OpenDB(*pgxConfig)}, err
}

func (d *db) Connect() error {

	return nil
}
