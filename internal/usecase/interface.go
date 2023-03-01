package usecase

import (
	"context"
	"net/http"
	"rtforum/internal/entity"
	"rtforum/internal/repository"
)

type Users interface {
	NewUser(ctx context.Context, user *entity.User) error
	SetCookie(ctx context.Context, user *entity.User) (*http.Cookie, error)
	IsCookieValid(ctx context.Context, token string) (string, error)
	GetUserName(ctx context.Context, userID string) (string, error)
}

type Post interface {
	Validate(post *entity.Post) error
	Create(post *entity.Post) error
	GetAllPosts(ctx context.Context) ([]entity.Post, error)
	GetPostsByCategory(ctx context.Context, category string) ([]entity.Post, error)
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
