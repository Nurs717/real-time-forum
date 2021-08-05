package usecase

import (
	"math/rand"
	"rtforum/errors"
	"rtforum/internal/entity"
)

func (*UseCases) Validate(post *entity.Post) error {
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

func (*UseCases) Create(post *entity.Post) error {
	post.ID = rand.Int()
	return nil
}

func (*UseCases) FindAll() ([]entity.Post, error) {
	return nil, nil
}
