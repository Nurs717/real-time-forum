package repository

import (
	"database/sql"
	"rtforum/internal/entity"
	"time"
)

type Users interface {
	NewUser(user *entity.User) error
	GetUser(mail string) (string, string, error)
	AddCookie(id string, cookieValue string, dt time.Time) error
	GetUserIDbyCookie(token string) (string, error)
}

type Post interface {
	CreatePost(post *entity.Post) error
	GetAllPosts() ([]entity.Post, error)
	GetPostsByCategory(category string) ([]entity.Post, error)
}

type Comment interface {
}

type Repository struct {
	Users
	Post
	Comment
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Users: NewUsersRepo(db),
		Post:  NewPostRepo(db),
	}
}
