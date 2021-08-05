package repository

import (
	"database/sql"
	"log"
	"rtforum/config"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", config.DATABASE)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	tables := []string{`CREATE TABLE IF NOT EXISTS Post (Comment_ID INTEGER PRIMARY KEY, Post TEXT);`}

	for _, v := range tables {
		_, err = db.Exec(v)
		if err != nil {
			log.Fatalf("DB ERROR EXEC: %q\n%v", v, err.Error())
		}
	}
	return db, nil
}
