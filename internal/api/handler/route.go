package handler

import (
	"net/http"
)

func Router() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", MainPage())
	mux.HandleFunc("/login", LogIn())
	mux.HandleFunc("/signup", SignUp())
	return mux
}
