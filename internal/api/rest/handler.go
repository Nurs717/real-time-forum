package rest

import (
	"rtforum/internal/usecase"
)

type Handler struct {
	UseCases *usecase.UseCases
}

func NewHandler(UseCases *usecase.UseCases) *Handler {
	return &Handler{
		UseCases: UseCases,
	}
}
