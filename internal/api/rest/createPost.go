package rest

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"rtforum/internal/entity"
)

func (h *Handler) CreatPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create post rest:", r.Context().Value(CtxReqIdKey))
	if r.Context().Value(CtxReqIdKey) == "Guest" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	switch r.Method {
	case "POST":
		var post *entity.Post
		data, err := io.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			log.Printf("error reading body %v\n", err)
		}
		err = json.Unmarshal(data, &post)
		if err != nil {
			log.Printf("error unmarshaling %v\n", err)
		}
		post.UserID = r.Context().Value(CtxReqIdKey).(string)
		err = h.UseCases.Post.Create(post)
		if err != nil {
			log.Printf("error adding post in rest: %v\n", err)
		}
		fmt.Println("post from client:", post)
	}
}
