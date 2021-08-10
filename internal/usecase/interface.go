package usecase

import (
	"rtforum/internal/entity"
	"rtforum/internal/repository"
)

type Users interface {
}

type Post struct {
	ID   int
	Post string
}

type Posts interface {
	Validate(post *entity.Post) error
	Create(post *entity.Post) (*entity.Post, error)
	FindAll() []entity.Post
	AddIDPost(post *repository.Post) error
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

func (p *Post) AddIDPost(post *repository.Post) error {
	return nil
}
