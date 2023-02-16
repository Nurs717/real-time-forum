package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"rtforum/internal/entity"
)

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var user *entity.User
		data, err := io.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			log.Printf("error reading body %v\n", err)
		}
		err = json.Unmarshal(data, &user)
		if err != nil {
			log.Printf("error unmarshaling %v\n", err)
		}
		if err := h.ValidateSignUp(user); err != nil {
			log.Printf("invalid data recieved: %s\n", err.Error())
		}
		err = h.UseCases.Users.NewUser(user)
		if err != nil {
			log.Printf("error occured while singing up: %v\n", err.Error())
			w.Write([]byte("invalid data"))
		}
		fmt.Println("workerd", user)
	}
}

func (h *Handler) ValidateSignUp(user *entity.User) error {
	return nil
}
