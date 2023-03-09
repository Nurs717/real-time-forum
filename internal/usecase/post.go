package usecase

import (
	"context"
	"rtforum/internal/cerror"
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

func (*PostUseCase) validate(post *entity.Post) error {
	if post == nil || post.Body == "" {
		err := cerror.ErrEmptyPost
		return err
	}
	if len(post.Categories) < 1 {
		return nil
	}
	if post.Title == "" || len(post.Title) > 50 {
		return nil
	}
	return nil

}

func (u *PostUseCase) Create(ctx context.Context, post *entity.Post) error {
	if err := u.validate(post); err != nil {
		return err
	}
	date := time.Now().Format("2006-01-02 15:04:05")
	post.PostDate = date
	err := u.repo.CreatePost(ctx, post)
	if err != nil {
		return err
	}
	return nil
}

func (u *PostUseCase) GetAllPosts(ctx context.Context) ([]entity.Post, error) {
	posts, err := u.repo.GetAllPosts(ctx)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (u *PostUseCase) GetPostsByCategory(ctx context.Context, category string) ([]entity.Post, error) {
	posts, err := u.repo.GetPostsByCategory(ctx, category)
	if err != nil {
		return nil, err
	}
	return posts, nil
}
