package rest

import (
	"encoding/json"
	"io"
	"net/http"
	"rtforum/internal/cerror"
	"rtforum/internal/entity"
)

func (h *Handler) creatPost(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if r.Context().Value(CtxReqUserIdKey) == "Guest" {
		renderResponse(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	switch r.Method {
	case "POST":
		var post *entity.Post
		data, err := io.ReadAll(r.Body)
		if err != nil {
			err = cerror.WrapErrorf(err, cerror.ErrorCodeInternal, cerror.DefaultType, "rest: CreatePost: reading body")
			renderErrorResponse(w, "internal error", err)
			return
		}
		err = json.Unmarshal(data, &post)
		if err != nil {
			err = cerror.WrapErrorf(err, cerror.ErrorCodeInternal, cerror.DefaultType, "rest: creatPost: unmarshal error")
			renderErrorResponse(w, "internal error", err)
			return
		}
		post.UserID = r.Context().Value(CtxReqUserIdKey).(string)
		err = h.UseCases.Post.Create(r.Context(), post)
		if err != nil {
			renderErrorResponse(w, "unable to create post", err)
			return
		}
		renderResponse(w, "post created", http.StatusCreated)
	}
}
