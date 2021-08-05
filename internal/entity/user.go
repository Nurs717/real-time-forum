package entity

import (
	"rtforum/errors"

	"golang.org/x/crypto/bcrypt"
)

// User data
type User struct {
	ID        ID
	Email     string
	Password  string
	FirstName string
	LastName  string
	Post      []string
	Like      []string
	Dislike   []string
}

// NewUser creates a new user
func NewUser(email, password, firstName, lastName string) (*User, error) {
	u := &User{
		ID:        NewID(),
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
	}
	pwd, err := generatePassword(password)
	if err != nil {
		return nil, err
	}
	u.Password = pwd
	err = u.Validate()
	if err != nil {
		return nil, err
	}
	return u, nil
}

// Validate validates data
func (u *User) Validate() error {
	if u.Email == "" || u.Password == "" || u.FirstName == "" || u.LastName == "" {
		return errors.ErrInvalidEntity
	}
	return nil
}

// generatePassword generates crypted password
func generatePassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
