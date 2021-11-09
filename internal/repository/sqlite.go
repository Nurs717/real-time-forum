package repository

import (
	"database/sql"
	"log"
	"rtforum/config"
	"time"

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
	tables := []string{
		`PRAGMA foreign_keys = ON;`,
		`CREATE TABLE IF NOT EXISTS Post (ID INTEGER PRIMARY KEY AUTOINCREMENT, Post TEXT);`,
		`CREATE TABLE IF NOT EXISTS Users (ID TEXT PRIMARY KEY, UserName TEXT, First_Name TEXT, Last_Name TEXT, Mail TEXT, Password TEXT);`,
		`CREATE TABLE IF NOT EXISTS Session (Token TEXT PRIMARY KEY, Expired_Date DATETIME, User_ID TEXT, FOREIGN KEY(User_ID)REFERENCES Users(ID));`,
	}

	for _, v := range tables {
		_, err = db.Exec(v)
		if err != nil {
			log.Fatalf("DB ERROR EXEC: %q\n%v", v, err.Error())
		}
	}
	go deleteExpiredSessions(db)
	return db, nil
}

func deleteExpiredSessions(db *sql.DB) {
	ticker := time.NewTicker(5 * time.Minute)
	for range ticker.C {
		_, err := db.Exec("DELETE FROM Session WHERE Expired_Date < DATETIME('now')")
		if err != nil {
			log.Printf("error while deleting expired sessions: %v", err)
		}
	}
}
