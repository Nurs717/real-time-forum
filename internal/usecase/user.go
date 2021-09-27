package usecase

import (
	"log"
	"rtforum/internal/entity"
	"rtforum/internal/repository"

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

func generatePassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
