package usecase

import (
	"context"
	"fmt"
	"net/http"
	"rtforum/internal/cerror"
	"rtforum/internal/entity"
	"rtforum/internal/repository"
	"strings"
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
		return err
	}
	// setting uuid number for user
	user.ID = uuid.NewV4().String()
	// setting generated encrypted password to user
	pwd, err := generatePassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = pwd
	user.Email = strings.ToLower(user.Email)
	// creates user in repo
	if err = u.repo.NewUser(ctx, user); err != nil {
		return err
	}
	return nil
}

func (u *UserUseCase) SetCookie(ctx context.Context, user *entity.User) (*http.Cookie, error) {
	user.Email = strings.ToLower(user.Email)
	id, password, err := u.repo.GetUser(ctx, user.Email)
	if err != nil {
		return nil, err
	}
	fmt.Println(password)
	hashedPassword := []byte(password)
	passwordToCheck := []byte(user.Password)
	if err := bcrypt.CompareHashAndPassword(hashedPassword, passwordToCheck); err != nil {
		return nil, cerror.WrapErrorf(err, cerror.ErrorCodeUnauthorized, cerror.DefaultType, "usecase: SetCookie: check password")
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
		return nil, err
	}
	return cookie, nil
}

func (u *UserUseCase) IsCookieValid(ctx context.Context, token string) (string, error) {
	userID, err := u.repo.GetUserIDbyCookie(ctx, token)
	if err != nil {
		return "", err
	}
	return userID, nil
}

func (u *UserUseCase) GetUserName(ctx context.Context, userID string) (string, error) {
	username, err := u.repo.GetUserName(ctx, userID)
	if err != nil {
		return "", err
	}
	return username, nil
}

func generatePassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", cerror.WrapErrorf(err, cerror.ErrorCodeInternal, cerror.DefaultType, "usecase: NewUser: generatePassword")
	}
	return string(hash), nil
}
