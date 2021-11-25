package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"rtforum/internal/entity"
)

func (h *Handler) MainPage(w http.ResponseWriter, r *http.Request) {
	// fmt.Println(r.Cookie("session"))
	fmt.Println("id from middleware", r.Context().Value(CtxReqIdKey))
	if r.Context().Value(CtxReqIdKey) == "Guest" {
		w.WriteHeader(http.StatusForbidden)
	}
	switch r.Method {
	case "GET":
		var posts []entity.Post
		var err error
		category := r.FormValue("category")
		fmt.Println("category:", category)
		if category == "" {
			posts, err = h.UseCases.GetAllPosts()
			if err != nil {
				log.Printf("Error occured in handler getAllPosts: %v\n", err)
				return
			}
		} else {
			posts, err = h.UseCases.GetPostsByCategory(category)
			if err != nil {
				log.Printf("Error occured in handler getPostsByCategory: %v\n", err)
				return
			}
		}
		result, err := json.Marshal(posts)
		if err != nil {
			log.Printf("Error occured when marshalling %v\n", err)
			return
		}
		// fmt.Println("get posts: ", string(result))
		w.Write(result)
	case "POST":

	}
}
