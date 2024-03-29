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
		`CREATE TABLE IF NOT EXISTS Post (ID INTEGER PRIMARY KEY AUTOINCREMENT, Body TEXT, Title TEXT, User_ID TEXT, Date TEXT, CHECK(length(Title) <= 50));`,
		`CREATE TABLE IF NOT EXISTS Category(Post_ID INTEGER, Category_Name TEXT NOT NULL, FOREIGN KEY(Post_ID)REFERENCES Post(ID), CHECK(length(trim(Category_Name)) > 0));`,
		`CREATE TABLE IF NOT EXISTS Users (ID TEXT PRIMARY KEY, UserName TEXT, Age Text, Gender Text, First_Name TEXT, Last_Name TEXT, Mail TEXT, Password TEXT, UNIQUE(UserName), UNIQUE (Mail), CHECK(length(UserName) <= 15));`,
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
