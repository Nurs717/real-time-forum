package repository

import (
	"context"
	"database/sql"
	"log"
	"rtforum/internal/cerror"
	"rtforum/internal/entity"
	"time"
)

type PostRepo struct {
	db      *sql.DB
	timeout time.Duration
}

func NewPostRepo(db *sql.DB, timeout time.Duration) *PostRepo {
	return &PostRepo{
		db:      db,
		timeout: timeout,
	}
}

func (r *PostRepo) CreatePost(post *entity.Post) error {
	postSQL, err := r.db.Exec("INSERT INTO Post (Body, Title, User_ID, Date) VALUES (?, ?, ?, ?)", post.Body, post.Title, post.UserID, post.PostDate)
	if err != nil {
		log.Printf("error occured adding post to db: %v", err)
		return err
	}

	id, err := postSQL.LastInsertId()
	if err != nil {
		log.Printf("error ocured when getting id from sql result table post: %v", err)
		return err
	}

	for _, category := range post.Categories {
		if category != "" {
			_, err = r.db.Exec("INSERT INTO Category (Post_ID, Category_Name) VALUES (?, ?)", id, category)
			if err != nil {
				log.Printf("error occured adding post to db: %v", err)
				return err
			}
		}
	}

	return nil
}

func (r *PostRepo) GetAllPosts(ctx context.Context) ([]entity.Post, error) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	var posts []entity.Post
	postRows, err := r.db.QueryContext(ctxWithTimeout,
		"SELECT P.ID, U.UserName, P.Body, P.Title, P.Date FROM Post as P INNER JOIN Users as U ON U.ID = P.User_ID")
	if err != nil {
		return nil, cerror.WrapErrorf(err, cerror.ErrorCodeInternal, cerror.DefaultType, "repo: GetAllPosts: Query")
	}
	for postRows.Next() {
		var p entity.Post
		err := postRows.Scan(&p.ID, &p.UserName, &p.Body, &p.Title, &p.PostDate)
		if err != nil {
			return nil, cerror.WrapErrorf(err, cerror.ErrorCodeInternal, cerror.DefaultType, "repo: GetAllPosts: Scan postRows")
		}
		categories, err := r.getCategoriesByPostID(ctx, p.ID)
		if err != nil {
			return nil, err
		}
		p.Categories = categories
		posts = append(posts, p)
	}
	if err = postRows.Err(); err != nil {
		return nil, cerror.WrapErrorf(err, cerror.ErrorCodeInternal, cerror.DefaultType, "repo: GetAllPosts: catching postRows error")
	}
	return posts, nil
}

func (r *PostRepo) GetPostsByCategory(ctx context.Context, category string) ([]entity.Post, error) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	var posts []entity.Post
	postRows, err := r.db.QueryContext(ctxWithTimeout,
		"SELECT P.ID, U.UserName, P.Title, P.Date FROM Post as P INNER JOIN Users as U ON U.ID = P.User_ID INNER JOIN Category as C ON C.Post_ID = P.ID WHERE C.Category_Name =?", category)
	if err != nil {
		return nil, cerror.WrapErrorf(err, cerror.ErrorCodeInternal, cerror.DefaultType, "repo: GetPostsByCategory: Query")
	}
	for postRows.Next() {
		p := entity.Post{}
		err := postRows.Scan(&p.ID, &p.UserName, &p.Title, &p.PostDate)
		if err != nil {
			return nil, cerror.WrapErrorf(err, cerror.ErrorCodeInternal, cerror.DefaultType, "repo: GetPostsByCategory: Scan postRows")
		}
		categories, err := r.getCategoriesByPostID(ctx, p.ID)
		if err != nil {
			return nil, err
		}
		p.Categories = categories
		posts = append(posts, p)
	}
	if err = postRows.Err(); err != nil {
		return nil, cerror.WrapErrorf(err, cerror.ErrorCodeInternal, cerror.DefaultType, "repo: GetPostsByCategory: catching postRows error")
	}
	return posts, nil
}

func (r *PostRepo) getCategoriesByPostID(ctx context.Context, id int) ([]string, error) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	var categories []string
	rows, err := r.db.QueryContext(ctxWithTimeout,
		"SELECT C.Category_Name FROM Category as C WHERE C.Post_ID = ?", id)
	if err != nil {
		return nil, cerror.WrapErrorf(err, cerror.ErrorCodeInternal, cerror.DefaultType, "repo: getCategoriesByPostID: Query")
	}
	for rows.Next() {
		var category string
		err := rows.Scan(&category)
		if err != nil {
			return nil, cerror.WrapErrorf(err, cerror.ErrorCodeInternal, cerror.DefaultType, "repo: getCategoriesByPostID: Scan rows")
		}
		categories = append(categories, category)
	}
	if err = rows.Err(); err != nil {
		return nil, cerror.WrapErrorf(err, cerror.ErrorCodeInternal, cerror.DefaultType, "repo: getCategoriesByPostID: catching rows error")
	}
	return categories, nil
}
