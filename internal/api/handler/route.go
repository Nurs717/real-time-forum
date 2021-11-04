package handler

import (
	"net/http"
)

func (h *Handler) Router() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", EnableCORS(h.CheckCookie(h.MainPage)))
	mux.HandleFunc("/login", EnableCORS(h.LogIn))
	mux.HandleFunc("/signup", EnableCORS(h.SignUp))
	return mux
}
