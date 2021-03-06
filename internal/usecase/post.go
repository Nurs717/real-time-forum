package usecase

import (
	"fmt"
	"log"
	"rtforum/errors"
	"rtforum/internal/entity"
	"rtforum/internal/repository"
	"time"
)

type PostUseCase struct {
	repo repository.Post
}

func NewPostUseCase(repo repository.Post) *PostUseCase {
	return &PostUseCase{
		repo: repo,
	}
}

func (*PostUseCase) Validate(post *entity.Post) error {
	if post == nil {
		err := errors.ErrEmptyPost
		return err
	}
	if post.Body == "" {
		err := errors.ErrEmptyPost
		return err
	}
	return nil

}

func (u *PostUseCase) Create(post *entity.Post) error {
	date := time.Now().Format("2006-01-02 15:04:05")
	post.PostDate = date
	err := u.repo.CreatePost(post)
	if err != nil {
		fmt.Println("error occured usecase:", err)
		return err

	}
	return nil
}

func (u *PostUseCase) GetAllPosts() ([]entity.Post, error) {
	posts, err := u.repo.GetAllPosts()
	if err != nil {
		log.Printf("Error occured usecase getposts: %v", err)
		return nil, err
	}
	return posts, nil
}

func (u *PostUseCase) GetPostsByCategory(category string) ([]entity.Post, error) {
	posts, err := u.repo.GetPostsByCategory(category)
	if err != nil {
		log.Printf("Error occured in usecase getpostbycategory: %v", err)
		return nil, err
	}
	return posts, nil
}
