package repository

import (
	"database/sql"
	"rtforum/config"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite", config.DATABASE)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
