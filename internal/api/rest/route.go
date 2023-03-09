package rest

import (
	"net/http"
)

func (h *Handler) Router() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/posts", enableCORS(h.checkCookie(h.posts)))
	mux.HandleFunc("/login", enableCORS(h.checkCookie(h.logIn)))
	mux.HandleFunc("/create-post", enableCORS(h.checkCookie(h.creatPost)))
	mux.HandleFunc("/signup", enableCORS(h.signUp))
	mux.HandleFunc("/post/", enableCORS(h.checkCookie(h.PostAndComments)))
	return mux
}
