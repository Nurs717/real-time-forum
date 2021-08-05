package handler

import "rtforum/internal/usecase"

type Post struct {
	Post string `json:"post"`
}

type Handler struct {
	UseCases *usecase.UseCases
}

func NewHandler(UseCases *usecase.UseCases) *Handler {
	return &Handler{UseCases: UseCases}
}
