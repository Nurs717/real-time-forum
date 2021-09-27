package usecase

import (
	"rtforum/internal/entity"
	"rtforum/internal/repository"
)

type Users interface {
	NewUser(user *entity.User) error
}

type Post interface {
	Validate(post *entity.Post) error
	Create(post *entity.Post) error
	FindAll() ([]entity.Post, error)
}

type Comments interface {
}

type UseCases struct {
	Users
	Post
	Comments
}

type UseCaseDeps struct {
	Repo *repository.Repository
}

func NewUseCases(deps *UseCaseDeps) *UseCases {
	return &UseCases{
		Users: NewUserUseCase(deps.Repo.Users),
		Post:  NewPostUseCase(deps.Repo.Post),
	}
}
