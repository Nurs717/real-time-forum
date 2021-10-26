package repository

import (
	"database/sql"
	"log"
	"rtforum/errors"
	"rtforum/internal/entity"
	"time"
)

type UsersRepo struct {
	db *sql.DB
}

func NewUsersRepo(db *sql.DB) *UsersRepo {
	return &UsersRepo{
		db: db,
	}
}

func (r *UsersRepo) NewUser(user *entity.User) error {
	_, err := r.db.Exec("INSERT INTO Users (ID, First_Name, Last_name, Mail, Password) VALUES (?, ?, ?, ?, ?)", user.ID, user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		log.Printf("error occured adding post to db: %v\n", err)
		return err
	}
	return nil
}

func (r *UsersRepo) GetUser(mail string) (string, string, error) {
	rows, err := r.db.Query("SELECT ID, Mail, Password from Users")
	if err != nil {
		log.Printf("error occured getUser from db: %v\n", err)
		return "", "", err
	}
	defer rows.Close()
	for rows.Next() {
		var id string
		var scanMail string
		var password string
		err := rows.Scan(&id, &scanMail, &password)
		if err != nil {
			log.Printf("error occured getUser scanning rows from db: %v\n", err)
			continue
		}
		if mail == scanMail {
			return id, password, nil
		}
	}
	return "", "", errors.ErrMailNotExist
}

func (r *UsersRepo) AddCookie(id string, cookieValue string, dt time.Time) error {
	_, err := r.db.Exec("INSERT INTO Session (Token, Expired_Date, User_ID) VALUES (?, ?, ?)", cookieValue, dt, id)
	if err != nil {
		log.Printf("error occured adding session to db: %v\n", err)
		return err
	}
	return nil
}
