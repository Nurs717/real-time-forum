package rest

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
		if err != nil {
			log.Printf("error reading body %v\n", err)
			renderErrorResponse(w, "", err)
		}
		defer r.Body.Close()

		err = json.Unmarshal(data, &user)
		if err != nil {
			log.Printf("error unmarshaling sign up body %v\n", err)
			renderErrorResponse(w, "", err)
			return
		}
		fmt.Printf("user: %+v\n", user)
		if err = h.UseCases.Users.NewUser(r.Context(), user); err != nil {
			log.Printf("rest: create user: %v\n", err)
			renderErrorResponse(w, "creat user", err)
			return
		}
		fmt.Printf("user created")
	}
}
