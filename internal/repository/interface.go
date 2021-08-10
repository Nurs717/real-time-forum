package repository

import (
	"database/sql"
	"rtforum/internal/entity"
)

type Users interface {
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
