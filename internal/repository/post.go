package repository

import (
	"database/sql"
	"fmt"
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

func (r *PostRepo) Create(post *entity.Post) error {
	_, err := r.db.Exec("INSERT INTO Post (ID, Post) VALUES (?, ?)", post.ID, post.Post)
	if err != nil {
		fmt.Println("error occured adding post to db:", err)
		return err
	}
	return nil
}

func (r *PostRepo) FindAll() ([]entity.Post, error) {
	return nil, nil
}
