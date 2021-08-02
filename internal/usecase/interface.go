package usecase

import "rtforum/internal/repository"

type Users interface {
}

type Posts interface {
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
