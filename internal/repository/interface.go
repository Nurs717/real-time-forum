package repository

import (
	"database/sql"
)

type Users interface {
}

type Post interface {
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
	}
}
