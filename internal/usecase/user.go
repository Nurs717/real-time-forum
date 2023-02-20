package usecase

import (
	"context"
	"log"
	"net/http"
	"rtforum/internal/cerror"
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

func (u *UserUseCase) NewUser(ctx context.Context, user *entity.User) error {
	//validating incoming data
	if err := user.Validate(); err != nil {
		log.Printf("usecase: creat user: %v\n", err)
		return err
	}
	// setting uuid number for user
	user.ID = uuid.NewV4().String()
	// setting generated encrypted password to user
	pwd, err := generatePassword(user.Password)
	if err != nil {
		log.Printf("uscase creat user: generate pwd: %v", err)
		return err
	}
	user.Password = pwd
	// creates user in repo
	if err = u.repo.NewUser(ctx, user); err != nil {
		log.Printf("usecase creat user: %v\n", err)
		return err
	}
	return nil
}

func (u *UserUseCase) SetCookie(ctx context.Context, user *entity.User) (*http.Cookie, error) {
	id, password, err := u.repo.GetUser(ctx, user.Email)
	if err != nil {
		log.Printf("error occured usecase SetCookie: %v\n", err)
		return nil, err
	}
	hashedPassword := []byte(password)
	passwordToCheck := []byte(user.Password)
	if err := bcrypt.CompareHashAndPassword(hashedPassword, passwordToCheck); err != nil {
		return nil, cerror.ErrWrongPassword
	}

	expire := time.Now().Add(1 * time.Hour)
	u1 := uuid.NewV4()
	cookie := &http.Cookie{
		Name:     "session",
		Value:    u1.String(),
		Expires:  expire,
		SameSite: http.SameSiteLaxMode,
	}
	if err = u.repo.AddCookie(ctx, id, cookie.Value, expire); err != nil {
		log.Printf("error occured in usecase when adding session: %v\n", err)
		return nil, err
	}
	return cookie, nil
}

func (u *UserUseCase) IsCookieValid(ctx context.Context, token string) (string, error) {
	userID, err := u.repo.GetUserIDbyCookie(ctx, token)
	if err != nil {
		log.Printf("error occured while checking cookie in usecase: %v", err)
		return "", err
	}
	return userID, nil
}

func generatePassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
