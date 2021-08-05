package handler

import (
	"net/http"
	"rtforum/internal/api/middleware"
)

func (h *Handler) Router() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", middleware.EnableCORS(h.MainPage()))
	mux.HandleFunc("/login", LogIn())
	mux.HandleFunc("/signup", SignUp())
	return mux
}
