package entity

import (
	"rtforum/errors"
)

// User data
type User struct {
	ID        string
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Post      []string
	Like      []string
	Dislike   []string
}

func (u *User) Validate() error {
	if u.Email == "" || u.Password == "" || u.FirstName == "" || u.LastName == "" {
		return errors.ErrInvalidEntity
	}
	return nil
}
