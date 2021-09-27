package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"rtforum/internal/entity"
)

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
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
		err = h.UseCases.Users.NewUser(user)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("workerd", user)
	}
}
