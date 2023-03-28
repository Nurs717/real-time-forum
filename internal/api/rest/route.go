package rest

import (
	"net/http"
)

func (h *Handler) Router() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/posts", enableCORS(h.checkCookie(h.reqLogger(h.posts))))
	mux.HandleFunc("/login", enableCORS(h.checkCookie(h.reqLogger(h.logIn))))
	mux.HandleFunc("/create-post", enableCORS(h.checkCookie(h.reqLogger(h.creatPost))))
	mux.HandleFunc("/signup", enableCORS(h.reqLogger(h.signUp)))
	mux.HandleFunc("/post", enableCORS(h.checkCookie(h.reqLogger(h.post))))
	return mux
}
