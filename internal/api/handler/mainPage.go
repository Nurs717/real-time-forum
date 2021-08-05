package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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
			var message Message
			data, err := ioutil.ReadAll(r.Body)
			defer r.Body.Close()
			if err != nil {
				log.Printf("error reading body %v\n", err)
			}
			err = json.Unmarshal(data, &message)
			if err != nil {
				log.Printf("error unmarshaling %v\n", err)
			}
			fmt.Println("message from client:", message.Message)
			fmt.Fprintf(w, "Server: %s\n", message.Message+" | "+time.Now().Format(time.RFC3339))
		}
	}
}
