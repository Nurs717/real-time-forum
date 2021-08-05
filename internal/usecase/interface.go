package usecase

import (
	"rtforum/internal/entity"
	"rtforum/internal/repository"
)

type Users interface {
}

type Posts interface {
	Validate(post *entity.Post) error
	Create(post *entity.Post) (*entity.Post, error)
	FindAll() []entity.Post
}

type Comments interface {
}

type UseCases struct {
	Users
	Posts
	Comments
}

type UseCaseDeps struct {
	Repo *repository.Repository
}

func NewUseCases(deps *UseCaseDeps) *UseCases {
	return &UseCases{}
}
