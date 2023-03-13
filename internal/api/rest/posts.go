package rest

import (
	"fmt"
	"net/http"
	"rtforum/internal/entity"
)

type ResponsePosts struct {
	Posts    []entity.Post `json:"posts"`
	UserName string        `json:"username"`
}

func (h *Handler) posts(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	switch r.Method {
	case "GET":
		userID := fmt.Sprintf("%v", r.Context().Value(CtxReqUserIdKey))
		statusCode := http.StatusOK
		if userID == "Guest" {
			statusCode = http.StatusForbidden
		}
		var posts []entity.Post
		var err error
		var username string
		if userID != "Guest" {
			username = r.Context().Value(CtxReqUsernameKey).(string)
		}
		category := r.FormValue("category")
		if category == "" {
			posts, err = h.UseCases.GetAllPosts(r.Context())
			if err != nil {
				renderErrorResponse(w, "can't get posts", err)
				return
			}
		} else {
			posts, err = h.UseCases.GetPostsByCategory(r.Context(), category)
			if err != nil {
				renderErrorResponse(w, "can't get posts by category", err)
				return
			}
		}
		res := ResponsePosts{
			posts,
			username,
		}
		renderResponse(w, res, statusCode)
	}
}
