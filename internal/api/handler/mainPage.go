package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (h *Handler) MainPage(w http.ResponseWriter, r *http.Request) {
	// fmt.Println(r.Cookie("session"))
	fmt.Println("id from middleware", r.Context().Value(CtxReqIdKey))
	// fmt.Println(CtxReqIdKey)
	switch r.Method {
	case "GET":
		posts, err := h.UseCases.Post.GetAllPosts()
		if err != nil {
			log.Printf("Error occured %v\n", err)
			return
		}
		result, err := json.Marshal(posts)
		if err != nil {
			log.Printf("Error occured when marshalling %v\n", err)
			return
		}
		// fmt.Println("get posts: ", string(result))
		w.Write(result)

		// case "POST":
		// 	var post *entity.Post
		// 	data, err := ioutil.ReadAll(r.Body)
		// 	defer r.Body.Close()
		// 	if err != nil {
		// 		log.Printf("error reading body %v\n", err)
		// 	}
		// 	err = json.Unmarshal(data, &post)
		// 	if err != nil {
		// 		log.Printf("error unmarshaling %v\n", err)
		// 	}
		// 	post.UserID = r.Context().Value(CtxReqIdKey).(string)
		// 	err = h.UseCases.Post.Create(post)
		// 	if err != nil {
		// 		log.Printf("error adding post in handler: %v\n", err)
		// 	}
		// 	// fmt.Println("post from client:", post)
	}
}
