package repository

import (
	"context"
	"database/sql"
	"rtforum/internal/entity"
	"time"
)

type Users interface {
	NewUser(ctx context.Context, user *entity.User) error
	GetUser(ctx context.Context, mail string) (string, string, error)
	AddCookie(ctx context.Context, id string, cookieValue string, dt time.Time) error
	GetUserIDbyCookie(ctx context.Context, token string) (string, error)
	GetUserName(ctx context.Context, userID string) (string, error)
}

type Post interface {
	CreatePost(post *entity.Post) error
	GetAllPosts(ctx context.Context) ([]entity.Post, error)
	GetPostsByCategory(category string) ([]entity.Post, error)
}

type Comment interface {
}

type Repository struct {
	Users
	Post
	Comment
}

func NewRepository(db *sql.DB, timeout time.Duration) *Repository {
	return &Repository{
		Users: NewUsersRepo(db, timeout),
		Post:  NewPostRepo(db, timeout),
	}
}
