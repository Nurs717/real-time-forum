package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"rtforum/internal/entity"
)

func (h *Handler) MainPage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			posts, err := h.UseCases.Post.FindAll()
			if err != nil {
				log.Printf("Error occured %v\n", err)
			}
			result, err := json.Marshal(posts)
			if err != nil {
				log.Printf("Error occured when marshalling %v\n", err)
			}
			fmt.Println("method get worked", string(result))
			w.Write(result)

		case "POST":
			var post *entity.Post
			data, err := ioutil.ReadAll(r.Body)
			defer r.Body.Close()
			if err != nil {
				log.Printf("error reading body %v\n", err)
			}
			err = json.Unmarshal(data, &post)
			if err != nil {
				log.Printf("error unmarshaling %v\n", err)
			}
			err = h.UseCases.Post.Create(post)
			if err != nil {
				log.Printf("error adding ID to post %v\n", err)
			}
			fmt.Println("post from client:", post.Post, post.ID)
		}
	}
}
