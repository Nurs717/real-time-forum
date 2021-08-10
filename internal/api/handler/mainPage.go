package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"rtforum/internal/usecase"
	"time"
)

func (h *Handler) MainPage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			path := r.URL.Path
			fmt.Println(path)
			if path == "/" {
				path = "./web/index.html"
			} else {
				path = "." + path
			}
			http.ServeFile(w, r, path)
		case "POST":
			var post Post
			data, err := ioutil.ReadAll(r.Body)
			defer r.Body.Close()
			if err != nil {
				log.Printf("error reading body %v\n", err)
			}
			err = json.Unmarshal(data, &post)
			if err != nil {
				log.Printf("error unmarshaling %v\n", err)
			}
			err = usecase.AddIDPost(&post)
			if err != nil {
				log.Fatalf("error adding ID to post %v\n", err)
			}
			fmt.Println("post from client:", post.Post)
			fmt.Fprintf(w, "Server: %s\n", post.Post+" | "+time.Now().Format(time.RFC3339))
		}
	}
}
