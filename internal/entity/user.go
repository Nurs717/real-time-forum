package entity

import (
	"net/mail"
	"rtforum/internal/cerror"
	"unicode"
)

// User data
type User struct {
	ID        string
	UserName  string `json:"username"`
	Age       string `json:"age"`
	Gender    string `json:"gender"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Post      []string
	Like      []string
	Dislike   []string
}

func (u *User) Validate() error {
	if u.UserName == "" || u.Email == "" || u.Password == "" || u.FirstName == "" || u.LastName == "" ||
		u.Age == "" || u.Gender == "" {
		return cerror.NewErrorf(cerror.ErrorCodeInvalidArgument, cerror.DefaultType, cerror.ErrEmptyRegisterData)
	}
	_, err := mail.ParseAddress(u.Email)
	if err != nil {
		return cerror.NewErrorf(cerror.ErrorCodeInvalidArgument, cerror.FormatMailType, cerror.ErrEmailInvalid)
	}
	if !isPasswordValid(u.Password) {
		return cerror.NewErrorf(cerror.ErrorCodeInvalidArgument, cerror.FormatPwdType, cerror.ErrInvalidPassword)
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
	if len(password) >= 8 {
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
