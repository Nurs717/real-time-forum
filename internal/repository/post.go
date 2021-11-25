package repository

import (
	"database/sql"
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

func (r *PostRepo) CreatePost(post *entity.Post) error {
	postSQL, err := r.db.Exec("INSERT INTO Post (Body, Title, User_ID, Date) VALUES (?, ?, ?, ?)", post.Body, post.Title, post.UserID, post.PostDate)
	if err != nil {
		log.Printf("error occured adding post to db: %v", err)
		return err
	}
	var id int64
	id, err = postSQL.LastInsertId()
	if err != nil {
		log.Printf("error ocured when getting id from sql result table post: %v", err)
		return err
	}
	for _, category := range post.Category {
		if category != "" {
			_, err = r.db.Exec("INSERT OR IGNORE INTO Category (NAME) VALUES (?)", category)
			if err != nil {
				log.Printf("error ocured when inserting categories in to table: %v", err)
			}
			_, err = r.db.Exec("INSERT INTO Category_Map (Post_ID, Category_ID) VALUES (?, ?)", id, category)
			if err != nil {
				log.Printf("error occured adding post to db: %v", err)
				return err
			}
		}
	}

	return nil
}

func (r *PostRepo) GetAllPosts() ([]entity.Post, error) {
	posts := []entity.Post{}
	post, err := r.db.Query("SELECT P.ID, U.UserName, P.Title, P.Date FROM Post as P INNER JOIN Users as U ON U.ID = P.User_ID")
	if err != nil {
		log.Printf("error occured querying db: %v", err)
		return nil, err
	}
	for post.Next() {
		p := entity.Post{}
		err := post.Scan(&p.ID, &p.UserName, &p.Title, &p.PostDate)
		if err != nil {
			log.Printf("Error occured scanning Query %v\n", err)
			return nil, err
		}
		posts = append(posts, p)
	}
	return posts, nil
}

func (r *PostRepo) GetPostsByCategory(category string) ([]entity.Post, error) {
	posts := []entity.Post{}
	post, err := r.db.Query("SELECT P.ID, U.UserName, P.Title, P.Date FROM Post as P INNER JOIN Users as U ON U.ID = P.User_ID INNER JOIN Category_Map as Map ON Map.Category_ID = Category.NAME INNER JOIN Category ON Map.Post_ID = P.ID WHERE Category.NAME = ?", category)
	if err != nil {
		log.Printf("error occured querying db: %v", err)
		return nil, err
	}
	for post.Next() {
		p := entity.Post{}
		err := post.Scan(&p.ID, &p.UserName, &p.Title, &p.PostDate)
		if err != nil {
			log.Printf("Error occured scanning Query %v\n", err)
			return nil, err
		}
		posts = append(posts, p)
	}
	return posts, nil
}
