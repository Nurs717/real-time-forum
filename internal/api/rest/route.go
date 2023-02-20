package rest

import (
	"net/http"
)

func (h *Handler) Router() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", EnableCORS(h.CheckCookie(h.MainPage)))
	mux.HandleFunc("/login", EnableCORS(h.CheckCookie(h.LogIn)))
	mux.HandleFunc("/create-post", EnableCORS(h.CheckCookie(h.CreatPost)))
	mux.HandleFunc("/signup", EnableCORS(h.SignUp))
	mux.HandleFunc("/post/", EnableCORS(h.CheckCookie(h.PostAndComments)))
	return mux
}
