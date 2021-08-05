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

func (r *PostRepo) Save(post entity.Post) (*entity.Post, error) {
	// _, err := r.db.Exec("INSERT INTO Post (ID, Post) VALUES (?, ?,)", entity.Post.ID, entity.Post.Post)
	// if err != nil {
	// 	return nil, err
	// }
	return nil, nil
}

func (*Repository) FindAll() ([]entity.Post, error) {
	return nil, nil
}
