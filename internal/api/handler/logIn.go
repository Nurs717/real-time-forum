package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"rtforum/internal/entity"
)

func (h *Handler) LogIn(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var user *entity.User
		data, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			log.Printf("error reading body %v\n", err)
		}
		err = json.Unmarshal(data, &user)
		if err != nil {
			log.Printf("error unmarshaling %v\n", err)
		}
		cookie, err := h.UseCases.SetCookie(user)
		if err != nil {
			log.Printf("Error occured in LogIn handler: %v\n", err)
			http.Error(w, "user not exist", http.StatusUnauthorized)
			return
		}
		fmt.Println(cookie)
		http.SetCookie(w, cookie)
	}
}
