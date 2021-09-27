package repository

import (
	"database/sql"
	"log"
	"rtforum/internal/entity"
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
