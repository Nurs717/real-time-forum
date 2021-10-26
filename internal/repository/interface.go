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
}

type Post interface {
	Create(post *entity.Post) error
	FindAll() ([]entity.Post, error)
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
