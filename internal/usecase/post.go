package usecase

import (
	"fmt"
	"math/rand"
	"rtforum/errors"
	"rtforum/internal/entity"
	"rtforum/internal/repository"
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
	if post.Post == "" {
		err := errors.ErrEmptyPost
		return err
	}
	return nil

}

func (u *PostUseCase) Create(post *entity.Post) error {
	post.ID = rand.Int()
	err := u.repo.Create(post)
	if err != nil {
		fmt.Println("error occured usecase:", err)
		return err

	}
	return nil
}

func (*PostUseCase) FindAll() ([]entity.Post, error) {
	return nil, nil
}
