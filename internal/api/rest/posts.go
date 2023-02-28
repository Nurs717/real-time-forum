package rest

import (
	"fmt"
	"log"
	"net/http"
	"rtforum/internal/entity"
)

type ResponseGetPosts struct {
	Posts    []entity.Post `json:"posts"`
	UserName string        `json:"username"`
}

func (h *Handler) MainPage(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("id from middleware", r.Context().Value(CtxReqIdKey))
	userID := fmt.Sprintf("%v", r.Context().Value(CtxReqIdKey))
	statusCode := http.StatusOK
	if userID == "Guest" {
		statusCode = http.StatusForbidden
	}
	switch r.Method {
	case "GET":
		var posts []entity.Post
		var err error
		var username string
		if userID != "Guest" {
			username, err = h.UseCases.GetUserName(r.Context(), userID)
			if err != nil {
				renderErrorResponse(w, "can't get username", err)
				return
			}
		}
		category := r.FormValue("category")
		if category == "" {
			posts, err = h.UseCases.GetAllPosts(r.Context())
			if err != nil {
				renderErrorResponse(w, "can't get posts", err)
				return
			}
		} else {
			posts, err = h.UseCases.GetPostsByCategory(category)
			if err != nil {
				log.Printf("Error occured in rest getPostsByCategory: %v\n", err)
				return
			}
		}
		res := ResponseGetPosts{
			posts,
			username,
		}
		renderResponse(w, res, statusCode)
	}
}
