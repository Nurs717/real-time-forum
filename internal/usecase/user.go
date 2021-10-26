package usecase

import (
	"log"
	"net/http"
	"rtforum/errors"
	"rtforum/internal/entity"
	"rtforum/internal/repository"
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	repo repository.Users
}

func NewUserUseCase(repo repository.Users) *UserUseCase {
	return &UserUseCase{
		repo: repo,
	}
}

func (u *UserUseCase) NewUser(user *entity.User) error {
	user.ID = uuid.NewV4().String()
	pwd, err := generatePassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = pwd
	err = user.Validate()
	if err != nil {
		return err
	}
	err = u.repo.NewUser(user)
	if err != nil {
		log.Printf("error occured usecase: %v\n", err)
		return err
	}
	return nil
}

func (u *UserUseCase) SetCookie(user *entity.User) (*http.Cookie, error) {
	id, password, err := u.repo.GetUser(user.Email)
	if err != nil {
		log.Printf("error occured usecase SetCookie: %v\n", err)
		return nil, err
	}
	hashedPassword := []byte(password)
	passwordToCheck := []byte(user.Password)
	if err := bcrypt.CompareHashAndPassword(hashedPassword, passwordToCheck); err != nil {
		return nil, errors.ErrWrongPassword
	}

	expire := time.Now().Add(24 * time.Hour)
	u1 := uuid.NewV4()

	cookie := &http.Cookie{
		Name:     "session",
		Value:    u1.String(),
		Expires:  expire,
		Path:     "/",
		HttpOnly: true,
	}
	err = u.repo.AddCookie(id, cookie.Value, expire)
	if err != nil {
		log.Printf("error occured in usecase when addin session: %v\n", err)
		return nil, err
	}
	return cookie, nil
}

func generatePassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
