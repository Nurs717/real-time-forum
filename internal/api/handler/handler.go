package handler

import "rtforum/internal/usecase"

type Message struct {
	Message string `json:"message"`
}

type Handler struct {
	UseCases *usecase.UseCases
}

func NewHandler(UseCases *usecase.UseCases) *Handler {
	return &Handler{UseCases: UseCases}
}
