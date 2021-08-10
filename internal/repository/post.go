package repository

import (
	"database/sql"
	"rtforum/internal/entity"
)

type PostRepo struct {
	db *sql.DB
}

func NewPostRepo(db *sql.DB) *PostRepo {
	return &PostRepo{
		db: db,
	}
}

func (r *PostRepo) Save(post *entity.Post) (*entity.Post, error) {
	_, err := r.db.Exec("INSERT INTO Post (ID, Post) VALUES (?, ?,)", post.ID, post.Post)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *PostRepo) FindAll() ([]entity.Post, error) {
	return nil, nil
}
