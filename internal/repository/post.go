package repository

import (
	"database/sql"
	"fmt"
	"log"
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
	posts := []entity.Post{}
	post, err := r.db.Query("SELECT P.Post FROM Post as P")
	if err != nil {
		log.Printf("error occured querying %v", err)
		return nil, err
	}
	for post.Next() {
		p := entity.Post{}
		err := post.Scan(&p.Post)
		if err != nil {
			log.Printf("Error occured scanning Query %v", err)
			return nil, err
		}
		posts = append(posts, p)
	}
	return posts, nil
}
