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
	if r.Context().Value(CtxReqIdKey) == "Guest" {
		w.WriteHeader(http.StatusForbidden)
	}
	switch r.Method {
	case "GET":
		posts, err := h.UseCases.GetAllPosts()
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
	}
}
