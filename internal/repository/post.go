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

func (r *PostRepo) Create(post *entity.Post) error {
	postSQL, err := r.db.Exec("INSERT INTO Post (Body, Title, User_ID, Date, Category) VALUES (?, ?, ?, ?, ?)", post.Body, post.Title, post.UserID, post.PostDate, post.Category)
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
	_, err = r.db.Exec("INSERT OR IGNORE INTO Category (NAME) VALUES (?)", post.Category)
	if err != nil {
		log.Printf("error ocured when inserting categories in to table: %v", err)
	}
	_, err = r.db.Exec("INSERT INTO Category_Map (Post_ID, Category_ID) VALUES (?, ?)", id, post.Category)
	if err != nil {
		log.Printf("error occured adding post to db: %v", err)
		return err
	}

	return nil
}

func (r *PostRepo) FindAll() ([]entity.Post, error) {
	posts := []entity.Post{}
	post, err := r.db.Query("SELECT P.ID, P.Body, P.Date, P.Category FROM Post as P")
	if err != nil {
		log.Printf("error occured querying %v", err)
		return nil, err
	}
	for post.Next() {
		p := entity.Post{}
		err := post.Scan(&p.ID, &p.Body, &p.PostDate, &p.Category)
		if err != nil {
			log.Printf("Error occured scanning Query %v\n", err)
			return nil, err
		}
		posts = append(posts, p)
	}
	return posts, nil
}
