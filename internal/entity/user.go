package entity

import (
	"net/mail"
	"rtforum/errors"
	"unicode"
)

// User data
type User struct {
	ID        string
	UserName  string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Post      []string
	Like      []string
	Dislike   []string
}

func (u *User) Validate() error {
	if u.UserName == "" || u.Email == "" || u.Password == "" || u.FirstName == "" || u.LastName == "" {
		return errors.ErrEmptyRegisterData
	}
	_, err := mail.ParseAddress(u.Email)
	if err != nil {
		return errors.ErrEmailInvalid
	}
	if isPasswordValid(u.Password) {
		return errors.ErrInvalidPassword
	}
	return nil
}

func isPasswordValid(password string) bool {
	var (
		hasMinLen = false
		hasUpper  = false
		hasLower  = false
		hasNumber = false
	)
	if len(password) >= 7 {
		hasMinLen = true
	}
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		}
	}
	return hasMinLen && hasUpper && hasLower && hasNumber
}
